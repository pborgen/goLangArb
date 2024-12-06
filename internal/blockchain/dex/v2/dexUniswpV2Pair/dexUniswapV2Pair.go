package dexUniswpV2Pair

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	abi_uniswapv2pair "github.com/paulborgen/goLangArb/abi/dexV2"
	"github.com/paulborgen/goLangArb/internal/blockchain"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
)

type Reserve struct {
	Reserve0 big.Int
	Reserve1 big.Int
}

func GetTokenAddressesForPair(dex dex.ModelDex, pairAddress common.Address) struct {
	Token0Address common.Address
	Token1Address common.Address
} {

	client := blockchain.GetClient()
	contract, err := abi_uniswapv2pair.NewAbiUniswapv2pairCaller(pairAddress, client)

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Call Contract
	token0Address, err := contract.Token0(nil)
	token1Address, err := contract.Token1(nil)

	outStruct := new(struct {
		Token0Address common.Address
		Token1Address common.Address
	})

	outStruct.Token0Address = token0Address
	outStruct.Token1Address = token1Address

	return *outStruct
}

func PopulateReserves(pairAddress common.Address) (*struct {
	Reserve0 big.Int
	Reserve1 big.Int
}, error) {
	const maxRetries = 5
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		client := blockchain.GetClient()
		contract, err := abi_uniswapv2pair.NewAbiUniswapv2pairCaller(pairAddress, client)

		if err != nil {
			lastErr = err
			continue
		}

		reserves, err := contract.GetReserves(nil)
		if err != nil {
			lastErr = err
			continue
		}

		outStruct := new(struct {
			Reserve0 big.Int
			Reserve1 big.Int
		})

		outStruct.Reserve0 = *reserves.Reserve0
		outStruct.Reserve1 = *reserves.Reserve1

		return outStruct, nil
	}

	return nil, fmt.Errorf("failed after %d retries, last error: %v", maxRetries, lastErr)
}

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
