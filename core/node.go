package core

// AST Node
type Node interface {
	Render() []byte
	RenderPretty(indent string) []byte
}
