package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	parsec "github.com/prataprc/goparsec"

	"go.vixal.xyz/esp/ledger/api"
	"go.vixal.xyz/esp/ledger/dblentry"
)

func argparse() ([]string, error) {
	var journals, outfile, finYear, beginDt, endDt string

	f := flag.NewFlagSet("ledger", flag.ExitOnError)
	f.Usage = func() {
		fmsg := "Usage of command: %v [OPTIONS] COMMAND [ARGS]\n"
		fmt.Printf(fmsg, os.Args[0])
		f.PrintDefaults()
	}

	f.StringVar(&api.Options.Dbname, "db", "devjournal",
		"Provide datastore name")
	f.StringVar(&journals, "f", "example/first.ldg",
		"Comma separated list of input files.")
	f.StringVar(&outfile, "o", "",
		"outfile to report")
	f.StringVar(&api.Options.CurrentDt, "current", "",
		"Display only transactions on or before the current date.")
	f.StringVar(&beginDt, "begin", "",
		"Display only transactions on or before the current date.")
	f.StringVar(&endDt, "end", "",
		"Display only transactions on or before the current date.")
	f.StringVar(&finYear, "fy", "",
		"financial year.")
	f.StringVar(&api.Options.Period, "period", "",
		"Limit the processing to transactions in PERIOD_EXPRESSION.")
	f.BoolVar(&api.Options.NoSubtotal, "nosubtotal", false,
		"Don't accumulate postings on sub-leger to parent ledger.")
	f.BoolVar(&api.Options.Subtotal, "subtotal", false,
		"all transactions to be collapsed into a single, transaction")
	f.BoolVar(&api.Options.Cleared, "cleared", true,
		"Display only cleared postings.")
	f.BoolVar(&api.Options.Uncleared, "uncleared", true,
		"Display only uncleared postings.")
	f.BoolVar(&api.Options.Pending, "pending", true,
		"Display only pending postings.")
	f.BoolVar(&api.Options.DcFormat, "dc", false,
		"Display only real postings.")
	f.BoolVar(&api.Options.Strict, "strict", false,
		"Accounts, tags or commodities not previously declared "+
			"will cause warnings.")
	f.BoolVar(&api.Options.Pedantic, "pedantic", false,
		"Accounts, tags or commodities not previously declared "+
			"will cause httperror.")
	f.BoolVar(&api.Options.CheckPayee, "checkpayee", false,
		"Payee not previously declared will cause error.")
	f.BoolVar(&api.Options.Stitch, "stitch", false,
		"Skip payees with `Opening Balance`")
	f.BoolVar(&api.Options.NoPl, "nopl", false,
		"skip income and expense accounts")
	f.BoolVar(&api.Options.OnlyPl, "onlypl", false,
		"skip accounts other than income and expense")
	f.BoolVar(&api.Options.Detailed, "detailed", false,
		"for register, passbook commands list details")
	f.BoolVar(&api.Options.ByPayee, "bypayee", false,
		"Group postings by common payee names")
	f.BoolVar(&api.Options.Daily, "daily", false,
		"Group postings by day")
	f.BoolVar(&api.Options.Weekly, "weekly", false,
		"Group postings by week")
	f.BoolVar(&api.Options.Monthly, "monthly", false,
		"Group postings by month")
	f.BoolVar(&api.Options.Quarterly, "quarterly", false,
		"Group postings by quarter")
	f.BoolVar(&api.Options.Yearly, "yearly", false,
		"Group postings by yearly")
	f.BoolVar(&api.Options.Dow, "dow", false,
		"Group postings by day of the week")
	f.BoolVar(&api.Options.Verbose, "v", false,
		"verbose reporting / listing")

	f.StringVar(&api.Options.Loglevel, "log", "info",
		"Console log level")
	f.Parse(os.Args[1:])

	//logsetts := map[string]interface{}{
	//	"log.level":      api.Options.Loglevel,
	//	"log.file":       "",
	//	"log.timeformat": "",
	//	"log.prefix":     "%v:",
	//	"log.colorfatal": "red",
	//	"log.colorerror": "hired",
	//	"log.colorwarn":  "yellow",
	//}
	//log.SetLogger(nil, logsetts)

	api.Options.Journals = gatherJournals(journals)
	api.Options.OutFd = argOutFd(outfile)

	endYear := argFinYear(finYear)
	if endYear > 0 {
		till := time.Date(endYear, 4, 1, 0, 0, 0, 0, time.Local)
		ok := api.ValidateDate(till, endYear, 4, 1, 0, 0, 0)
		if ok == false {
			err := fmt.Errorf("invalid finYear %v", endYear)
			log.Printf("%v\n", err)
			return nil, err
		}
		from := time.Date(endYear-1, 4, 1, 0, 0, 0, 0, time.Local)
		ok = api.ValidateDate(from, endYear-1, 4, 1, 0, 0, 0)
		if ok == false {
			err := fmt.Errorf("invalid begin year for finYear %v", endYear)
			log.Printf("%v\n", err)
			return nil, err
		}
		// BeginDt is inclusive, but not TillDt
		api.Options.BeginDt, api.Options.EndDt = &from, &till
	}

	if beginDt != "" {
		scanner := parsec.NewScanner([]byte(beginDt))
		node, _ := dblentry.Ydate(time.Now().Year())(scanner)
		tm, ok := node.(time.Time)
		if ok == false {
			err := fmt.Errorf("invalid date %q: %v\n", beginDt, node)
			log.Printf("%v\n", err)
			return nil, err
		}
		api.Options.BeginDt = &tm
	}

	if endDt != "" {
		scanner := parsec.NewScanner([]byte(endDt))
		node, _ := dblentry.Ydate(time.Now().Year())(scanner)
		tm, ok := node.(time.Time)
		if ok == false {
			err := fmt.Errorf("invalid date %q: %v\n", endDt, node)
			log.Printf("%v\n", err)
			return nil, err
		}
		api.Options.EndDt = &tm
	}
	return f.Args(), nil
}

func gatherJournals(journals string) (files []string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.Getwd(): %v\n", err)
		os.Exit(1)
	}

	if journals == "list" {
		files, err = listJournals(cwd)
		if err != nil {
			os.Exit(1)
		}

	} else if journals == "find" {
		files, err = findJournals(cwd)
		if err != nil {
			os.Exit(1)
		}

	} else {
		files = api.ParseCsv(journals)
	}
	files = append(files, coveringJournals(cwd)...)
	return files
}

func argOutFd(outfile string) *os.File {
	outFd := os.Stdout
	if outfile != "" {
		fd, err := os.Create(outfile)
		if err != nil {
			log.Printf("%v\n", err)
			os.Exit(1)
		}
		outFd = fd
	}
	return outFd
}

func argFinYear(finYear string) int {
	if finYear == "" {
		return 0
	}

	fy, err := strconv.Atoi(finYear)
	if err != nil {
		log.Printf("arg `-fy` %v\n", err)
		os.Exit(1)
	}
	return fy
}
