package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)


func GetClient() *ethclient.Client {

	client, err := ethclient.Dial("https://rpc.pulsechain.com")

	if err != nil {
		log.Fatal().Msgf("Error in GetClient")
	} else {
		log.Log().Msgf("Success! you are connected to the Network")
	}

	return client
}

func GetClientAuth() *ethclient.Client {

	client, err := ethclient.Dial("https://rpc.pulsechain.com")

	if err != nil {

		log.Fatal().Msgf("Error in GetClient")
	} else {
		log.Log().Msgf("Success! you are connected to the Network")
	}

	//log.Log().Msgf("Creating blockchain Client single instance now.")
	//singleInstance = client

	return client
}

func GetBlockNumber() uint64 {
	var client = GetClient()
	header, err := client.HeaderByNumber(context.Background(), nil)
	var blockNumber uint64 = 0

	if err != nil {
		log.Err(err)
	} else {

		blockNumber = uint64(header.Number.Uint64())
	}

	return blockNumber
}
