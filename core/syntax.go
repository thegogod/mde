package core

type Syntax interface {
	IsBlock() bool
	IsInline() bool
	Name() string
	Select(iter Iterator) bool
	Parse(parser Parser, iter Iterator) (Node, error)
}
