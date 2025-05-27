package ast

type Hr struct{}

func (self Hr) Render() ([]byte, error) {
	return []byte("<hr>"), nil
}
