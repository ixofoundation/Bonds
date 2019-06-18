package pricing

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryCosmicBonds = "cosmic-bonds"
	QueryCosmicBond  = "cosmic-bond"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryCosmicBonds:
			return queryCosmicBonds(ctx, req, keeper)
		case QueryCosmicBond:
			return queryCosmicBond(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

func queryCosmicBonds(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var namesList QueryResCosmicBonds

	iterator := keeper.GetCosmicBondIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		namesList = append(namesList, name)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// nolint: unparam
func queryCosmicBond(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	name := path[0]

	cb := keeper.GetCosmicBond(ctx, name)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, cb)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

type QueryResCosmicBonds []string

func (n QueryResCosmicBonds) String() string {
	return strings.Join(n[:], "\n")
}
