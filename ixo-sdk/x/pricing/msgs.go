package pricing

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgCreateCosmicBond struct {
	Moniker        string
	Creator        sdk.AccAddress
	ReserveToken   string
	ReserveAddress sdk.AccAddress
	MaxSupply      sdk.Coin
	FunctionType   string
	M              string
	N              string
	AllowSell      string
}

func NewMsgCreateCosmicBond(moniker string, creator sdk.AccAddress, reserveToken string, reserveAddress sdk.AccAddress, maxSupply sdk.Coin, functionType, m, n, allowSell string) MsgCreateCosmicBond {
	return MsgCreateCosmicBond{
		Moniker:        moniker,
		Creator:        creator,
		ReserveToken:   reserveToken,
		ReserveAddress: reserveAddress,
		MaxSupply:      maxSupply,
		FunctionType:   functionType,
		M:              m,
		N:              n,
		AllowSell:      allowSell,
	}
}

func (msg MsgCreateCosmicBond) ValidateBasic() sdk.Error {
	if msg.Creator.Empty() {
		return sdk.ErrInvalidAddress(msg.Creator.String())
	}
	if len(msg.Moniker) == 0 {
		return sdk.ErrUnknownRequest("Moniker cannot be empty")
	}
	return nil
}

func (msg MsgCreateCosmicBond) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgCreateCosmicBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

func (msg MsgCreateCosmicBond) Route() string { return RouterKey }

func (msg MsgCreateCosmicBond) Type() string { return "create_cosmic_bond" }

type MsgBuy struct {
	Moniker  string
	Buyer    sdk.AccAddress
	Amount   sdk.Coin
	MaxPrice string
}

func NewMsgBuy(moniker string, accAddress sdk.AccAddress, amount sdk.Coin, maxPrice string) MsgBuy {
	return MsgBuy{
		Moniker:  moniker,
		Buyer:    accAddress,
		Amount:   amount,
		MaxPrice: maxPrice,
	}
}

func (msg MsgBuy) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Moniker) == 0 || len(msg.MaxPrice) == 0 {
		return sdk.ErrUnknownRequest("Moniker cannot be empty")
	}
	return nil
}

func (msg MsgBuy) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

func (msg MsgBuy) Route() string { return RouterKey }

func (msg MsgBuy) Type() string { return "buy" }

type MsgSell struct {
	Moniker string
	Seller  sdk.AccAddress
	Amount  sdk.Coin
}

func NewMsgSell(moniker string, accAddress sdk.AccAddress, amount sdk.Coin) MsgSell {
	return MsgSell{
		Moniker: moniker,
		Seller:  accAddress,
		Amount:  amount,
	}
}

func (msg MsgSell) ValidateBasic() sdk.Error {
	if msg.Seller.Empty() {
		return sdk.ErrInvalidAddress(msg.Seller.String())
	}
	if len(msg.Moniker) == 0 {
		return sdk.ErrUnknownRequest("Moniker cannot be empty")
	}
	return nil
}

func (msg MsgSell) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Seller}
}

func (msg MsgSell) Route() string { return RouterKey }

func (msg MsgSell) Type() string { return "sell" }
