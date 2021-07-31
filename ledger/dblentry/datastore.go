package dblentry

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	parsec "github.com/prataprc/goparsec"

	"go.vixal.xyz/esp/ledger/api"
)

// ParsePhase state to be tracked at datastore.
type ParsePhase int

const (
	// DBSTART is parsing started.
	DBSTART ParsePhase = iota + 1
	// DBFIRSTPASS is all FirstPass() calls are completed.
	DBFIRSTPASS
	// DBSECONDPASS is all SecondPass() calls are completed.
	DBSECONDPASS
)

// Datastore managing accounts, commodities, transactions and posting and
// every thing else that are related.
type Datastore struct {
	// immutable from first initialization.
	name string

	// immutable once firstPass is ok
	journals    map[uint64]string
	currjournal string
	firstPass

	// changes with every second pass.
	reporter    api.Reporter
	pass        ParsePhase
	commodities map[string]*Commodity
	accntdb     map[string]*Account // full account-name -> account
	dclrdacc    []string
	dclrdcomm   []string
	dclrdpayee  []string
	de          *DoubleEntry
	transdb     *DB
	pricedb     *DB

	// configuration
	periodtill *time.Time
}

// NewDatastore return a new datastore.
func NewDatastore(name string, reporter api.Reporter) *Datastore {
	db := &Datastore{
		name:     name,
		journals: make(map[uint64]string),
		reporter: reporter,

		pass:        DBSTART,
		transdb:     NewDB(fmt.Sprintf("%v-transactions", name)),
		pricedb:     NewDB(fmt.Sprintf("%v-pricedb", name)),
		accntdb:     map[string]*Account{},
		commodities: map[string]*Commodity{},
		de:          NewDoubleEntry("master"),
	}
	db.initFirstPass()
	db.defaultprices()
	return db
}

//---- local accessors

func (db *Datastore) assertfirstpass() {
	if db.pass < DBFIRSTPASS {
		panic("impossible situation")
	}
}

func (db *Datastore) getCommodity(name string, defcomm *Commodity) *Commodity {
	defaultcomm := db.getDefaultComm()
	if name == "" && defaultcomm != "" {
		return db.commodities[defaultcomm]
	}
	if defaultcomm == "" && name != "" {
		db.setDefaultComm(name)
	}
	if comm, ok := db.commodities[name]; ok {
		return comm
	}
	log.Printf("commodity %q added\n", name)
	if defcomm == nil {
		defcomm = NewCommodity(name)
	}
	db.commodities[name] = defcomm
	return defcomm
}

//---- exported accessors

func (db *Datastore) Addjournal(journalfile string, data []byte) {
	hash := api.Crc64(data)
	db.journals[hash] = journalfile
	db.currjournal = journalfile
}

func (db *Datastore) Hasjournal(data []byte) bool {
	hash := api.Crc64(data)
	_, ok := db.journals[hash]
	return ok
}

func (db *Datastore) CurrentJournal() string {
	return db.currjournal
}

func (db *Datastore) GetCommodity(name string) api.Commoditiser {
	return db.getCommodity(name, nil)
}

func (db *Datastore) Applytill(till time.Time) {
	periodtill := till
	db.periodtill = &periodtill
}

// Firstpassok to track parsephase
func (db *Datastore) FirstPassOk() {
	db.pass = DBFIRSTPASS
}

// Secondpassok to track parsephase
func (db *Datastore) SecondPassOk() {
	db.pass = DBSECONDPASS
}

// PrintAccounts is a debug api for application.
func (db *Datastore) PrintAccounts() {
	for _, accname := range db.AccountNames() {
		log.Printf("-- %v\n", db.accntdb[accname])
	}
}

//---- api.Datastorer methods

func (db *Datastore) IsCommodityDeclared(name string) bool {
	for _, xname := range db.dclrdcomm {
		if xname == name {
			return true
		}
	}
	return false
}

func (db *Datastore) IsAccountDeclared(name string) bool {
	for _, xname := range db.dclrdacc {
		if xname == name {
			return true
		}
	}
	return false
}

func (db *Datastore) IsPayeeDeclared(name string) bool {
	for _, xname := range db.dclrdpayee {
		if xname == name {
			return true
		}
	}
	return false
}

func (db *Datastore) GetAccount(name string) api.Accounter {
	if name == "" {
		return (*Account)(nil)
	}
	account, ok := db.accntdb[name]
	if ok == false {
		account = NewAccount(name)
	}
	db.accntdb[name] = account
	return account
}

