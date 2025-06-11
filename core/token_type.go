package core

type TokenType struct {
	Id   rune
	Name string
	Scan func(ptr *Pointer) (*Token, error)
}
