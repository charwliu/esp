package dblentry

import (
	"fmt"
	"log"
	"strings"
	"time"

	parsec "github.com/prataprc/goparsec"

	"go.vixal.xyz/esp/ledger/api"
)

const (
	// PostUncleared notion for a posting.
	PostUncleared = "uncleared"
	// PostCleared notion for a posting.
	PostCleared = "cleared"
	// PostPending notion for a posting.
	PostPending = "pending"
)

var prefix2state = map[rune]string{
	'*': PostCleared,
	'!': PostPending,
}

// Posting instance for every single posting within a transaction.
type Posting struct {
	trans     *Transaction
	account   *Account
	virtual   bool
	balanced  bool
	commodity *Commodity

	lotPrice  *Commodity
	lotDate   time.Time
	costPrice *Commodity
	balPrice  *Commodity

	tags     []string
	metadata map[string]interface{}
	note     string
}

// NewPosting create a new posting instance.
func NewPosting(trans *Transaction) *Posting {
	return &Posting{
		trans:    trans,
		tags:     []string{},
		metadata: map[string]interface{}{},
	}
}

//---- local accessors

func (p *Posting) getMetadata(key string) interface{} {
	if value, ok := p.metadata[key]; ok {
		return value
	}
	return p.trans.getMetadata(key)
}

func (p *Posting) setMetadata(key string, value interface{}) {
	p.metadata[strings.ToLower(key)] = value
}

func (p *Posting) isVirtual() bool {
	return p.virtual
}

func (p *Posting) isBalanced() bool {
	return p.balanced
}

//---- api.Poster methods.

func (p *Posting) Account() api.Accounter {
	return p.account
}

func (p *Posting) Commodity() api.Commoditiser {
	return p.commodity
}

func (p *Posting) LotPrice() api.Commoditiser {
	return p.lotPrice
}

func (p *Posting) CostPrice() api.Commoditiser {
	return p.costPrice
}

func (p *Posting) BalancePrice() api.Commoditiser {
	return p.balPrice
}

func (p *Posting) Payee() string {
	payee := p.getMetadata("payee")
	if payee == nil {
		return ""
	}
	return payee.(string)
}

func (p *Posting) IsCredit() bool {
	if p.commodity == nil {
		panic("impossible situation")
	}
	return p.commodity.IsCredit()
}

func (p *Posting) IsDebit() bool {
	if p.commodity == nil {
		panic("impossible situation")
	}
	return p.commodity.IsDebit()
}

func (p *Posting) getState() string {
	state := p.getMetadata("state")
	if state == nil {
		return ""
	}
	return state.(string)
}

//---- ledger parser

