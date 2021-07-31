lexer grammar BeanCountLexer;

tokens { INDENT, DEDENT }
@lexer::members {
var Denter *DenterHelper

func (p *BeanCountLexer) NextToken() antlr.Token {
	return Denter.NextToken()
}
}

// Tokens
NEWLINE
 : ( '\r'? '\n' | '\r' | '\f' ) SPACES?;

WHITESPACE
  : [ \r\t]+ -> skip;

COMMENTS
  : [;].*? -> skip
  ;

BOOL
  :'TRUE' | 'FALSE'
  ;

NONE
  :'NULL'
  ;

CURRENCY
  : [A-Z][A-Z0-9'._-]*[A-Z0-9]?;

PIPE: '|';
ATAT: '@@';
AT: '@';
LCURLCURL: '{{';
RCURLCURL: '}}';
LCURL:'{';
RCURL:'}';
COMMA:',';
TILDE:'~';
ASTERISK: '*';
SLASH: '/';
PLUS: '+';
MINUS: '-';
COLON:':';

LPAREN          : '(';
RPAREN          : ')';
LBRACK          : '[';
RBRACK          : ']';
SEMI            : ';';
DOT             : '.';

FLAG
  : FLAGS
  | ['][A-Z][']
  ;

NUMBER: (SIGN)? ('.' DIGIT+ | NON_ZERO_DIGIT DIGIT* ('.' DIGIT*)?)
  ;

TXN
  : 'txn'
  ;
BALANCE
  : 'balance'
  ;

OPEN
  : 'open'
  ;

CLOSE
  : 'close'
  ;
COMMODITY
  : 'commodity'
  ;

PAD
  : 'pad'
  ;

EVENT
  : 'event'
  ;

QUERY
  :'query'
  ;

CUSTOM
:'custom'
;

PRICE
  :'price'
  ;

NOTE
  :'note'
  ;

DOCUMENT
  :'document'
  ;

PUSHTAG
  :'pushtag'
  ;

POPTAG
  :'poptag'
  ;

PUSHMETA
  :'pushmeta'
  ;

POPMETA
  :'popmeta'
  ;

INCLUDE
  :'include'
  ;

OPTION
  :'option'
  ;

HASH
  : '#'
  ;

DATETIME
  : DATE ([ \t]+ TIME)?
  ;

ACCOUNT
  : ACCOUNTTYPE (':' ACCOUNTNAME)+
  ;

LINK
  : '^'[A-Za-z0-9\-_/.]+
  ;

TAG
  : '#'[A-Za-z0-0\-_/.]+;
KEY
  : [a-z][a-zA-Z0-9\-_]+;

STRING
  : '"' ('\\"'|.)*? '"';

fragment
LETTER
  : UNICODE_LETTER
  | '_'
  ;

fragment
UNICODE_DIGIT
  : [\p{Nd}];

fragment
UNICODE_LETTER
  : [\p{L}];

fragment
HEX_DIGIT
  : [0-9a-fA-F]
  ;

fragment
NON_ZERO_DIGIT
 : [1-9]
 ;

fragment
INT4
  : DIGIT DIGIT DIGIT DIGIT
  ;

fragment
DIGIT
  : [0-9]
  ;

fragment
SIGN
  : MINUS
  | PLUS
  ;

fragment
SEPARATOR
  : [/\\\-]
  ;

fragment
SPACES
 : [ \t]+
 ;


fragment
ACCOUNTTYPE
  : [\p{Lu}\p{Lo}][\p{L}\p{N}-]*
  ;

fragment
ACCOUNTNAME
  : [\p{Lu}\p{Lo}\p{N}][\p{L}\p{N}-]*
  ;

// Characters that may be used as flags. Single-character letters may also be
// used as flags, but must be prefixed by a quote, as in "'P' for "P".
fragment
FLAGS
  : [!&#?%]
  ;

fragment
TIME
  : ([01]? DIGIT | '2' [0-4] ) ':' [0-5]? DIGIT (':' [0-5]? DIGIT)?
  ;

fragment
DATE
  : INT4 SEPARATOR [01]?DIGIT SEPARATOR [0-3]? DIGIT
  ;
