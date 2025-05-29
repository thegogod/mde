package html

import (
	"github.com/thegogod/mde/core"
)

type Paragraph struct {
	element *Element
}

func P() *Paragraph {
	return &Paragraph{Elem("p")}
}

func (self *Paragraph) Id(value string) *Paragraph {
	self.element.Id(value)
	return self
}

func (self *Paragraph) Style(styles ...core.KeyValue[string, string]) *Paragraph {
	self.element.Style(styles...)
	return self
}

func (self *Paragraph) Class(classes ...string) *Paragraph {
	self.element.Class(classes...)
	return self
}

func (self *Paragraph) Attr(name string, value string) *Paragraph {
	self.element.Attr(name, value)
	return self
}

func (self *Paragraph) Add(children ...any) *Paragraph {
	self.element.Add(children...)
	return self
}

func (self Paragraph) String() string {
	return self.element.String()
}

func (self Paragraph) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}
