package ast

type Parser interface {
	Parse() ([]Stmt, error)
	ParseLine() (Stmt, error)
}
