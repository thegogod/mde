package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/p
type ParagraphElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/p
func P(children ...any) *ParagraphElement {
	return &ParagraphElement{Elem("p").Push(children...)}
}

func (self *ParagraphElement) Id(value string) *ParagraphElement {
	self.element.Id(value)
	return self
}

func (self *ParagraphElement) Style(styles ...maps.KeyValue[string, string]) *ParagraphElement {
	self.element.Style(styles...)
	return self
}

func (self ParagraphElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *ParagraphElement) Class(classes ...string) *ParagraphElement {
	self.element.Class(classes...)
	return self
}

func (self ParagraphElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self ParagraphElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *ParagraphElement) Attr(name string, value string) *ParagraphElement {
	self.element.Attr(name, value)
	return self
}

func (self ParagraphElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *ParagraphElement) DelAttr(name string) *ParagraphElement {
	self.element.DelAttr(name)
	return self
}

func (self *ParagraphElement) Push(children ...any) *ParagraphElement {
	self.element.Push(children...)
	return self
}

func (self *ParagraphElement) Pop() *ParagraphElement {
	self.element.Pop()
	return self
}

func (self ParagraphElement) Children() []core.Node {
	return self.element.children
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
	return []byte(self.PrettyString(indent))
}

func (self ParagraphElement) GetById(id string) core.Node {
	return self.element.GetById(id)
}

func (self ParagraphElement) GetByClass(classes ...string) []core.Node {
	return self.element.GetByClass(classes...)
}
