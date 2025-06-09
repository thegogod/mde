package core

type Token interface {
	Kind() byte
	KindString() string
	Start() Position
	End() Position
	Bytes() []byte
	String() string
	Error(message string) error
}
