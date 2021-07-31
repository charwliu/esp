package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"go.vixal.xyz/esp/ledger"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Aliases: []string{"bal"},
	Use:     "balance [account-substring-filter]...",
	Short:   "Print account balances",
	Run: func(cmd *cobra.Command, args []string) {
		generalLedger, err := cliTransactions()
		if err != nil {
			log.Fatalln(err)
		}
		if period == "" {
			PrintBalances(ledger.GetBalances(generalLedger, args), showEmptyAccounts, transactionDepth, columnWidth)
		} else {
			lperiod := ledger.Period(period)
			rbalances := ledger.BalancesByPeriod(generalLedger, lperiod, ledger.RangePartition)
			for rIdx, rb := range rbalances {
				if rIdx > 0 {
					fmt.Println("")
					fmt.Println(strings.Repeat("=", columnWidth))
				}
				fmt.Println(rb.Start.Format(transactionDateFormat), "-", rb.End.Format(transactionDateFormat))
				fmt.Println(strings.Repeat("=", columnWidth))
				PrintBalances(rb.Balances, showEmptyAccounts, transactionDepth, columnWidth)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	var startDate, endDate time.Time
	startDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
	endDate = time.Now().Add(time.Hour * 24)
	balanceCmd.Flags().StringVarP(&startString, "begin-date", "b", startDate.Format(transactionDateFormat), "Begin date of transaction processing.")
	balanceCmd.Flags().StringVarP(&endString, "end-date", "e", endDate.Format(transactionDateFormat), "End date of transaction processing.")
	balanceCmd.Flags().StringVar(&payeeFilter, "payee", "", "Filter output to payees that contain this string.")
	balanceCmd.Flags().IntVar(&columnWidth, "columns", 80, "Set a column width for output.")
	balanceCmd.Flags().BoolVar(&columnWide, "wide", false, "Wide output (same as --columns=132).")

	balanceCmd.Flags().StringVar(&period, "period", "", "Split output into periods (Monthly,Quarterly,SemiYearly,Yearly).")
	balanceCmd.Flags().BoolVar(&showEmptyAccounts, "empty", false, "Show empty (zero balance) accounts.")
	balanceCmd.Flags().IntVar(&transactionDepth, "depth", -1, "Depth of transaction output (balance).")
}
