package pricing

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateCosmicBond{}, "pricing/CreateCosmicBond", nil)
	cdc.RegisterConcrete(MsgBuy{}, "pricing/Buy", nil)
	cdc.RegisterConcrete(MsgSell{}, "pricing/Sell", nil)
	cdc.RegisterConcrete(CosmicBond{}, "pricing/CosmicBond", nil)
}

var ModuleCdc = codec.New()
