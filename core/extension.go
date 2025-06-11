package core

type Extension interface {
	Name() string
	TokenTypes() []TokenType
	Syntax() []Syntax
}