func (db *Datastore) AccountNames() []string {
	db.assertfirstpass()

	accnames := []string{}
	for name := range db.accntdb {
		accnames = append(accnames, name)
	}
	sort.Strings(accnames)
	return accnames
}

func (db *Datastore) CommodityNames() []string {
	db.assertfirstpass()

	cnames := []string{}
	for name := range db.commodities {
		cnames = append(cnames, name)
	}
	sort.Strings(cnames)
	return cnames
}

func (db *Datastore) Balance(obj interface{}) (balance api.Commoditiser) {
	db.assertfirstpass()
	return db.de.Balance(obj)
}

func (db *Datastore) Balances() []api.Commoditiser {
	db.assertfirstpass()
	return db.de.Balances()
}

func (db *Datastore) Debit(obj interface{}) (balance api.Commoditiser) {
	db.assertfirstpass()
	return db.de.Debit(obj)
}

func (db *Datastore) Debits() []api.Commoditiser {
	db.assertfirstpass()
	return db.de.Debits()
}

func (db *Datastore) Credit(obj interface{}) (balance api.Commoditiser) {
	db.assertfirstpass()
	return db.de.Credit(obj)
}

func (db *Datastore) Credits() []api.Commoditiser {
	db.assertfirstpass()
	return db.de.Credits()
}

func (db *Datastore) AggregateTotal(trans api.Transactor, p api.Poster) error {
	posting := p.(*Posting)
	names := SplitAccount(posting.account.name)
	parts := []string{}
	for _, name := range names[:len(names)-1] {
		parts = append(parts, name)
		fullname := strings.Join(parts, ":")
		consacc := db.GetAccount(fullname).(*Account)
		if err := consacc.addBalance(posting.commodity); err != nil {
			return err
		}
		err := db.reporter.BubblePosting(db, trans, posting, consacc)
		if err != nil {
			return err
		}
	}
	return nil
}

//---- engine

func (db *Datastore) FirstPass(obj interface{}) (err error) {
	if trans, ok := obj.(*Transaction); ok {
		if err := trans.Firstpass(db); err != nil {
			return err
		}
		db.setCurrentDate(trans.date)
		db.transdb.Insert(trans.date, trans)

	} else if price, ok := obj.(*Price); ok {
		err = db.pricedb.Insert(price.when, price)

	} else if directive, ok := obj.(*Directive); ok {
		err = directive.Firstpass(db)

	} else if comment, ok := obj.(*Comment); ok {
		err = comment.FirstPass(db)
		db.addComment(comment.line)
	}
	return err
}

func (db *Datastore) SecondPass() error {
	entries := []api.TimeEntry{}

	for _, entry := range db.transdb.Range(nil, nil, "both", entries) {
		trans := entry.Value().(*Transaction)
		if db.periodtill == nil || trans.Date().Before(*db.periodtill) {
			if err := trans.Secondpass(db); err != nil {
				return err
			}
		}
	}
	return nil
}

func (db *Datastore) Clone(nreporter api.Reporter) api.Datastorer {
	ndb := *db
	ndb.reporter = nreporter

	ndb.commodities = map[string]*Commodity{}
	for name, commodity := range db.commodities {
		ndb.commodities[name] = commodity.Clone(&ndb).(*Commodity)
	}

	ndb.accntdb = map[string]*Account{}
	for name, account := range db.accntdb {
		ndb.accntdb[name] = account.Clone(&ndb)
	}

	ndb.de = db.de.Clone()

	ndb.transdb = NewDB(fmt.Sprintf("%v-transactions", ndb.name))
	entries := []api.TimeEntry{}
	for _, entry := range db.transdb.Range(nil, nil, "both", entries) {
		k, ntrans := entry.Key(), entry.Value().(*Transaction).Clone(&ndb)
		ndb.transdb.Insert(k, ntrans)
	}

	ndb.pricedb = NewDB(fmt.Sprintf("%v-pricedb", ndb.name))
	entries = []api.TimeEntry{}
	for _, entry := range db.pricedb.Range(nil, nil, "both", entries) {
		k, nprice := entry.Key(), entry.Value().(*Price).Clone(&ndb)
		ndb.pricedb.Insert(k, nprice)
	}
	return &ndb
}

func (db *Datastore) addBalance(commodity *Commodity) error {
	return db.de.AddBalance(commodity)
}

