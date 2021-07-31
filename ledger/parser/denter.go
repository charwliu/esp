package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phf/go-queue/queue"
)

type DenterHelper struct {
	dentsBuffer  *queue.Queue
	indentations *queue.Queue
	nlToken      int
	identToken   int
	dedentToken  int
	reachedEof   bool

	BaseLexer *antlr.BaseLexer
}

func NewDenter(nlToken, identToken, dedentToken int) *DenterHelper {
	return &DenterHelper{
		nlToken:      nlToken,
		identToken:   identToken,
		dedentToken:  dedentToken,
		dentsBuffer:  queue.New(),
		indentations: queue.New(),
	}
}

func (d *DenterHelper) NextToken() antlr.Token {
	d.initIfFirstRun()
	var t antlr.Token
	if d.dentsBuffer.Len() == 0 {
		t = d.pullToken()
	} else {
		t = d.dentsBuffer.PopFront().(antlr.Token)
	}
	if d.reachedEof {
		return t
	}
	var r antlr.Token
	if t.GetTokenType() == d.nlToken {
		r = d.handleNewLineToken(t)
	} else if t.GetTokenType() == antlr.TokenEOF {
		r = d.handleEofToken(t)
	} else {
		r = t
	}
	return r
}

func (d *DenterHelper) initIfFirstRun() {
	if d.indentations.Len() == 0 {
		d.indentations.PushFront(0)

		var firstRealToken antlr.Token
		for {
			firstRealToken = d.pullToken()
			if firstRealToken.GetTokenType() != d.nlToken {
				break
			}
		}
		if firstRealToken.GetColumn() > 0 {
			d.indentations.PushFront(firstRealToken.GetColumn())
			d.dentsBuffer.PushBack(d.createToken(d.identToken, firstRealToken))
		}
		d.dentsBuffer.PushBack(firstRealToken)
	}
}

func (d *DenterHelper) createToken(tokenType int, copyFrom antlr.Token) antlr.Token {
	var tokenTypeStr string
	if tokenType == d.nlToken {
		tokenTypeStr = "newline"
	} else if tokenType == d.identToken {
		tokenTypeStr = "indent"
	} else if tokenType == d.dedentToken {
		tokenTypeStr = "dedent"
	}
	return NewInjectedToken(copyFrom, tokenType, tokenTypeStr)
}

type InjectedToken struct {
	*antlr.CommonToken
	tokenType string
}

func NewInjectedToken(oldToken antlr.Token, tokenType int, tokenTypeString string) antlr.Token {
	t := antlr.NewCommonToken(oldToken.GetSource(), tokenType, oldToken.GetChannel(),
		oldToken.GetStart(), oldToken.GetStart())
	return &InjectedToken{
		CommonToken: t,
		tokenType:   tokenTypeString,
	}
}

func (t *InjectedToken) GetText() string {
	if len(t.tokenType) > 0 {
		t.CommonToken.SetText(t.tokenType)
	}
	return t.CommonToken.GetText()
}

func (d *DenterHelper) handleEofToken(token antlr.Token) antlr.Token {
	var r antlr.Token
	// when we reach EOF, unwind all indentations. If there aren't any, insert a NL. This lets the grammar treat
	// un-indented expressions as just being NL-terminated, rather than NL|EOF.
	if d.indentations.Len() == 0 {
		r = d.createToken(d.nlToken, token)
		d.dentsBuffer.PushBack(token)
	} else {
		r = d.unwindTo(0, token)
		d.dentsBuffer.PushBack(token)
	}
	d.reachedEof = true
	return r
}

//  Returns a DEDENT token, and also queues up additional DEDENT as necessary.
func (d *DenterHelper) unwindTo(targetIdent int, copyFrom antlr.Token) antlr.Token {
	d.dentsBuffer.PushBack(d.createToken(d.nlToken, copyFrom))
	for {
		prevIndent := d.indentations.PopFront().(int)
		if prevIndent == targetIdent {
			break
		}
		if targetIdent > prevIndent {
			d.indentations.PushFront(prevIndent)
			d.dentsBuffer.PushBack(d.createToken(d.identToken, copyFrom))
			break
		}
		d.dentsBuffer.PushBack(d.createToken(d.dedentToken, copyFrom))
	}
	d.indentations.PushFront(targetIdent)
	return d.dentsBuffer.PopFront().(antlr.Token)
}

func (d *DenterHelper) handleNewLineToken(t antlr.Token) antlr.Token {
	nextNext := d.pullToken()
	for nextNext.GetTokenType() == d.nlToken {
		t = nextNext
		nextNext = d.pullToken()
	}
	if nextNext.GetTokenType() == antlr.TokenEOF {
		return d.handleEofToken(nextNext)
	}
	// nextNext is now a non-NL token; we'll queue it up after any possible dents
	nlText := t.GetText()
	indent := len(nlText) - 1 // every NL has one \n char, so shorten the length to account for it
	if indent > 0 && nlText[0] == '\r' {
		indent -= 1 // If the NL also has a \r char, we should account for that as well
	}
	prevIndent := d.indentations.Front().(int)
	var r antlr.Token
	if indent == prevIndent {
		r = t // just a newline
	} else if indent > prevIndent {
		r = d.createToken(d.nlToken, t)
		d.dentsBuffer.PushBack(d.createToken(d.identToken, r))
		d.indentations.PushFront(indent)
	} else {
		r = d.unwindTo(indent, t)
	}
	d.dentsBuffer.PushBack(nextNext)
	return r
}

func (d *DenterHelper) pullToken() antlr.Token {
	return d.BaseLexer.NextToken()
}
