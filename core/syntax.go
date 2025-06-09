package core

type Syntax interface {
	IsBlock() bool
	IsInline() bool
	Name() string
	Select(parser Parser, iter *Iterator) bool
	Parse(parser Parser, iter *Iterator) (Node, error)
}
