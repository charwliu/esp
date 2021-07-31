package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	parsec "github.com/prataprc/goparsec"
	"go.vixal.xyz/esp/ledger/api"
	"go.vixal.xyz/esp/ledger/dblentry"
	"go.vixal.xyz/esp/ledger/reports"
)

func doFirstPass(reporter api.Reporter,
	db *dblentry.Datastore, journalFile string) (err error) {

	data, err := os.ReadFile(journalFile)
	if err != nil {
		log.Printf("%v\n", err)
		return err
	}
	if db.Hasjournal(data) {
		log.Printf("%q already processed !!\n", journalFile)
		return nil
	}
	db.Addjournal(journalFile, data)

	var lineno int
	var block []string
	var eof bool

	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic at %q:%v : %v\n", journalFile, lineno, r)
			log.Printf("\n%s", api.GetStacktrace(2, debug.Stack()))
			err = fmt.Errorf("%s", r)
		}
	}()

	log.Printf("firstpass %v\n", journalFile)
	var node parsec.ParsecNode

	lines, err := readLines(journalFile)
	if err != nil {
		return err
	}

	iterate := blockIterate(lines)
	lineno, block, eof, err = iterate()
	for len(block) > 0 {
		lineno -= len(block)
		if err != nil {
			log.Printf("iterate at %q:%v : %v\n", journalFile, lineno, err)
			return err
		}
		log.Printf("parsing block: %v\n", block[0])

		node, lineno, err = parseEntry(lineno, block, journalFile, db)
		if err != nil {
			log.Printf("parsec at %q:%v : %v\n", journalFile, lineno, err)
			return err
		} else if node != nil {
			if err := db.FirstPass(node); err != nil {
				log.Printf("%T at %q:%v : %v\n", node, journalFile, lineno-len(block)+1, err)
				return err
			}
		}

		tryInclude(reporter, db, node, journalFile)
		lineno, block, eof, err = iterate()
	}
	if err != nil {
		log.Printf("%q : %v\n", journalFile, err)
	} else if eof == false {
		log.Printf("%q : expected eof\n", journalFile)
	}
	return nil
}

func parseEntry(lineno int, block []string, journalFile string,
	db *dblentry.Datastore) (parsec.ParsecNode, int, error) {

	var err error
	var index int

	scanner := parsec.NewScanner([]byte(block[0]))
	yTrans := dblentry.NewTransaction(journalFile).Yledger(db)
	yPrice := dblentry.NewPrice().Yledger(db)
	yDirective := dblentry.NewDirective().Yledger(db)
	yComment := dblentry.NewComment().Yledger(db)
	y := parsec.OrdChoice(
		dblentry.Vector2scalar,
		yTrans, yPrice, yDirective, yComment,
	)
	node, _ := y(scanner)
	switch obj := node.(type) {
	case *dblentry.Transaction:
		obj.Addlines(block[0])
		payee := strings.Trim(obj.Payee(), " \t")
		if api.Options.Stitch && payee == reports.PayeeOpeningBalance {
			return nil, lineno, nil
		}
		index, err = obj.Yledgerblock(db, block[1:])
		lineno += 1 + index
		obj.SetLineno(lineno)
		obj.Addlines(block[1:]...)

	case *dblentry.Directive:
		index, err = obj.Yledgerblock(db, block[1:])
		lineno += 1 + index

	case error:
		err = obj

	case *dblentry.Comment, *dblentry.Price:
	}
	return node, lineno, err
}

func tryInclude(
	reporter api.Reporter, db *dblentry.Datastore, node parsec.ParsecNode,
	includedBy string) bool {

	if node == nil {
		return false
	}

	d, ok := node.(*dblentry.Directive)
	if ok && d.Type() == "include" {
		journalFile := d.IncludeFile()
		journalFile = strings.Trim(journalFile, "/")
		journalFile = filepath.Join(filepath.Dir(includedBy), journalFile)
		reporter.StartJournal(journalFile, true /*included*/)
		err := doFirstPass(reporter, db, journalFile)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func blockIterate(lines []string) func() (int, []string, bool, error) {
	row := 0

	parseBlock := func() []string {
		var blockLines []string
		for ; row < len(lines); row++ {
			line := lines[row]
			if (len(line) > 0) && (line[0] == ' ' || line[0] == '\t') {
				if line1 := strings.TrimLeft(line, " \t"); line1 == "" {
					break
				}
				blockLines = append(blockLines, line)
			} else {
				break
			}
		}
		return blockLines
	}

	return func() (int, []string, bool, error) {
		var blockLines []string
		for ; row < len(lines); row++ {
			line := lines[row]
			if len(line) == 0 {
				continue
			}
			if line[0] == ' ' || line[0] == '\t' {
				line1 := strings.TrimLeft(line, " \t")
				if line1 == "" { // empty line
					continue
				} else {
					return row + 1, nil, false, fmt.Errorf("must be at the beginning: row:%v column: 0", row+1)
				}

			} else { // begin block
				row++
				blockLines = append(blockLines, line)
				blockLines = append(blockLines, parseBlock()...)
				return row + 1, blockLines, row >= len(lines), nil
			}
		}
		return row + 1, blockLines, true, nil
	}
}
