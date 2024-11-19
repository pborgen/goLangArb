package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}
var dbConnection *sql.DB
var myHost = os.Getenv("POSTGRES_HOST")

func init() {

	if myHost == "" {
		log.Warn().Msg("ENV POSTGRES_HOST is empty. Defaulting to localhost")
		myHost = "localhost"
	}
	log.Info().Msg("------------------------------")
	log.Info().Msg("-------- Postgress -----------")
	log.Info().Msg("Host: " + myHost)
	log.Info().Msg("------------------------------")
}
func GetDBConnection() *sql.DB {

	if dbConnection == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbConnection == nil {

			var (
				host     = myHost
				port     = 5432
				user     = "postgres"
				password = "9217"
				dbname   = "bot"
			)
			psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
				"password=%s dbname=%s sslmode=disable",
				host, port, user, password, dbname)

			db, err := sql.Open("postgres", psqlInfo)

			if err != nil {
				panic("Could not get db connection" + err.Error())
			}
			dbConnection = db
		}
	}

	return dbConnection

}
