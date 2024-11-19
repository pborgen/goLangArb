package account

import (
	"github.com/hexlivelive/goBot/internal/database"
)

type ModelAccount struct {
	AccountId   int    `postgres.Table:"ACCOUNT_ID"`
	Name        string `postgres.Table:"NAME"`
	Description string `postgres.Table:"DESCRIPTION"`
	PublicKey   string `postgres.Table:"PUBLIC_KEY"`
	PrivateKey  string `postgres.Table:"PRIVATE_KEY"`
}

func GetById(id int) (*ModelAccount, error) {

	var account ModelAccount

	db := database.GetDBConnection()

	err := db.QueryRow("SELECT * FROM ACCOUNT WHERE account_id = $1", id).Scan(&account.AccountId, &account.Name, &account.Description, &account.PublicKey, &account.PrivateKey)

	if err != nil {
		return nil, err
	}
	return &account, nil
}

func GetAll() []ModelAccount {

	db := database.GetDBConnection()

	results := make([]ModelAccount, 0)
	rows, err := db.Query("SELECT account_id, name, description, public_key, private_key FROM ACCOUNT")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var account ModelAccount
	for rows.Next() {
		err := rows.Scan(&account.AccountId, &account.Name, &account.Description, &account.PublicKey, &account.PrivateKey)
		if err != nil {
			panic(err)
		}
		results = append(results, account)
	}

	return results

}
