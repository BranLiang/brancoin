package cmd

import (
	"fmt"
	"log"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

// listAddressesCmd represents the listAddresses command
var listAddressesCmd = &cobra.Command{
	Use:   "listAddresses",
	Short: "List all addresses under a specific node id",
	Run: func(cmd *cobra.Command, args []string) {
		wallets, err := core.NewWallets(nodeID)
		if err != nil {
			log.Panic(err)
		}
		addresses := wallets.GetAddresses()

		for _, address := range addresses {
			fmt.Println(address)
		}
	},
}

func init() {
	rootCmd.AddCommand(listAddressesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAddressesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAddressesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
