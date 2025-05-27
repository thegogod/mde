package ast

type Br struct{}

func (self Br) Render() ([]byte, error) {
	return []byte("<br>"), nil
}
