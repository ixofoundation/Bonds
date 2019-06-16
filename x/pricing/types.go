package pricing

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"strings"
)

const (
	PowerFunction = "power_function"
)

type CosmicBond struct {
	Moniker        string         `json:"moniker"`
	Creator        sdk.AccAddress `json:"creator"`
	ReserveToken   string         `json:"reserve_token"`
	ReserveAddress sdk.AccAddress `json:"reserve_address"`
	MaxSupply      sdk.Coin       `json:"max_supply"`
	CurrentSupply  sdk.Coin       `json:"current_supply"`
	FunctionType   string         `json:"function_type"`
	M              string         `json:"m"`
	N              string         `json:"n"`
	AllowSells     string         `json:"allow_sells"`
}

func NewCosmicBond(moniker string) CosmicBond {
	return CosmicBond{
		Moniker:       moniker,
		CurrentSupply: sdk.NewInt64Coin(moniker, 0),
	}
}

func (c CosmicBond) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Moniker: %s
Creator: %s
ReserveToken: %s
ReserveAddress: %s
MaxSupply: %s
CurrentSupply: %s
FunctionType: %s
M: %s
N: %s
AllowSells: %s`, c.Moniker, c.Creator, c.ReserveToken, c.ReserveAddress,
		c.MaxSupply, c.CurrentSupply, c.FunctionType, c.M, c.N, c.AllowSells))
}

func (c CosmicBond) GetPriceForSupply(supply sdk.Int) sdk.Int {
	switch c.FunctionType {
	case PowerFunction:
		m, _ := strconv.ParseInt(c.M, 10, 64)
		return supply.MulRaw(m)
	default:
		return sdk.Int{} // TODO: error
	}
}

func (c CosmicBond) GetCurrentPrice() sdk.Int {
	switch c.FunctionType {
	case PowerFunction:
		return c.GetPriceForSupply(c.CurrentSupply.Amount)
	default:
		return sdk.Int{} // TODO: error
	}
}

func (c CosmicBond) GetReserveNecessaryForSupply(supply sdk.Int) sdk.Int {
	switch c.FunctionType {
	case PowerFunction:
		m, _ := strconv.ParseInt(c.M, 10, 64)
		n, _ := strconv.ParseInt(c.N, 10, 64)
		temp := supply.Int64()
		for i := n + 1; i > 1; i-- {
			temp = temp * supply.Int64()
		}
		a := m * temp / (n + 1)
		return sdk.NewInt(a)
	default:
		return sdk.Int{} // TODO: error
	}
}

func (c CosmicBond) GetReserveNecessaryForIncreaseInSupply(increase sdk.Int) sdk.Int {
	switch c.FunctionType {
	case PowerFunction:
		price1 := c.GetReserveNecessaryForSupply(c.CurrentSupply.Amount)
		price2 := c.GetReserveNecessaryForSupply(c.CurrentSupply.Amount.Add(increase))
		return price2.Sub(price1)
	default:
		return sdk.Int{} // TODO: error
	}
}

func (c CosmicBond) GetReserveAvailableAfterDecreaseInSupply(decrease sdk.Int) sdk.Int {
	switch c.FunctionType {
	case PowerFunction:
		price1 := c.GetReserveNecessaryForSupply(c.CurrentSupply.Amount)
		price2 := c.GetReserveNecessaryForSupply(c.CurrentSupply.Amount.Sub(decrease))
		return price1.Sub(price2)
	default:
		return sdk.Int{} // TODO: error
	}
}
