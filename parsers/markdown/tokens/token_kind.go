package tokens

type TokenKind = byte

const (
	Eof TokenKind = iota
	Text

	// whitespace

	NewLine
	Tab

	// singles

	Colon        // :
	Bang         // !
	LeftBracket  // [
	RightBracket // ]
	LeftParen    // (
	RightParen   // )

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
	Highlight  // ==
	Br         // line break (<br>) (two spaces)
	Ol         // ordered list (<ol>) (1. test)
	Ul         // unordered list (<ul>) (- test)
	BlockQuote // '>'
	Code       // `test`
	CodeBlock  // ```test```
	Hr         // horizontal rule (<hr>) (---)
)
