package lexer

import "../token"

//Lexer type.
type Lexer struct {
	input        string
	positon      int  //curren positon in input.
	readPosition int  //curren reading positon in input.
	ch           byte //curren char	under examination.
}

//New Explained structure.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//readChar is to read a char at current positon and positon++.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.positon = l.readPosition
	l.readPosition++
}

//peakChar is like to read a char at current positon but without positon++.
func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

//NextToken to get the next_token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	//eat no meanings whitespace we get.
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.LESSTHAN_EQ, Literal: "<="}
		} else {
			tok = newToken(token.LESSTHAN, l.ch)
		}
	case '>':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.GREATHAN_EQ, Literal: ">="}
		} else {
			tok = newToken(token.GREATHAN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

//newToken is a function to make a new token.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

//readIdentifier is to read a Identifier from current positon.
func (l *Lexer) readIdentifier() string {
	positon := l.positon
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[positon:l.positon]
}

//readNumber is a function to read a NUMBER from current positon.
func (l *Lexer) readNumber() string {
	positon := l.positon
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[positon:l.positon]
}

//skipWhitespace is a function to skip all the whitespace we meet.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

//isLetter is to judge a letter can be a legal LETTER.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

//isDigit is to judge a letter can be a legal NUMBER.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
