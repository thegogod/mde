package core

type Scanner interface {
	Stateful

	Pointer() *Pointer
	Scan() (*Token, error)
}
