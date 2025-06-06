package core

type Token interface {
	Kind() byte
	KindString() string
	Position() Position
	Bytes() []byte
	String() string
	Error(message string) error
}
