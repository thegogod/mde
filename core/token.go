package core

type Token interface {
	GetKind() TokenKind
	GetPosition() Position
	GetBytes() []byte
	String() string
}
