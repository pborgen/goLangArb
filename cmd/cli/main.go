package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paulborgen/goLangArb/internal/blockchain/dex/v2/dexUniswapV2Router"
	"github.com/paulborgen/goLangArb/internal/blockchain/dex/v2/dexUniswpV2Pair"
	"github.com/paulborgen/goLangArb/internal/database/model/pair"
	"github.com/paulborgen/goLangArb/internal/mylogger"
	"github.com/paulborgen/goLangArb/internal/service/pairFinder"
	"github.com/rs/zerolog/log"
)

func main() {
	mylogger.Init()

	pair.GetAllPairs()
	reserves, _ := dexUniswpV2Pair.PopulateReserves(common.HexToAddress("0x718Dd8B743ea19d71BDb4Cb48BB984b73a65cE06"))

	log.Print(reserves.Reserve0.String())
	router := common.HexToAddress("0x05d5F20500eD8d9E012647E6CFe1b2Bf89f5b926")
	factory := dexUniswapV2Router.V2RouterGetFactory(router)

	pairFinder.UpdateAllReserves()

	log.Info().Msgf(factory.String())

	log.Print("hello world")

	//pair := model.Pair{
	//	DexId:               1,
	//	PairContractAddress: "",
	//	Token0Erc20Id:       1,
	//	Token1Erc20Id:       1,
	//	Token0Reserves:      *big.NewInt(1),
	//	Token1Reserves:      *big.NewInt(1),
	//	ShouldFindArb:       false,
	//}
	//model.PairInsert(pair)

	//var diaContractAddress = common.HexToAddress("0xefd766ccb38eaf1dfd701853bfce31359239f305")
	//var wplsContractAddress common.Address = common.HexToAddress("0xA1077a294dDE1B09bB078844df40758a5D0f9a27")
	////var usdl common.Address = common.HexToAddress("0x0dEEd1486bc52aA0d3E6f8849cEC5adD6598A162")
	////client := blockchain.GetClient()
	////
	////ctx := context.Background()
	////ctx = context.WithValue(ctx, "myClient", client)
	////
	////var myClient ethclient.Client = ctx.Value("myClient").(ethclient.Client)
	////fmt.Println(myClient)
	//var PulseXV2FactoryAddress = common.HexToAddress("0x29eA7545DEf87022BAdc76323F373EA1e707C523")
	//
	//diaErc20 := erc20Helper.GetErc20(diaContractAddress)
	//wplsErc20 := erc20Helper.GetErc20(wplsContractAddress)
	//
	//fmt.Println(diaErc20)
	//fmt.Println(wplsErc20)
	//
	//pair := dexUniswapV2.GetPair(PulseXV2FactoryAddress, wplsErc20, diaErc20)
	//fmt.Println(pair)
	//pair2 := dexUniswapV2.PopulateReserves(pair)
	//
	//fmt.Println(pair2)
	//fmt.Println(pair2.Pair.Token0Reserves.String())
	//fmt.Println(pair2.Pair.Token1Reserves.String())
	//bla := dexuniswapv2.GetPair(client, PulseXV2FactoryAddress, dia, wpls)
	//
	//fmt.Println(bla)
	//dex.GetPairAddress(wpls, dia)

	//var bla = dex.GeneratePairAddress(wpls, dia)

	fmt.Println("here")
	//fmt.Println(bla)

}
