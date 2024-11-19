package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/hexlivelive/goBot/internal/blockchain"
	myUtil "github.com/hexlivelive/goBot/internal/util"
)

func main() {

	client := blockchain.GetClient()
	hash := common.HexToHash("0xf67aeb9f8c7bfca6e8f50a6eb2b5922942a26644e556dc6baf0204594d523710")
	myUtil.GetRevertReason(client, hash)
	//0xf67aeb9f8c7bfca6e8f50a6eb2b5922942a26644e556dc6baf0204594d523710
}
