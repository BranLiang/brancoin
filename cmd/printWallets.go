package cmd

import (
	"fmt"
	"log"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

// listAddressesCmd represents the listAddresses command
var printWalletsCmd = &cobra.Command{
	Use:   "printWallets",
	Short: "List all wallet under a specific node id",
	Run: func(cmd *cobra.Command, args []string) {
		wallets, err := core.NewWallets(nodeID)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(wallets)
	},
}

func init() {
	rootCmd.AddCommand(printWalletsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAddressesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAddressesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
