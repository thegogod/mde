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
	case H1:
		fallthrough
	case H2:
		fallthrough
	case H3:
		fallthrough
	case H4:
		fallthrough
	case H5:
		fallthrough
	case H6:
		fallthrough
	case Ol:
		fallthrough
	case Ul:
		fallthrough
	case CodeBlock:
		fallthrough
	case BlockQuote:
		fallthrough
	case Hr:
		fallthrough
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
	case Tab:
		return "tab"
	case H1:
		return "h1"
	case H2:
		return "h2"
	case H3:
		return "h3"
	case H4:
		return "h4"
	case H5:
		return "h5"
	case H6:
		return "h6"
	case Bold:
		return "bold"
	case BoldAlt:
		return "bold-alt"
	case Italic:
		return "italic"
	case ItalicAlt:
		return "italic-alt"
	case Strike:
		return "strike"
	case StrikeAlt:
		return "strike-alt"
	case Br:
		return "br"
	case Ol:
		return "ol"
	case Ul:
		return "ul"
	case BlockQuote:
		return "block-quote"
	case Code:
		return "code"
	case CodeBlock:
		return "code-block"
	case Hr:
		return "hr"
	case Colon:
		return "colon"
	case Bang:
		return "bang"
	case LeftParen:
		return "("
	case RightParen:
		return ")"
	case LeftBracket:
		return "["
	case RightBracket:
		return "]"
	default:
		panic("unsupported token type")
	}
}
