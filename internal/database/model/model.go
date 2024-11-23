package model

import (
	"github.com/paulborgen/goLangArb/internal/database/model/account"
	"github.com/paulborgen/goLangArb/internal/database/model/dex"
	"github.com/paulborgen/goLangArb/internal/database/model/erc20"
	"github.com/paulborgen/goLangArb/internal/database/model/network"
	"github.com/paulborgen/goLangArb/internal/database/model/pair"
	"github.com/paulborgen/goLangArb/internal/database/model/simpleV2Arb"
)

type MyModel interface {
	account.ModelAccount | dex.ModelDex | erc20.ModelERC20 | network.NetworkModel | pair.ModelPair | simpleV2Arb.SimpleV2ArbModel
}
