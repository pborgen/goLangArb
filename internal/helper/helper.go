package helper

import (
	"fmt"
)

func Here() {
	fmt.Print("Hello, ")
	fmt.Print("World!")
	//_, err := ethclient.Dial("ADD_YOUR_ETHEREUM_NODE_URL")
	//
	//if err != nil {
	//	log.Fatalf("Oops! There was a problem", err)
	//} else {
	//	fmt.Println("Success! you are connected to the Ethereum Network")
	//}
}
