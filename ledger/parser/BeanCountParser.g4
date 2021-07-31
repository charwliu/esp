// BeanCountParser.g4
parser grammar BeanCountParser;

options {tokenVocab=BeanCountLexer;}

// Rules
/* Start token for parsing an entire file of the DSL. */
ledger
  : declarations EOF
  ;

eol
  : NEWLINE
  ;

declarations
  : declarations NEWLINE
  | declarations directive
  | declarations pragma
  |
  ;

/*- Pragmas -----------------------------------------------------------------*/

/* pragmas: Anything but a directive. Those constructs apply side-effects to the
   builder object. */

pragma
  : pushtag
  | poptag
  | pushmeta
  | popmeta
  | option
  | include
  ;

pushtag
  : PUSHTAG TAG eol;

poptag
  : POPTAG TAG eol;

pushmeta
  : PUSHMETA KEY COLON metaValue eol;

popmeta
  : POPMETA KEY COLON eol;

option
  : OPTION key=STRING value=STRING eol;

include
  : INCLUDE STRING eol;


/*- Directives --------------------------------------------------------------*/

/* A rule to reduce any of the Beancount directives we offer. */

directive
  : transaction   # transactionDir
  | price         # priceDir
  | balance       # balanceDir
  | open          # openDir
  | closeDirective # closeDir
  | commodity     # commodityDir
  | pad           # padDir
  | document      # documentDir
  | note          # noteDir
  | event         # eventDir
  | query         # queryDir
  | custom        # customDir
  ;



transaction
  : transactionLine (INDENT metadata postingList DEDENT)?
  ;

/* Matches just the first line of a transaction, without its metadata. */
transactionLine
  : DATETIME txn payeeNarration tagsOrLinks eol
  ;

/* A Price directive. */
price
  : DATETIME PRICE CURRENCY amount tagsOrLinks eol indentedMetadata
  ;

/* A Balance directive. */
balance
  : DATETIME BALANCE account amountTolerance tagsOrLinks eol indentedMetadata
  ;

/* An Open directive. */
open
  : DATETIME OPEN account currencyList booking tagsOrLinks eol indentedMetadata
  ;

/* A Close directive. */
closeDirective
  : DATETIME CLOSE account tagsOrLinks eol indentedMetadata
  ;

/* A Commodity directive. */
commodity
  : DATETIME COMMODITY CURRENCY tagsOrLinks eol indentedMetadata
  ;

/* A Pad directive. */
pad
  : DATETIME PAD account source=account tagsOrLinks eol indentedMetadata
  ;

/* A Document directive. */
document
  : DATETIME DOCUMENT account filename tagsOrLinks eol indentedMetadata
  ;

/* A Note directive. */
note
  : DATETIME NOTE account comment=STRING tagsOrLinks eol indentedMetadata
  ;

/* An Event directive. */
event
  : DATETIME EVENT eventType=STRING description=STRING tagsOrLinks eol indentedMetadata
  ;

/* A Query directive. */
query
  : DATETIME QUERY name=STRING qstr=STRING tagsOrLinks eol indentedMetadata
  ;



/* A Custom directive. Note that this directive does not support the dedicated
   tags and links because those can be part of metadata. It would be ambiguous
   which they are. */
custom
  : DATETIME CUSTOM customType=STRING metaValueList eol indentedMetadata
  ;

/* A list of metadata (optionally empty), including its indentation. */
indentedMetadata
  : INDENT metadata DEDENT
  |
  ;

/*- Metadata ----------------------------------------------------------------*/

/* A list of metadata (optionally empty). */
metadata
  : (metadataLine)*
  ;

/* A single line of a metadata declaration. */
metadataLine
  : KEY COLON metaValue eol
  | TAG eol
  | LINK eol
  ;

metaValue
  : STRING
  | CURRENCY
  | account
  | TAG
  | LINK
  | DATETIME
  | BOOL
  | NONE
  | numberExpr
  | amount
  ;

/* A list of variant values. */
metaValueList
  : (metaValue)*
  ;

/* A container for tags and links, which show up at the end of a transaction's
   first line or in some of the directives. */
tagsOrLinks
  : (TAG | LINK)*
  ;

/*- String matchers ---------------------------------------------------------*/

/* Payee and narration strings. */
payeeNarration
  : payee=STRING narration=STRING
  | narration=STRING
  |
  ;

/* An  account name. */
account
  : ACCOUNT
  ;

/* A filename. Consider specializing this. */
filename
  : STRING
  ;

/* A comma-separated list of currencies. */
currencyList
  : (CURRENCY COMMA?)*
  ;

/*- Arithmetic expressions --------------------------------------------------*/

/* Arithmetic expression, reduced to a number. The reduction is carried out in
   the parser. */
numberExpr
  : numberExpr (PLUS | MINUS) numberExpr     # BinaryOp
  | numberExpr (ASTERISK | SLASH) numberExpr # BinaryOp
  | LPAREN numberExpr RPAREN                 # UniaryOp
  | NUMBER                                   # UniaryOp
  ;


txn
  : TXN
  | FLAG
  | ASTERISK
  | HASH
  ;

optflag
  : (ASTERISK
  | HASH
  | FLAG)?
  ;

booking
  : STRING?
  ;

postingList
  :
  | postingList eol
  | postingList postingAndMetadata
  ;

postingAndMetadata
  : posting indentedMetadata
  ;

posting
  : optflag account partialAmount costSpec (AT|ATAT) priceAnnotation eol    # postingPrice
  | optflag account partialAmount costSpec eol                              # postingAmount
  | optflag account eol                                                     # postingAccount
  ;

priceAnnotation
  : partialAmount
  ;

amount
  : numberExpr CURRENCY
  ;

amountTolerance
  : amount
  | numberExpr TILDE numberExpr CURRENCY
  ;


partialAmount
  : maybeNumber maybeCurrency
  ;

maybeNumber
  : numberExpr?
  ;

maybeCurrency
  : CURRENCY?
  ;

costSpec
  : LCURL costCompList RCURL
  | LCURLCURL costCompList RCURLCURL
  |
  ;

costCompList
  :
  | costComp
  | costCompList COMMA costComp
  ;

costComp
  : compoundAmount
  | DATETIME
  | STRING
  | ASTERISK
  ;

compoundAmount
  : maybeNumber CURRENCY
  | numberExpr maybeCurrency
  | maybeNumber HASH maybeNumber CURRENCY
  ;
