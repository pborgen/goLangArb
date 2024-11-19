package network

import (
	"github.com/hexlivelive/goBot/internal/database"
	"github.com/hexlivelive/goBot/internal/database/model/orm"
	"strconv"
)

type NetworkModel struct {
	NetworkId        int    `postgres.Table:"NETWORK_ID"`
	Name             string `postgres.Table:"NAME"`
	ChainId          int    `postgres.Table:"CHAIN_ID"`
	Symbol           string `postgres.Table:"CURRENCY_SYMBOL"`
	BlockExplorerUrl string `postgres.Table:"BLOCK_EXPLORER_URL"`
	BaseTokenAddress string `postgres.Table:"BASE_TOKEN_CONTRACT_ADDRESS"`
}

const networkTableName = "NETWORK"
const networkPrimaryKey = "NETWORK_ID"

var networkColumnNames = orm.GetColumnNames(NetworkModel{})

func Insert(networkModel NetworkModel) NetworkModel {
	db := database.GetDBConnection()

	sqlStatement := orm.CreateInsertStatement(NetworkModel{}, networkTableName, networkPrimaryKey)

	id := 0

	err := db.QueryRow(sqlStatement, networkModel.Name, networkModel.ChainId, networkModel.Symbol, networkModel.BlockExplorerUrl, networkModel.BaseTokenAddress).Scan(&id)
	if err != nil {
		panic(err)
	}

	return GetById(id)
}

func GetById(id int) NetworkModel {

	db := database.GetDBConnection()

	rows := db.QueryRow(
		"SELECT "+networkColumnNames+" FROM "+networkTableName+" WHERE NETWORK_ID = $1", id)
	if rows == nil {
		panic("NetworkModel does not exists with id:" + strconv.Itoa(id))
	}

	networkModel, err := scan(rows)

	if err != nil {
		panic("Error")
	}
	return *networkModel
}

func GetByName(name string) NetworkModel {

	db := database.GetDBConnection()

	rows := db.QueryRow(
		"SELECT "+networkColumnNames+" FROM "+networkTableName+" WHERE NAME = $1", name)
	if rows == nil {
		panic("NetworkModel with name does not exists. Name:" + name)
	}

	network, err := scan(rows)

	if err != nil {
		panic("Error")
	}
	return *network
}

func GetAll() []NetworkModel {

	db := database.GetDBConnection()

	results := make([]NetworkModel, 0)
	rows, err := db.Query("SELECT " + networkColumnNames + " FROM " + networkTableName)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//var network NetworkModel
	for rows.Next() {
		network, err := scan(rows)
		if err != nil {
			panic(err)
		}
		results = append(results, *network)
	}

	return results
}

func scan(rows orm.Scannable) (*NetworkModel, error) {
	var networkModel NetworkModel

	err := rows.Scan(&networkModel.NetworkId, &networkModel.Name, &networkModel.ChainId, &networkModel.Symbol, &networkModel.BlockExplorerUrl, &networkModel.BaseTokenAddress)
	if err != nil {
		return nil, err
	}
	return &networkModel, nil
}
