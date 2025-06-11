package core

// AST Node
type Node interface {
	GetById(id string) Node
	GetByClass(classes ...string) []Node

	String() string
	PrettyString(indent string) string

	Bytes() []byte
	PrettyBytes(indent string) []byte
}
