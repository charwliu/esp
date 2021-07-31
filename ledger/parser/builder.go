package parser

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/ericlagergren/decimal"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"go.vixal.xyz/esp/ledger"
	"go.vixal.xyz/esp/ledger/options"
)

type MetaValues []*ledger.MetaValue

type MetaMap map[string]MetaValues

type Builder struct {
	ActiveTags map[string]struct{}
	ActiveMeta MetaMap
	Accounts   map[string]struct{}

	Ctx         *decimal.Context
	directives_ []*ledger.Directive
	options_    *options.Options
	info_       *options.ProcessingInfo
	errors_     []*ledger.Error
}

func NewBuilder() *Builder {
	return &Builder{
		ActiveTags:  make(map[string]struct{}),
		Accounts:    make(map[string]struct{}),
		ActiveMeta:  make(map[string]MetaValues),
		Ctx:         &decimal.Context{},
		directives_: make([]*ledger.Directive, 0),
		errors_:     make([]*ledger.Error, 0),
		options_:    &options.Options{},
		info_:       &options.ProcessingInfo{},
	}
}

func (b *Builder) AddOption(key string, value string, loc *ledger.Location) {
	descriptor := b.options_.ProtoReflect().Descriptor()
	field := descriptor.Fields().ByName(protoreflect.Name(key))
	if field == nil {
		b.AddError(fmt.Sprintf("invalid options: '%s'", key), loc)
		return
	}

}

func (b *Builder) AddInclude(filename string) {
	b.info_.Include = append(b.info_.Include, filename)
}

func (b *Builder) Account(account string) string {
	_, found := b.Accounts[account]
	if found {
		return account
	}
	b.Accounts[account] = struct{}{}
	return account
}

func (b *Builder) SetTagsAndLinks(tagsLinks *TagsLinks, message *ledger.Directive) {
	if len(b.ActiveTags) > 0 {
		for tag, _ := range b.ActiveTags {
			message.Tags = append(message.Tags, tag)
		}
	}

	if tagsLinks != nil {
		// Add the new tags
		if len(tagsLinks.Tags) > 0 {
			for _, tag := range tagsLinks.Tags {
				_, found := b.ActiveTags[tag]
				if !found {
					message.Tags = append(message.Tags, tag)
				}
			}
		}
		// Add the new links.
		if len(tagsLinks.Links) > 0 {
			for _, link := range tagsLinks.Links {
				message.Links = append(message.Links, link)
			}
		}
	}
}

func (b *Builder) PushTag(tag string) {
	b.ActiveTags[tag] = struct{}{}
}

func (b *Builder) PopTag(tag string, loc *ledger.Location) {
	_, found := b.ActiveTags[tag]
	if found {
		delete(b.ActiveTags, tag)
	} else {
		b.AddError(fmt.Sprintf("attempting to pop absent tag: '%s'", tag), loc)
	}
}

func (b *Builder) PushMeta(key string, value *ledger.MetaValue) {
	_ = append(b.ActiveMeta[key], value)
}

func (b *Builder) PopMeta(key string, loc *ledger.Location) {
	iter, found := b.ActiveMeta[key]
	if found {
		iter[len(iter)-1] = nil
		iter = iter[:len(iter)-1]
		if len(iter) == 0 {
			delete(b.ActiveMeta, key)
		}
	} else {
		b.AddError(fmt.Sprintf("attempting to pop absent metadata key: '%s'", key), loc)
	}
}

func (b *Builder) AddActiveMetadata(meta *ledger.Meta, dir *ledger.Directive) {
	if len(b.ActiveMeta) > 0 {
		for key, vl := range b.ActiveMeta {
			dst := new(ledger.MetaValue)
			proto.Merge(dst, vl[len(vl)-1])
			dir.Meta.Kv = append(dir.Meta.Kv, &ledger.Meta_KV{
				Key:   &key,
				Value: dst,
			})
		}
	}
	if meta != nil {
		proto.Merge(dir.Meta, meta)
	}
}

func (b *Builder) MergeCost(newCostSpec, accumulator *ledger.CostSpec) error {
	if newCostSpec.NumberPer != nil && accumulator.NumberPer != nil {
		return errors.New("duplicate `number_per` cost spec field")
	}
	if newCostSpec.NumberTotal != nil && accumulator.NumberTotal != nil {
		return errors.New("duplicate `number_total` cost spec field")
	}
	if newCostSpec.Currency != nil && accumulator.Currency != nil {
		return errors.New("duplicate `currency` cost spec field")
	}
	if newCostSpec.Date != nil && accumulator.Date != nil {
		return errors.New("duplicate `date` cost spec field")
	}
	if newCostSpec.Label != nil && accumulator.Label != nil {
		return errors.New("duplicate `label` cost spec field")
	}
	if newCostSpec.MergeCost != nil && accumulator.MergeCost != nil {
		return errors.New("duplicate `merge_cost` cost spec field")
	}
	proto.Merge(accumulator, newCostSpec)
	return nil
}

