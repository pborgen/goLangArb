package pairFinder

import (
	"math/big"

	dexUniswapV2Factory "github.com/paulborgen/goLangArb/internal/blockchain/dex/v2"
	"github.com/paulborgen/goLangArb/internal/blockchain/dex/v2/dexUniswpV2Pair"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/paulborgen/goLangArb/internal/database/model/pair"
	"github.com/paulborgen/goLangArb/internal/service/erc20Service"
	"github.com/rs/zerolog/log"
)

func PopulatePairsInDb(dex dex.ModelDex) {

	currentLength := dexUniswapV2Factory.GetAllPairsLength(dex)
	largestPairIndex, _ := pair.GetLargestPairIndex(dex)

	// Add the new pairs
	for pairIndex := largestPairIndex + 1; pairIndex < currentLength; pairIndex++ {

		iAsBigInt := *big.NewInt(int64(pairIndex))

		pairAddress := dexUniswapV2Factory.GetPairAddress(dex, iAsBigInt)

		// Get the 2 taken Addresses for the pair
		result := dexUniswpV2Pair.GetTokenAddressesForPair(dex, pairAddress)
		token0Address := result.Token0Address
		token1Address := result.Token1Address

		erc20Token0, errToken0 := erc20Service.GetByContractAddress(token0Address, dex.NetworkId)
		erc20Token1, errToken1 := erc20Service.GetByContractAddress(token1Address, dex.NetworkId)

		if errToken0 == nil && errToken1 == nil {
			pairModel := pair.ModelPair{
				DexId:               dex.DexId,
				PairIndex:           pairIndex,
				PairContractAddress: pairAddress,
				Token0Erc20Id:       erc20Token0.Erc20Id,
				Token1Erc20Id:       erc20Token1.Erc20Id,
				Token0Reserves:      big.Int{},
				Token1Reserves:      big.Int{},
				ShouldFindArb:       true, // set the correct value
			}

			pair.Insert(pairModel)

		} else {
			if errToken0 != nil {
				log.Warn().Msgf("Could not get Erc20 with Address:" + token0Address.String())
			}
			if errToken1 != nil {
				log.Warn().Msgf("Could not get Erc20 with Address:" + token1Address.String())
			}
		}
	}

	log.Info().Msgf("Completed gathering pairs for " + dex.Name)
}

func UpdateAllReserves() {
	dexs := dex.GetAll()

	for _, dex := range dexs {
		UpdateReserves(dex)
	}
}

func UpdateReservesForPairs(modelPairs []pair.ModelPair) {
	for _, modelPair := range modelPairs {
		UpdateReservesForPair(modelPair)
	}
}

func UpdateReservesForPair(modelPair pair.ModelPair) {
	reserves, err := dexUniswpV2Pair.PopulateReserves(modelPair.PairContractAddress)

	if err != nil {
		panic("bla")
	}

	err2 := pair.UpdateReserves(modelPair.PairId, reserves.Reserve0, reserves.Reserve1)

	if err2 != nil {
		panic("")
	}
}

func UpdateReserves(dex dex.ModelDex) {

	allPairs, _ := pair.GetAllPairsOnDex(dex.DexId)

	// Add the new pairs
	for _, tempPair := range allPairs {

		reserves, err := dexUniswpV2Pair.PopulateReserves(tempPair.PairContractAddress)

		if err != nil {
			panic("bla")
		}

		err2 := pair.UpdateReserves(tempPair.PairId, reserves.Reserve0, reserves.Reserve1)

		if err2 != nil {
			panic("")
		}
	}

	log.Info().Msgf("Completed gathering pairs for " + dex.Name)
}
