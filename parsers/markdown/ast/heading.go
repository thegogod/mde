package ast

import (
	"bytes"
	"fmt"
	"unicode"

	"github.com/thegogod/mde/core"
)

type Heading struct {
	Depth   int
	Content []core.Node
}

func (self *Heading) Add(items ...core.Node) *Heading {
	self.Content = append(self.Content, items...)
	return self
}

func (self Heading) Render() ([]byte, error) {
	id := []byte{}
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)

		for _, b := range value {
			if unicode.IsSpace(rune(b)) {
				id = append(id, '-')
			} else if unicode.IsNumber(rune(b)) {
				id = append(id, b)
			} else if unicode.IsLetter(rune(b)) {
				id = append(id, bytes.ToLower([]byte{b})...)
			}
		}
	}

	value := fmt.Appendf(nil, `<h%d id="%s">%s</h%d>`, self.Depth, id, content, self.Depth)
	return value, nil
}