func (b *Builder) MakeDirective(date *ledger.Date, meta **ledger.Meta, tagsLinks **TagsLinks) *ledger.Directive {
	dir := new(ledger.Directive)
	dir.Date = date
	var ourMeta *ledger.Meta
	if meta != nil {
		ourMeta = *meta
		meta = nil
	}
	b.AddActiveMetadata(ourMeta, dir)
	ourMeta = nil
	if tagsLinks != nil {
		b.SetTagsAndLinks(*tagsLinks, dir)
		if *tagsLinks != nil {
			*tagsLinks = nil
		}
	}
	return dir
}

func (b *Builder) AppendDirective(directive *ledger.Directive) {
	b.directives_ = append(b.directives_, directive)
}

func (b *Builder) PreparePosting(posting *ledger.Posting, unitSpec *ledger.Amount,
	flag []byte, account string, isTotalPrice bool, loc *ledger.Location) {
	if posting == nil {
		panic("posting is null")
	}
	if unitSpec != nil {
		if posting.Spec == nil {
			posting.Spec = &ledger.Spec{}
		}
		if posting.Spec.Units == nil {
			posting.Spec.Units = &ledger.UnitSpec{}
		}
		spec := posting.Spec.Units
		if unitSpec.Number != nil {
			spec.Number = unitSpec.Number
		}
		if unitSpec.Currency != nil {
			spec.Currency = unitSpec.Currency
		}
	}
	if flag[0] != 0 {
		posting.Flag = flag
	}
	posting.Account = account
	if posting.Spec != nil && posting.Spec.Price != nil {
		spec := posting.Spec
		price := spec.Price

		// If the price is specified for the entire amount, compute the effective
		// price here and forget about that detail of the input syntax.
		if isTotalPrice {
			if spec.Units == nil || spec.Units.Number == nil {
				b.AddError(fmt.Sprintf("total price on a posting without units: %s", price.String()), loc)
				posting.Spec.Price = nil
			} else if price.Number != nil {
				var dunits *decimal.Big = ledger.ProtoToDecimal(spec.Units.Number)

				var dprice *decimal.Big
				if dunits.Sign() == 0 {
					dprice = dunits
				} else {
					b.Ctx.Quo(dprice, ledger.ProtoToDecimal(price.Number), dunits)
				}
				ledger.DecimalToProto(dprice, price.Number)
			}
		}
		// If both cost and price are specified, the currencies must match, or
		// that is an error.
		if spec.Cost != nil && spec.Cost.Currency != nil && price.Currency != nil &&
			spec.Cost.GetCurrency() != price.GetCurrency() {
			b.AddError(fmt.Sprintf("cost and price currencies must match: %s != %s",
				spec.Cost.GetCurrency(), price.GetCurrency()), loc)
		}
	}
}

func (b *Builder) ValidateAccountNames() {

}

func (b *Builder) Finalize(loc *ledger.Location) {
	b.ValidateAccountNames()
	for _, tag := range b.ActiveTags {
		b.AddError(fmt.Sprintf("unbalanced pushed tag: '%s'", tag), loc)
	}
	for key, _ := range b.ActiveMeta {
		b.AddError(fmt.Sprintf("unbalanced metadata key '%s' has leftover metadata", key), loc)
	}
}

func (b *Builder) MakeLedger() *ledger.Ledger {
	lgr := &ledger.Ledger{
		Directives: b.directives_,
		Options:    b.options_,
		Info:       b.info_,
		Errors:     b.errors_,
	}
	return lgr
}

func (b *Builder) AddError(msg string, loc *ledger.Location) {
	err := new(ledger.Error)
	err.Location = loc
	err.Message = Capitalize(msg)
	b.errors_ = append(b.errors_, err)
}

func Capitalize(str string) string {
	if len(str) > 0 {
		r, n := utf8.DecodeRuneInString(str)
		if r != utf8.RuneError || n > 1 {
			upper := unicode.ToUpper(r)
			if upper != r {
				str = string(unicode.ToUpper(r)) + str[n:]
			}
		}
	}
	return str
}

func (b *Builder) DecimalProto(dec *decimal.Big, proto *ledger.Number) {
	ledger.DecimalToProto(dec, proto)
}

type PayeeNarration struct {
	Payee     string
	Narration string
}
