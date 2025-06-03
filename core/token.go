package core

type Token interface {
	Kind() byte
	Position() Position
	Bytes() []byte
	String() string
	Error(message string) error
	IsInline() bool
}
