package core

// Token Iterator
type Iterator interface {
	Stateful

	Curr() Token
	Prev() Token
	Next() bool
	Reset(src []byte)
	Match(kind byte) bool
	MatchCount(kind byte, count int) bool
	Consume(kind byte, message string) (Token, error)
}
