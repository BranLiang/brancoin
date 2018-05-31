package cmd

import (
	"fmt"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

// createWalletCmd represents the createWallet command
var createWalletCmd = &cobra.Command{
	Use:   "createWallet",
	Short: "create a new wallet and return wallet address.",
	Run: func(cmd *cobra.Command, args []string) {
		wallets, _ := core.NewWallets(nodeID)
		address := wallets.CreateWallet()
		wallets.SaveToFile(nodeID)

		fmt.Println(address)
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)
}
