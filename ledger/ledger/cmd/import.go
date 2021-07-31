package cmd

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jbrukh/bayesian"
	"github.com/spf13/cobra"
	"go.vixal.xyz/esp/ledger"
)

var csvDateFormat string
var destAccSearch string
var negateAmount bool
var allowMatching bool
var fieldDelimiter string
var scaleFactor float64

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import <account-substring> <csv-file>",
	Args:  cobra.ExactArgs(2),
	Short: "Import transactions from csv to ledger format",
	Run: func(cmd *cobra.Command, args []string) {
		var accountSubstring, csvFileName string
		accountSubstring = args[0]
		csvFileName = args[1]

		ratScale := big.NewRat(1, 1)
		ratScale.SetFloat64(scaleFactor)

		csvFileReader, err := os.Open(csvFileName)
		if err != nil {
			fmt.Println("CSV: ", err)
			return
		}
		defer csvFileReader.Close()

		ledgerFileReader, err := ledger.NewLedgerReader(ledgerFilePath)
		if err != nil {
			fmt.Println("Ledger: ", err)
			return
		}

		generalLedger, parseError := ledger.ParseLedger(ledgerFileReader)
		if parseError != nil {
			fmt.Printf("%s:%s\n", ledgerFilePath, parseError.Error())
			return
		}

		var matchingAccount string
		matchingAccounts := ledger.GetBalances(generalLedger, []string{accountSubstring})
		if len(matchingAccounts) < 1 {
			fmt.Println("Unable to find matching account.")
			return
		}
		matchingAccount = matchingAccounts[len(matchingAccounts)-1].Name

		allAccounts := ledger.GetBalances(generalLedger, []string{})

		csvReader := csv.NewReader(csvFileReader)
		csvReader.Comma, _ = utf8.DecodeRuneInString(fieldDelimiter)
		csvRecords, cerr := csvReader.ReadAll()
		if cerr != nil {
			fmt.Println("CSV parse error:", cerr.Error())
			return
		}

		classes := make([]bayesian.Class, len(allAccounts))
		for i, bal := range allAccounts {
			classes[i] = bayesian.Class(bal.Name)
		}
		classifier := bayesian.NewClassifier(classes...)
		for _, tran := range generalLedger {
			payeeWords := strings.Split(tran.Payee, " ")
			for _, accChange := range tran.AccountChanges {
				if strings.Contains(accChange.Name, destAccSearch) {
					classifier.Learn(payeeWords, bayesian.Class(accChange.Name))
				}
			}
		}

		// Find columns from header
		var dateColumn, payeeColumn, amountColumn, commentColumn int
		dateColumn, payeeColumn, amountColumn, commentColumn = -1, -1, -1, -1
		for fieldIndex, fieldName := range csvRecords[0] {
			fieldName = strings.ToLower(fieldName)
			if strings.Contains(fieldName, "date") {
				dateColumn = fieldIndex
			} else if strings.Contains(fieldName, "description") {
				payeeColumn = fieldIndex
			} else if strings.Contains(fieldName, "payee") {
				payeeColumn = fieldIndex
			} else if strings.Contains(fieldName, "amount") {
				amountColumn = fieldIndex
			} else if strings.Contains(fieldName, "expense") {
				amountColumn = fieldIndex
			} else if strings.Contains(fieldName, "note") {
				commentColumn = fieldIndex
			} else if strings.Contains(fieldName, "comment") {
				commentColumn = fieldIndex
			}
		}

		if dateColumn < 0 || payeeColumn < 0 || amountColumn < 0 {
			fmt.Println("Unable to find columns required from header field names.")
			return
		}

		expenseAccount := ledger.Account{Name: "unknown:unknown", Balance: new(big.Rat)}
		csvAccount := ledger.Account{Name: matchingAccount, Balance: new(big.Rat)}
		for _, record := range csvRecords[1:] {
			inputPayeeWords := strings.Split(record[payeeColumn], " ")
			csvDate, _ := time.Parse(csvDateFormat, record[dateColumn])
			if allowMatching || !existingTransaction(generalLedger, csvDate, inputPayeeWords[0]) {
				// Classify into expense account
				_, likely, _ := classifier.LogScores(inputPayeeWords)
				if likely >= 0 {
					expenseAccount.Name = string(classifier.Classes[likely])
				}

				// Negate amount if required
				expenseAccount.Balance.SetString(record[amountColumn])
				if negateAmount {
					expenseAccount.Balance.Neg(expenseAccount.Balance)
				}

				// handleEofToken scale
				expenseAccount.Balance = expenseAccount.Balance.Mul(expenseAccount.Balance, ratScale)

				// Csv amount is the negative of the expense amount
				csvAccount.Balance.Neg(expenseAccount.Balance)

				// Create valid transaction for print in ledger format
				trans := &ledger.Transaction{Date: csvDate, Payee: record[payeeColumn]}
				trans.AccountChanges = []ledger.Account{csvAccount, expenseAccount}

				// Comment
				if commentColumn >= 0 && record[commentColumn] != "" {
					trans.Comments = []string{";" + record[commentColumn]}
				}
				PrintTransaction(trans, 80)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().BoolVar(&negateAmount, "neg", false, "Negate amount column value.")
	importCmd.Flags().BoolVar(&allowMatching, "allow-matching", false, "Have output include imported transactions that\nmatch existing ledger transactions.")
	importCmd.Flags().Float64Var(&scaleFactor, "scale", 1.0, "Scale factor to multiply against every imported amount.")
	importCmd.Flags().StringVar(&destAccSearch, "set-search", "Expense", "Search string used to find set of accounts for classification.")
	importCmd.Flags().StringVar(&csvDateFormat, "date-format", "01/02/2006", "Date format.")
	importCmd.Flags().StringVar(&fieldDelimiter, "delimiter", ",", "Field delimiter.")
}

func existingTransaction(generalLedger []*ledger.Transaction, transDate time.Time, payee string) bool {
	for _, trans := range generalLedger {
		if trans.Date == transDate && strings.HasPrefix(trans.Payee, payee) {
			return true
		}
	}
	return false
}
