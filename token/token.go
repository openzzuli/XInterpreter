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
	IDEN = "IDEN"
	INT  = "INT"

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
