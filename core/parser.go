package core

// Parses tokens into an AST
type Parser interface {
	ParseBlock(parser Parser, iter *Iterator) (Node, error)
	ParseInline(parser Parser, iter *Iterator) (Node, error)
	ParseSyntax(extension string, name string, parser Parser, iter *Iterator) (Node, error)

	ParseText(parser Parser, iter *Iterator) ([]byte, error)
	ParseTextUntil(kind rune, parser Parser, iter *Iterator) ([]byte, error)
}
