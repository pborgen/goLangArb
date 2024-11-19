package erc20

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hexlivelive/goBot/internal/database"
	"github.com/hexlivelive/goBot/internal/database/model/orm"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ModelERC20 struct {
	Erc20Id         int            `postgres.Table:"ERC20_ID"`
	NetworkId       int            `postgres.Table:"NETWORK_ID"`
	ContractAddress common.Address `postgres.Table:"CONTRACT_ADDRESS"`
	Name            string         `postgres.Table:"NAME"`
	Symbol          string         `postgres.Table:"SYMBOL"`
	Decimal         uint8          `postgres.Table:"DECIMAL"`
	ShouldFindArb   bool           `postgres.Table:"SHOULD_FIND_ARB"`
	IsValidated     bool           `postgres.Table:"IS_VALIDATED"`
}

const primaryKey = "ERC20_ID"
const tableName = "ERC20"

var columnNames = orm.GetColumnNames(ModelERC20{})

func Insert(erc20Model ModelERC20) (ModelERC20, error) {
	db := database.GetDBConnection()

	sqlStatement := orm.CreateInsertStatement(ModelERC20{}, tableName, primaryKey)

	id := 0
	var returnValue ModelERC20

	err := db.QueryRow(sqlStatement, erc20Model.NetworkId, erc20Model.ContractAddress.String(), erc20Model.Name, erc20Model.Symbol, erc20Model.Decimal, erc20Model.ShouldFindArb, erc20Model.IsValidated).Scan(&id)
	if err == nil {
		returnValue = GetById(id)
	}

	return returnValue, err
}

func Update(modelErc20 ModelERC20) (ModelERC20, error) {
	db := database.GetDBConnection()
	sqlStatement := `
	UPDATE ERC20 
	SET NETWORK_ID=$2, CONTRACT_ADDRESS=$3, NAME=$4, SYMBOL=$5, DECIMAL=$6, SHOULD_FIND_ARB=$7, IS_VALIDATED=$8
	WHERE ERC20_ID=$1
	`
	_, err := db.Exec(sqlStatement, modelErc20.Erc20Id, modelErc20.NetworkId, modelErc20.ContractAddress.String(), modelErc20.Name, modelErc20.Symbol, modelErc20.Decimal, modelErc20.ShouldFindArb, modelErc20.IsValidated)
	return modelErc20, err
}

func GetById(id int) ModelERC20 {

	db := database.GetDBConnection()

	sql := "SELECT " + columnNames + " FROM " + tableName + " WHERE ERC20_ID = $1"
	row := db.QueryRow(sql, id)

	erc20Model, err := scan(row)

	if err != nil {
		log.Fatal().Msgf(err.Error())
		panic("Erc20 does not exists with id:" + strconv.Itoa(id))
	}

	return *erc20Model
}

func GetByContractAddress(contractAddress common.Address) ModelERC20 {

	db := database.GetDBConnection()

	row := db.QueryRow(
		"SELECT "+columnNames+" FROM "+tableName+" WHERE CONTRACT_ADDRESS = $1", contractAddress.String())

	erc20, err := scan(row)

	if err != nil {
		panic("Error in Erc20GetByContractAddress")
	}

	return *erc20
}

func GetAll() []ModelERC20 {

	db := database.GetDBConnection()

	results := make([]ModelERC20, 0)
	rows, err := db.Query("SELECT " + columnNames + " FROM " + tableName)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//var erc20 ERC20Model
	for rows.Next() {
		erc20, err := scan(rows)
		if err != nil {

			panic(err)
		}
		results = append(results, *erc20)
	}

	return results

}

func ExistsByContractAddress(contractAddress common.Address) bool {
	return orm.MyExists(tableName, "CONTRACT_ADDRESS", contractAddress.String())
}

func scan(rows orm.Scannable) (*ModelERC20, error) {
	erc20 := ModelERC20{}

	var contractAddressTemp string
	err := rows.Scan(&erc20.Erc20Id, &erc20.NetworkId, &contractAddressTemp, &erc20.Name, &erc20.Symbol, &erc20.Decimal, &erc20.ShouldFindArb, &erc20.IsValidated)
	if err != nil {
		return &ModelERC20{}, err
	}

	erc20.ContractAddress = common.HexToAddress(contractAddressTemp)

	return &erc20, nil
}

func ToString(modelErc20 ModelERC20) string {
	var returnValue = fmt.Sprintf("Erc20Id: %d ", modelErc20.Erc20Id)
	returnValue = returnValue + fmt.Sprintf("NetworkId: %d ", modelErc20.NetworkId)
	returnValue = returnValue + fmt.Sprintf("ContractAddress: %s ", modelErc20.ContractAddress.String())
	returnValue = returnValue + fmt.Sprintf("Name: %s ", modelErc20.Name)
	returnValue = returnValue + fmt.Sprintf("Symbol: %s ", modelErc20.Symbol)
	returnValue = returnValue + fmt.Sprintf("Decimal: %d ", modelErc20.Decimal)
	returnValue = returnValue + fmt.Sprintf("ShouldFindArb: %t ", modelErc20.ShouldFindArb)
	returnValue = returnValue + fmt.Sprintf("IsValidated: %t ", modelErc20.IsValidated)

	return returnValue
}
