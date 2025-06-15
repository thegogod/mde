package core

// AST Node
type Node interface {
	Render(scope Scope) []byte
	RenderPretty(scope Scope, indent string) []byte
}
