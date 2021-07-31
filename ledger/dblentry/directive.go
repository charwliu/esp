package dblentry

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	parsec "github.com/prataprc/goparsec"

	"go.vixal.xyz/esp/ledger/api"
)

// Directive can handle all directives in ledger journal.
type Directive struct {
	dType             string
	year              int      // year
	note              string   // account, commodity
	comments          []string // account, commodity
	nDefault          bool     // account, commodity
	accName           string   // account, alias, apply
	accAlias          string   // account
	accPayee          string   // account
	accCheck          string   // account
	accAssert         string   // account
	accEval           string   // account
	accTypes          []string // account
	aliasName         string   // alias
	expression        string   // assert, check
	capture           string   // capture pattern
	commodityName     string   // commodity
	commodityFmt      string   // commodity
	commodityNmrkt    bool     // commodity
	commodityCurrency bool     // commodity
	includeFile       string   // include
	dPayee            string   // payee
	dPayeeAlias       []string // payee
	dPayeeUuid        []string // payee
	endArgs           []string // end
}

// NewDirective create a new Directive instance, one instance to be created
// for handling each directive in the journal file.
func NewDirective() *Directive {
	return &Directive{}
}

//---- exported methods

func (d *Directive) Type() string {
	return d.dType
}

func (d *Directive) IncludeFile() string {
	if d.dType == "include" {
		return d.includeFile
	}
	panic("impossible situation")
}

//---- ledger parser

// Yledger return a parser-combinator that can parse directives.
func (d *Directive) Yledger(db *Datastore) parsec.Parser {
	y := parsec.OrdChoice(
		Vector2scalar,
		d.yaccount(db),
		d.yapply(db),
		d.yalias(db),
		d.yassert(db),
		d.ybucket(db),
		d.ycheck(db),
		d.ycapture(db),
		d.ycheck(db),
		d.ycomment(db),
		d.ycommodity(db),
		d.ydefine(db),
		d.yfixed(db),
		d.yinclude(db),
		d.ypayee(db),
		d.ytest(db),
		d.yend(db),
		d.yyear(db),
	)
	return y
}

// Yledgerblock return a parser-combinator that can parse sub directives under
// account directive.
func (d *Directive) Yledgerblock(db *Datastore, block []string) (int, error) {
	trimstr := func(n parsec.ParsecNode) string {
		return strings.Trim(n.(*parsec.Terminal).Value, " \t")
	}

	var node parsec.ParsecNode
	switch d.dType {
	case "account":
		for index, line := range block {
			scanner := parsec.NewScanner([]byte(line))
			parser := d.yaccountdirectives(db)
			if parser == nil {
				continue
			}
			node, _ = parser(scanner)
			if node == nil {
				return index, fmt.Errorf("parsing %q", line)
			}
			nodes := node.([]parsec.ParsecNode)
			t := nodes[0].(*parsec.Terminal)
			switch t.Name {
			case "DRTV_NOTE":
				d.note = trimstr(nodes[2])
			case "DRTV_SHORTNOTE":
				d.addBlockComment(trimstr(nodes[0]))
			case "DRTV_ACCOUNT_ALIAS":
				d.accAlias = trimstr(nodes[2])
			case "DRTV_ACCOUNT_PAYEE":
				d.accPayee = trimstr(nodes[2])
			case "DRTV_ACCOUNT_CHECK":
				d.accCheck = trimstr(nodes[2])
			case "DRTV_ACCOUNT_ASSERT":
				d.accAssert = trimstr(nodes[2])
			case "DRTV_ACCOUNT_EVAL":
				d.accEval = trimstr(nodes[2])
			case "DRTV_ACCOUNT_TYPE":
				acctypes := api.ParseCsv(trimstr(nodes[2]))
				d.accTypes = d.addAccounttype(acctypes, d.accTypes)
			case "DRTV_DEFAULT":
				d.nDefault = true
			}
		}
		return len(block), nil

	case "commodity":
		for index, line := range block {
			scanner := parsec.NewScanner([]byte(line))
			parser := d.ycommoditydirectives(db)
			if parser == nil {
				continue
			}
			node, _ = parser(scanner)
			if node == nil {
				return index, fmt.Errorf("parsing %q", line)
			}
			nodes := node.([]parsec.ParsecNode)
			t := nodes[0].(*parsec.Terminal)
			switch t.Name {
			case "DRTV_NOTE":
				d.note = nodes[2].(*parsec.Terminal).Value
			case "DRTV_SHORTNOTE":
				d.addBlockComment(trimstr(nodes[0]))
			case "DRTV_COMMODITY_FORMAT":
				d.commodityFmt = nodes[2].(*parsec.Terminal).Value
			case "DRTV_COMMODITY_NOMARKET":
				d.commodityNmrkt = true
			case "DRTV_COMMODITY_CURRENCY":
				d.commodityCurrency = true
			case "DRTV_DEFAULT":
				d.nDefault = true
			}
		}
		return len(block), nil

	case "payee":
		d.dPayeeAlias = []string{}
		d.dPayeeUuid = []string{}
		for index, line := range block {
			scanner := parsec.NewScanner([]byte(line))
			parser := d.ypayeedirectives(db)
			if parser == nil {
				continue
			}
			node, _ = parser(scanner)
			if node == nil {
				return index, fmt.Errorf("parsing %q", line)
			}
			nodes := node.([]parsec.ParsecNode)
			t := nodes[0].(*parsec.Terminal)
			switch t.Name {
			case "DRTV_SHORTNOTE":
				d.addBlockComment(trimstr(nodes[0]))
			case "DRTV_PAYEE_ALIAS":
				aliasv := nodes[2].(*parsec.Terminal).Value
				d.dPayeeAlias = append(d.dPayeeAlias, aliasv)
			case "DRTV_PAYEE_UUID":
				uuidv := nodes[2].(*parsec.Terminal).Value
				d.dPayeeUuid = append(d.dPayeeUuid, uuidv)
			}
		}
		return len(block), nil

	case "apply", "alias", "assert", "bucket", "capture", "check", "comment",
		"define", "fixed", "include", "test", "end", "year":
		return len(block), nil
	}
	panic(fmt.Errorf("unreachable code"))
}

