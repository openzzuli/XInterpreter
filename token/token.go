package token

//TokenType Type as string.
type TokenType string

//Token Structure.
type Token struct {
	Type    TokenType
	Literal string
}

//ALL Tokens defines here.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + Literals.
	IDENT = "IDENT"
	INT   = "INT"

	//Operators.
	COMMA     = ","
	SEMICOLON = ";"
	ASSIGN    = "ASSIGN"
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"

	LESSTHAN = "<"
	GREATHAN = ">"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// KEYWORDS is a map to REMEMBER all the Keywords in this context.
var KEYWORDS = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

//LookupIdent is a FUNCTION to ensure that there is such a Keywords in this context.
func LookupIdent(ident string) TokenType {
	if tok, ok := KEYWORDS[ident]; ok {
		return tok
	}
	return IDENT
}
