package dexpairgather

import (
	"encoding/json"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/paulborgen/goLangArb/internal/database/model/pair"
	"github.com/paulborgen/goLangArb/internal/mylogger"
	"github.com/paulborgen/goLangArb/internal/service/pairFinder"
	"github.com/rs/zerolog/log"
)

type PossibleSandwich struct {
    DexName  string
	DexRouterAddress string
	DexFactoryAddress string

	Token0Address string
	Token1Address string
	PairContractAddress string
}

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

func plsPairWithHighAmountOfPls(dexIds []int, minAmountOfPls *big.Int) []pair.ModelPair {

	returnValue := []pair.ModelPair{}

	allPairs := []pair.ModelPair{}
	for _, dexId := range dexIds {
		allPairsForDexId, err := pair.GetAllPairsThatHavePlsByDexId(dexId)
		if err != nil {
			log.Error().Msgf(err.Error())
		}

		allPairs = append(allPairs, allPairsForDexId...)
	}

	for _, pair := range allPairs {
		shouldAdd := false
		if pair.Token0Erc20Id == 1 {
			if pair.Token0Reserves.Cmp(minAmountOfPls) >= 0 {
				shouldAdd = true
			}
		} else if pair.Token1Erc20Id == 1 {
			if pair.Token1Reserves.Cmp(minAmountOfPls) >= 0 {
				shouldAdd = true
			}
		} else {
			panic("No pls pair found for dexId: " + strconv.Itoa(pair.PairId))
		}

		if shouldAdd {
			returnValue = append(returnValue, pair)
		}

		shouldAdd = false
	}

	return returnValue
}

func WriteToFilePlsPairsByDexId(dexIds []int) {

	millionPls := big.NewInt(0)
    millionPls.SetString("1000000000000000000000000", 10)
	tenMillionPls := new(big.Int).Mul(millionPls, big.NewInt(10))


	var allPairs []pair.ModelPair = plsPairWithHighAmountOfPls(dexIds, tenMillionPls)

	log.Info().Msgf("Found " + strconv.Itoa(len(allPairs)) + " pairs with high amount of pls")

	var myMap = make(map[string]PossibleSandwich)
	for _, pair := range allPairs {
		dex := dex.GetById(pair.DexId)
		myMap[pair.PairContractAddress.String()] = PossibleSandwich{
			DexName: dex.Name,
			DexRouterAddress: dex.RouterContractAddress.String(),
			DexFactoryAddress: dex.FactoryContractAddress.String(),

			Token0Address: pair.Token0Erc20.ContractAddress.String(),
			Token1Address: pair.Token1Erc20.ContractAddress.String(),
			PairContractAddress: pair.PairContractAddress.String(),
		}
	}

	// Convert map to json
	jsonData, err := json.MarshalIndent(myMap, "", "  ")
	if err != nil {
		log.Error().Msgf(err.Error())
	}


	// Write to file
	err = os.WriteFile("possibleSandwiches.json", jsonData, 0644)
	if err != nil {
		log.Error().Msgf(err.Error())
	}

	log.Info().Msgf(string(jsonData))
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
