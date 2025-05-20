package core

type Scanner interface {
	Scan() (Token, error)
}
