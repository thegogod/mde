package core

// Token Iterator
type Iterator interface {
	Stateful

	Curr() Node
	Prev() Node
	Next() bool
	Reset(src []byte)
	Match(kind TokenKind) bool
	MatchCount(kind TokenKind, count int) bool
	Consume(kind TokenKind, message string) (Token, error)
}
