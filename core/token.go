package mde

type Token interface {
	GetTokenKind() uint8
	GetPosition() Position
	GetBytes() []byte
	String() string
}
