package ast

type NodeKind uint8

const (
	Statement NodeKind = iota
	Expression
)

func (self NodeKind) String() string {
	switch self {
	case Statement:
		return "statement"
	case Expression:
		return "expression"
	default:
		panic("unsupported ast node kind")
	}
}
