package ledger

import (
	"regexp"
	"strings"
)

var (
	kAssets      = "Assets"
	kLiabilities = "Liabilities"
	kEquity      = "Equity"
	kIncome      = "Income"
	kExpenses    = "Expenses"
)

func GetDefaultAccountTypes() *AccountTypes {
	return &AccountTypes{
		Assets:      &kAssets,
		Liabilities: &kLiabilities,
		Equity:      &kEquity,
		Income:      &kIncome,
		Expenses:    &kExpenses,
	}
}

var DefaultAccountTypes = &AccountTypes{
	Assets:      &kAssets,
	Liabilities: &kLiabilities,
	Equity:      &kEquity,
	Income:      &kIncome,
	Expenses:    &kExpenses,
}

func GetAccountType(account string) string {
	return AccountRoot(account, 1)
}

func GetAccountSortKey(accountTypes *AccountTypes, accountName string) map[int]string {
	atype := GetAccountType(accountName)
	itype := -1
	if atype == accountTypes.GetAssets() {
		itype = 0
	} else if atype == accountTypes.GetLiabilities() {
		itype = 1
	} else if atype == accountTypes.GetEquity() {
		itype = 2
	} else if atype == accountTypes.GetIncome() {
		itype = 3
	} else if atype == accountTypes.GetExpenses() {
		itype = 4
	}
	return map[int]string{itype: accountName}
}

func IsAccountType(accountType string, accountName string) bool {
	return strings.HasPrefix(accountName, accountType) &&
		(len(accountType) == len(accountName) ||
			len(accountType) <= len(accountName)) &&
		strings.HasPrefix(accountName[len(accountType):], kSep)
}

var reRoot *regexp.Regexp

func init() {
	reRoot = regexp.MustCompile(kAccountCompTypeRe + "$")
}
func IsRootAccount(accountName string) bool {
	return reRoot.MatchString(accountName)
}

func IsBalanceSheetAccount(accountName string, accountTypes *AccountTypes) bool {
	atype := GetAccountType(accountName)

	return atype == accountTypes.GetAssets() ||
		atype == accountTypes.GetLiabilities() ||
		atype == accountTypes.GetEquity()
}

func IsIncomeStatementAccount(accountName string, accountTypes *AccountTypes) bool {
	atype := GetAccountType(accountName)

	return atype == accountTypes.GetIncome() ||
		atype == accountTypes.GetExpenses()
}

func IsEquityAccount(accountName string, accountTypes *AccountTypes) bool {
	atype := GetAccountType(accountName)
	return atype == accountTypes.GetEquity()
}

func IsInvertedAccount(accountName string, accountTypes *AccountTypes) bool {
	atype := GetAccountType(accountName)

	return atype == accountTypes.GetIncome() ||
		atype == accountTypes.GetLiabilities() ||
		atype == accountTypes.GetEquity()
}

func GetAccountSign(accountName string, types *AccountTypes) int {
	atype := GetAccountType(accountName)
	if atype == types.GetAssets() || atype == types.GetExpenses() {
		return +1
	} else {
		return -1
	}
}
