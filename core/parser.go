package core

type Parser interface {
	Parse(src []byte) (Node, error)
}
