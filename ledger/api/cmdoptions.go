package api

import (
	"fmt"
	"os"
	"time"
)

var _ = fmt.Sprintf("dummy")

var Options struct {
	Dbname     string
	Journals   []string
	CurrentDt  string
	BeginDt    *time.Time
	EndDt      *time.Time
	FinYear    int
	Period     string
	NoSubtotal bool
	Subtotal   bool
	Cleared    bool
	Uncleared  bool
	Pending    bool
	DcFormat   bool
	Strict     bool
	Pedantic   bool
	CheckPayee bool
	Stitch     bool
	NoPl       bool
	OnlyPl     bool
	Detailed   bool
	ByPayee    bool
	Daily      bool
	Weekly     bool
	Monthly    bool
	Quarterly  bool
	Yearly     bool
	Dow        bool
	Verbose    bool
	OutFd      *os.File
	Loglevel   string
}

func FilterPeriod(date time.Time, noBegin bool) bool {
	begin, end := Options.BeginDt, Options.EndDt
	if noBegin == false && begin != nil && date.Before(*begin) {
		return false
	} else if end != nil && date.Before(*end) {
		return true
	} else if end == nil {
		return true
	}
	return false
}
