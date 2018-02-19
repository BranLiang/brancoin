package cmd

import (
	"fmt"
	"strconv"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

// printChainCmd represents the printChain command
var printChainCmd = &cobra.Command{
	Use:   "printChain",
	Short: "Print all block info in chain.",
	Run: func(cmd *cobra.Command, args []string) {
		bc := core.NewBlockchain(nodeID)
		defer bc.DB.Close()

		bci := bc.Iterator()

		for {
			block := bci.Next()

			fmt.Printf("============ Block %x ============\n", block.Hash)
			fmt.Printf("Height: %d\n", block.Height)
			fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
			pow := core.NewProofOfWork(block)
			fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
			for _, tx := range block.Transactions {
				fmt.Println(tx)
			}
			fmt.Printf("\n\n")

			if len(block.PrevBlockHash) == 0 {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(printChainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printChainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printChainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
