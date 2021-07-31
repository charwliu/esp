package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAccountValid(t *testing.T) {
	assert.True(t, IsAccountValid("Expenses:Financial:Fees"))
	assert.False(t, IsAccountValid("assets:Financial:Fees"))
	assert.True(t, IsAccountValid("Expenses:1financial:1fees"))
	assert.False(t, IsAccountValid("1assets:Financial:Fees"))
	assert.True(t, IsAccountValid("Expenses:Financial-a:Fees"))
	assert.False(t, IsAccountValid("Expenses:Financial -a:Fees"))
	assert.True(t, IsAccountValid("Expenses-1:Financial-a:Fees"))
	assert.False(t, IsAccountValid("Assets"))
	assert.False(t, IsAccountValid("Invalid"))
	assert.False(t, IsAccountValid("Other"))
	assert.False(t, IsAccountValid("Assets:US:RBS*Checking"))
	assert.False(t, IsAccountValid("Assets:US:RBS:Checking&"))
	assert.False(t, IsAccountValid("Assets:US:RBS:checking"))
	assert.False(t, IsAccountValid("Assets:us:RBS:checking"))
}

func TestJoinAccount(t *testing.T) {
	assert.Equal(t, "Expenses:Toys:Computer", JoinAccount([]string{"Expenses", "Toys", "Computer"}))
	assert.Equal(t, "Expenses", JoinAccount([]string{"Expenses"}))
	assert.Equal(t, "", JoinAccount([]string{}))
}

func TestSplitAccount(t *testing.T) {
	assert.Equal(t, []string{"Expenses", "Toys", "Computer"}, SplitAccount("Expenses:Toys:Computer"))
	assert.Equal(t, []string{"Expenses"}, SplitAccount("Expenses"))
	assert.Equal(t, []string{""}, SplitAccount(""))
}

func TestLeafAccount(t *testing.T) {
	assert.Equal(t, "Fees", LeafAccount("Expenses:Financial:Fees"))
	assert.Equal(t, "Financial", LeafAccount("Expenses:Financial"))
	assert.Equal(t, "Expenses", LeafAccount("Expenses"))
	assert.Equal(t, "", LeafAccount(""))
}

func TestParentAccount(t *testing.T) {
	assert.Equal(t, "Expenses:Financial", ParentAccount("Expenses:Financial:Fees"))
	assert.Equal(t, "Expenses", ParentAccount("Expenses:Financial"))
	assert.Equal(t, "", ParentAccount("Expenses"))
	assert.Equal(t, "", ParentAccount(""))
}

func TestAccountSansRoot(t *testing.T) {
	assert.Equal(t, "Financial:Fees", AccountSansRoot("Expenses:Financial:Fees"))
	assert.Equal(t, "US:BofA:Checking", AccountSansRoot("Assets:US:BofA:Checking"))
	assert.Equal(t, "", AccountSansRoot("Assets"))
}

func TestAccountRoot(t *testing.T) {
	name := "Liabilities:US:Credit-Card:Blue"
	assert.Equal(t, "", AccountRoot(name, 0))
	assert.Equal(t, "Liabilities", AccountRoot(name, 1))
	assert.Equal(t, "Liabilities:US", AccountRoot(name, 2))
	assert.Equal(t, "Liabilities:US:Credit-Card", AccountRoot(name, 3))
	assert.Equal(t, "Liabilities:US:Credit-Card:Blue", AccountRoot(name, 4))
	assert.Equal(t, "Liabilities:US:Credit-Card:Blue", AccountRoot(name, 5))
}

func TestHasAccountComponent(t *testing.T) {
	assert.True(t, HasAccountComponent("Liabilities:US:Credit-Card", "US"))
	assert.False(t, HasAccountComponent("Liabilities:US:Credit-Card", "CA"))
	assert.True(t, HasAccountComponent("Liabilities:US:Credit-Card", "Credit-Card"))
	assert.True(t, HasAccountComponent("Liabilities:US:Credit-Card", "Liabilities"))
	assert.False(t, HasAccountComponent("Liabilities:US:Credit-Card", "Credit"))
	assert.False(t, HasAccountComponent("Liabilities:US:Credit-Card", "Card"))
}

func TestCommonPrefix(t *testing.T) {
	var testSuite = []struct {
		data   []string
		expect string
	}{
		{
			data: []string{
				"Assets:US:Bank",
				"Liabilities:US:Mortgage",
			},
			expect: "",
		},
		{
			data: []string{
				"Assets:US:BofA:Checking",
				"Assets:US:BofA:Cash",
			},
			expect: "Assets:US:BofA",
		},
		{
			data: []string{
				"Assets:US:BofA:Checking",
				"Assets:US:BofA:Cash",
				"Assets:US:BofA:Saving",
			},
			expect: "Assets:US:BofA",
		},
		{
			data: []string{
				"Liabilities:US:Mortgage",
				"Liabilities",
			},
			expect: "Liabilities",
		},
		{
			data: []string{
				"Liabilities",
				"Liabilit",
			},
			expect: "",
		},
		{
			data: []string{
				"Liabilities:US:Mort",
				"Liabilities:US:Mortgage",
			},
			expect: "Liabilities:US",
		},
	}
	for _, s := range testSuite {
		prefix := CommonPrefix(s.data)
		assert.Equal(t, s.expect, prefix)
	}

}
