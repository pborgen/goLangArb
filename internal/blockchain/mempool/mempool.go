package mempool

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/paulborgen/goLangArb/internal/blockchain"
)

func GetAllTransactions2() {
    client := blockchain.GetClientWebSocket()
    ctx := context.Background()
    
    pendingTxs := make(chan *types.Transaction)
    
    // Remove the 'true' argument
    sub, err := gethclient.New(client.Client()).SubscribeFullPendingTransactions(ctx, pendingTxs)
    if err != nil {
        log.Fatal().Err(err)
    }
    defer sub.Unsubscribe()


    // Start processing incoming pending transactions
    // for {
    //     select {
    //     case err := <-sub.Err():
    //         log.Fatal(err)
    //     case txs := <-pendingTxSub:
    //         hash := txs.Hash()
    //         fmt.Printf("Pending Transaction Hash: %s\n", hash)

    //     }
    // }

}

func GetAllTransactions() {
 
	client := blockchain.GetClientWebSocket()

	ctx := context.Background()

	pendingTxSub := make(chan common.Hash)

    sub, err := gethclient.New(client.Client()).SubscribePendingTransactions(ctx,pendingTxSub)
    if err != nil {
        log.Fatal().Err(err)
    }
    defer sub.Unsubscribe()

    // Start processing incoming pending transactions
    for {
        select {
        case err := <-sub.Err():
            log.Fatal().Err(err)
        case txs := <-pendingTxSub:
           
            transaction, _, err := client.TransactionByHash(ctx, txs)
            if err != nil {
              log.Fatal().Err(err)
            }
            fmt.Printf("Pending Transaction Hash: %s\n", transaction)

        }
    }

    
}

// getPendingTransactions retrieves pending transactions from the mempool
// func getPendingTransactions(client *ethclient.Client) ([]*ethereum.Transaction, error) {

// 	// Call the eth_pendingTransactions method
// 	var pendingTxs []*ethereum.Transaction
// 	err = rpcClient.CallContext(context.Background(), &pendingTxs, "eth_pendingTransactions")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to call eth_pendingTransactions: %v", err)
// 	}

// 	return pendingTxs, nil
// }
