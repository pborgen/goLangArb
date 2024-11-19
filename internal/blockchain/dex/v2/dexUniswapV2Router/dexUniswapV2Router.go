package dexUniswapV2Router

import (
	"github.com/ethereum/go-ethereum/common"
	uniswapv2router "github.com/hexlivelive/goBot/abi/V2Router"
	"github.com/hexlivelive/goBot/internal/blockchain"
	"github.com/rs/zerolog/log"
	"math/big"
)

func V2RouterGetFactory(routerAddress common.Address) common.Address {

	client := blockchain.GetClient()
	router, err := uniswapv2router.NewUniswapv2routerCaller(routerAddress, client)

	if err != nil {
		log.Fatal().Msgf("Error creating router instance: %v", err)
	}
	factoryAddress, err := router.Factory(nil)

	return factoryAddress
}

func GetAmountsOut(routerAddress common.Address, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	client := blockchain.GetClient()
	router, err := uniswapv2router.NewUniswapv2routerCaller(routerAddress, client)

	if err != nil {
		log.Fatal().Msgf("Error creating router instance: %v", err)
	}

	amountsOut, err := router.GetAmountsOut(nil, amountIn, path)

	return amountsOut, err
}
