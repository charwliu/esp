package dblentry

import (
	"regexp"
)

type Payee struct {
	name      string
	aliases   []string // regular expression that can match transactions.
	reAliases []*regexp.Regexp
	uuids     []string
}

func NewPayee(name string) *Payee {
	return &Payee{
		name:      name,
		aliases:   []string{},
		reAliases: []*regexp.Regexp{},
		uuids:     []string{},
	}
}

func (payee *Payee) addAlias(pattern string) error {
	if pattern == "" {
		return nil
	}
	payee.aliases = append(payee.aliases, pattern)
	regC, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	payee.reAliases = append(payee.reAliases, regC)
	return nil
}

func (payee *Payee) addUuid(uuid string) {
	if uuid == "" {
		return
	}
	payee.uuids = append(payee.uuids, uuid)
}

func (payee *Payee) matchAlias(alias string) (string, bool) {
	for _, regC := range payee.reAliases {
		if regC.MatchString(alias) {
			return payee.name, true
		}
	}
	return "", false
}

func (payee *Payee) matchUuid(uuid string) (string, bool) {
	for _, xUuid := range payee.uuids {
		if xUuid == uuid {
			return payee.name, true
		}
	}
	return "", false
}
