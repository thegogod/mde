package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/p
type ParagraphElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/p
func P(children ...any) *ParagraphElement {
	return &ParagraphElement{Elem("p").Add(children...)}
}

func (self *ParagraphElement) Id(value string) *ParagraphElement {
	self.element.Id(value)
	return self
}

func (self *ParagraphElement) Style(styles ...core.KeyValue[string, string]) *ParagraphElement {
	self.element.Style(styles...)
	return self
}

func (self *ParagraphElement) Class(classes ...string) *ParagraphElement {
	self.element.Class(classes...)
	return self
}

func (self *ParagraphElement) Attr(name string, value string) *ParagraphElement {
	self.element.Attr(name, value)
	return self
}

func (self *ParagraphElement) DelAttr(name string) *ParagraphElement {
	self.element.DelAttr(name)
	return self
}

func (self *ParagraphElement) Add(children ...any) *ParagraphElement {
	self.element.Add(children...)
	return self
}

func (self ParagraphElement) String() string {
	return self.element.String()
}

func (self ParagraphElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self ParagraphElement) Bytes() []byte {
	return []byte(self.String())
}

func (self ParagraphElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
