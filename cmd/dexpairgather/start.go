package dexpairgather

import (
	"strconv"
	"sync"
	"time"

	"github.com/paulborgen/goLangArb/internal/database/model/dex"
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
