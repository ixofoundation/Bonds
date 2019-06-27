package cli

import (
	"fmt"
	"github.com/cosmos/cosmic/ixo-sdk/x/pricing"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func GetCmdCreateCosmicBond(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-cosmic-bond [moniker] [reserve-token] [reserve-address] [max-supply] [function-type] [m] [n] [allow-sells]",
		Short: "Create cosmic bond",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			reserveAddress, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			// TODO: check that function-type is a valid one

			maxSupply, err := sdk.ParseCoin(fmt.Sprintf("%s%s", args[3], args[0]))
			if err != nil {
				return err
			}

			_, err = strconv.ParseInt(args[5], 10, 64) // Check that m parsable to integer
			if err != nil {
				return err
			}

			_, err = strconv.ParseInt(args[6], 10, 64) // Check that n parsable to integer
			if err != nil {
				return err
			}

			msg := pricing.NewMsgCreateCosmicBond(args[0], cliCtx.GetFromAddress(), args[1], reserveAddress, maxSupply, args[4], args[5], args[6], args[7])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdBuy(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy [moniker] [amount] [max-price]",
		Short: "Buy into a cosmic bond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			coins, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			msg := pricing.NewMsgBuy(args[0], cliCtx.GetFromAddress(), coins, args[2])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdSell(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "sell [moniker] [amount]",
		Short: "Sell from a cosmic bond",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			coins, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			msg := pricing.NewMsgSell(args[0], cliCtx.GetFromAddress(), coins)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