// Yledger return parser-combinator that can parse a posting line within
// a transaction.
func (p *Posting) Yledger(db *Datastore) parsec.Parser {
	account := NewAccount("") // kept alive till FirstPass.
	comm := NewCommodity("")
	lotprice := NewCommodity("")
	costprice := NewCommodity("")
	balprice := NewCommodity("")

	ylotdate := parsec.And(
		nil,
		ytokOpensqrt, Ydate(db.getYear()), ytokClosesqrt,
	)

	yposting := parsec.And(
		nil,
		parsec.Maybe(maybeNode, ytokPrefix),
		account.Ypostaccn(db),
		parsec.Maybe(maybeNode, comm.Yledger(db)),
		parsec.Maybe(maybeNode, lotprice.Ylotprice(db)),
		parsec.Maybe(maybeNode, ylotdate),
		parsec.Maybe(maybeNode, costprice.Ycostprice(db)),
		parsec.Maybe(maybeNode, balprice.Ybalprice(db)),
		parsec.Maybe(maybeNode, ytokPostnote),
	)

	var err error

	y := parsec.OrdChoice(
		func(nodes []parsec.ParsecNode) parsec.ParsecNode {
			switch items := nodes[0].(type) {
			case []parsec.ParsecNode:
				// prefix
				if t, ok := items[0].(*parsec.Terminal); ok {
					p.setMetadata("state", prefix2state[[]rune(t.Value)[0]])
				}

				p.account, p.virtual, p.balanced = p.fixAccount(db, items[1])
				p.commodity, err = p.fixCommodity(db, items[2]) // commodity
				if err != nil {
					return err
				}
				p.lotPrice, err = p.fixLotPrice(db, items[3]) // lot price
				if err != nil {
					return err
				}
				p.lotDate, err = p.fixLotDate(items[4]) // lot date
				if err != nil {
					return err
				}
				p.costPrice, err = p.fixcostprice(db, items[5]) // cost price
				if err != nil {
					return err
				}
				p.balPrice, err = p.fixbalprice(db, items[6]) // balance price
				if err != nil {
					return err
				}

				if p.lotPrice != nil && lotprice.currency == false {
					return fmt.Errorf("lot price must be currency")
				} else if p.costPrice != nil && costprice.currency == false {
					return fmt.Errorf("cost price must be currency")
				}
				if x, y := p.balPrice, p.commodity; x != nil && y != nil {
					if x.name != y.name {
						fmsg := "balance-commodity(%v) != posting-commodity(%v)"
						return fmt.Errorf(fmsg, x.name, y.name)
					}
				}

				// optionally tags or tagkv or note
				if note, ok := items[7].(*parsec.Terminal); ok {
					input := strings.Trim(note.Value, "; ")
					scanner := parsec.NewScanner([]byte(input))
					if node, _ := NewTag().Yledger(db)(scanner); node == nil {
						p.note = note.Value

					} else {
						tag := node.(*Tags)
						p.tags = append(p.tags, tag.tags...)
						for k, v := range tag.tagm {
							p.setMetadata(k, v)
						}
					}
				}

				fmsg := "posting.yledger account:%v commodity:%v %v\n"
				log.Printf(fmsg, p.account, p.commodity, p.costPrice)
				return p

			case *parsec.Terminal:
				inp := []byte(strings.TrimLeft(items.Value, ";"))
				scanner := parsec.NewScanner(inp)
				node, _ := NewTag().Yledger(db)(scanner)
				if node == nil {
					log.Printf("posting.yledger %v\n", string(items.Value))
					return typeTransnote(items.Value)
				}
				log.Printf("posting.yledger %v\n", node)
				return node.(*Tags)
			}
			fmsg := "unreachable code posting: len(nodes): %v"
			panic(fmt.Errorf(fmsg, len(nodes)))
		},
		yposting,
		ytokTransnote,
	)
	return y
}

//---- engine

func (p *Posting) fixAccount(
	db *Datastore, item interface{}) (*Account, bool, bool) {

	account := item.(*Account)
	return account, account.virtual, account.balanced
}

func (p *Posting) fixCommodity(
	db *Datastore, item interface{}) (*Commodity, error) {

	if commodity, ok := item.(*Commodity); ok {
		cname := commodity.name
		c := db.getCommodity(cname, commodity).makeSimilar(commodity.amount)
		return c, nil
	}
	return nil, nil
}

func (p *Posting) fixLotPrice(
	db *Datastore, item interface{}) (*Commodity, error) {

	if lotPrice, ok := item.(*Commodity); ok {
		return lotPrice, nil
	}
	return nil, nil
}

func (p *Posting) fixLotDate(item interface{}) (time.Time, error) {
	var tm time.Time

	if lotnodes, ok := item.([]parsec.ParsecNode); ok {
		if tm, ok = lotnodes[1].(time.Time); ok {
			return tm, nil
		}
		return tm, lotnodes[1].(error)
	}
	return tm, nil
}

func (p *Posting) fixcostprice(
	db *Datastore, item interface{}) (*Commodity, error) {

	if costprice, ok := item.(*Commodity); ok {
		return costprice, nil
	}
	return nil, nil
}

func (p *Posting) fixbalprice(
	db *Datastore, item interface{}) (*Commodity, error) {

	if balprice, ok := item.(*Commodity); ok {
		return balprice, nil
	}
	return nil, nil
}

