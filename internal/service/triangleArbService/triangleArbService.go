package triangleArbService

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/paulborgen/goLangArb/internal/database/model/erc20"
	"github.com/paulborgen/goLangArb/internal/database/model/pair"
	"github.com/paulborgen/goLangArb/internal/database/model/triangleArb"
	clientservice "github.com/paulborgen/goLangArb/internal/service/clientService"
	"github.com/rs/zerolog/log"
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
	pair2 pair.ModelPair
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

func Find() error {

	var foundArray []foundArbsStruct

	plsPairs, err := pair.GetAllPairsThatHavePls()
	nonPlsPairs, err := pair.GetAllPairsWithOutPls()

	if err != nil {
		panic("Must be a issue with the DB")
	}

	for _, pair0 := range plsPairs {
		// Get the side of the pair that is not pls

		var pair0ToPair1CommonToken erc20.ModelERC20
		if pair0.Token0Erc20.ContractAddress != wplsAddress {
			pair0ToPair1CommonToken = pair0.Token0Erc20
		} else {
			pair0ToPair1CommonToken = pair0.Token1Erc20
		}

		// Find Pairs on other dexs
		for _, pair1 := range nonPlsPairs {

			// Check if we found a pair that links with the previous one
			if pair1.Token0Erc20.ContractAddress == pair0ToPair1CommonToken.ContractAddress ||
				pair1.Token1Erc20.ContractAddress == pair0ToPair1CommonToken.ContractAddress {

				var pair1ToPair2CommonToken erc20.ModelERC20
				if pair1.Token0Erc20.ContractAddress != pair0ToPair1CommonToken.ContractAddress {
					pair1ToPair2CommonToken = pair1.Token0Erc20
				} else if pair1.Token1Erc20.ContractAddress != pair0ToPair1CommonToken.ContractAddress {
					pair1ToPair2CommonToken = pair1.Token1Erc20
				} else {
					panic("this state should not happen")
				}
				// We now have 2 pairs that are related Pair0: (WPLS / ERC20)  --> (ERC20 / ERC20)
				// Try to find another pair (ERC20 / WPLS) to complete our triangle arb
				for _, pair2 := range plsPairs {
					var nonPlsSidePair2 common.Address
					if pair2.Token0Erc20.ContractAddress != wplsAddress {
						nonPlsSidePair2 = pair2.Token0Erc20.ContractAddress
					} else {
						nonPlsSidePair2 = pair2.Token1Erc20.ContractAddress
					}

					// Did we find an arb
					if pair1ToPair2CommonToken.ContractAddress == nonPlsSidePair2 {
						foundArray = append(foundArray, foundArbsStruct{
							pair0: pair0,
							pair1: pair1,
							pair2: pair2,
						})
						log.Log().Msgf("Found triangle Arb:")
						log.Log().Msgf("([%s] - %s / %s) --> ([%s] - %s / %s) --> ([%s] - %s / %s)",
							dex.GetById(pair0.DexId).Name,
							pair0.Token0Erc20.Name,
							pair0.Token1Erc20.Name,
							dex.GetById(pair1.DexId).Name,
							pair1.Token0Erc20.Name,
							pair1.Token1Erc20.Name,
							dex.GetById(pair2.DexId).Name,
							pair2.Token0Erc20.Name,
							pair2.Token1Erc20.Name)
					}
				}
			}
		}
	}

	for _, found := range foundArray {
		modelTriangleArb := triangleArb.ModelTriangleArb{
			Pair0Id:         found.pair0.PairId,
			Pair1Id:         found.pair1.PairId,
			Pair2Id:         found.pair2.PairId,
			ShouldFindArb:   true,
			FailedArbCount:  0,
			SuccessArbCount: 0,
			Note:            "",
		}

		exists, err := triangleArb.Exists(modelTriangleArb.Pair0Id, modelTriangleArb.Pair1Id, modelTriangleArb.Pair2Id)

		if err != nil {
			panic("should not happen" + err.Error())
		}

		if !*exists {

			// Pulse rate is giving issues because a contract is not ble to call it

			err := triangleArb.Insert(modelTriangleArb)

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
	routerAddress0 common.Address,
	routerAddress1 common.Address,
	token0Address common.Address,
	token1Address common.Address) {

	auth, err := clientservice.GetAndLockClient()

	if err != nil {
		log.Error().Err(err).Msg("Failed to get client")
		return
	}

	tip := big.NewInt(10000000000) // maxPriorityFeePerGas = 2 Gwei
	//feeCap := big.NewInt(20000000000)
	gasLimit := uint64(400000)

	auth.GasLimit = gasLimit
	auth.GasTipCap = tip
	auth.GasFeeCap = auth.GasPrice
	auth.GasPrice = nil

	//err := simpleV2ArbHelper.MakeArbitrageSimple(auth, amountToken0, expectedFinalAmount, routerAddress0, routerAddress1, token0Address, token1Address)
	//if err != nil {
	//	clientservice.ReleaseClient(auth)
	//	log.Error().Err(err).Msg("Failed to execute arbitrage")
	//	return
	//} else {
	//	clientservice.ReleaseClient(auth)
	//}

}