func (d *Directive) yaccount(db *Datastore) parsec.Parser {
	account := NewAccount("") // local scope.
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "account"
			d.accName = nodes[1].(*Account).name
			log.Printf("directive %q %v\n", d.dType, d.accName)
			return d
		},
		ytokAccount, account.Yledger(db),
	)
}

func (d *Directive) yapply(db *Datastore) parsec.Parser {
	account := NewAccount("") // local scope.
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "apply"
			d.accName = nodes[2].(*Account).name
			return d
		},
		ytokApply, ytokAccount, account.Yledger(db),
	)
}

func (d *Directive) yalias(db *Datastore) parsec.Parser {
	account := NewAccount("") // local scope.
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "alias"
			d.aliasName = string(nodes[1].(*parsec.Terminal).Value)
			d.accName = nodes[3].(*Account).name
			return d
		},
		ytokAlias, ytokAliasname, ytokEqual, account.Yledger(db),
	)
}

func (d *Directive) yassert(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "assert"
			d.expression = string(nodes[1].(*parsec.Terminal).Value)
			return d
		},
		ytokAssert, ytokExpr,
	)
}

func (d *Directive) ybucket(db *Datastore) parsec.Parser {
	account := NewAccount("") // local scope
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "bucket"
			d.accName = nodes[1].(*Account).name
			return d
		},
		ytokBucket, account.Yledger(db),
	)
}

func (d *Directive) ycapture(db *Datastore) parsec.Parser {
	account := NewAccount("") // local scope
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "capture"
			d.accName = nodes[1].(*Account).name
			d.capture = nodes[2].(*parsec.Terminal).Value
			return d
		},
		ytokCapture, account.Ypostaccn(db), ytokValue,
	)
}

func (d *Directive) ycheck(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "check"
			d.expression = nodes[1].(*parsec.Terminal).Value
			return d
		},
		ytokCheck, ytokExpr,
	)
}

func (d *Directive) ycomment(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "comment"
			return d
		},
		ytokComment,
	)
}

func (d *Directive) ycommodity(db *Datastore) parsec.Parser {
	commodity := NewCommodity("") // local scope.
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "commodity"
			d.commodityName = nodes[1].(*parsec.Terminal).Value
			log.Printf("directive %q %v\n", d.dType, d.commodityName)
			return d
		},
		ytokDirtCommodity, commodity.Yname(db),
	)
}

func (d *Directive) ydefine(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "define"
			return d
		},
		ytokDirtDefine,
	)
}

func (d *Directive) yfixed(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "fixed"
			return d
		},
		ytokDirtFixed,
	)
}

func (d *Directive) yinclude(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "include"
			d.includeFile = nodes[1].(*parsec.Terminal).Value
			d.includeFile = strings.Trim(d.includeFile, " \t")
			return d
		},
		ytokDirtInclude, ytokValue,
	)
}

func (d *Directive) ypayee(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "payee"
			d.dPayee = nodes[1].(*parsec.Terminal).Value
			return d
		},
		ytokPayee, ytokPayeestr,
	)
}

