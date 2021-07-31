package dblentry

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

type firstPass struct {
	defaultComm string
	comments    []string

	currDate         time.Time
	rootAccount      string
	balancingAccount string
	aliases          map[string]string // alias, account-alias
	payees           map[string]string // account-payee map[regex]->accountName
	rePayees         map[string]*regexp.Regexp
	captures         map[string]string
	recaptures       map[string]*regexp.Regexp
	dPayees          map[string]*Payee
}

func (fp *firstPass) initFirstPass() {
	fp.comments = []string{}
	fp.currDate = time.Now()
	fp.aliases = map[string]string{}
	fp.payees = map[string]string{}
	fp.rePayees = map[string]*regexp.Regexp{}
	fp.captures = map[string]string{}
	fp.recaptures = map[string]*regexp.Regexp{}
}

//---- local accessors

func (fp *firstPass) setDefaultComm(name string) {
	fp.defaultComm = name
}

func (fp *firstPass) getDefaultComm() string {
	return fp.defaultComm
}

func (fp *firstPass) addComment(comment string) {
	fp.comments = append(fp.comments, comment)
}

func (fp *firstPass) setYear(year int) error {
	fp.currDate = time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	if fp.currDate.Year() != year {
		return fmt.Errorf("invalid year %v", year)
	}
	return nil
}

func (fp *firstPass) getYear() int {
	return fp.currDate.Year()
}

func (fp *firstPass) setCurrentDate(date time.Time) {
	fp.currDate = date
}

func (fp *firstPass) currentDate() time.Time {
	return fp.currDate
}

func (fp *firstPass) setRootAccount(name string) error {
	if fp.rootAccount != "" {
		return fmt.Errorf("previous `apply` directive(%v) not closed", fp.rootAccount)
	}
	fp.rootAccount = name
	return nil
}

func (fp *firstPass) clearRootAccount() error {
	if fp.rootAccount == "" {
		return fmt.Errorf("dangling `end` directive")
	}
	fp.rootAccount = ""
	return nil
}

func (fp *firstPass) applyRoot(name string) string {
	if fp.rootAccount != "" {
		return fp.rootAccount + ":" + name
	}
	return name
}

func (fp *firstPass) setBalancingAccount(name string) {
	if name != "" {
		fp.balancingAccount = name
	}
}

func (fp *firstPass) getBalancingAccount() string {
	return fp.balancingAccount
}

func (fp *firstPass) addAlias(aliasName, accountName string) {
	if aliasName != "" {
		if _, ok := fp.aliases[aliasName]; ok {
			log.Printf("alias %q already defined for account %q\n", aliasName, accountName)
		}
		fp.aliases[aliasName] = accountName
	}
}

func (fp *firstPass) lookupAlias(name string) string {
	if accountName, ok := fp.aliases[name]; ok {
		return accountName
	}
	return name
}

func (fp *firstPass) addPayee(regex, accountName string) error {
	if regex != "" && accountName != "" {
		fp.payees[regex] = accountName
		regexC, err := regexp.Compile(regex)
		if err != nil {
			return err
		}
		fp.rePayees[regex] = regexC
	}
	return nil
}

func (fp *firstPass) matchAccPayee(payee string) (string, bool) {
	for regex, regexC := range fp.rePayees {
		if regexC.MatchString(payee) {
			return fp.payees[regex], true
		}
	}
	return "", false
}

func (fp *firstPass) addCapture(regex, accountName string) error {
	if regex != "" && accountName != "" {
		fp.captures[regex] = accountName
		regexc, err := regexp.Compile(regex)
		if err != nil {
			return err
		}
		fp.recaptures[regex] = regexc
	}
	return nil
}

func (fp *firstPass) matchCapture(accountName string) (string, bool) {
	for regex, regexC := range fp.recaptures {
		if regexC.MatchString(accountName) {
			return fp.captures[regex], true
		}
	}
	return "", false
}

func (fp *firstPass) findPayee(payee string) *Payee {
	if pe, ok := fp.dPayees[payee]; ok {
		return pe
	}
	fp.dPayees[payee] = NewPayee(payee)
	return fp.dPayees[payee]
}

func (fp *firstPass) matchPayee(payee string) (string, bool) {
	for _, py := range fp.dPayees {
		if name, ok := py.matchAlias(payee); ok {
			return name, true
		}
	}
	return "", false
}

func (fp *firstPass) matchUuid(uuid string) (string, bool) {
	for _, payee := range fp.dPayees {
		if name, ok := payee.matchUuid(uuid); ok {
			return name, true
		}
	}
	return "", false
}
