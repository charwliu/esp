package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ericlagergren/decimal"
	"go.vixal.xyz/esp/ledger"
	"go.vixal.xyz/esp/ledger/parser"
	"google.golang.org/protobuf/proto"
)

type Visitor struct {
	parser.BeanCountParser
	builder *parser.Builder
}

func NewVisitor(builder *parser.Builder) *Visitor {
	return &Visitor{
		builder: builder,
	}
}

func (v *Visitor) VisitLedger(ctx *parser.LedgerContext) interface{} {

	if c := ctx.Declarations(); c != nil {
		v.Visit(c)
	}
	v.builder.Finalize(nil)
	return v.builder
}

func (v *Visitor) VisitDeclarations(ctx *parser.DeclarationsContext) interface{} {
	var dir interface{}
	if c := ctx.NEWLINE(); c != nil {
		fmt.Printf("Declarations(%d): \n", c.GetSymbol().GetLine())
		v.Visit(ctx.Declarations())
		return dir
	} else if c := ctx.Directive(); c != nil {
		v.Visit(ctx.Declarations())
		fmt.Printf("Declarations(%d): \n", c.GetStart().GetLine())
		d := v.Visit(c).(*ledger.Directive)
		v.builder.AppendDirective(d)
		fmt.Printf("directive: %s\n", d.String())
		return dir
	} else if c := ctx.Pragma(); c != nil {
		v.Visit(ctx.Declarations())
		fmt.Printf("Declarations(%d): \n", c.GetStart().GetLine())
		dir = v.Visit(c)
		return dir
	}
	return dir
}

