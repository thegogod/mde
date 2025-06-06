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
	Dash         // -
	Underscore   // _
	Tilde        // ~
	Equals       // =
	GreaterThan  // >
	BackQuote    // `
	Period       // .
	Pipe         // |

	// compounds

	Integer // 123
	Decimal // 123.4

	// elements

	Ol // ordered list (<ol>) (1. test)
	Ul // unordered list (<ul>) (- test)
)
