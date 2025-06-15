package core

type TokenKind = rune

const (
	Eof TokenKind = iota

	// whitespace

	NewLine
	Space
	Tab

	// singles

	Colon             // :
	Bang              // !
	Hash              // #
	At                // @
	LeftBracket       // [
	RightBracket      // ]
	LeftParen         // (
	RightParen        // )
	LeftBrace         // {
	RightBrace        // }
	Asterisk          // *
	Plus              // +
	Percent           // %
	Dash              // -
	Underscore        // _
	Tilde             // ~
	Equals            // =
	EqualsEquals      // ==
	GreaterThan       // >
	GreaterThanEquals // >=
	LessThan          // <
	LessThanEquals    // <=
	Quote             // '
	DoubleQuote       // "
	BackQuote         // `
	Period            // .
	Pipe              // |
	Ampersand         // &
	Slash             // /
	BackSlash         // \

	// compounds

	Integer // 123
	Decimal // 123.4
	Text    // text
)
