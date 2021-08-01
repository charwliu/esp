package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.vixal.xyz/esp/ledger"
)

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Check ledger for httperror",
	Run: func(cmd *cobra.Command, args []string) {
		ledgerFileReader, err := ledger.NewLedgerReader(ledgerFilePath)
		if err != nil {
			fmt.Println("Ledger: ", err)
			return
		}

		c, e := ledger.ParseLedgerAsync(ledgerFileReader)
		errorCount := 0
		for {
			select {
			case <-c:
				continue
			case err := <-e:
				if err == nil {
					os.Exit(errorCount)
				}
				fmt.Println("Ledger: ", err)
				errorCount++
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
