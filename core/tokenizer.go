package core

type Tokenizer struct {
	Id   rune
	Name string
	Scan func(ptr *Pointer) (*Token, error)
}
