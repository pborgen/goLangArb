package simpleV2ArbService

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hexlivelive/goBot/internal/blockchain/dex/v2/dexUniswapV2Router"
	simpleV2ArbHelper "github.com/hexlivelive/goBot/internal/blockchain/simplev2arbhelper"
	"github.com/hexlivelive/goBot/internal/database/model/dex"
	"github.com/hexlivelive/goBot/internal/database/model/pair"
	"github.com/hexlivelive/goBot/internal/database/model/simpleV2Arb"
	clientservice "github.com/hexlivelive/goBot/internal/service/clientService"
	myUtil "github.com/hexlivelive/goBot/internal/util"
	"github.com/rs/zerolog/log"
	"math/big"
	"strings"
	"sync"
)

var wplsAddress = common.HexToAddress("0xa1077a294dde1b09bb078844df40758a5d0f9a27")

var ten = new(big.Int).SetUint64(10)
var onePls = new(big.Int)
var tenPls = new(big.Int)
var hundredPls = new(big.Int)
var thousandPls = new(big.Int)
var tenThousandPls = new(big.Int)

type foundArbsStruct struct {
	pair0 pair.ModelPair
	pair1 pair.ModelPair
}

func init() {
	tempOnePls, _ := new(big.Int).SetString("1000000000000000000", 0)
	onePls = tempOnePls

	log.Log().Msgf("1 PLS: %d", onePls)

	tenPls.Mul(onePls, ten)
	log.Log().Msgf("10 PLS: %d", tenPls)

	hundredPls.Mul(tenPls, ten)
	log.Log().Msgf("100 PLS: %d", tenPls)

	thousandPls.Mul(hundredPls, ten)
	log.Log().Msgf("1000 PLS: %d", tenPls)

	tenThousandPls.Mul(thousandPls, ten)
	log.Log().Msgf("10000 PLS: %d", tenPls)

}
func Monitor() {

	amountIn := new(big.Int)
	amountIn.Mul(hundredPls, new(big.Int).SetUint64(500))

	for {
		modelSimpleArbAll, err := simpleV2Arb.GetAll()
		if err != nil {
			log.Error().Err(err).Msg("Failed to get SimpleArbV2 with Id: 1")
			return
		}

		var totalElements = len(modelSimpleArbAll)
		const elementsToRunInParallel = 30

		var gatheredElements []simpleV2Arb.ModelSimpleV2Arb

		for i, modelSimpleArb := range modelSimpleArbAll {

			isLastElement := i == totalElements

			if modelSimpleArb.ShouldFindArb {
				gatheredElements = append(gatheredElements, modelSimpleArb)
			}

			if len(gatheredElements) == elementsToRunInParallel || isLastElement {
				var wg sync.WaitGroup

				for _, element := range gatheredElements {

					wg.Add(1)

					tempElement := element

					go func() {
						defer wg.Done()
						err := tryArb(tempElement, amountIn, hundredPls)

						if err != nil {

							var cusErr *simpleV2ArbHelper.ArbError

							if errors.As(err, &cusErr) {
								revertReason := cusErr.RevertReason

								if strings.Contains(revertReason, "TRANSFER_FROM_FAILED") ||
									strings.Contains(revertReason, "TransferHelper::transferFrom: transferFrom failed") {

									simpleV2Arb.IncrementFailedCount(&tempElement)
									simpleV2Arb.SetShouldFindArb(&tempElement, false)
								}
							}

						}

					}()

				}

				wg.Wait()

				// Clear the array
				gatheredElements = []simpleV2Arb.ModelSimpleV2Arb{}

			} else {
				log.Log().Msgf("done with group")
			}
		}
		log.Log().Msgf("Scanned all Simple Arbs")
	}
}

func tryArb(modelSimpleV2Arb simpleV2Arb.ModelSimpleV2Arb, amountIn *big.Int, profitThreshold *big.Int) error {
	modelPair0 := modelSimpleV2Arb.Pair0
	modelPair1 := modelSimpleV2Arb.Pair1

	// Dexes to use
	modleDexPair0 := dex.GetById(modelPair0.DexId)
	modleDexPair1 := dex.GetById(modelPair1.DexId)

	pair0AmountIn := amountIn

	nonPlsERC20ModelPair0 := pair.GetNonPlsERC20(modelPair0)

	// Start First Swap
	var pathModelPair0 []common.Address
	pathModelPair0 = append(pathModelPair0, wplsAddress, nonPlsERC20ModelPair0.ContractAddress)

	log.Log().Msgf(simpleV2Arb.ToString(modelSimpleV2Arb))

	amountsOutPair0, err := dexUniswapV2Router.GetAmountsOut(modleDexPair0.RouterContractAddress, pair0AmountIn, pathModelPair0)
	if err != nil {
		log.Log().Msgf("Error Finding arb Error:" + err.Error())
		return err
	}

	pair0AmountOut := amountsOutPair0[1]

	log.Log().Msgf("Pair0 In Amount: %s ", amountsOutPair0[0].String())
	log.Log().Msgf("Pair0 Out Amount: %s ", pair0AmountOut.String())

	// Start Second swap
	pair1AmountIn := amountsOutPair0[1]

	var pathModelPair1 []common.Address
	pathModelPair1 = append(pathModelPair1, nonPlsERC20ModelPair0.ContractAddress, wplsAddress)

	amountsOutPair1, err := dexUniswapV2Router.GetAmountsOut(modleDexPair1.RouterContractAddress, pair1AmountIn, pathModelPair1)
	if err != nil {
		return err
	}

	pair1AmountOut := amountsOutPair1[1]

	log.Log().Msgf(simpleV2Arb.ToString(modelSimpleV2Arb))

	hasProfit := pair1AmountOut.Cmp(pair0AmountIn) == 1
	log.Log().Msgf("pair0AmountIn : %d", pair0AmountIn)
	log.Log().Msgf("pair1AmountOut: %d", pair1AmountOut)

	//hasProfit = true

	if hasProfit {
		profit := new(big.Int)
		profit.Sub(pair1AmountOut, pair0AmountIn)

		log.Log().Msgf("profit         : %d", profit)
		log.Log().Msgf("profitThreshold: %d", profitThreshold)

		isProfitGreaterThenThreshold := profit.Cmp(profitThreshold) == 1

		//isProfitGreaterThenThreshold = true
		if isProfitGreaterThenThreshold {
			log.Log().Msgf("----------------> There IS profit of %s", myUtil.Wei2EthAsString(profit))

			err = executeArb(
				pair0AmountIn,
				pair1AmountOut,
				modelSimpleV2Arb)

			if err != nil {
				return err
			}
		} else {
			log.Log().Msgf("There is NOT Enough profit")
		}

	} else {
		log.Log().Msgf("There is NOT profit")
	}

	return nil
}

