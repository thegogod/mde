package core

type Token interface {
	GetKind() uint8
	GetPosition() Position
	GetBytes() []byte
	String() string
}
