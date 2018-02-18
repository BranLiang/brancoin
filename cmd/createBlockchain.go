package cmd

import (
	"fmt"
	"log"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

// createBlockchainCmd represents the createBlockchain command
var createBlockchainCmd = &cobra.Command{
	Use:   "createBlockchain",
	Short: "Create a new blockchain with coinbase transaction.",
	Run: func(cmd *cobra.Command, args []string) {
		if !core.ValidateAddress(args[0]) {
			log.Panic("ERROR: Address is not valid")
		}
		bc := core.CreateBlockchain(args[0], nodeID)

		UTXOset := core.UTXOSet{Blockchain: bc}
		UTXOset.Reindex()

		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(createBlockchainCmd)
}
