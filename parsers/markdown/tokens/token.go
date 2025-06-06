package tokens

import "github.com/thegogod/mde/core"

type Token struct {
	kind     TokenKind
	position core.Position
	value    []byte
}

func NewToken(kind TokenKind, position core.Position, value []byte) *Token {
	return &Token{
		kind:     kind,
		position: position,
		value:    value,
	}
}

func (self Token) Kind() byte {
	return self.kind
}

func (self Token) Position() core.Position {
	return self.position
}

func (self Token) Bytes() []byte {
	return self.value
}

func (self Token) String() string {
	return string(self.value)
}

func (self Token) Error(message string) error {
	return NewError(self.position, message)
}

func (self Token) IsInline() bool {
	switch self.kind {
	case Ol:
		fallthrough
	case Ul:
		fallthrough
	case GreaterThan:
		return false
	case Eof:
		return false
	default:
		return true
	}
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
	case Ol:
		return "ol"
	case Ul:
		return "ul"
	case Dash:
		return "dash"
	case BackQuote:
		return "back-quote"
	case Colon:
		return "colon"
	case Bang:
		return "bang"
	case Hash:
		return "hash"
	case LeftParen:
		return "left-paren"
	case RightParen:
		return "right-paren"
	case LeftBracket:
		return "left-bracket"
	case RightBracket:
		return "right-bracket"
	case Asterisk:
		return "asterisk"
	case Tilde:
		return "tilde"
	case Equals:
		return "equals"
	case GreaterThan:
		return "greater-than"
	case Period:
		return "period"
	case Pipe:
		return "pipe"
	case Integer:
		return "integer"
	case Decimal:
		return "decimal"
	default:
		panic("unsupported token type")
	}
}
