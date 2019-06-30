package cli

import (
	"fmt"
	"github.com/ixofoundation/cosmic/ixo-sdk/x/pricing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetCmdCosmicBonds(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "cosmic-bonds",
		Short: "List of all cosmic bonds",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cosmic-bonds", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query cosmic-bonds\n")
				return nil
			}

			var out pricing.QueryResCosmicBonds
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdCosmicBond(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "cosmic-bond [name]",
		Short: "Query info of a cosmic bond",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cosmic-bond/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve cosmic-bond - %s \n", name)
				return nil
			}

			var out pricing.CosmicBond
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