func (p *Posting) getCostprice() *Commodity {
	checkdebit := p.IsDebit() && p.commodity.currency == false
	if checkdebit && p.costPrice != nil {
		if p.costPrice.isTotal() { // first compute per unit price
			p.costPrice.amount /= p.commodity.amount
		}
		return p.costPrice.makeSimilar(p.commodity.amount * p.costPrice.amount)

	}

	checkcredit := p.IsCredit() && p.commodity.currency == false
	if checkcredit && p.lotPrice != nil {
		if p.lotPrice.isTotal() { // first compute per unit price
			p.lotPrice.amount /= p.commodity.amount
		}
		return p.lotPrice.makeSimilar(p.commodity.amount * p.lotPrice.amount)
	}

	return p.commodity.makeSimilar(p.commodity.amount)
}

func (p *Posting) FirstPass(db *Datastore, trans *Transaction) error {
	// payee-rewrite
	if val, ok := p.metadata["payee"]; ok {
		payee := val.(string)
		if payee, ok = db.matchPayee(payee); ok {
			p.setMetadata("payee", payee)
		}
	}

	accname := p.account.name

	// if account is Unknown, try rewrite !!
	if p.account.isUnknown() {
		// fetch the declared account name with payee
		daccname, ok := db.matchAccPayee(trans.Payee())
		if ok == false {
			fmsg := "Unknown account %q has no matching payee %q"
			return fmt.Errorf(fmsg, p.account.name, trans.Payee())
		}
		prefix := p.account.name[:len(p.account.name)-len("Unknown")]
		if strings.HasPrefix(daccname, prefix) == false {
			fmsg := "Unknown account %q has no matching prefix with %q"
			return fmt.Errorf(fmsg, p.account.name, daccname)
		}
		accname = prefix + daccname[len(prefix):]

	} else if accname1, ok := db.matchCapture(accname); ok {
		accname = accname1

	} else {
		accname = db.applyRoot(db.lookupAlias(accname))
	}

	p.account = db.GetAccount(accname).(*Account)

	if err := p.account.FirstPass(db, trans, p); err != nil {
		return err
	}
	if err := p.commodity.Firstpass(db, trans, p); err != nil {
		return err
	}

	db.reporter.FirstPass(db, trans, p)

	return nil
}

func (p *Posting) SecondPass(db *Datastore, trans *Transaction) error {
	db.addBalance(p.commodity)
	p.account.setPosting()

	if err := p.account.SecondPass(db, trans, p); err != nil {
		return err
	}
	if err := p.commodity.SecondPass(db, trans, p); err != nil {
		return err
	}

	return db.reporter.Posting(db, trans, p)
}

func (p *Posting) Clone(ndb *Datastore, ntrans *Transaction) *Posting {
	np := *p
	np.trans = ntrans
	np.account = ndb.GetAccount(p.account.name).(*Account)
	if p.commodity != nil {
		np.commodity = p.commodity.Clone(ndb).(*Commodity)
	}
	if p.lotPrice != nil {
		np.lotPrice = p.lotPrice.Clone(ndb).(*Commodity)
	}
	if p.costPrice != nil {
		np.costPrice = p.costPrice.Clone(ndb).(*Commodity)
	}
	if p.balPrice != nil {
		np.balPrice = p.balPrice.Clone(ndb).(*Commodity)
	}
	return &np
}

//---- api.Reporter methods

func (p *Posting) FmtBalances(
	db api.Datastorer, trans api.Transactor, _ api.Poster,
	_ api.Accounter) [][]string {

	panic("not supported")
}

func (p *Posting) FmtRegister(
	db api.Datastorer, trans api.Transactor, _ api.Poster,
	_ api.Accounter) [][]string {

	panic("not supported")
}

func (p *Posting) FmtEquity(
	db api.Datastorer, trans api.Transactor, _ api.Poster,
	_ api.Accounter) [][]string {

	panic("not supported")
}

func (p *Posting) FmtPassbook(
	db api.Datastorer, trans api.Transactor, _ api.Poster,
	_ api.Accounter) [][]string {

	panic("not supported")
}
