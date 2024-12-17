package mempool

import (
	"context"

	"errors"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/paulborgen/goLangArb/internal/blockchain"
	"github.com/paulborgen/goLangArb/internal/blockchain/node"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
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

func SubscribeToPendingTransactions() error {
 
    bla := dex.GetById(1)
    log.Info().Msg(bla.Name)
    
	client := blockchain.GetClientWebSocket()

	ctx := context.Background()

	pendingTxSub := make(chan common.Hash)

    sub, err := gethclient.New(client.Client()).SubscribePendingTransactions(ctx, pendingTxSub)
    if err != nil {
        log.Fatal().Err(err)
    }
    defer sub.Unsubscribe()

    isNodeSynced := node.CheckNodeSync(client)

    if !isNodeSynced {
        log.Fatal().Msg("Node is not fully synced")
        return errors.New("node is not fully synced")
    }

    checkNodeSyncCounter := 0


    dexes:= dex.GetAllByNetworkId(1);
   

    routerMap := dex.ModelDexToMapWithRouterAddressAsKey(dexes);

    processor := NewTransactionProcessor(client, routerMap, 10, log.Logger)

    processor.Start(ctx)

    // Start processing incoming pending transactions
    for {

        if checkNodeSyncCounter > 100 {
            isNodeSynced = node.CheckNodeSync(client)
            if !isNodeSynced {
                log.Fatal().Msg("Node is not fully synced")
                return errors.New("node is not fully synced")
            }
            checkNodeSyncCounter = 0
        } else {
            checkNodeSyncCounter++
        }

    }

    log.Info().Msgf("Done")

    return nil
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
