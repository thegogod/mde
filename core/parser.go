package core

// Parses tokens into an AST
type Parser interface {
	Parse(src []byte) (Node, error)
	ParseBlock(iter Iterator) (Node, error)
	ParseInline(iter Iterator) (Node, error)
}
