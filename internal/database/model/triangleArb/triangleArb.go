package triangleArb

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/paulborgen/goLangArb/internal/database"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/paulborgen/goLangArb/internal/database/model/orm"
	"github.com/paulborgen/goLangArb/internal/database/model/pair"
	"github.com/rs/zerolog/log"
)

type ModelTriangleArb struct {
	TriangleArbId   int `postgres.Table:"TRIANGLE_ARB_ID"`
	Pair0Id         int `postgres.Table:"PAIR0_ID"`
	Pair1Id         int `postgres.Table:"PAIR1_ID"`
	Pair2Id         int `postgres.Table:"PAIR2_ID"`
	Pair0           pair.ModelPair
	Pair1           pair.ModelPair
	Pair2           pair.ModelPair
	ShouldFindArb   bool   `postgres.Table:"SHOULD_FIND_ARB"`
	FailedArbCount  int    `postgres.Table:"FAILED_ARB_COUNT"`
	SuccessArbCount int    `postgres.Table:"SUCCESS_ARB_COUNT"`
	ErrorArbCount   int    `postgres.Table:"ERROR_ARB_COUNT"`
	Note            string `postgres.Table:"NOTE"`
}

const primaryKey = "TRIANGLE_ARB_ID"
const tableName = "TRIANGLE_ARB"

var columnNames = orm.GetColumnNames(ModelTriangleArb{})

func Insert(modelTriangleArb ModelTriangleArb) error {
	db := database.GetDBConnection()

	sqlStatement := orm.CreateInsertStatement(ModelTriangleArb{}, tableName, primaryKey)

	id := 0

	err := db.QueryRow(sqlStatement, modelTriangleArb.Pair0Id, modelTriangleArb.Pair1Id, modelTriangleArb.Pair2Id, modelTriangleArb.ShouldFindArb, modelTriangleArb.FailedArbCount, modelTriangleArb.SuccessArbCount, modelTriangleArb.Note).Scan(&id)

	if err != nil {
		return err
	}

	log.Info().Msgf("triangleArb Inserted ")
	return nil
}

func GetById(id int) (*ModelTriangleArb, error) {

	db := database.GetDBConnection()

	rows := db.QueryRow(
		"SELECT "+columnNames+" FROM "+tableName+" WHERE "+primaryKey+" = $1", id)

	if rows == nil {
		return nil, errors.New("Could not get ModelTriangleArb with Id: " + strconv.Itoa(id))
	}

	element, err := scan(rows)

	if err != nil {
		return nil, err
	}

	return element, nil
}

func GetAll() ([]ModelTriangleArb, error) {

	db := database.GetDBConnection()

	results := make([]ModelTriangleArb, 0)
	rows, err := db.Query("SELECT " + columnNames + " FROM " + tableName)

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

func Exists(pair0Id int, pair1Id int, pair2Id int) (*bool, error) {

	db := database.GetDBConnection()
	var exists = false
	sql := "SELECT EXISTS(SELECT 1 FROM " + tableName + " WHERE PAIR0_ID=$1 AND PAIR1_ID=$2 AND PAIR2_ID=$3)"

	err := db.QueryRow(sql, pair0Id, pair1Id, pair2Id).Scan(&exists)
	if err != nil {
		return nil, err
	}

	return &exists, nil
}

func GetByPairIds(pair0Id int, pair1Id int, pair2Id int) (*ModelTriangleArb, error) {

	db := database.GetDBConnection()

	rows := db.QueryRow(
		"SELECT "+columnNames+" FROM "+tableName+" WHERE PAIR0_ID=$1 AND PAIR1_ID=$2 AND PAIR2_ID=$3 ", pair0Id, pair1Id, pair2Id)

	if rows == nil {
		return nil, errors.New("Could not find triangleArb with PairId0=" + strconv.Itoa(pair0Id) + " and PairId1=" + strconv.Itoa(pair1Id) + " and PairId2=" + strconv.Itoa(pair2Id))
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

func ToString(modelTriangleArb ModelTriangleArb) string {

	dexPair0 := dex.GetById(modelTriangleArb.Pair0.DexId)
	dexPair1 := dex.GetById(modelTriangleArb.Pair1.DexId)
	dexPair2 := dex.GetById(modelTriangleArb.Pair2.DexId)
	var returnValue = fmt.Sprintf("[(Dex0: %s (%s)) (Pair Address: %s) ", dexPair0.Name, dexPair0.RouterContractAddress.String(), modelTriangleArb.Pair0.PairContractAddress.String())
	returnValue += fmt.Sprintf("%s(%s) / %s(%s) ", modelTriangleArb.Pair0.Token0Erc20.Name, modelTriangleArb.Pair0.Token0Erc20.ContractAddress.String(), modelTriangleArb.Pair0.Token1Erc20.Name, modelTriangleArb.Pair0.Token1Erc20.ContractAddress.String())

	returnValue += fmt.Sprintf("] --> [(Dex1: %s (%s)) ", dexPair1.Name, dexPair1.RouterContractAddress.String())
	returnValue += fmt.Sprintf("%s / %s ]", modelTriangleArb.Pair1.Token0Erc20.Name, modelTriangleArb.Pair1.Token1Erc20.Name)

	returnValue += fmt.Sprintf("] --> [(Dex2: %s (%s)) ", dexPair2.Name, dexPair2.RouterContractAddress.String())
	returnValue += fmt.Sprintf("%s / %s ]", modelTriangleArb.Pair2.Token0Erc20.Name, modelTriangleArb.Pair2.Token1Erc20.Name)
	return returnValue
}

func scan(rows orm.Scannable) (*ModelTriangleArb, error) {
	var element ModelTriangleArb

	err := rows.Scan(
		&element.TriangleArbId,
		&element.Pair0Id,
		&element.Pair1Id,
		&element.Pair2Id,
		&element.ShouldFindArb,
		&element.FailedArbCount,
		&element.SuccessArbCount,
		&element.ErrorArbCount,
		&element.Note)

	pair0, err := pair.GetById(element.Pair0Id)
	if err != nil {
		return nil, err
	}
	pair1, err := pair.GetById(element.Pair1Id)
	if err != nil {
		return nil, err
	}
	pair2, err := pair.GetById(element.Pair2Id)
	if err != nil {
		return nil, err
	}

	element.Pair0 = *pair0
	element.Pair1 = *pair1
	element.Pair2 = *pair2

	return &element, nil
}
