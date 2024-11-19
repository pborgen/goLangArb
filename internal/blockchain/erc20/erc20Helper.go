package erc20Helper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/hexlivelive/goBot/abi/erc20"
	"github.com/hexlivelive/goBot/internal/blockchain"
	"github.com/rs/zerolog/log"
)

func GetErc20(contractAddress common.Address) (struct {
	Name    string
	Symbol  string
	Decimal uint8
}, error) {

	log.Info().Msgf("Getting Erc20 with address:" + contractAddress.String())
	client := blockchain.GetClient()

	contract, err := erc20.NewErc20(contractAddress, client)

	name, err := contract.Name(nil)
	symbol, err := contract.Symbol(nil)
	decimal, err := contract.Decimals(nil)

	outStruct := new(struct {
		Name    string
		Symbol  string
		Decimal uint8
	})

	if err == nil {
		outStruct.Name = name
		outStruct.Symbol = symbol
		outStruct.Decimal = decimal
	} else {
		log.Error().Msgf("Could not get blockchain info for token with address" + contractAddress.String())
	}

	return *outStruct, err
}
