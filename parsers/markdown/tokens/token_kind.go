package tokens

type TokenKind = byte

const (
	Eof TokenKind = iota
	Text

	// whitespace

	NewLine
	Space
	Tab

	// singles

	Colon        // :
	Bang         // !
	Hash         // #
	LeftBracket  // [
	RightBracket // ]
	LeftParen    // (
	RightParen   // )
	Asterisk     // *
	Underscore   // _
	Tilde        // ~
	// Equals       // =

	// elements

	Highlight  // ==
	Br         // line break (<br>) (two spaces)
	Ol         // ordered list (<ol>) (1. test)
	Ul         // unordered list (<ul>) (- test)
	BlockQuote // '>'
	Code       // `test`
	CodeBlock  // ```test```
	Hr         // horizontal rule (<hr>) (---)
)
