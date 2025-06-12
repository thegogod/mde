package core

import (
	"fmt"
)

type Token struct {
	kind  TokenKind
	start Position
	end   Position
	value []byte
}

func NewToken(kind TokenKind, start Position, end Position, value []byte) *Token {
	return &Token{
		kind:  kind,
		start: start,
		end:   end,
		value: value,
	}
}

func (self Token) Kind() rune {
	return self.kind
}

func (self Token) Start() Position {
	return self.start
}

func (self Token) End() Position {
	return self.end
}

func (self Token) Bytes() []byte {
	return self.value
}

func (self Token) String() string {
	return string(self.value)
}

func (self Token) Error(message string) error {
	return Err(self.start, self.end, message)
}

func (self Token) KindString() string {
	switch self.kind {
	case Eof:
		return "eof"
	case Text:
		return "text"
	case NewLine:
		return "newline"
	case Space:
		return "space"
	case Tab:
		return "tab"
	case Dash:
		return "dash"
	case Quote:
		return "quote"
	case DoubleQuote:
		return "double-quote"
	case BackQuote:
		return "back-quote"
	case Colon:
		return "colon"
	case Bang:
		return "bang"
	case Hash:
		return "hash"
	case At:
		return "at"
	case LeftParen:
		return "left-paren"
	case RightParen:
		return "right-paren"
	case LeftBracket:
		return "left-bracket"
	case RightBracket:
		return "right-bracket"
	case LeftBrace:
		return "left-brace"
	case RightBrace:
		return "right-brace"
	case Asterisk:
		return "asterisk"
	case Underscore:
		return "underscore"
	case Slash:
		return "slash"
	case BackSlash:
		return "back-slash"
	case Tilde:
		return "tilde"
	case Equals:
		return "equals"
	case GreaterThan:
		return "greater-than"
	case LessThan:
		return "less-than"
	case Period:
		return "period"
	case Pipe:
		return "pipe"
	case Ampersand:
		return "ampersand"
	case Integer:
		return "integer"
	case Decimal:
		return "decimal"
	default:
		panic(fmt.Sprintf("'%d': unsupported token type", self.kind))
	}
}
