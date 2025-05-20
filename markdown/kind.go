package markdown

type Kind uint8

const (
	Eof Kind = iota
	Text

	// whitespace

	NewLine
	Space

	// elements

	H1         // # test
	H2         // ## test
	H3         // ### test
	H4         // #### test
	H5         // ##### test
	H6         // ###### test
	Bold       // ** or __
	Italic     // * or _
	Br         // line break (<br>) (two spaces)
	Ol         // ordered list (<ol>) (1. test)
	Ul         // unordered list (<ul>) (- test)
	BlockQuote // >
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
	case BlockQuote:
		return "block-quote"
	default:
		panic("unsupported token type")
	}
}
