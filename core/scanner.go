package mde

type Scanner interface {
	Scan() (Token, error)
}
