package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

// getBalanceCmd represents the getBalance command
var getBalanceCmd = &cobra.Command{
	Use:   "getBalance",
	Short: "Get amount of balance for address.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Address is not given!")
			os.Exit(1)
		}
		if !core.ValidateAddress(args[0]) {
			log.Panic("ERROR: Address is not valid")
		}
		bc := core.NewBlockchain(nodeID)
		UTXOSet := core.UTXOSet{bc}
		defer bc.DB.Close()

		balance := 0
		pubKeyHash := core.Base58Decode([]byte(args[0]))
		pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
		UTXOs := core.UTXOSet.FindUTXO(pubKeyHash)

		for _, out := range UTXOs {
			balance += out.Value
		}

		fmt.Printf("Balance of '%s': %d\n", args[0], balance)
	},
}

func init() {
	rootCmd.AddCommand(getBalanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getBalanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getBalanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
