package pair

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hexlivelive/goBot/internal/database"
	"github.com/hexlivelive/goBot/internal/database/model/dex"
	"github.com/hexlivelive/goBot/internal/database/model/erc20"
	"github.com/hexlivelive/goBot/internal/database/model/orm"
	myUtil "github.com/hexlivelive/goBot/internal/util"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/semaphore"
	"math/big"
	"strconv"
	"sync"
)

type ModelPair struct {
	PairId              int            `postgres.Table:"PAIR_ID"`
	DexId               int            `postgres.Table:"DEX_ID"`
	PairIndex           int            `postgres.Table:"PAIR_INDEX"`
	PairContractAddress common.Address `postgres.Table:"PAIR_CONTRACT_ADDRESS"`
	Token0Erc20Id       int            `postgres.Table:"TOKEN0_ID"`
	Token1Erc20Id       int            `postgres.Table:"TOKEN1_ID"`
	Token0Erc20         erc20.ModelERC20
	Token1Erc20         erc20.ModelERC20
	ModelDex            dex.ModelDex
	Token0Reserves      big.Int `postgres.Table:"TOKEN0_RESERVES"`
	Token1Reserves      big.Int `postgres.Table:"TOKEN1_RESERVES"`
	ShouldFindArb       bool    `postgres.Table:"SHOULD_FIND_ARB"`
}

type PairsOrganized struct {
	plsPairs    []ModelPair
	nonPlsPairs []ModelPair
}

const primaryKey = "PAIR_ID"
const tableName = "PAIR"

var wplsAddress = common.HexToAddress("0xa1077a294dde1b09bb078844df40758a5d0f9a27")

var pairColumnNames = orm.GetColumnNames(ModelPair{})

func Insert(pairModel ModelPair) (*ModelPair, error) {
	db := database.GetDBConnection()

	// Check if ERC20's Exists
	erc20.GetById(pairModel.Token0Erc20Id)
	erc20.GetById(pairModel.Token1Erc20Id)

	sqlStatement := orm.CreateInsertStatement(ModelPair{}, tableName, primaryKey)

	id := 0

	err := db.QueryRow(sqlStatement, pairModel.DexId, pairModel.PairIndex, pairModel.PairContractAddress.String(), pairModel.Token0Erc20Id, pairModel.Token1Erc20Id, pairModel.Token0Reserves.String(), pairModel.Token1Reserves.String(), pairModel.ShouldFindArb).Scan(&id)
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("Pair Inserted with address + " + pairModel.PairContractAddress.String())
	return GetById(id)
}

// A function that check if a pair exists
func ExistsByContractAddress(contractAddress common.Address) bool {

	exists := false
	count := 0

	db := database.GetDBConnection()
	row := db.QueryRow(
		"SELECT COUNT(1) FROM "+tableName+" WHERE PAIR_CONTRACT_ADDRESS = $1", contractAddress.String())

	row.Scan(&count)

	if count == 1 {
		exists = true
	}

	return exists
}

func UpdateReserves(pairId int, reserve0 big.Int, reserve1 big.Int) error {
	db := database.GetDBConnection()

	sql := "UPDATE " + tableName + " SET TOKEN0_RESERVES = $1, TOKEN1_RESERVES = $2 WHERE PAIR_ID = $3"

	_, err := db.Exec(sql, reserve0.String(), reserve1.String(), pairId)
	if err != nil {
		return err
	}

	return nil
}

func GetById(id int) (*ModelPair, error) {

	db := database.GetDBConnection()

	rows := db.QueryRow(
		"SELECT "+pairColumnNames+" FROM "+tableName+" WHERE PAIR_ID = $1", id)

	if rows == nil {
		return nil, errors.New("Could not find pairId:" + strconv.Itoa(id))
	}

	pairModel, err := scan(rows)

	if err != nil {
		return nil, err
	}

	return pairModel, nil
}

func GetByContractAddress(contractAddress common.Address) (*ModelPair, error) {

	//var pairModel ModelPair

	db := database.GetDBConnection()
	contractAddressAsString := contractAddress.String()
	sql := "SELECT " + pairColumnNames + " FROM " + tableName + " WHERE PAIR_CONTRACT_ADDRESS = $1"
	rows := db.QueryRow(sql, contractAddressAsString)

	pairModel := &ModelPair{}

	pairModel, err := scan(rows)

	if err != nil {
		return nil, err
	}

	return pairModel, nil
}

