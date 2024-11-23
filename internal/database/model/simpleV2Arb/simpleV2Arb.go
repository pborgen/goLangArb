package simpleV2Arb

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/paulborgen/goLangArb/internal/database"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/paulborgen/goLangArb/internal/database/model/orm"
	"github.com/paulborgen/goLangArb/internal/database/model/pair"
	myUtil "github.com/paulborgen/goLangArb/internal/util"
	"github.com/rs/zerolog/log"
)

type ModelSimpleV2Arb struct {
	SimpleV2ArbeId int `postgres.Table:"SIMPLE_V2_ARB_ID"`
	Pair0Id        int `postgres.Table:"PAIR0_ID"`
	Pair1Id        int `postgres.Table:"PAIR1_ID"`
	Pair0          pair.ModelPair
	Pair1          pair.ModelPair
	ShouldFindArb  bool   `postgres.Table:"SHOULD_FIND_ARB"`
	FailedArbCount int    `postgres.Table:"FAILED_ARB_COUNT"`
	Note           string `postgres.Table:"NOTE"`
}

const primaryKey = "SIMPLE_V2_ARB_ID"
const tableName = "SIMPLE_V2_ARB"

var simpleV2ArbColumnNames = orm.GetColumnNames(ModelSimpleV2Arb{})

func Insert(simpleV2ArbModel ModelSimpleV2Arb) (*ModelSimpleV2Arb, error) {
	db := database.GetDBConnection()

	sqlStatement := orm.CreateInsertStatement(ModelSimpleV2Arb{}, tableName, primaryKey)

	id := 0

	err := db.QueryRow(sqlStatement, simpleV2ArbModel.Pair0Id, simpleV2ArbModel.Pair1Id, simpleV2ArbModel.ShouldFindArb, simpleV2ArbModel.FailedArbCount, simpleV2ArbModel.Note).Scan(&id)

	if err != nil {
		return nil, err
	}

	log.Info().Msgf("simpleV2Arb Inserted ")
	return GetById(id)
}

func GetById(id int) (*ModelSimpleV2Arb, error) {

	db := database.GetDBConnection()

	rows := db.QueryRow(
		"SELECT "+simpleV2ArbColumnNames+" FROM "+tableName+" WHERE "+primaryKey+" = $1", id)

	if rows == nil {
		return nil, errors.New("Could not get ModelSimpleV2Arb with Id: " + strconv.Itoa(id))
	}

	element, err := scan(rows)

	if err != nil {
		return nil, err
	}

	return element, nil
}

func GetAll() ([]ModelSimpleV2Arb, error) {

	db := database.GetDBConnection()

	results := make([]ModelSimpleV2Arb, 0)
	rows, err := db.Query("SELECT " + simpleV2ArbColumnNames + " FROM SIMPLE_V2_ARB")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		element, err := scan(rows)

		if err != nil {
			return nil, err
		}

		results = append(results, *element)
	}

	return results, nil
}

func IncrementFailedCount(modelSimpleV2Arb *ModelSimpleV2Arb) error {
	db := database.GetDBConnection()

	sql := "UPDATE " + tableName + " SET FAILED_ARB_COUNT = FAILED_ARB_COUNT + 1 WHERE SIMPLE_V2_ARB_ID = $1"

	_, err := db.Exec(sql, modelSimpleV2Arb.SimpleV2ArbeId)
	if err != nil {
		return err
	}

	return nil
}

func SetShouldFindArb(modelSimpleV2Arb *ModelSimpleV2Arb, shouldProcessArb bool) error {
	db := database.GetDBConnection()

	sql := "UPDATE " + tableName + " SET SHOULD_FIND_ARB = $1 WHERE SIMPLE_V2_ARB_ID = $2"

	_, err := db.Exec(sql, shouldProcessArb, modelSimpleV2Arb.SimpleV2ArbeId)
	if err != nil {
		return err
	}

	return nil
}

func Exists(pair0Id int, pair1Id int) (*bool, error) {

	db := database.GetDBConnection()
	var exists = false
	sql := "SELECT EXISTS(SELECT 1 FROM " + tableName + " WHERE PAIR0_ID=$1 AND PAIR1_ID=$2)"

	err := db.QueryRow(sql, pair0Id, pair1Id).Scan(&exists)
	if err != nil {
		return nil, err
	}

	return &exists, nil
}

func GetByPairIds(pair0Id int, pair1Id int) (*ModelSimpleV2Arb, error) {

	db := database.GetDBConnection()

	rows := db.QueryRow(
		"SELECT "+simpleV2ArbColumnNames+" FROM "+tableName+" WHERE PAIR0_ID=$1 AND PAIR1_ID=$2 ", pair0Id, pair1Id)

	if rows == nil {
		return nil, errors.New("Could not find SimpleArbV2 with PairId0=" + strconv.Itoa(pair0Id) + " and PairId1=" + strconv.Itoa(pair1Id))
	}

	element, err := scan(rows)

	if err != nil {
		return nil, err
	}

	return element, nil
}

func scanId(rows orm.Scannable) (*int, error) {
	var id int

	err := rows.Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ToString(modelSimpleArb ModelSimpleV2Arb) string {

	dexPair0 := dex.GetById(modelSimpleArb.Pair0.DexId)
	dexPair1 := dex.GetById(modelSimpleArb.Pair1.DexId)

	var returnValue = fmt.Sprintf("[(Dex0: %s (%s)) (Pair Address: %s) ", dexPair0.Name, dexPair0.RouterContractAddress.String(), modelSimpleArb.Pair0.PairContractAddress.String())
	returnValue += fmt.Sprintf("%s(%s) / %s(%s) ", modelSimpleArb.Pair0.Token0Erc20.Name, modelSimpleArb.Pair0.Token0Erc20.ContractAddress.String(), modelSimpleArb.Pair0.Token1Erc20.Name, modelSimpleArb.Pair0.Token1Erc20.ContractAddress.String())

	returnValue += fmt.Sprintf("] --> [(Dex1: %s (%s)) ", dexPair1.Name, dexPair1.RouterContractAddress.String())
	returnValue += fmt.Sprintf("%s / %s ]", modelSimpleArb.Pair1.Token0Erc20.Name, modelSimpleArb.Pair1.Token1Erc20.Name)

	return returnValue
}

func scan(rows orm.Scannable) (*ModelSimpleV2Arb, error) {
	var element ModelSimpleV2Arb

	err := rows.Scan(
		&element.SimpleV2ArbeId,
		&element.Pair0Id,
		&element.Pair1Id,
		&element.ShouldFindArb,
		&element.FailedArbCount,
		&element.Note)

	if err != nil {
		return nil, err
	}

	// Hydrate
	out1 := myUtil.MyAsync(func() pair.ModelPair {
		modelPair, _ := pair.GetById(element.Pair0Id)
		return *modelPair
	})
	out2 := myUtil.MyAsync(func() pair.ModelPair {
		modelPair, _ := pair.GetById(element.Pair1Id)
		return *modelPair
	})
	element.Pair0 = <-out1
	element.Pair1 = <-out2

	return &element, nil
}
