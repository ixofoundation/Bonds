package client

import (
	pricingcmd "github.com/cosmos/cosmic/ixo-sdk/x/pricing/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	pricingQueryCmd := &cobra.Command{
		Use:   "pricing",
		Short: "Querying commands for the pricing module",
	}

	pricingQueryCmd.AddCommand(client.GetCommands(
		pricingcmd.GetCmdCosmicBonds(mc.storeKey, mc.cdc),
		pricingcmd.GetCmdCosmicBond(mc.storeKey, mc.cdc),
	)...)

	return pricingQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	pricingTxCmd := &cobra.Command{
		Use:   "pricing",
		Short: "Pricing transactions subcommands",
	}

	pricingTxCmd.AddCommand(client.PostCommands(
		pricingcmd.GetCmdCreateCosmicBond(mc.cdc),
		pricingcmd.GetCmdBuy(mc.cdc),
		pricingcmd.GetCmdSell(mc.cdc),
	)...)

	return pricingTxCmd
}