// directive-account

func (db *Datastore) declare(value interface{}) error {
	switch v := value.(type) {
	case *Directive:
		d := v
		switch d.dType {
		case "account":
			db.addAlias(d.accAlias, d.accName)
			err := db.addPayee(d.accPayee, d.accName)
			if err != nil {
				return err
			}
			account := db.GetAccount(d.accName).(*Account)
			if len(d.accTypes) > 0 {
				account.types = d.addAccounttype(d.accTypes, account.types)
			}
			account.addNote(d.note)
			account.addAlias(d.accAlias)
			account.addPayee(d.accPayee)
			account.addComments(d.comments...)
			if d.nDefault {
				db.setBalancingAccount(account.name)
			}
			db.dclrdacc = append(db.dclrdacc, d.accName)

		case "commodity":
			scanner := parsec.NewScanner([]byte(d.commodityFmt))
			node, _ := NewCommodity("").Yledger(db)(scanner)
			commodity := node.(*Commodity)
			if commodity.name != "" && commodity.name != d.commodityName {
				x, y := commodity.name, d.commodityName
				return fmt.Errorf("name mismatching %q vs %q", x, y)
			} else if commodity.name == "" {
				commodity.name = d.commodityName
				commodity.noname = true
			}

			commodity.addNote(d.note)
			if d.nDefault {
				db.setDefaultComm(commodity.name)
			}
			if d.commodityNmrkt {
				commodity.nomarket = true
			}
			if d.commodityCurrency {
				commodity.currency = true
			}
			// now finally update the datastore.commodity db.
			db.commodities[commodity.name] = commodity
			db.dclrdcomm = append(db.dclrdcomm, d.commodityName)

		case "payee":
			payee := db.findPayee(d.dPayee)
			for _, alias := range d.dPayeeAlias {
				if err := payee.addAlias(alias); err != nil {
					return err
				}
			}
			for _, uuid := range d.dPayeeUuid {
				payee.addUuid(uuid)
			}
			db.dclrdpayee = append(db.dclrdpayee, d.dPayee)

		}
		return nil
	}
	panic("unreachable code")
}

//---- api.Reporter methods

func (db *Datastore) FmtBalances(
	_ api.Datastorer, trans api.Transactor, p api.Poster,
	acc api.Accounter) [][]string {

	var rows [][]string

	if len(db.Balances()) == 0 {
		return append(rows, []string{"", "", "-"})
	}

	for _, balance := range db.Balances() {
		rows = append(rows, []string{"", "", balance.String()})
	}
	if len(rows) > 0 { // last row to include date.
		lastrow := rows[len(rows)-1]
		date := trans.Date().Format("2006/Jan/02")
		lastrow[0] = date
	}
	return rows
}

func (db *Datastore) FmtDCBalances(
	_ api.Datastorer, trans api.Transactor, p api.Poster,
	acc api.Accounter) [][]string {

	var rows [][]string

	if len(db.Balances()) == 0 {
		return append(rows, []string{"", "", "-", "-", "-"})
	}

	for _, bal := range db.Balances() {
		name := bal.Name()
		dr, cr := db.Debit(name), db.Credit(name)
		cols := []string{"", "", dr.String(), cr.String(), bal.String()}
		rows = append(rows, cols)
	}
	if len(rows) > 0 { // last row to include date.
		lastrow := rows[len(rows)-1]
		date := trans.Date().Format("2006/Jan/02")
		lastrow[0] = date
	}
	return rows
}

func (db *Datastore) FmtEquity(
	_ api.Datastorer, trans api.Transactor, _ api.Poster,
	_ api.Accounter) [][]string {

	panic("not supported")
}

func (db *Datastore) FmtPassbook(
	_ api.Datastorer, trans api.Transactor, _ api.Poster,
	_ api.Accounter) [][]string {

	panic("not supported")
}

func (db *Datastore) FmtRegister(
	_ api.Datastorer, trans api.Transactor, p api.Poster,
	acc api.Accounter) [][]string {

	panic("not supported")
}

func (db *Datastore) defaultprices() {
	_ = []string{
		"P 01/01/2000 kb 1024b",
		"P 01/01/2000 mb 1024kb",
		"P 01/01/2000 gb 1024mb",
		"P 01/01/2000 tb 1024gb",
		"P 01/01/2000 pb 1024tb",

		"P 01/01/2000 m 60s",
		"P 01/01/2000 h 60m",
	}
}
