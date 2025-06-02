package core

type TokenKind uint8

const (
	Eof TokenKind = iota

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
)
