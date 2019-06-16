package pricing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	CosmicBondRecords []CosmicBond `json:"cosmic_bonds"`
}

func NewGenesisState(cosmicBondRecords []CosmicBond) GenesisState {
	return GenesisState{CosmicBondRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	// TODO
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		CosmicBondRecords: []CosmicBond{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.CosmicBondRecords {
		keeper.SetCosmicBond(ctx, record.Moniker, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []CosmicBond
	iterator := k.GetCosmicBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		var ct CosmicBond
		ct = k.GetCosmicBond(ctx, name)
		records = append(records, ct)
	}
	return GenesisState{CosmicBondRecords: records}
}
