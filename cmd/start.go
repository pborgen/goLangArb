package main

import (
	"github.com/hexlivelive/goBot/cmd/dexpairgather"
	"github.com/hexlivelive/goBot/internal/mylogger"
	"github.com/hexlivelive/goBot/internal/service/simpleV2ArbService"
	"github.com/hexlivelive/goBot/internal/service/triangleArbService"
	"github.com/hexlivelive/goBot/internal/test"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	mylogger.Init()

	log.Info().Msgf("Starting...")

	log.Info().Msgf("-----------------------------")
	log.Info().Msgf("POSTGRES_HOST: " + os.Getenv("POSTGRES_HOST"))
	log.Info().Msgf("-----------------------------")

	args := os.Args
	log.Info().Msgf("-----------------------------")
	log.Info().Msgf("Type of Args = %T\n", args)

	if len(args) < 2 {
		panic("Invalid params passed")
	}
	log.Info().Msgf(args[0], args[1])
	log.Info().Msgf("-----------------------------")

	processName := args[1]
	log.Info().Msgf("About to run process with Name: " + processName)

	if processName == "test" {
		test.Start(processName)
	} else if processName == "gatherPairs" {
		dexpairgather.Start()
	} else if processName == "findSimpleArb" {
		simpleV2ArbService.Find()
	} else if processName == "findTriangleArb" {
		triangleArbService.Find()
	} else if processName == "monitorSimpleArb" {
		simpleV2ArbService.Monitor()
	} else {
		log.Error().Msgf("Invalid process Name: " + processName)
	}
	log.Info().Msgf("Completed...")

}
