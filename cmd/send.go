package cmd

import (
	"fmt"
	"log"

	"github.com/branLiang/brancoin/core"
	"github.com/spf13/cobra"
)

var from, to string
var amount int
var mineNow bool

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if !core.ValidateAddress(from) {
			log.Panic("Error: Sender address is not valid")
		}
		if !core.ValidateAddress(to) {
			log.Panic("Error: Recipient address is not valid")
		}

		bc := core.NewBlockchain(nodeID)
		UTXOSet := core.UTXOSet{Blockchain: bc}
		defer bc.DB.Close()

		wallets, err := core.NewWallets(nodeID)
		if err != nil {
			log.Panic(err)
		}
		wallet := wallets.GetWallet(from)

		tx := core.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

		if mineNow {
			cbTx := core.NewCoinbaseTX(from, "")
			txs := []*core.Transaction{cbTx, tx}

			newBlock := bc.MineBlock(txs)
			UTXOSet.Update(newBlock)
		} else {
			core.SendTx(core.KnownNodes[0], tx)
		}

		fmt.Println("Success!")
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	sendCmd.Flags().StringVarP(&from, "from", "f", "", "Brancoin sender")
	sendCmd.Flags().StringVarP(&to, "to", "t", "", "Brancoin receiver")
	sendCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Amount of brancoin to send with")
	sendCmd.Flags().BoolVarP(&mineNow, "mineNow", "m", false, "Decided if should mine the block now")
}
