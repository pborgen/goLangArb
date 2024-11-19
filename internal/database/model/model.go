package model

import (
	"github.com/hexlivelive/goBot/internal/database/model/account"
	"github.com/hexlivelive/goBot/internal/database/model/dex"
	"github.com/hexlivelive/goBot/internal/database/model/erc20"
	"github.com/hexlivelive/goBot/internal/database/model/network"
	"github.com/hexlivelive/goBot/internal/database/model/pair"
	"github.com/hexlivelive/goBot/internal/database/model/simpleV2Arb"
)

type MyModel interface {
	account.ModelAccount | dex.ModelDex | erc20.ModelERC20 | network.NetworkModel | pair.ModelPair | simpleV2Arb.SimpleV2ArbModel
}