func (d *Directive) ytest(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "test"
			return d
		},
		ytokDirtTest,
	)
}

func (d *Directive) yend(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "end"
			d.endArgs = []string{"apply", "account"}
			return d
		},
		ytokEnd, ytokApply, ytokAccount,
	)
}

func (d *Directive) yyear(db *Datastore) parsec.Parser {
	return parsec.And(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			d.dType = "year"
			d.year, _ = strconv.Atoi(string(nodes[1].(*parsec.Terminal).Value))
			return d
		},
		ytokYear, ytokYearval,
	)
}

func (d *Directive) yaccountdirectives(db *Datastore) parsec.Parser {
	ynote := parsec.And(nil, ytokNote, ytokHardSpace, ytokValue)
	yalias := parsec.And(nil, ytokAlias, ytokHardSpace, ytokValue)
	ypayee := parsec.And(nil, ytokPayee, ytokHardSpace, ytokValue)
	ycheck := parsec.And(nil, ytokCheck, ytokHardSpace, ytokValue)
	yassert := parsec.And(nil, ytokAssert, ytokHardSpace, ytokValue)
	yeval := parsec.And(nil, ytokEval, ytokHardSpace, ytokValue)
	ytype := parsec.And(nil, ytokType, ytokHardSpace, ytokValue)
	ydefault := parsec.And(nil, ytokDefault)
	yshortnote := parsec.And(nil, ytokDirectivenote)
	y := parsec.OrdChoice(
		Vector2scalar,
		ynote, yalias, ypayee, ycheck, yassert, yeval, ytype, ydefault,
		yshortnote,
	)
	return y
}

func (d *Directive) ycommoditydirectives(db *Datastore) parsec.Parser {
	ynote := parsec.And(nil, ytokNote, ytokHardSpace, ytokValue)
	yformat := parsec.And(nil, ytokFormat, ytokHardSpace, ytokValue)
	ynomarket := parsec.And(nil, ytokNomarket)
	ycurrency := parsec.And(nil, ytokCommCurrency)
	ydefault := parsec.And(nil, ytokDefault)
	y := parsec.OrdChoice(
		Vector2scalar, ynote, yformat, ynomarket, ycurrency, ydefault,
	)
	return y
}

func (d *Directive) ypayeedirectives(db *Datastore) parsec.Parser {
	yalias := parsec.And(nil, ytokPayeeAlias, ytokHardSpace, ytokValue)
	yuuid := parsec.And(nil, ytokPayeeUuid, ytokHardSpace, ytokValue)
	y := parsec.OrdChoice(Vector2scalar, yalias, yuuid)
	return y
}

//---- engine

func (d *Directive) Firstpass(db *Datastore) error {
	switch d.dType {
	case "account":
		return db.declare(d)

	case "apply":
		return db.setRootAccount(d.accName)

	case "alias":
		account := db.GetAccount(d.accName).(*Account)
		account.addAlias(d.aliasName)
		db.addAlias(d.aliasName, d.accName)
		return nil

	case "assert":
		return fmt.Errorf("assert directive not-implemented")

	case "bucket":
		db.setBalancingAccount(d.accName)
		return nil

	case "capture":
		db.addCapture(d.capture, d.accName)
		return nil

	case "check":
		return fmt.Errorf("check directive not-implemented")

	case "comment":
		return fmt.Errorf("comment directive not-implemented")

	case "commodity":
		return db.declare(d)

	case "define":
		return fmt.Errorf("define directive not-implemented")

	case "fixed":
		return fmt.Errorf("fixed directive not-implemented")

	case "include":
		return nil

	case "payee":
		return db.declare(d)

	case "test":
		return fmt.Errorf("test directive not-implemented")

	case "end":
		return db.clearRootAccount()

	case "year":
		return db.setYear(d.year)
	}
	panic("unreachable code")
}

func (d *Directive) Secondpass(db *Datastore) error {
	return nil
}

func (d *Directive) addAccounttype(typenames []string, acc []string) []string {
	if acc == nil {
		acc = []string{}
	}
	for _, typename := range typenames {
		typename = strings.ToLower(typename)
		if api.HasString(acc, typename) {
			continue
		}
		switch typename {
		case "income":
			implies := []string{"credit"}
			acc = d.addAccounttype(implies, append(acc, typename))
		case "expense":
			implies := []string{"debit"}
			acc = d.addAccounttype(implies, append(acc, typename))
		default:
			acc = append(acc, typename)
		}
	}
	return acc
}

func (d *Directive) addBlockComment(comment string) {
	if d.comments == nil {
		d.comments = make([]string, 0)
	}
	d.comments = append(d.comments, comment)
}
