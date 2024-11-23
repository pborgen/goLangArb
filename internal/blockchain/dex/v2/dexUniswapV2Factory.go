package dexUniswapV2Factory

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	pulsexV2Factory "github.com/paulborgen/goLangArb/abi/pulseXV2"
	"github.com/paulborgen/goLangArb/internal/blockchain"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/rs/zerolog/log"
)

func GetAllPairsLength(dexModel dex.ModelDex) int {

	factoryAddress := dexModel.FactoryContractAddress

	client := blockchain.GetClient()
	contract, err := pulsexV2Factory.NewPulsexV2Factory(factoryAddress, client)

	length, err := contract.AllPairsLength(nil)
	if err != nil {
		log.Fatal().Msgf("Error", err)
	}
	return int(length.Uint64())
}

func GetPairAddress(dexModel dex.ModelDex, idOfPair big.Int) common.Address {
	factoryAddress := dexModel.FactoryContractAddress

	client := blockchain.GetClient()
	contract, err := pulsexV2Factory.NewPulsexV2Factory(factoryAddress, client)

	pairAddress, err := contract.AllPairs(nil, &idOfPair)
	if err != nil {
		log.Fatal().Msgf("Error", err)
	}
	return pairAddress

}

//
//func PopulateReserves(uniSwapV2Pair mytypes.UniSwapV2Pair) mytypes.UniSwapV2Pair {
//
//	pairAddress := uniSwapV2Pair.Pair.PairAddress
//	client := blockchain.GetClient()
//
//	contract, err := abi_uniswapv2pair.NewAbiUniswapv2pairCaller(pairAddress, client)
//
//	if err != nil {
//		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
//	}
//
//	reserves, err := contract.GetReserves(nil)
//
//	uniSwapV2Pair.Pair.Token0Reserves = *reserves.Reserve0
//	uniSwapV2Pair.Pair.Token1Reserves = *reserves.Reserve1
//
//	return uniSwapV2Pair
//
//}
//
//func GetPair(factoryAddress common.Address, token0 mytypes.Erc20, token1 mytypes.Erc20) mytypes.UniSwapV2Pair {
//
//	client := blockchain.GetClient()
//	contract, err := pulsexV2Factory.NewPulsexV2Factory(factoryAddress, client)
//
//	if err != nil {
//		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
//	}
//	//pairAddress, err := pulsexhelperInstance.PairFor(nil, factoryAddress, token0.ContractAddress, token1.ContractAddress)
//
//	pairAddress, err := contract.GetPair(nil, token0.ContractAddress, token1.ContractAddress)
//
//	pair := mytypes.Pair{
//		PairAddress:    pairAddress,
//		Token0:         token0,
//		Token1:         token1,
//		Token0Reserves: *big.NewInt(-1),
//		Token1Reserves: *big.NewInt(-1),
//	}
//
//	returnValue := mytypes.UniSwapV2Pair{
//		pair,
//	}
//
//	return returnValue
//}
