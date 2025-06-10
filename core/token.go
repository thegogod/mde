package core

type Token interface {
	Kind() rune
	KindString() string
	Start() Position
	End() Position
	Bytes() []byte
	String() string
	Error(message string) error
}