func Find() error {

	var foundArray []foundArbsStruct

	plsPairs, err := pair.GetAllPairsThatHavePls()
	if err != nil {
		log.Error().Msgf("Error getting pairs that have pls: %v", err)
		return err
	}

	for _, pair0 := range plsPairs {
		// Get the side of the pair that is not pls

		var nonPlsSidePair0 common.Address
		if pair0.Token0Erc20.ContractAddress != wplsAddress {
			nonPlsSidePair0 = pair0.Token0Erc20.ContractAddress
		} else {
			nonPlsSidePair0 = pair0.Token1Erc20.ContractAddress
		}
		var pair0DexId = pair0.DexId
		// Find Pairs on other dexs
		for _, pair1 := range plsPairs {

			var nonPlsSidePair1 common.Address
			if pair1.Token0Erc20.ContractAddress != wplsAddress {
				nonPlsSidePair1 = pair1.Token0Erc20.ContractAddress
			} else {
				nonPlsSidePair1 = pair1.Token1Erc20.ContractAddress
			}

			// Check if we found a arb
			onDifferentDex := pair1.DexId != pair0DexId
			hasCommonNonPulseToken := nonPlsSidePair0 == nonPlsSidePair1

			if hasCommonNonPulseToken && onDifferentDex {
				// Found a arb
				foundArray = append(foundArray, foundArbsStruct{
					pair0: pair0,
					pair1: pair1,
				})

				log.Log().Msgf("Found Simple Arb. Pair0: %s ---------- Pair1: %s", pair.ToString(pair0), pair.ToString(pair1))

			}
		}
	}

	for _, found := range foundArray {
		modelSimpleV2Arb := simpleV2Arb.ModelSimpleV2Arb{
			Pair0Id:        found.pair0.PairId,
			Pair1Id:        found.pair1.PairId,
			ShouldFindArb:  true,
			FailedArbCount: 0,
			Note:           "",
		}

		exists, err := simpleV2Arb.Exists(modelSimpleV2Arb.Pair0Id, modelSimpleV2Arb.Pair1Id)

		if err != nil {
			panic("should not happen" + err.Error())
		}

		if !*exists {

			_, err := simpleV2Arb.Insert(modelSimpleV2Arb)

			if err != nil {
				panic("should not happen." + err.Error())
			}

		}
	}

	log.Log().Msgf("Completed finding simple arbs. Found %d arbc", len(foundArray))

	return nil
}

func executeArb(

	amountToken0 *big.Int,
	expectedFinalAmount *big.Int,
	modelSimpleV2Arb simpleV2Arb.ModelSimpleV2Arb) error {

	auth, err := clientservice.GetAndLockClient()

	if err != nil {
		log.Error().Err(err).Msg("Failed to get client")
		return err
	}

	tip := big.NewInt(10000000000) // maxPriorityFeePerGas = 2 Gwei
	//feeCap := big.NewInt(20000000000)
	gasLimit := uint64(400000)

	auth.GasLimit = gasLimit
	auth.GasTipCap = tip
	auth.GasFeeCap = nil
	auth.GasPrice = nil

	var nonPlsAddress = common.HexToAddress("")

	if modelSimpleV2Arb.Pair1.Token0Erc20.ContractAddress != wplsAddress {
		nonPlsAddress = modelSimpleV2Arb.Pair1.Token0Erc20.ContractAddress
	} else {
		nonPlsAddress = modelSimpleV2Arb.Pair1.Token1Erc20.ContractAddress
	}

	err = simpleV2ArbHelper.MakeArbitrageSimple(
		auth,
		amountToken0,
		expectedFinalAmount,
		modelSimpleV2Arb.Pair0.ModelDex.RouterContractAddress,
		modelSimpleV2Arb.Pair1.ModelDex.RouterContractAddress,
		wplsAddress,
		nonPlsAddress)
	if err != nil {
		clientservice.ReleaseClient(auth)
		log.Error().Err(err).Msg("Failed to execute arbitrage")
		return err
	} else {
		clientservice.ReleaseClient(auth)
	}

	return nil
}
