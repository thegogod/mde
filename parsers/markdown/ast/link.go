package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
)

type Link struct {
	Text []core.Node
	Href []core.Node
}

func (self Link) Render() ([]byte, error) {
	text := []byte{}

	for _, item := range self.Text {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		text = append(text, value...)
	}

	href := []byte{}

	for _, item := range self.Href {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		href = append(href, value...)
	}

	value := fmt.Appendf(nil, `<a href="%s">%s</a>`, href, text)
	return value, nil
}
