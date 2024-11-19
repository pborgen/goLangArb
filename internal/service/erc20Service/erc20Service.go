package erc20Service

import (
	"github.com/ethereum/go-ethereum/common"
	erc20Helper "github.com/hexlivelive/goBot/internal/blockchain/erc20"
	"github.com/hexlivelive/goBot/internal/database/model/erc20"
)

func GetByContractAddress(contractAddress common.Address, networkId int) (erc20.ModelERC20, error) {
	var erc20Token0 erc20.ModelERC20
	var returnError error
	// Get the erc20 token from db. Create if does not exists
	if erc20.ExistsByContractAddress(contractAddress) {
		erc20Token0 = erc20.GetByContractAddress(contractAddress)
	} else {
		erc20Token0Info, err := erc20Helper.GetErc20(contractAddress)

		if err == nil {

			erc20Token0 = erc20.ModelERC20{
				NetworkId:       networkId,
				Name:            erc20Token0Info.Name,
				Symbol:          erc20Token0Info.Symbol,
				ContractAddress: contractAddress,
				Decimal:         erc20Token0Info.Decimal,
				ShouldFindArb:   true,
				IsValidated:     false,
			}
			erc20Token0, err = erc20.Insert(erc20Token0)
			returnError = err
		} else {
			returnError = err
		}
	}

	return erc20Token0, returnError
}
