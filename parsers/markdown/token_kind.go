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
	BlockQuote // >
	Code       // `test`
	CodeBlock  // ```test```
	Hr         // horizontal rule (<hr>) (---)
	Link       // link (<a>) ([test](https://test.com))
	Image      // image (<img>) (![test](https://test.com/image.png))
)

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
	default:
		panic("unsupported token type")
	}
}
