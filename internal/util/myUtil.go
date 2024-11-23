package myUtil

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/paulborgen/goLangArb/internal/blockchain"
)

func MyAsync[T any](f func() T) chan T {
	ch := make(chan T)
	go func() {
		ch <- f()
	}()
	return ch
}

func Wei2EthAsString(amount *big.Int) string {

	compact_amount := big.NewInt(0)
	reminder := big.NewInt(0)
	divisor := big.NewInt(1e18)
	compact_amount.QuoRem(amount, divisor, reminder)
	return fmt.Sprintf("%v.%018s", compact_amount.String(), reminder.String())
}

func IsTransactionSuccess(receipt *types.Receipt) (bool, error) {

	if receipt.Status == types.ReceiptStatusSuccessful {
		return true, nil
	} else {
		return false, nil
	}
}

func FromAddressFromTransaction(transaction *types.Transaction) (common.Address, error) {

	from, err := types.Sender(types.NewEIP155Signer(transaction.ChainId()), transaction)
	if err != nil {
		return common.Address{}, err
	}

	return from, nil
}

func GetRevertReason(from common.Address, tx *types.Transaction) (string, error) {

	client := blockchain.GetClient()

	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		if strings.Contains(err.Error(), "reverted") {
			return err.Error(), nil
		}
		return "", err
	}

	return string(res), nil
}

func parseRevertReason(input []byte) (string, error) {
	if len(input) < 4 {
		return "", fmt.Errorf("invalid input")
	}

	// methodID := input[:4]
	inputData := input[4:]

	parsedRevert, err := abi.JSON(strings.NewReader(abiRevert))
	if err != nil {
		return "", err
	}

	var reason string
	err = parsedRevert.UnpackIntoInterface(&reason, "Error", inputData)
	if err != nil {
		return "", err
	}

	return reason, nil
}

const abiRevert = `[{ "name": "Error", "type": "function", "inputs": [ { "type": "string" } ] }]`
