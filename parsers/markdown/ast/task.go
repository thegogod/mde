package ast

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	Label   string
	Checked bool
}

func (self Task) Render() ([]byte, error) {
	id := uuid.NewString()
	label := fmt.Sprintf(`<label for="%s">%s</label>`, id, self.Label)
	input := fmt.Sprintf(`<input id="%s" type="checkbox" />`, id)

	if self.Checked {
		input = fmt.Sprintf(`<input id="%s" type="checkbox" checked />`, id)
	}

	value := fmt.Appendf(nil, "%s%s", input, label)
	return value, nil
}
