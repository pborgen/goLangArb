package dex

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paulborgen/goLangArb/internal/database"
	"github.com/paulborgen/goLangArb/internal/database/model/orm"
	"github.com/rs/zerolog/log"
)

type ModelDex struct {
	DexId                  int            `postgres.Table:"DEX_ID"`
	Name                   string         `postgres.Table:"NAME"`
	NetworkId              int            `postgres.Table:"NETWORK_ID"`
	RouterContractAddress  common.Address `postgres.Table:"ROUTER_ADDRESS"`
	FactoryContractAddress common.Address `postgres.Table:"FACTORY_ADDRESS"`
}

var dexColumnNames = orm.GetColumnNames(ModelDex{})

const tableName = "DEX"

func GetById(id int) ModelDex {

	db := database.GetDBConnection()

	rows := db.QueryRow(
		"SELECT "+dexColumnNames+" FROM "+tableName+" WHERE DEX_ID = $1", id)

	dex, err := scan(rows)

	if err != nil {
		panic("Error getting DexModel with id:" + strconv.Itoa(id))
	}
	if rows != nil {
		return *dex
	} else {
		panic("Could not find dex with id:" + strconv.Itoa(id))
	}

}

func GetAllByNetworkId(networkId int) []ModelDex {
	allDexes := GetAll()

	filteredDexes := make([]ModelDex, 0)
	for _, dex := range allDexes {
		if dex.NetworkId == networkId {
			filteredDexes = append(filteredDexes, dex)
		}
	}
	return filteredDexes
}
func GetAll() []ModelDex {

	db := database.GetDBConnection()

	results := make([]ModelDex, 0)
	rows, err := db.Query("SELECT " + dexColumnNames + " FROM " + tableName)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		dex, err := scan(rows)
		if err != nil {

			panic(err)
		}
		results = append(results, *dex)
	}

	return results

}

func scan(rows orm.Scannable) (*ModelDex, error) {
	dex := ModelDex{}

	var routerContractAddressString string
	var factoryContractAddressString string

	err := rows.Scan(&dex.DexId, &dex.Name, &dex.NetworkId, &routerContractAddressString, &factoryContractAddressString)
	if err != nil {
		log.Error().Msgf("Could not scan rows for dex", err)

	}

	dex.RouterContractAddress = common.HexToAddress(routerContractAddressString)
	dex.FactoryContractAddress = common.HexToAddress(factoryContractAddressString)

	return &dex, nil
}
