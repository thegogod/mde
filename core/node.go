package core

type Node interface {
	Render() ([]byte, error)
}
