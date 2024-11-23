package clientservice

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	blockchainAccount "github.com/paulborgen/goLangArb/internal/blockchain/account"
	"github.com/paulborgen/goLangArb/internal/database/model/account"
)

// MAP[publicKey]isLocked
var isClientLockedMap map[string]bool
var modelAccountLookMap map[string]account.ModelAccount

func init() {
	all := account.GetAll()

	isClientLockedMap = make(map[string]bool)
	modelAccountLookMap = make(map[string]account.ModelAccount)

	for _, modelAccount := range all {
		publicKeyLowerCase := strings.ToLower(modelAccount.PublicKey)

		isClientLockedMap[publicKeyLowerCase] = false
		modelAccountLookMap[publicKeyLowerCase] = modelAccount
	}
}

func GetAndLockClient() (*bind.TransactOpts, error) {

	var foundClient = false
	var foundClientPublicKey = ""

	for publicKey, isLocked := range isClientLockedMap {
		if !isLocked {
			foundClient = true
			foundClientPublicKey = publicKey
			isClientLockedMap[publicKey] = true
			break
		}
	}

	if foundClient {
		modelAccount := modelAccountLookMap[foundClientPublicKey]
		auth, err := createAuth(modelAccount)

		if err != nil {
			isClientLockedMap[foundClientPublicKey] = false
			return nil, err
		}
		return auth, nil
	} else {
		return nil, errors.New("could not get a client they are all locked")
	}
}

func ReleaseClient(auth *bind.TransactOpts) {
	publicKey := strings.ToLower(auth.From.String())
	isClientLockedMap[publicKey] = false
}

func createAuth(modelAccount account.ModelAccount) (*bind.TransactOpts, error) {
	auth, err := blockchainAccount.GetAuthAccount(modelAccount)

	if err != nil {
		return nil, err
	}
	return auth, nil
}
