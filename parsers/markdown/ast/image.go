package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
)

type Image struct {
	Alt []core.Node
	Src []core.Node
}

func (self Image) Render() ([]byte, error) {
	alt := []byte{}

	for _, item := range self.Alt {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		alt = append(alt, value...)
	}

	src := []byte{}

	for _, item := range self.Src {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		src = append(src, value...)
	}

	value := fmt.Appendf(nil, `<img alt="%s" src="%s" />`, alt, src)
	return value, nil
}
