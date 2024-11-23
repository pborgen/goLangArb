package simpleV2ArbHelper

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	simplev2arbabi "github.com/paulborgen/goLangArb/abi/arbSimple"
	"github.com/paulborgen/goLangArb/internal/blockchain"
	myUtil "github.com/paulborgen/goLangArb/internal/util"
	"github.com/rs/zerolog/log"
)

var simpleArbContractAddress = common.HexToAddress("0x7021DaEaBDaacfe7512d607F1B3b509e08cf441E")

type ArbError struct {
	Err                     error
	RevertReason            string
	DidExecuteTransaction   bool
	IsTransactionSuccessful bool
}

func (w *ArbError) Error() string {
	return fmt.Sprintf("DidExecuteTransaction: %t, IsTransactionSuccessful: %t,  %v", w.DidExecuteTransaction, w.IsTransactionSuccessful, w.Err)
}

func MakeArbitrageSimple(
	opts *bind.TransactOpts,
	amountToken0 *big.Int,
	expectedFinalAmount *big.Int,
	routerAddress0 common.Address,
	routerAddress1 common.Address,
	token0Address common.Address,
	token1Address common.Address) error {

	client := blockchain.GetClient()
	contract, err := simplev2arbabi.NewSimplev2arbabiTransactor(simpleArbContractAddress, client)

	if err != nil {
		return err
	}
	log.Log().Msgf("------------- Executing Arb With Contract ----------------")
	log.Log().Msgf(
		"Calling Contract MakeArbitrageSimple: Params: amountToken0: %s, expectedFinalAmount: %s, routerAddress0: %s, routerAddress1: %s, token0Address: %s, token1Address: %s",
		amountToken0.String(), expectedFinalAmount.String(), routerAddress0.String(), routerAddress1.String(), token0Address.String(), token1Address.String())
	log.Log().Msgf("-------------______________________________----------------")

	//	transaction, err := contract.MakeArbitrageSimple(opts, amountToken0, expectedFinalAmount, routerAddress0, routerAddress1, token0Address, token1Address)

	log.Log().Msgf("About to eceute a transaction with address %s", opts.From.String())
	transaction, err := contract.MakeArbitrageSimpleNoCheck(opts, amountToken0, routerAddress0, routerAddress1, token0Address, token1Address)

	log.Log().Msgf("txHash: %s", transaction.Hash().String())

	if err != nil {
		myError := &ArbError{
			DidExecuteTransaction: true,
			Err:                   err,
		}

		return myError
	}

	receipt, err := bind.WaitMined(context.Background(), client, transaction)

	if err != nil {
		log.Log().Msgf("Error waiting for transaction to be mined: %v", err)

		myError := &ArbError{
			DidExecuteTransaction: true,
			Err:                   err,
		}

		return myError
	}

	isSuccessful, err := myUtil.IsTransactionSuccess(receipt)

	if err != nil {
		myError := &ArbError{
			DidExecuteTransaction:   true,
			IsTransactionSuccessful: isSuccessful,
			Err:                     err,
		}

		return myError
	}

	if isSuccessful {
		log.Log().Msgf("We got One!!!")
	} else {
		revertReason, err := myUtil.GetRevertReason(opts.From, transaction)

		myError := &ArbError{
			DidExecuteTransaction:   true,
			RevertReason:            revertReason,
			IsTransactionSuccessful: isSuccessful,
			Err:                     err,
		}

		return myError
	}
	log.Log().Msgf("TxHash:", receipt.TxHash.String())

	return err
}
