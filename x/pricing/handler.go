package pricing

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateCosmicBond:
			return handleMsgCreateCosmicBond(ctx, keeper, msg)
		case MsgBuy:
			return handleMsgBuy(ctx, keeper, msg)
		case MsgSell:
			return handleMsgSell(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized pricing Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateCosmicBond(ctx sdk.Context, keeper Keeper, msg MsgCreateCosmicBond) sdk.Result {
	keeper.GetCosmicBond(ctx, msg.Moniker)
	keeper.SetCreator(ctx, msg.Moniker, msg.Creator)
	keeper.SetReserveToken(ctx, msg.Moniker, msg.ReserveToken)
	keeper.SetReserveAddress(ctx, msg.Moniker, msg.ReserveAddress)
	keeper.SetMaxSupply(ctx, msg.Moniker, msg.MaxSupply)
	keeper.SetFunctionType(ctx, msg.Moniker, msg.FunctionType)
	keeper.SetM(ctx, msg.Moniker, msg.M)
	keeper.SetN(ctx, msg.Moniker, msg.N)
	keeper.SetAllowSells(ctx, msg.Moniker, msg.AllowSell)
	return sdk.Result{}
}

func handleMsgBuy(ctx sdk.Context, keeper Keeper, msg MsgBuy) sdk.Result {

	cb := keeper.GetCosmicBond(ctx, msg.Moniker)
	cost := cb.GetReserveNecessaryForIncreaseInSupply(msg.Amount.Amount)

	err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetReserveAddress(ctx, msg.Moniker),
		sdk.Coins{sdk.NewCoin(cb.ReserveToken, cost)})
	if err != nil {
		return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
	}

	_, err = keeper.coinKeeper.AddCoins(ctx, msg.Buyer, sdk.Coins{msg.Amount})
	if err != nil {
		return err.Result()
	}

	newCurrentSupply := keeper.GetCurrentSupply(ctx, msg.Moniker).Add(msg.Amount)
	keeper.SetCurrentSupply(ctx, msg.Moniker, newCurrentSupply)

	return sdk.Result{}
}

func handleMsgSell(ctx sdk.Context, keeper Keeper, msg MsgSell) sdk.Result {

	cb := keeper.GetCosmicBond(ctx, msg.Moniker)
	if cb.AllowSells != "true" {
		return sdk.ErrInternal("Coin does not allow selling.").Result()
	}
	cost := cb.GetReserveAvailableAfterDecreaseInSupply(msg.Amount.Amount)

	err := keeper.coinKeeper.SendCoins(ctx, keeper.GetReserveAddress(ctx, msg.Moniker), msg.Seller,
		sdk.Coins{sdk.NewCoin(cb.ReserveToken, cost)})
	if err != nil {
		return sdk.ErrInsufficientCoins("Reserve does not have enough coins").Result()
	}

	_, err = keeper.coinKeeper.SubtractCoins(ctx, msg.Seller, sdk.Coins{msg.Amount})
	if err != nil {
		return err.Result()
	}

	newCurrentSupply := keeper.GetCurrentSupply(ctx, msg.Moniker).Sub(msg.Amount)
	keeper.SetCurrentSupply(ctx, msg.Moniker, newCurrentSupply)

	return sdk.Result{}
}
