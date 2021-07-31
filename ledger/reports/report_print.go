package reports

import (
	"fmt"

	"go.vixal.xyz/esp/ledger/api"
	"go.vixal.xyz/esp/ledger/dblentry"
)

// ReportPrint for balance reporting.
type ReportPrint struct {
	transdb *dblentry.DB
}

// NewReportPrint creates an instance for balance reporting
func NewReportPrint(args []string) *ReportPrint {
	report := &ReportPrint{transdb: dblentry.NewDB("report_print")}
	return report
}

//---- api.Reporter methods

func (report *ReportPrint) FirstPass(
	db api.Datastorer, trans api.Transactor, p api.Poster) error {

	return nil
}

func (report *ReportPrint) Transaction(
	_ api.Datastorer, trans api.Transactor) error {

	date := trans.Date()
	if dt := api.Options.BeginDt; dt != nil && date.Before(*dt) {
		return nil
	} else if dt = api.Options.EndDt; dt != nil && date.Before(*dt) {
		report.transdb.Insert(date, trans)
	} else {
		report.transdb.Insert(date, trans)
	}
	return nil
}

func (report *ReportPrint) Posting(
	_ api.Datastorer, _ api.Transactor, _ api.Poster) error {

	return nil
}

func (report *ReportPrint) BubblePosting(
	_ api.Datastorer, _ api.Transactor, _ api.Poster, _ api.Accounter) error {

	return nil
}

func (report *ReportPrint) Render(args []string, ndb api.Datastorer) {
	outfd := api.Options.OutFd
	var entries []api.TimeEntry
	for _, entry := range report.transdb.Range(nil, nil, "both", entries) {
		trans := entry.Value().(api.Transactor)
		for _, line := range trans.PrintLines() {
			fmt.Fprintln(outfd, line)
		}
		fmt.Fprintln(outfd)
	}
}

func (report *ReportPrint) Clone() api.Reporter {
	nreport := *report
	nreport.transdb = report.transdb.Clone()
	return &nreport
}

func (report *ReportPrint) StartJournal(fname string, included bool) {
	panic("not implemented")
}
