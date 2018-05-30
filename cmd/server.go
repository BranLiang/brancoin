package cmd

import (
	"fmt"
	"log"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

var minerAddress string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server for node",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting node %s\n", nodeID)
		if len(minerAddress) > 0 {
			if core.ValidateAddress(minerAddress) {
				fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
			} else {
				log.Panic("Wrong miner address!")
			}
		}
		core.StartServer(nodeID, minerAddress)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&minerAddress, "miner", "m", "", "Miner Address")
}
