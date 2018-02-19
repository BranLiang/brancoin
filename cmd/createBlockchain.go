package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

// createBlockchainCmd represents the createBlockchain command
var createBlockchainCmd = &cobra.Command{
	Use:   "createBlockchain",
	Short: "Create a new blockchain with coinbase transaction.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("ERROR: Address is not given")
			os.Exit(1)
		}
		if !core.ValidateAddress(args[0]) {
			log.Panic("ERROR: Address is not valid")
		}
		bc := core.CreateBlockchain(args[0], nodeID)
		defer bc.DB.Close()

		UTXOset := core.UTXOSet{Blockchain: bc}
		UTXOset.Reindex()

		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(createBlockchainCmd)
}
