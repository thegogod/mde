package ast

type NewLine struct{}

func (self NewLine) Render() ([]byte, error) {
	return []byte("\n"), nil
}
