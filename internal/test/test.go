package test

import "github.com/rs/zerolog/log"

func Start(param string) {
	log.Info().Msgf("Test processes started")
	log.Info().Msgf(param)
}
