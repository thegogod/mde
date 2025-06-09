package core

// Parses tokens into an AST
type Parser interface {
	ParseBlock(iter *Iterator) (Node, error)
	ParseInline(iter *Iterator) (Node, error)
	ParseSyntax(name string, iter *Iterator) (Node, error)

	ParseText(iter *Iterator) ([]byte, error)
	ParseTextUntil(kind byte, iter *Iterator) ([]byte, error)
}
