package tokens

type Kind uint8

const (
	Eof Kind = iota
	Text

	// whitespace

	NewLine
	Space

	// elements

	H1     // # test
	H2     // ## test
	H3     // ### test
	H4     // #### test
	H5     // ##### test
	H6     // ###### test
	Bold   // ** or __
	Italic // * or _
	Br     // line break (<br>) (two spaces)
	Ol     // ordered list (<ol>) (1. test)
	Ul     // unordered list (<ul>) (- test)

	// singles

	Dot          // .
	LeftParen    // (
	RightParen   // )
	LeftBracket  // [
	RightBracket // ]
	Hash         // #
	Underscore   // _
	Hyphen       // -
	BackTick     // `
	Asterisk     // *
	Bang         // !
	Gt           // >
	Lt           // <
)

func (self Kind) String() string {
	switch self {
	case Eof:
		return "eof"
	case Text:
		return "text"
	case NewLine:
		return "newline"
	case Space:
		return "space"
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
	case Italic:
		return "italic"
	case Br:
		return "br"
	case Ol:
		return "ol"
	case Ul:
		return "ul"
	case Dot:
		return "dot"
	case LeftParen:
		return "left-paren"
	case RightParen:
		return "right-paren"
	case LeftBracket:
		return "left-bracket"
	case RightBracket:
		return "right-bracket"
	case Hash:
		return "hash"
	case Underscore:
		return "underscore"
	case BackTick:
		return "back-tick"
	case Asterisk:
		return "star"
	case Hyphen:
		return "hyphen"
	case Bang:
		return "bang"
	case Gt:
		return "gt"
	case Lt:
		return "lt"
	default:
		panic("unsupported token type")
	}
}