func GetAll() ([]ModelPair, error) {

	db := database.GetDBConnection()

	results := make([]ModelPair, 0)
	rows, err := db.Query("SELECT " + pairColumnNames + " FROM " + tableName)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pair, err := scan(rows)

		if err != nil {

			return nil, err
		}
		results = append(results, *pair)
	}

	return results, nil
}

func GetAllPairsThatHavePls() ([]ModelPair, error) {

	db := database.GetDBConnection()

	wplsModelErc20 := erc20.GetByContractAddress(wplsAddress)

	results := make([]ModelPair, 0)
	rows, err := db.Query("SELECT "+pairColumnNames+" FROM "+tableName+" WHERE TOKEN0_ID = $1 OR TOKEN1_ID = $1", wplsModelErc20.Erc20Id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pair, err := scan(rows)

		if err != nil {

			return nil, err
		}
		results = append(results, *pair)
	}

	return results, nil
}
func GetAllPairs() ([]ModelPair, error) {

	db := database.GetDBConnection()
	totalRows := 0
	row := db.QueryRow("SELECT COUNT(1) FROM " + tableName)

	row.Scan(&totalRows)

	log.Print(totalRows)
	limit := 1000
	currentPage := 1
	offset := limit * (currentPage - 1)

	ch := make(chan []ModelPair, 1) // Creating a buffered channel with a capacity of 2
	var wg sync.WaitGroup
	var sem = semaphore.NewWeighted(int64(50))
	ctx := context.Background()

	results := make([]ModelPair, 0)

	sql := "SELECT " + pairColumnNames + " FROM " + tableName + " ORDER BY PAIR_ID asc LIMIT $1 OFFSET $2"

	stop := false

	for offset < totalRows {
		//for count < 3 {
		wg.Add(1)

		sem.Acquire(ctx, 1)
		go query(sem, ch, &wg, sql, limit, offset)

		if stop {
			break
		}

		currentPage = currentPage + 1
		offset = limit * (currentPage - 1)

		// Is the next execution going to get all the rows
		if offset+limit >= totalRows {
			stop = true
		}
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for result := range ch {

		results = append(results, result...) // Process the results
	}

	return results, nil
}

func GetAllPairsOnDex(dexId int) ([]ModelPair, error) {

	db := database.GetDBConnection()

	results := make([]ModelPair, 0)
	rows, err := db.Query("SELECT "+pairColumnNames+" FROM "+tableName+" WHERE DEX_ID != $1", dexId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pair, err := scan(rows)

		if err != nil {

			return nil, err
		}
		results = append(results, *pair)
	}

	return results, nil
}

func GetAllPairsWithOutPls() ([]ModelPair, error) {

	db := database.GetDBConnection()

	wplsModelErc20 := erc20.GetByContractAddress(wplsAddress)

	results := make([]ModelPair, 0)
	rows, err := db.Query("SELECT "+pairColumnNames+" FROM "+tableName+" WHERE TOKEN0_ID != $1 AND TOKEN1_ID != $1", wplsModelErc20.Erc20Id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pair, err := scan(rows)

		if err != nil {

			return nil, err
		}
		results = append(results, *pair)
	}

	return results, nil
}

//func GetAllPairsOrganized() ([]ModelPair, []ModelPair, error) {
//	// Get all the pairs that have pls in it
//	pairs, err := GetAll()
//
//	var plsPairs []ModelPair
//	var nonPlsPairs []ModelPair
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	for _, modelPair := range pairs {
//		if modelPair.Token0Erc20.ContractAddress == wplsAddress || modelPair.Token1Erc20.ContractAddress == wplsAddress {
//			plsPairs = append(plsPairs, modelPair)
//		} else {
//			nonPlsPairs = append(nonPlsPairs, modelPair)
//		}
//	}
//
//	return plsPairs, nonPlsPairs, nil
//}

func GetLargestPairIndex(dexModel dex.ModelDex) (int, error) {
	db := database.GetDBConnection()

	var returnValue int = -1
	var count sql.NullInt64

	row := db.QueryRow("SELECT MAX(PAIR_INDEX) FROM "+tableName+" WHERE DEX_ID=$1", dexModel.DexId)
	err := row.Scan(&count)

	if err != nil {
		return -1, err
	}

	if count.Valid {
		returnValue = int(count.Int64)
	}

	return returnValue, nil
}

func GetNonPlsAddress(modelPair ModelPair) common.Address {

	return GetNonPlsERC20(modelPair).ContractAddress
}

func GetNonPlsERC20(modelPair ModelPair) erc20.ModelERC20 {
	token0ContractAddress := modelPair.Token0Erc20.ContractAddress
	token1ContractAddress := modelPair.Token1Erc20.ContractAddress

	if token0ContractAddress != wplsAddress {
		return modelPair.Token0Erc20
	} else if token1ContractAddress != wplsAddress {
		return modelPair.Token1Erc20
	} else {
		panic("modelPair did not have a wpls address in it. ModlePareId:" + strconv.Itoa(modelPair.PairId))
	}
}

func query(sem *semaphore.Weighted, ch chan []ModelPair, wg *sync.WaitGroup, sql string, limit int, offset int) {
	defer wg.Done()
	db := database.GetDBConnection()

	results := make([]ModelPair, 0)

	rows, err := db.Query(sql, limit, offset)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		pair, err := scan(rows)

		if err != nil {

			panic(err)
		}
		results = append(results, *pair)
	}

	log.Log().Msgf("Bock Complete")
	sem.Release(1)
	ch <- results
}

func scan(rows orm.Scannable) (*ModelPair, error) {
	pairModel := ModelPair{}

	var tempPairContractAddress string
	var tempToken0Reserves []uint8
	var tempToken1Reserves []uint8

	err := rows.Scan(
		&pairModel.PairId,
		&pairModel.DexId,
		&pairModel.PairIndex,
		&tempPairContractAddress,
		&pairModel.Token0Erc20Id,
		&pairModel.Token1Erc20Id,
		&tempToken0Reserves,
		&tempToken1Reserves,
		&pairModel.ShouldFindArb)

	if err != nil {
		return &ModelPair{}, err
	}

	pairModel.PairContractAddress = common.HexToAddress(tempPairContractAddress)
	pairModel.Token0Reserves = *new(big.Int).SetBytes(tempToken0Reserves)
	pairModel.Token1Reserves = *new(big.Int).SetBytes(tempToken1Reserves)

	// Hydrate
	out1 := myUtil.MyAsync(func() erc20.ModelERC20 {
		return erc20.GetById(pairModel.Token0Erc20Id)
	})
	out2 := myUtil.MyAsync(func() erc20.ModelERC20 {
		return erc20.GetById(pairModel.Token1Erc20Id)
	})
	out3 := myUtil.MyAsync(func() dex.ModelDex {
		return dex.GetById(pairModel.DexId)
	})
	pairModel.Token0Erc20 = <-out1
	pairModel.Token1Erc20 = <-out2
	pairModel.ModelDex = <-out3

	return &pairModel, nil
}

func ToString(modelPair ModelPair) string {
	var returnValue = fmt.Sprintf("PairId: %d ", modelPair.PairId)
	returnValue = returnValue + fmt.Sprintf("DexId: %d ", modelPair.DexId)
	returnValue = returnValue + fmt.Sprintf("PairIndex: %d ", modelPair.PairIndex)
	returnValue = returnValue + fmt.Sprintf("PairContractAddress: %s ", modelPair.PairContractAddress.String())
	returnValue = returnValue + fmt.Sprintf("Token0Erc20Id: %d ", modelPair.Token0Erc20Id)
	returnValue = returnValue + fmt.Sprintf("Token1Erc20Id: %d ", modelPair.Token1Erc20Id)
	returnValue = returnValue + fmt.Sprintf("Token0: %s ", erc20.ToString(modelPair.Token0Erc20))
	returnValue = returnValue + fmt.Sprintf("Token1: %s ", erc20.ToString(modelPair.Token1Erc20))
	returnValue = returnValue + fmt.Sprintf("Token0Reserves: %s ", modelPair.Token0Reserves.String())
	returnValue = returnValue + fmt.Sprintf("Token1Reserves: %s ", modelPair.Token1Reserves.String())
	returnValue = returnValue + fmt.Sprintf("ShouldFindArb: %t ", modelPair.ShouldFindArb)
	return returnValue

}
