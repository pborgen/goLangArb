package dex

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	pulsexhelper "github.com/hexlivelive/goBot/abi"
	"github.com/hexlivelive/goBot/abi/erc20"
	pulsexV2Factory "github.com/hexlivelive/goBot/abi/pulseXV2"
	"log"
	"math/big"
)

// FactoryAddress points to the uniswap factory.
var PulseXV2FactoryAddress = common.HexToAddress("0x29eA7545DEf87022BAdc76323F373EA1e707C523")
var PulseXV1FactoryAddress = common.HexToAddress("0x1715a3E4A142d8b698131108995174F37aEBA10D")
var PulseXHelperContractAddress = common.HexToAddress("0x00a4a4b202e91ab3556a4776ca9be653e97b1e5e")

// Router02Address points to the uniswap v2 02 router.
var Router02Address = common.HexToAddress("0x165C3410fC91EF562C50559f7d2289fEbed552d9")

const pairAddressSuffix = "96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f"

//const pairAddressSuffix = "59fffffddd756cba9095128e53f3291a6ba38b21e3df744936e7289326555d62"

func GetReserves(pairAddress string) {

}

func GetPairAddress(token0, token1 common.Address) {
	//url := "https://rpc-pulsechain.g4mm4.io"
	url := "https://rpc.pulsechain.com"
	//url := "https://mainnet.infura.io/v3/2b9982d45dc84b68b83854cd2a5d0450"
	//client := blockchain.GetClient()m

	ctx := context.Background()

	client, err := ethclient.Dial(url)

	defer client.Close()

	var accountAddress = common.HexToAddress("0xf39FD575BdfD7cf15D6B0fBb674eEe569cfda353")

	balance, _ := client.BalanceAt(ctx, accountAddress, nil)
	//

	fmt.Println("Balance: ", balance)
	//ec := ethclient.NewClient(client.Client())

	callOpts := &bind.CallOpts{Context: ctx, Pending: false}

	//chainId, _ := client.ChainID(context.Background())
	//fmt.Println(chainId)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	//var wpls common.Address = common.HexToAddress("0xA1077a294dDE1B09bB078844df40758a5D0f9a27")
	var hex common.Address = common.HexToAddress("0x2b591e99afe9f32eaa6214f7b7629768c40eeb39")
	//var myAddress = common.HexToAddress("0x6029811b4EEcC1db60c20D32Ad2102dE88E36652")
	contract, err := erc20.NewErc20(hex, client)

	fmt.Println("-----------")
	totalSupply, err := contract.TotalSupply(nil)
	fmt.Print("TotalSupply:", totalSupply)

	bal, err := contract.TotalSupply(callOpts)
	fmt.Println(bal)

	fmt.Println("-----------")
	//bal, err := erc20Instance.Name(&bind.CallOpts{})
	//fmt.Println(bal)
	////balance, err := erc20Instance.BalanceOf(nil, common.HexToAddress("0x6029811b4EEcC1db60c20D32Ad2102dE88E36652"))
	//balance, err := erc20Instance.Erc20Caller.Name(nil)
	//fmt.Println(balance)
	//header, err := client.HeaderByNumber(context.Background(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(header.Number.String()) // 5671744

	pulseXV2FactoryInstance, err := pulsexV2Factory.NewPulsexV2FactoryCaller(PulseXV2FactoryAddress, client)
	////
	//if err != nil {
	//	log.Fatal(err)
	//}
	////callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	//
	//opts := bind.New

	bla4, err := pulseXV2FactoryInstance.AllPairsLength(nil)
	fmt.Println(bla4)

	bla5, err := pulseXV2FactoryInstance.GetPair(nil, token0, token1)
	fmt.Println(bla5.String())
	//bla3, err := pulseXV2FactoryInstance.AllPairs(nil, big.NewInt(1))
	//
	//fmt.Println(bla3)
	//
	pulsexhelperInstance, err := pulsexhelper.NewPulsexhelper(PulseXHelperContractAddress, client)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var diaAddress = common.HexToAddress("0xefd766ccb38eaf1dfd701853bfce31359239f305")
	//var wplsAddress = common.HexToAddress("0xa1077a294dde1b09bb078844df40758a5d0f9a27")

	pairAddress, err := pulsexhelperInstance.PairFor(nil, PulseXV1FactoryAddress, token0, token1)
	//pairAddress, err := pulsexhelperInstance.ConvertUintToAddress(nil, big.NewInt(1))
	////pairAddress, err := pulsexhelperInstance.PairFor(nil, PulseXV2FactoryAddress, token0, token1)

	fmt.Println("PairAddress", pairAddress.String())
	//
	//if err != nil {
	//	fmt.Println("Oops! There was a problem", err)
	//} else {
	//	fmt.Println("Success! you are connected to the Ethereum Network")
	//}

}

func GeneratePairAddress(token0, token1 common.Address) common.Address {
	// addresses need to be sorted in an ascending order for proper behaviour
	token0, token1 = sortAddressess(token0, token1)

	// 255 is required as a prefix for this to work
	// see: https://uniswap.org/docs/v2/javascript-SDK/getting-pair-addresses/
	message := []byte{255}

	message = append(message, PulseXV2FactoryAddress.Bytes()...)

	addrSum := token0.Bytes()
	addrSum = append(addrSum, token1.Bytes()...)

	message = append(message, crypto.Keccak256(addrSum)...)

	b, _ := hex.DecodeString(pairAddressSuffix)
	message = append(message, b...)
	hashed := crypto.Keccak256(message)
	addressBytes := big.NewInt(0).SetBytes(hashed)
	addressBytes = addressBytes.Abs(addressBytes)
	return common.BytesToAddress(addressBytes.Bytes())
}

func sortAddressess(tkn0, tkn1 common.Address) (common.Address, common.Address) {
	token0Rep := big.NewInt(0).SetBytes(tkn0.Bytes())
	token1Rep := big.NewInt(0).SetBytes(tkn1.Bytes())

	if token0Rep.Cmp(token1Rep) > 0 {
		tkn0, tkn1 = tkn1, tkn0
	}

	return tkn0, tkn1
}
