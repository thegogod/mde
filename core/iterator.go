package core

// Token Iterator
type Iterator interface {
	Stateful

	Curr() Token
	Prev() Token

	Next() bool
	Match(kind ...byte) bool
	MatchCount(kind byte, count int) bool
	MatchBytes(value ...byte) bool
	Consume(kind byte, message string) (Token, error)
}
