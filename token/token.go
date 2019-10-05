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

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	PLUS   = "PLUS"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// KEYWORDS is a map to REMEMBER all the Keywords in this context.
var KEYWORDS = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

//LookupIdent is a FUNCTION to ensure that there is such a Keywords in this context.
func LookupIdent(ident string) TokenType {
	if tok, ok := KEYWORDS[ident]; ok {
		return tok
	}
	return IDENT
}
