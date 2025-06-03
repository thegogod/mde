package core

// Parses tokens into an AST
type Parser interface {
	Parse(src []byte) (Node, error)
	ParseBlock(iter Iterator) (Node, error)
	ParseInline(iter Iterator) (Node, error)
	ParseSyntax(name string, iter Iterator) (Node, error)

	ParseText(iter Iterator) ([]byte, error)
	ParseTextUntil(iter Iterator, kind byte) ([]byte, error)
}