func (v *Visitor) VisitTransactionDir(ctx *parser.TransactionDirContext) interface{} {
	dir := v.Visit(ctx.Transaction()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitPriceDir(ctx *parser.PriceDirContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitBalanceDir(ctx *parser.BalanceDirContext) interface{} {
	dir := v.Visit(ctx.Balance()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitOpenDir(ctx *parser.OpenDirContext) interface{} {
	dir := v.Visit(ctx.Open()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitCloseDir(ctx *parser.CloseDirContext) interface{} {
	dir := v.Visit(ctx.CloseDirective()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitCommodityDir(ctx *parser.CommodityDirContext) interface{} {
	dir := v.Visit(ctx.Commodity()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitPadDir(ctx *parser.PadDirContext) interface{} {
	dir := v.Visit(ctx.Pad()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitDocumentDir(ctx *parser.DocumentDirContext) interface{} {
	dir := v.Visit(ctx.Document()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitNoteDir(ctx *parser.NoteDirContext) interface{} {
	dir := v.Visit(ctx.Note()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitEventDir(ctx *parser.EventDirContext) interface{} {
	dir := v.Visit(ctx.Event()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitQueryDir(ctx *parser.QueryDirContext) interface{} {
	dir := v.Visit(ctx.Query()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitCustomDir(ctx *parser.CustomDirContext) interface{} {
	dir := v.Visit(ctx.Custom()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitPrice(ctx *parser.PriceContext) interface{} {
	dir := v.Visit(ctx.PRICE()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitOpen(ctx *parser.OpenContext) interface{} {
	dir := v.Visit(ctx.OPEN()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitCloseDirective(ctx *parser.CloseDirectiveContext) interface{} {
	dir := v.Visit(ctx.CLOSE()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitCommodity(ctx *parser.CommodityContext) interface{} {
	dir := v.Visit(ctx.COMMODITY()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitPad(ctx *parser.PadContext) interface{} {
	dir := v.Visit(ctx.PAD()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitDocument(ctx *parser.DocumentContext) interface{} {
	dir := v.Visit(ctx.DOCUMENT()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitNote(ctx *parser.NoteContext) interface{} {
	dir := v.Visit(ctx.NOTE()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitEvent(ctx *parser.EventContext) interface{} {
	dir := v.Visit(ctx.EVENT()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitQuery(ctx *parser.QueryContext) interface{} {
	dir := v.Visit(ctx.QUERY()).(*ledger.Directive)
	return dir
}

func (v *Visitor) VisitCustom(ctx *parser.CustomContext) interface{} {
	customType := ctx.GetCustomType()
	metaValueList := v.Visit(ctx.MetaValueList())
	metadata := v.Visit(ctx.IndentedMetadata())
	return append([]interface{}{}, customType, metaValueList, metadata)
}

func (v *Visitor) VisitIndentedMetadata(ctx *parser.IndentedMetadataContext) interface{} {
	meta := ctx.Metadata()
	if meta != nil {
		return v.Visit(ctx.Metadata())
	}
	return nil
}

func (v *Visitor) VisitMetadata(ctx *parser.MetadataContext) interface{} {
	meta := ledger.Meta{}
	for _, line := range ctx.AllMetadataLine() {
		kv := v.Visit(line).(*ledger.Meta_KV)
		meta.Kv = append(meta.Kv, kv)
	}
	return &meta
}

func (v *Visitor) VisitMetadataLine(ctx *parser.MetadataLineContext) interface{} {
	if ctx.KEY() != nil {
		key := ctx.KEY().GetSymbol().GetText()
		kv := ledger.Meta_KV{
			Key: &key,
		}
		kv.Value = v.Visit(ctx.MetaValue()).(*ledger.MetaValue)
		return &kv
	}
	if node := ctx.LINK(); node != nil {
		value := &ledger.MetaValue{
			Value: &ledger.MetaValue_Link{
				Link: node.GetSymbol().GetText(),
			},
		}
		kv := ledger.Meta_KV{
			Value: value,
		}
		return &kv
	}
	if node := ctx.TAG(); node != nil {
		value := &ledger.MetaValue{
			Value: &ledger.MetaValue_Tag{
				Tag: node.GetSymbol().GetText(),
			},
		}
		kv := ledger.Meta_KV{
			Value: value,
		}
		return &kv
	}
	return nil
}

func (v *Visitor) VisitMetaValue(ctx *parser.MetaValueContext) interface{} {
	if ctx.Account() != nil {
		mv := new(ledger.MetaValue)
		mv.Value = &ledger.MetaValue_Account{
			Account: v.Visit(ctx.Account()).(string),
		}
		return mv
	}
	if ctx.Amount() != nil {
		mv := new(ledger.MetaValue)
		mv.Value = &ledger.MetaValue_Amount{
			Amount: v.Visit(ctx.Amount()).(*ledger.Amount),
		}
		return mv
	}
	if ctx.NumberExpr() != nil {
		mv := new(ledger.MetaValue)
		num := new(ledger.Number)
		dec := v.Visit(ctx.NumberExpr()).(*decimal.Big)
		v.builder.DecimalProto(dec, num)
		mv.Value = &ledger.MetaValue_Number{
			Number: num,
		}
		return mv
	}
	if node := ctx.LINK(); node != nil {
		mv := new(ledger.MetaValue)
		mv.Value = &ledger.MetaValue_Link{
			Link: node.GetSymbol().GetText(),
		}
		return mv
	}
	if node := ctx.TAG(); node != nil {
		mv := new(ledger.MetaValue)
		mv.Value = &ledger.MetaValue_Tag{
			Tag: node.GetSymbol().GetText(),
		}
		return mv
	}
	if node := ctx.STRING(); node != nil {
		mv := new(ledger.MetaValue)
		mv.Value = &ledger.MetaValue_Text{
			Text: node.GetSymbol().GetText(),
		}
		return mv
	}
	if node := ctx.DATETIME(); node != nil {
		mv := new(ledger.MetaValue)
		dt, err := time.Parse("2006-01-02", node.GetSymbol().GetText())
		if err != nil {
			return nil
		}
		mv.Value = &ledger.MetaValue_Date{
			Date: &ledger.Date{
				Year:  int32(dt.Year()),
				Month: int32(dt.Month()),
				Day:   int32(dt.Day()),
			},
		}
		return mv
	}
	if node := ctx.NONE(); node != nil {
		mv := new(ledger.MetaValue)
		return mv
	}
	if node := ctx.BOOL(); node != nil {
		mv := new(ledger.MetaValue)
		b, _ := strconv.ParseBool(node.GetSymbol().GetText())
		mv.Value = &ledger.MetaValue_Boolean{
			Boolean: b,
		}
		return mv
	}
	if node := ctx.CURRENCY(); node != nil {
		mv := new(ledger.MetaValue)
		mv.Value = &ledger.MetaValue_Currency{
			Currency: node.GetSymbol().GetText(),
		}
		return mv
	}
	return nil
}

func (v *Visitor) VisitTagsOrLinks(ctx *parser.TagsOrLinksContext) interface{} {
	tl := &parser.TagsLinks{}
	for _, node := range ctx.AllTAG() {
		tl.Tags = append(tl.Tags, node.GetSymbol().GetText())
	}
	for _, node := range ctx.AllLINK() {
		tl.Links = append(tl.Links, node.GetSymbol().GetText())
	}
	return tl
}

func (v *Visitor) VisitPayeeNarration(ctx *parser.PayeeNarrationContext) interface{} {
	pn := parser.PayeeNarration{}
	if c := ctx.GetPayee(); c != nil {
		pn.Payee = c.GetText()
	}
	if c := ctx.GetNarration(); c != nil {
		pn.Narration = c.GetText()
	}
	return &pn
}

func (v *Visitor) VisitFilename(ctx *parser.FilenameContext) interface{} {
	return ctx.STRING().GetSymbol().GetText()
}

func (v *Visitor) VisitCurrencyList(ctx *parser.CurrencyListContext) interface{} {
	var cl []string
	for _, node := range ctx.AllCURRENCY() {
		cl = append(cl, node.GetSymbol().GetText())
	}
	return cl
}

func (v *Visitor) VisitUniaryOp(ctx *parser.UniaryOpContext) interface{} {
	if c := ctx.NUMBER(); c != nil {
		dec := new(decimal.Big)
		dec.SetString(c.GetSymbol().GetText())
		return dec
	}
	if c := ctx.NumberExpr(); c != nil {
		return v.Visit(c)
	}
	return &decimal.Big{}
}

func (v *Visitor) VisitBinaryOp(ctx *parser.BinaryOpContext) interface{} {
	right := v.Visit(ctx.NumberExpr(0)).(*decimal.Big)
	left := v.Visit(ctx.NumberExpr(1)).(*decimal.Big)
	dec := decimal.Big{}
	if ctx.ASTERISK() != nil {
		v.builder.Ctx.Mul(&dec, right, left)
	}
	if ctx.PLUS() != nil {
		v.builder.Ctx.Add(&dec, right, left)
	}
	return &dec
}

func (v *Visitor) VisitTxn(ctx *parser.TxnContext) interface{} {
	if c := ctx.ASTERISK(); c != nil {
		return []byte("*")
	}
	if c := ctx.TXN(); c != nil {
		return []byte("*")
	}
	if c := ctx.HASH(); c != nil {
		return []byte("#")
	}
	if c := ctx.FLAG(); c != nil {
		return []byte(c.GetSymbol().GetText())
	}
	return []byte{0}
}

func (v *Visitor) VisitOptflag(ctx *parser.OptflagContext) interface{} {
	if c := ctx.ASTERISK(); c != nil {
		return []byte{'*'}
	}
	if c := ctx.FLAG(); c != nil {
		return []byte(c.GetSymbol().GetText())
	}

	if c := ctx.HASH(); c != nil {
		return []byte{'#'}
	}
	return []byte{0}
}

func (v *Visitor) VisitBooking(ctx *parser.BookingContext) interface{} {
	var result []antlr.Token
	if node := ctx.STRING(); node != nil {
		result = append(result, node.GetSymbol())
	}
	return result
}

func (v *Visitor) VisitPostingList(ctx *parser.PostingListContext) interface{} {
	var pl []*ledger.Posting
	if c := ctx.PostingAndMetadata(); c != nil {
		pl = append(pl, v.Visit(c).(*ledger.Posting))
	}
	if c := ctx.PostingList(); c != nil {
		pl = append(pl, v.Visit(c).([]*ledger.Posting)...)
	}
	return pl
}

func (v *Visitor) VisitPostingAndMetadata(ctx *parser.PostingAndMetadataContext) interface{} {
	posting := v.Visit(ctx.Posting()).(*ledger.Posting)
	if c := ctx.IndentedMetadata(); c != nil {
		meta := v.Visit(c)
		if meta != nil {
			proto.Merge(posting.Meta, meta.(*ledger.Meta))
		}
	}
	lineno := int32(ctx.GetStart().GetLine())
	linenoEnd := int32(ctx.GetStop().GetLine())
	posting.Location = &ledger.Location{
		Lineno:    &lineno,
		LinenoEnd: &linenoEnd,
	}
	return posting
}

func (v *Visitor) VisitPosting(ctx *parser.PostingContext) interface{} {
	return nil
}

func (v *Visitor) VisitPostingPrice(ctx *parser.PostingPriceContext) interface{} {
	posting := new(ledger.Posting)
	posting.Spec = new(ledger.Spec)
	if cs := v.Visit(ctx.CostSpec()); cs != nil {
		posting.Spec.Cost = cs.(*ledger.CostSpec)
	}
	optflag := v.Visit(ctx.Optflag()).([]byte)
	acct := v.Visit(ctx.Account()).(string)
	pa := v.Visit(ctx.PartialAmount()).(*ledger.Amount)
	if ctx.AT() != nil {
		price := v.Visit(ctx.PriceAnnotation()).(*ledger.PriceSpec)
		posting.Spec.Price = &ledger.PriceSpec{}
		proto.Merge(posting.Spec.Price, price)
	}
	if ctx.ATAT() != nil {
		price := v.Visit(ctx.PriceAnnotation()).(*ledger.PriceSpec)
		posting.Spec.Price = &ledger.PriceSpec{}
		proto.Merge(posting.Spec.Price, price)
		*posting.Spec.Price.IsTotal = true
	}

	loc := &ledger.Location{}
	v.builder.PreparePosting(posting, pa, optflag, acct, true, loc)
	return posting
}

func (v *Visitor) VisitPostingAmount(ctx *parser.PostingAmountContext) interface{} {
	posting := new(ledger.Posting)

	if cs := v.Visit(ctx.CostSpec()); cs != nil {
		posting.Spec = &ledger.Spec{
			Cost: cs.(*ledger.CostSpec),
		}
	}
	optflag := v.Visit(ctx.Optflag()).([]byte)
	acct := v.Visit(ctx.Account()).(string)
	pa := v.Visit(ctx.PartialAmount()).(*ledger.Amount)

	loc := &ledger.Location{}
	v.builder.PreparePosting(posting, pa, optflag, acct, true, loc)
	return posting
}

func (v *Visitor) VisitPostingAccount(ctx *parser.PostingAccountContext) interface{} {
	posting := new(ledger.Posting)
	optflag := v.Visit(ctx.Optflag()).([]byte)
	acct := v.Visit(ctx.Account()).(string)
	loc := &ledger.Location{}
	v.builder.PreparePosting(posting, nil, optflag, acct, true, loc)
	return posting
}

func (v *Visitor) VisitPriceAnnotation(ctx *parser.PriceAnnotationContext) interface{} {
	ps := new(ledger.PriceSpec)
	pa := v.Visit(ctx.PartialAmount()).(*ledger.Amount)
	ps.Number = new(ledger.Number)
	proto.Merge(ps.Number, pa.Number)
	ps.Currency = pa.Currency
	if ps.Number != nil {
		dec := ledger.ProtoToDecimal(pa.Number)
		if dec.Sign() == -1 {
			//v.builder.AddError("")
			z := &decimal.Big{}
			v.builder.DecimalProto(z.Neg(dec), ps.Number)
		}
	}
	return ps
}

func (v *Visitor) VisitAmount(ctx *parser.AmountContext) interface{} {
	amount := new(ledger.Amount)
	number := v.Visit(ctx.NumberExpr()).(*decimal.Big)
	amount.Number = &ledger.Number{}
	v.builder.DecimalProto(number, amount.Number)
	currency := ctx.CURRENCY().GetText()
	amount.Currency = &currency
	return amount
}

func (v *Visitor) VisitAmountTolerance(ctx *parser.AmountToleranceContext) interface{} {
	at := &AmountTolerance{}
	if c := ctx.Amount(); c != nil {
		amount := v.Visit(c).(*ledger.Amount)
		at.Amount = amount
	}
	if ctx.TILDE() != nil {
		n1 := v.Visit(ctx.NumberExpr(0)).(*decimal.Big)
		n2 := v.Visit(ctx.NumberExpr(1)).(*decimal.Big)
		currency := ctx.CURRENCY().GetText()

		amount := new(ledger.Amount)
		amount.Number = &ledger.Number{}
		v.builder.DecimalProto(n1, amount.Number)
		amount.Currency = &currency
		at.Amount = amount
		at.Tolerance = n2
	}
	return at
}

type AmountTolerance struct {
	Amount    *ledger.Amount
	Tolerance *decimal.Big
}

func (v *Visitor) VisitMaybeNumber(ctx *parser.MaybeNumberContext) interface{} {
	return v.Visit(ctx.NumberExpr())
}

func (v *Visitor) VisitMaybeCurrency(ctx *parser.MaybeCurrencyContext) interface{} {
	return ctx.CURRENCY().GetText()
}

func (v *Visitor) VisitPartialAmount(ctx *parser.PartialAmountContext) interface{} {
	amt := new(ledger.Amount)
	if c := ctx.MaybeNumber(); c != nil {
		num := v.Visit(c).(*decimal.Big)
		amt.Number = new(ledger.Number)
		v.builder.DecimalProto(num, amt.Number)
	}
	if c := ctx.MaybeCurrency(); c != nil {
		cur := v.Visit(c).(string)
		amt.Currency = &cur
	}
	return amt
}

func (v *Visitor) VisitCompoundAmount(ctx *parser.CompoundAmountContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitCostSpec(ctx *parser.CostSpecContext) interface{} {
	if ctx.LCURL() != nil {
		costs := v.Visit(ctx.CostCompList())
		return costs
	}
	if ctx.LCURLCURL() != nil {
		costs := v.Visit(ctx.CostCompList())
		return costs
	}
	return nil
}

func (v *Visitor) VisitCostCompList(ctx *parser.CostCompListContext) interface{} {
	cs := new(ledger.CostSpec)
	var s *ledger.CostSpec
	if c := ctx.CostCompList(); c != nil {
		s = v.Visit(c).(*ledger.CostSpec)
	}
	if c := ctx.CostComp(); c != nil {
		s = v.Visit(c).(*ledger.CostSpec)
	}
	err := v.builder.MergeCost(s, cs)
	if err != nil {
		loc := &ledger.Location{
			Lineno:    new(int32),
			LinenoEnd: new(int32),
		}
		*loc.Lineno = int32(ctx.GetStart().GetLine())
		*loc.LinenoEnd = int32(ctx.GetStop().GetLine())
		v.builder.AddError(err.Error(), loc)
	}
	return cs
}

func (v *Visitor) VisitCostComp(ctx *parser.CostCompContext) interface{} {
	cs := new(ledger.CostSpec)
	if c := ctx.DATETIME(); c != nil {
		cs.Date = new(ledger.Date)
		dt, _ := time.Parse("2006-01-02", c.GetSymbol().GetText())
		cs.Date.Day = int32(dt.Day())
		cs.Date.Month = int32(dt.Month())
		cs.Date.Year = int32(dt.Year())
		return cs
	}
	if c := ctx.STRING(); c != nil {
		str := c.GetSymbol().GetText()
		cs.Label = &str
	}
	if c := ctx.ASTERISK(); c != nil {
		cs.MergeCost = new(bool)
		*cs.MergeCost = true
	}
	if c := ctx.CompoundAmount(); c != nil {
		return v.Visit(c)
	}
	return cs
}

func (v *Visitor) VisitAccount(ctx *parser.AccountContext) interface{} {
	return v.builder.Account(ctx.ACCOUNT().GetSymbol().GetText())
}

func (v *Visitor) VisitTransaction(ctx *parser.TransactionContext) interface{} {
	var dir *ledger.Directive
	if c := ctx.TransactionLine(); c != nil {
		dir = v.Visit(c).(*ledger.Directive)
	}
	if c := ctx.Metadata(); c != nil {
		metadata := v.Visit(c)
		if metadata != nil {
			dir.Meta = &ledger.Meta{}
			proto.Merge(dir.Meta, metadata.(*ledger.Meta))
		}
	}
	if c := ctx.PostingList(); c != nil {
		pl := v.Visit(c).([]*ledger.Posting)
		if pl != nil {
			txn := dir.GetTransaction()
			for _, posting := range pl {
				txn.Postings = append(txn.Postings, posting)
			}
		}
	}
	return dir
}

func (v *Visitor) VisitTransactionLine(ctx *parser.TransactionLineContext) interface{} {
	date, _ := time.Parse("2006-01-02", ctx.DATETIME().GetText())
	dt := ledger.Date{Year: int32(date.Year()), Month: int32(date.Month()), Day: int32(date.Day())}
	var tagsLink *parser.TagsLinks

	if c := v.Visit(ctx.TagsOrLinks()); c != nil {
		tagsLink = c.(*parser.TagsLinks)
	}
	dir := v.builder.MakeDirective(&dt, nil, &tagsLink)
	txn := &ledger.Transaction{}
	dir.Body = &ledger.Directive_Transaction{Transaction: txn}
	if c := ctx.Txn(); c != nil {
		flag := v.Visit(c).([]byte)
		if flag[0] != 0 {
			txn.Flag = flag
		}
	}
	if c := ctx.PayeeNarration(); c != nil {
		pn := v.Visit(c).(*parser.PayeeNarration)
		if pn.Payee != "" {
			txn.Payee = &pn.Payee
		}
		if pn.Narration != "" {
			txn.Narration = &pn.Narration
		}
	}
	return dir
}

func (v *Visitor) VisitBalance(ctx *parser.BalanceContext) interface{} {
	var dir *ledger.Directive
	date, _ := time.Parse("2006-01-02", ctx.DATETIME().GetText())
	dt := ledger.Date{Year: int32(date.Year()), Month: int32(date.Month()), Day: int32(date.Day())}
	var tagsLink *parser.TagsLinks

	if c := v.Visit(ctx.TagsOrLinks()); c != nil {
		tagsLink = c.(*parser.TagsLinks)
	}
	var metadata *ledger.Meta
	if c := v.Visit(ctx.IndentedMetadata()); c != nil {
		metadata = c.(*ledger.Meta)
	}
	dir = v.builder.MakeDirective(&dt, &metadata, &tagsLink)
	bal := &ledger.Balance{}
	dir.Body = &ledger.Directive_Balance{Balance: bal}
	acc := v.Visit(ctx.Account()).(string)
	bal.Account = &acc
	if c := ctx.AmountTolerance(); c != nil {
		at := v.Visit(c).(*AmountTolerance)
		bal.Amount = &ledger.Amount{}
		proto.Merge(bal.Amount, at.Amount)
		if at.Tolerance != nil {
			bal.Tolerance = &ledger.Number{}
			v.builder.DecimalProto(at.Tolerance, bal.Tolerance)
		}
	}

	return dir
}

func (v *Visitor) VisitEol(ctx *parser.EolContext) interface{} {
	return nil
}

func (v *Visitor) VisitMetaValueList(ctx *parser.MetaValueListContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	return nil
}

func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitPragma(ctx *parser.PragmaContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitPushtag(ctx *parser.PushtagContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitPoptag(ctx *parser.PoptagContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitPushmeta(ctx *parser.PushmetaContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitPopmeta(ctx *parser.PopmetaContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitOption(ctx *parser.OptionContext) interface{} {
	panic("implement me")
}

func (v *Visitor) VisitInclude(ctx *parser.IncludeContext) interface{} {
	panic("implement me")
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	switch ctx := tree.(type) {
	case *parser.LedgerContext:
		return v.VisitLedger(ctx)
	case *parser.DeclarationsContext:
		return v.VisitDeclarations(ctx)
	case *parser.AccountContext:
		return v.VisitAccount(ctx)
	case *parser.TransactionDirContext:
		return v.VisitTransactionDir(ctx)
	case *parser.TransactionContext:
		return v.VisitTransaction(ctx)
	case *parser.TransactionLineContext:
		return v.VisitTransactionLine(ctx)
	case *parser.BalanceContext:
		return v.VisitBalance(ctx)
	case *parser.PostingListContext:
		return v.VisitPostingList(ctx)
	case *parser.PostingContext:
		return v.VisitPosting(ctx)
	case *parser.PostingAccountContext:
		return v.VisitPostingAccount(ctx)
	case *parser.PostingAmountContext:
		return v.VisitPostingAmount(ctx)
	case *parser.PostingPriceContext:
		return v.VisitPostingPrice(ctx)
	case *parser.OptflagContext:
		return v.VisitOptflag(ctx)
	case *parser.PayeeNarrationContext:
		return v.VisitPayeeNarration(ctx)
	case *parser.PriceContext:
		return v.VisitPrice(ctx)
	case *parser.PriceAnnotationContext:
		return v.VisitPriceAnnotation(ctx)
	case *parser.PartialAmountContext:
		return v.VisitPartialAmount(ctx)
	case *parser.MaybeCurrencyContext:
		return v.VisitMaybeCurrency(ctx)
	case *parser.MaybeNumberContext:
		return v.VisitMaybeNumber(ctx)
	case *parser.BinaryOpContext:
		return v.VisitBinaryOp(ctx)
	case *parser.UniaryOpContext:
		return v.VisitUniaryOp(ctx)
	case *parser.CostSpecContext:
		return v.VisitCostSpec(ctx)
	case *parser.CompoundAmountContext:
		return v.VisitCompoundAmount(ctx)
	case *parser.PostingAndMetadataContext:
		return v.VisitPostingAndMetadata(ctx)
	case *parser.IndentedMetadataContext:
		return v.VisitIndentedMetadata(ctx)
	case *parser.TxnContext:
		return v.VisitTxn(ctx)
	case *parser.BalanceDirContext:
		return v.VisitBalanceDir(ctx)
	case *parser.AmountToleranceContext:
		return v.VisitAmountTolerance(ctx)
	case *parser.AmountContext:
		return v.VisitAmount(ctx)
	case *parser.EolContext:
		return v.VisitEol(ctx)
	case *parser.MetaValueListContext:
		return v.VisitMetaValueList(ctx)
	case antlr.RuleNode:
		return v.VisitChildren(ctx)
	case antlr.TerminalNode:
		return v.VisitTerminal(ctx)
	case *parser.PushtagContext:
		return v.VisitPushtag(ctx)
	case *parser.PoptagContext:
		return v.VisitPoptag(ctx)
	case *parser.PushmetaContext:
		return v.VisitPushmeta(ctx)
	case *parser.PopmetaContext:
		return v.VisitPopmeta(ctx)
	case *parser.IncludeContext:
		return v.VisitInclude(ctx)
	case *parser.OptionContext:
		return v.VisitOption(ctx)
	case *parser.PriceDirContext:
		return v.VisitPriceDir(ctx)
	case *parser.OpenDirContext:
		return v.VisitOpenDir(ctx)
	case *parser.CloseDirContext:
		return v.VisitCloseDir(ctx)
	case *parser.CloseDirectiveContext:
		return v.VisitCloseDirective(ctx)
	case *parser.PadDirContext:
		return v.VisitPadDir(ctx)
	case *parser.PadContext:
		return v.VisitPad(ctx)
	case *parser.PragmaContext:
		return v.VisitPragma(ctx)
	case *parser.DocumentContext:
		return v.VisitDocument(ctx)
	case *parser.DocumentDirContext:
		return v.VisitDocumentDir(ctx)
	case *parser.CommodityDirContext:
		return v.VisitCommodityDir(ctx)
	case *parser.CommodityContext:
		return v.VisitCommodity(ctx)
	case *parser.NoteDirContext:
		return v.VisitNoteDir(ctx)
	case *parser.NoteContext:
		return v.VisitNote(ctx)
	default:
		return nil
	}
}

var (
	_ parser.BeanCountParserVisitor = (*Visitor)(nil)
)
