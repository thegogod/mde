package core

// Parses tokens into an AST
type Parser interface {
	ParseBlock() (Node, error)
	ParseInline() (Node, error)
}
