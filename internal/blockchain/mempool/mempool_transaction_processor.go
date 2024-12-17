package mempool

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	

	"github.com/paulborgen/goLangArb/internal/blockchain"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type MethodInfo struct {
    ID   string
    Name string
    Description string
}

var PulseXMethods = map[string]MethodInfo{
    "18cbafe5": {ID: "18cbafe5", Name: "swapExactTokensForETH", Description: "swapExactTokensForETH(uint256,uint256,address[],address,uint256)"},
    "7ff36ab5": {ID: "7ff36ab5", Name: "swapExactETHForTokens", Description: "swapExactETHForTokens(uint256,address[],address,uint256)"},
    "5c11d795": {ID: "5c11d795", Name: "swapExactTokensForTokensSupportingFeeOnTransferTokens", Description: "swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)"},
    "791ac947": {ID: "791ac947", Name: "swapExactTokensForETHSupportingFeeOnTransferTokens", Description: "swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)"},
    "f305d719": {ID: "f305d719", Name: "addLiquidityETH", Description: "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)"},
    "e8e33700": {ID: "e8e33700", Name: "addLiquidity", Description: "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)"},
    "022c0d9f": {ID: "022c0d9f", Name: "removeLiquidityETHWithPermit", Description: "removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)"},
}

type TransactionProcessor struct {
    client     *ethclient.Client
    routerMap  map[common.Address]dex.ModelDex
    numWorkers int
    jobQueue   chan common.Hash
    logger     zerolog.Logger
}

func NewTransactionProcessor(client *ethclient.Client, routerMap map[common.Address]dex.ModelDex, numWorkers int, logger zerolog.Logger) *TransactionProcessor {
    return &TransactionProcessor{
        client: client,
        routerMap: routerMap,
        numWorkers: numWorkers,
        jobQueue: make(chan common.Hash, 1000),
        logger: logger,
    }
}

func (tp *TransactionProcessor) Start(ctx context.Context) {
    // Start worker pool
    var wg sync.WaitGroup
    for i := 0; i < tp.numWorkers; i++ {
        wg.Add(1)
        go tp.worker(ctx, &wg, i)
    }

	client := blockchain.GetClientWebSocket()

    // Subscribe to pending transactions
    pendingTxs := make(chan common.Hash)
    sub, err := gethclient.New(client.Client()).SubscribePendingTransactions(ctx, pendingTxs)
    if err != nil {
        tp.logger.Fatal().Err(err).Msg("Failed to subscribe to pending transactions")
    }
    defer sub.Unsubscribe()

    // Process incoming transactions
    for {
        select {
        case <-ctx.Done():
            close(tp.jobQueue)
            wg.Wait()
            return
        case err := <-sub.Err():
            tp.logger.Error().Err(err).Msg("Subscription error")
            return
        case hash := <-pendingTxs:
            tp.jobQueue <- hash
        }
    }
}

func (tp *TransactionProcessor) worker(ctx context.Context, wg *sync.WaitGroup, workerID int) {
    defer wg.Done()

    logger := tp.logger.With().Int("worker_id", workerID).Logger()
    logger.Info().Msg("Worker started")

    for {
        select {
        case <-ctx.Done():
            logger.Info().Msg("Worker shutting down")
            return
        case hash, ok := <-tp.jobQueue:
            if !ok {
                logger.Info().Msg("Job queue closed")
                return
            }
            tp.processTransaction(ctx, hash, logger)
        }
    }
}

func (tp *TransactionProcessor) processTransaction(ctx context.Context, hash common.Hash, logger zerolog.Logger) {
    tx, isPending, err := tp.client.TransactionByHash(ctx, hash)
    if err != nil {
        logger.Debug().Err(err).Str("hash", hash.String()).Msg("Failed to get transaction")
        return
    }

    if !isPending || tx == nil || tx.To() == nil {
        return
    }

    toAddress := *tx.To()
    
    if dex, ok := tp.routerMap[toAddress]; ok && dex.Name == "PulseChain_Network_PulseX_V2" {
        
		now := time.Now()
		transactionTime := tx.Time()
		timeDiff := now.Sub(transactionTime)
		logger.Debug().Msg("Time difference between now and transaction time: " + timeDiff.String())
		logger.Debug().Msg("Transaction : " + tx.Hash().String())
		tp.processPulseXTransaction(tx, dex, hash, logger)
    }
}

func (tp *TransactionProcessor) processPulseXTransaction(tx *types.Transaction, dex dex.ModelDex, hash common.Hash, logger zerolog.Logger) {
    data := tx.Data()
    if len(data) < 4 {
        return
    }
	//abiString := dex.RouterAbi
    // TODO: Implement ABI decoding
    // Need to properly import and initialize ABI decoder
    // Example:
    //  abiDecoder, err := abi.JSON(strings.NewReader(abiString)) 
    // if err != nil {
    //    logger.Error().Err(err).Msg("Failed to parse ABI")
    //    return
    // }

	// method, err := abiDecoder.MethodById(data[:4])
	// if err != nil {
	// 	logger.Error().Err(err).Msg("Failed to decode method")
	// 	return
	// }

	// logger.Debug().Msg("Method: " + method.Name)

	// bla, err := method.Inputs.Unpack(data)
	// if err != nil {
	// 	logger.Error().Err(err).Msg("Failed to unpack method")
	// 	return
	// }

	// logger.Debug().Interface("bla", bla).Msg("bla")


    methodID := hex.EncodeToString(data[:4])
    
    logger = logger.With().
        Str("tx", hash.String()).
        Str("to", tx.To().String()).
        Str("methodID", methodID).
        Logger()

    if method, ok := PulseXMethods[methodID]; ok {
        logger.Debug().
            Str("method", method.Name).
            Str("description", method.Description).
            Msg("Processing PulseX transaction")

        switch methodID {
        case "7ff36ab5": // swapExactETHForTokens
            tp.processSwapExactETHForTokens(method.Name, data, tx.Value())
        case "18cbafe5": // swapExactTokensForETH
            //tp.processSwapExactTokensForETH(data, logger)
        }
    }
}

func (tp *TransactionProcessor) processSwapExactETHForTokens(methodName string, data []byte, value *big.Int) {
   

	inputsSigData := data[4:]

	abiString := `[{
    "inputs": [
      {
        "internalType": "uint256",
        "name": "amountOutMin",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "path",
        "type": "address[]"
      },
      {
        "internalType": "address",
        "name": "to",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "deadline",
        "type": "uint256"
      }
    ],
    "name": "swapExactETHForTokens",
    "outputs": [
      {
        "internalType": "uint256[]",
        "name": "amounts",
        "type": "uint256[]"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  }]`

	abiDecoder, err := abi.JSON(strings.NewReader(abiString)) 
    if err != nil {
       log.Error().Err(err).Msg("Failed to parse ABI")
       return
    }

	method2, err := abiDecoder.MethodById(data[:4])
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode method")
		return
	}

	inputsMap := make(map[string]interface{})
	if err := method2.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Error().Err(err).Msg("Failed to unpack into map")
	} else {
		fmt.Println(inputsMap)
	}

	amountOutMin := (inputsMap["amountOutMin"]).(big.Int)
	path := inputsMap["path"]
	//to := inputsMap["to"]
	//deadline := inputsMap["deadline"]
	
    log.Debug().
        Str("value", value.String()).
		Str("amountOutMin", amountOutMin.String()).
        Interface("path", path).
        Msg("Decoded swapExactETHForTokens")
}