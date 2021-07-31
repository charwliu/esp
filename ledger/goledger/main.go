package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"go.vixal.xyz/esp/ledger/parser"
)

func main() {
	expression := `

2013-01-01 * "Opening Balance for checking account"
  Assets:US:BofA:Checking                         3219.17 USD
  Equity:Opening-Balances                        -3219.17 USD

2013-01-02 balance Assets:US:BofA:Checking        3219.17 USD


2013-01-04 * "BANK FEES" "Monthly bank fee"
  Assets:US:BofA:Checking                           -4.00 USD
  Expenses:Financial:Fees                            4.00 USD


`
	input := antlr.NewInputStream(expression)

	lexer := parser.NewBeanCountLexer(input)
	parser.Denter = parser.NewDenter(
		parser.BeanCountLexerNEWLINE, parser.BeanCountLexerINDENT, parser.BeanCountLexerDEDENT)
	parser.Denter.BaseLexer = lexer.BaseLexer

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBeanCountParser(stream)

	p.BuildParseTrees = true
	tree := p.Ledger()

	trees := tree.ToStringTree(nil, p)
	fmt.Printf("AST: %s\n", trees)

	var visitor = NewVisitor(parser.NewBuilder())
	var result = visitor.Visit(tree).(*parser.Builder)
	fmt.Println(result.Accounts)
}
