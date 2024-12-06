package dexpairgather

import (
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/paulborgen/goLangArb/internal/database/model/pair"
	"github.com/paulborgen/goLangArb/internal/mylogger"
	"github.com/paulborgen/goLangArb/internal/service/pairFinder"
	"github.com/rs/zerolog/log"
)

func Start() {
	mylogger.Init()

	allDexes := dex.GetAllByNetworkId(1)

	var wg sync.WaitGroup
	wg.Add(len(allDexes))

	for _, modelDex := range allDexes {
		go Gather(modelDex.DexId)
	}

	wg.Wait()
}

func UpdateReservesForPairs(dexIds []int) []pair.ModelPair {
	var allPairs []pair.ModelPair

	for _, dexId := range dexIds {
		allPairsForDexId, err := pair.GetAllPairsThatHavePlsByDexId(dexId)

		pairFinder.UpdateReservesForPairs(allPairsForDexId)
		if err != nil {
			log.Error().Msgf(err.Error())
		}

		allPairs = append(allPairs, allPairsForDexId...)
		
	}

	return allPairs
}

func plsPairWithHighAmountOfPls(dexId int) []pair.ModelPair {

	returnValue := []pair.ModelPair{}

	millionPls := big.NewInt(0)
    millionPls.SetString("1000000000000000000000000", 10)
	//tenMillionPls := new(big.Int).Mul(millionPls, big.NewInt(10))
	fiftyMillionPls := new(big.Int).Mul(millionPls, big.NewInt(50))



	allPairsForDexId, err := pair.GetAllPairsThatHavePlsByDexId(dexId)
	if err != nil {
		log.Error().Msgf(err.Error())
	}

	for _, pair := range allPairsForDexId {
		shouldAdd := false
		if pair.Token0Erc20Id == 1 {
			if pair.Token0Reserves.Cmp(fiftyMillionPls) >= 0 {
				shouldAdd = true
			}
		} else if pair.Token1Erc20Id == 1 {
			if pair.Token1Reserves.Cmp(fiftyMillionPls) >= 0 {
				shouldAdd = true
			}
		} else {
			panic("No pls pair found for dexId: " + strconv.Itoa(dexId))
		}

		if shouldAdd {
			returnValue = append(returnValue, pair)
		}

		shouldAdd = false
	}

	return returnValue
}

func WriteToFilePlsPairsByDexId(dexIds []int) {


	var allPairs []pair.ModelPair = plsPairWithHighAmountOfPls(3)

	log.Info().Msgf("Found " + strconv.Itoa(len(allPairs)) + " pairs with high amount of pls")

	for _, pair := range allPairs {
		log.Info().Msgf(pair.PairContractAddress.String())
	}
}

func Gather(dexId int) {
	log.Info().Msgf("Starting dexpairgather.Start with Id:" + strconv.Itoa(dexId))
	dex := dex.GetById(dexId)

	log.Info().Msgf("Gathering new pairs for dex with name:" + dex.Name)

	for {
		pairFinder.PopulatePairsInDb(dex)

		log.Info().Msgf("Completed gathering new pairs for dex with name:" + dex.Name)

		log.Info().Msgf("Sleeping 5 min.....")

		time.Sleep(time.Minute * 5)
	}
}
