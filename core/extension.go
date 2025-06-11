package core

type Extension interface {
	Name() string
	Tokenizers() []Tokenizer
	Syntax() []Syntax
}
