package pricing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const StoreKey = "pricing"

type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Gets the entire Whois metadata struct for a name
func (k Keeper) GetCosmicBond(ctx sdk.Context, name string) CosmicBond {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(name)) {
		return NewCosmicBond(name)
	}
	bz := store.Get([]byte(name))
	var cb CosmicBond
	k.cdc.MustUnmarshalBinaryBare(bz, &cb)
	return cb
}

func (k Keeper) SetCosmicBond(ctx sdk.Context, name string, cb CosmicBond) {
	if cb.Creator.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(cb))
}

func (k Keeper) SetCreator(ctx sdk.Context, moniker string, creator sdk.AccAddress) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.Creator = creator
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetCreator(ctx sdk.Context, moniker string) sdk.AccAddress {
	return k.GetCosmicBond(ctx, moniker).Creator
}

func (k Keeper) SetReserveToken(ctx sdk.Context, moniker, reserveToken string) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.ReserveToken = reserveToken
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetReserveToken(ctx sdk.Context, moniker string) string {
	return k.GetCosmicBond(ctx, moniker).ReserveToken
}

func (k Keeper) SetReserveAddress(ctx sdk.Context, moniker string, reserveAddress sdk.AccAddress) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.ReserveAddress = reserveAddress
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetReserveAddress(ctx sdk.Context, moniker string) sdk.AccAddress {
	return k.GetCosmicBond(ctx, moniker).ReserveAddress
}

func (k Keeper) SetCurrentSupply(ctx sdk.Context, moniker string, currentSupply sdk.Coin) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.CurrentSupply = currentSupply
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetCurrentSupply(ctx sdk.Context, moniker string) sdk.Coin {
	return k.GetCosmicBond(ctx, moniker).CurrentSupply
}

func (k Keeper) SetMaxSupply(ctx sdk.Context, moniker string, max_supply sdk.Coin) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.MaxSupply = max_supply
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetMaxSupply(ctx sdk.Context, moniker string) sdk.Coin {
	return k.GetCosmicBond(ctx, moniker).MaxSupply
}

func (k Keeper) SetFunctionType(ctx sdk.Context, moniker string, functionType string) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.FunctionType = functionType
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetFunctionType(ctx sdk.Context, moniker string) string {
	return k.GetCosmicBond(ctx, moniker).M
}

func (k Keeper) SetM(ctx sdk.Context, moniker string, m string) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.M = m
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetM(ctx sdk.Context, moniker string) string {
	return k.GetCosmicBond(ctx, moniker).M
}

func (k Keeper) SetN(ctx sdk.Context, moniker string, n string) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.N = n
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetN(ctx sdk.Context, moniker string) string {
	return k.GetCosmicBond(ctx, moniker).N
}

func (k Keeper) SetAllowSells(ctx sdk.Context, moniker string, allowSells string) {
	cb := k.GetCosmicBond(ctx, moniker)
	cb.AllowSells = allowSells
	k.SetCosmicBond(ctx, moniker, cb)
}

func (k Keeper) GetAllowSells(ctx sdk.Context, moniker string) string {
	return k.GetCosmicBond(ctx, moniker).AllowSells
}

func (k Keeper) GetCosmicBondIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
