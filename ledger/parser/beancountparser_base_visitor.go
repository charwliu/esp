// Code generated from BeanCountParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // BeanCountParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseBeanCountParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBeanCountParserVisitor) VisitLedger(ctx *LedgerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitEol(ctx *EolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitDeclarations(ctx *DeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPragma(ctx *PragmaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPushtag(ctx *PushtagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPoptag(ctx *PoptagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPushmeta(ctx *PushmetaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPopmeta(ctx *PopmetaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitInclude(ctx *IncludeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitTransactionDir(ctx *TransactionDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPriceDir(ctx *PriceDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitBalanceDir(ctx *BalanceDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitOpenDir(ctx *OpenDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCloseDir(ctx *CloseDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCommodityDir(ctx *CommodityDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPadDir(ctx *PadDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitDocumentDir(ctx *DocumentDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitNoteDir(ctx *NoteDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitEventDir(ctx *EventDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitQueryDir(ctx *QueryDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCustomDir(ctx *CustomDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitTransaction(ctx *TransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitTransactionLine(ctx *TransactionLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPrice(ctx *PriceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitBalance(ctx *BalanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitOpen(ctx *OpenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCloseDirective(ctx *CloseDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCommodity(ctx *CommodityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPad(ctx *PadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitNote(ctx *NoteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitEvent(ctx *EventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCustom(ctx *CustomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitIndentedMetadata(ctx *IndentedMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitMetadata(ctx *MetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitMetadataLine(ctx *MetadataLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitMetaValue(ctx *MetaValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitMetaValueList(ctx *MetaValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitTagsOrLinks(ctx *TagsOrLinksContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPayeeNarration(ctx *PayeeNarrationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitAccount(ctx *AccountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitFilename(ctx *FilenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCurrencyList(ctx *CurrencyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitUniaryOp(ctx *UniaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitBinaryOp(ctx *BinaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitTxn(ctx *TxnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitOptflag(ctx *OptflagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitBooking(ctx *BookingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPostingList(ctx *PostingListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPostingAndMetadata(ctx *PostingAndMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPostingPrice(ctx *PostingPriceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPostingAmount(ctx *PostingAmountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPostingAccount(ctx *PostingAccountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPriceAnnotation(ctx *PriceAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitAmount(ctx *AmountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitAmountTolerance(ctx *AmountToleranceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitPartialAmount(ctx *PartialAmountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitMaybeNumber(ctx *MaybeNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitMaybeCurrency(ctx *MaybeCurrencyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCostSpec(ctx *CostSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCostCompList(ctx *CostCompListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCostComp(ctx *CostCompContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanCountParserVisitor) VisitCompoundAmount(ctx *CompoundAmountContext) interface{} {
	return v.VisitChildren(ctx)
}
