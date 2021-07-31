// Code generated from BeanCountParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // BeanCountParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by BeanCountParser.
type BeanCountParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BeanCountParser#ledger.
	VisitLedger(ctx *LedgerContext) interface{}

	// Visit a parse tree produced by BeanCountParser#eol.
	VisitEol(ctx *EolContext) interface{}

	// Visit a parse tree produced by BeanCountParser#declarations.
	VisitDeclarations(ctx *DeclarationsContext) interface{}

	// Visit a parse tree produced by BeanCountParser#pragma.
	VisitPragma(ctx *PragmaContext) interface{}

	// Visit a parse tree produced by BeanCountParser#pushtag.
	VisitPushtag(ctx *PushtagContext) interface{}

	// Visit a parse tree produced by BeanCountParser#poptag.
	VisitPoptag(ctx *PoptagContext) interface{}

	// Visit a parse tree produced by BeanCountParser#pushmeta.
	VisitPushmeta(ctx *PushmetaContext) interface{}

	// Visit a parse tree produced by BeanCountParser#popmeta.
	VisitPopmeta(ctx *PopmetaContext) interface{}

	// Visit a parse tree produced by BeanCountParser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Visit a parse tree produced by BeanCountParser#include.
	VisitInclude(ctx *IncludeContext) interface{}

	// Visit a parse tree produced by BeanCountParser#transactionDir.
	VisitTransactionDir(ctx *TransactionDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#priceDir.
	VisitPriceDir(ctx *PriceDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#balanceDir.
	VisitBalanceDir(ctx *BalanceDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#openDir.
	VisitOpenDir(ctx *OpenDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#closeDir.
	VisitCloseDir(ctx *CloseDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#commodityDir.
	VisitCommodityDir(ctx *CommodityDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#padDir.
	VisitPadDir(ctx *PadDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#documentDir.
	VisitDocumentDir(ctx *DocumentDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#noteDir.
	VisitNoteDir(ctx *NoteDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#eventDir.
	VisitEventDir(ctx *EventDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#queryDir.
	VisitQueryDir(ctx *QueryDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#customDir.
	VisitCustomDir(ctx *CustomDirContext) interface{}

	// Visit a parse tree produced by BeanCountParser#transaction.
	VisitTransaction(ctx *TransactionContext) interface{}

	// Visit a parse tree produced by BeanCountParser#transactionLine.
	VisitTransactionLine(ctx *TransactionLineContext) interface{}

	// Visit a parse tree produced by BeanCountParser#price.
	VisitPrice(ctx *PriceContext) interface{}

	// Visit a parse tree produced by BeanCountParser#balance.
	VisitBalance(ctx *BalanceContext) interface{}

	// Visit a parse tree produced by BeanCountParser#open.
	VisitOpen(ctx *OpenContext) interface{}

	// Visit a parse tree produced by BeanCountParser#closeDirective.
	VisitCloseDirective(ctx *CloseDirectiveContext) interface{}

	// Visit a parse tree produced by BeanCountParser#commodity.
	VisitCommodity(ctx *CommodityContext) interface{}

	// Visit a parse tree produced by BeanCountParser#pad.
	VisitPad(ctx *PadContext) interface{}

	// Visit a parse tree produced by BeanCountParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by BeanCountParser#note.
	VisitNote(ctx *NoteContext) interface{}

	// Visit a parse tree produced by BeanCountParser#event.
	VisitEvent(ctx *EventContext) interface{}

	// Visit a parse tree produced by BeanCountParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by BeanCountParser#custom.
	VisitCustom(ctx *CustomContext) interface{}

	// Visit a parse tree produced by BeanCountParser#indentedMetadata.
	VisitIndentedMetadata(ctx *IndentedMetadataContext) interface{}

	// Visit a parse tree produced by BeanCountParser#metadata.
	VisitMetadata(ctx *MetadataContext) interface{}

	// Visit a parse tree produced by BeanCountParser#metadataLine.
	VisitMetadataLine(ctx *MetadataLineContext) interface{}

	// Visit a parse tree produced by BeanCountParser#metaValue.
	VisitMetaValue(ctx *MetaValueContext) interface{}

	// Visit a parse tree produced by BeanCountParser#metaValueList.
	VisitMetaValueList(ctx *MetaValueListContext) interface{}

	// Visit a parse tree produced by BeanCountParser#tagsOrLinks.
	VisitTagsOrLinks(ctx *TagsOrLinksContext) interface{}

	// Visit a parse tree produced by BeanCountParser#payeeNarration.
	VisitPayeeNarration(ctx *PayeeNarrationContext) interface{}

	// Visit a parse tree produced by BeanCountParser#account.
	VisitAccount(ctx *AccountContext) interface{}

	// Visit a parse tree produced by BeanCountParser#filename.
	VisitFilename(ctx *FilenameContext) interface{}

	// Visit a parse tree produced by BeanCountParser#currencyList.
	VisitCurrencyList(ctx *CurrencyListContext) interface{}

	// Visit a parse tree produced by BeanCountParser#UniaryOp.
	VisitUniaryOp(ctx *UniaryOpContext) interface{}

	// Visit a parse tree produced by BeanCountParser#BinaryOp.
	VisitBinaryOp(ctx *BinaryOpContext) interface{}

	// Visit a parse tree produced by BeanCountParser#txn.
	VisitTxn(ctx *TxnContext) interface{}

	// Visit a parse tree produced by BeanCountParser#optflag.
	VisitOptflag(ctx *OptflagContext) interface{}

	// Visit a parse tree produced by BeanCountParser#booking.
	VisitBooking(ctx *BookingContext) interface{}

	// Visit a parse tree produced by BeanCountParser#postingList.
	VisitPostingList(ctx *PostingListContext) interface{}

	// Visit a parse tree produced by BeanCountParser#postingAndMetadata.
	VisitPostingAndMetadata(ctx *PostingAndMetadataContext) interface{}

	// Visit a parse tree produced by BeanCountParser#postingPrice.
	VisitPostingPrice(ctx *PostingPriceContext) interface{}

	// Visit a parse tree produced by BeanCountParser#postingAmount.
	VisitPostingAmount(ctx *PostingAmountContext) interface{}

	// Visit a parse tree produced by BeanCountParser#postingAccount.
	VisitPostingAccount(ctx *PostingAccountContext) interface{}

	// Visit a parse tree produced by BeanCountParser#priceAnnotation.
	VisitPriceAnnotation(ctx *PriceAnnotationContext) interface{}

	// Visit a parse tree produced by BeanCountParser#amount.
	VisitAmount(ctx *AmountContext) interface{}

	// Visit a parse tree produced by BeanCountParser#amountTolerance.
	VisitAmountTolerance(ctx *AmountToleranceContext) interface{}

	// Visit a parse tree produced by BeanCountParser#partialAmount.
	VisitPartialAmount(ctx *PartialAmountContext) interface{}

	// Visit a parse tree produced by BeanCountParser#maybeNumber.
	VisitMaybeNumber(ctx *MaybeNumberContext) interface{}

	// Visit a parse tree produced by BeanCountParser#maybeCurrency.
	VisitMaybeCurrency(ctx *MaybeCurrencyContext) interface{}

	// Visit a parse tree produced by BeanCountParser#costSpec.
	VisitCostSpec(ctx *CostSpecContext) interface{}

	// Visit a parse tree produced by BeanCountParser#costCompList.
	VisitCostCompList(ctx *CostCompListContext) interface{}

	// Visit a parse tree produced by BeanCountParser#costComp.
	VisitCostComp(ctx *CostCompContext) interface{}

	// Visit a parse tree produced by BeanCountParser#compoundAmount.
	VisitCompoundAmount(ctx *CompoundAmountContext) interface{}
}
