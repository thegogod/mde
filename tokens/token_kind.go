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
	At           // @
	LeftBracket  // [
	RightBracket // ]
	LeftParen    // (
	RightParen   // )
	LeftBrace    // {
	RightBrace   // }
	Asterisk     // *
	Dash         // -
	Underscore   // _
	Tilde        // ~
	Equals       // =
	GreaterThan  // >
	LessThan     // <
	BackQuote    // `
	Period       // .
	Pipe         // |
	Ampersand    // &
	Slash        // /
	BackSlash    // \

	// compounds

	Integer // 123
	Decimal // 123.4
)
