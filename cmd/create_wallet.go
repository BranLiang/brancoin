package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createWallet)
}

var createWallet = &cobra.Command{
	Use:   "createwallet",
	Short: "Generate a new wallet with pair of private and public keys",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test createwallet")
	},
}
