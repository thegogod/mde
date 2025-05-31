package markdown

type TokenKind uint8

const (
	Eof TokenKind = iota
	Text

	// whitespace

	NewLine

	// elements

	H1         // # test
	H2         // ## test
	H3         // ### test
	H4         // #### test
	H5         // ##### test
	H6         // ###### test
	Bold       // **
	BoldAlt    // __
	Italic     // *
	ItalicAlt  // _
	Strike     // ~
	StrikeAlt  // ~~
	Br         // line break (<br>) (two spaces)
	Ol         // ordered list (<ol>) (1. test)
	Ul         // unordered list (<ul>) (- test)
	BlockQuote // '>'
	Code       // `test`
	CodeBlock  // ```test```
	Hr         // horizontal rule (<hr>) (---)
	Link       // link (<a>) ([test](https://test.com))
	Image      // image (<img>) (![test](https://test.com/image.png))

	// singles
	Bang         // !
	LeftBracket  // [
	RightBracket // ]
	LeftParen    // (
	RightParen   // )
)

func (self TokenKind) IsInline() bool {
	switch self {
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

func (self TokenKind) String() string {
	switch self {
	case Eof:
		return "eof"
	case Text:
		return "text"
	case NewLine:
		return "newline"
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
	case Link:
		return "link"
	case Image:
		return "image"
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
