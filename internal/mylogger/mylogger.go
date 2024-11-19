package mylogger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func Init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
