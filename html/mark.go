package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/mark
type MarkElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/mark
func Mark(children ...any) *MarkElement {
	return &MarkElement{Elem("mark").Push(children...)}
}

func (self *MarkElement) Id(value string) *MarkElement {
	self.element.Id(value)
	return self
}

func (self *MarkElement) Style(styles ...maps.KeyValue[string, string]) *MarkElement {
	self.element.Style(styles...)
	return self
}

func (self MarkElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *MarkElement) Class(classes ...string) *MarkElement {
	self.element.Class(classes...)
	return self
}

func (self MarkElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self MarkElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *MarkElement) Attr(name string, value string) *MarkElement {
	self.element.Attr(name, value)
	return self
}

func (self MarkElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *MarkElement) DelAttr(name string) *MarkElement {
	self.element.DelAttr(name)
	return self
}

func (self *MarkElement) Push(children ...any) *MarkElement {
	self.element.Push(children...)
	return self
}

func (self *MarkElement) Pop() *MarkElement {
	self.element.Pop()
	return self
}

func (self MarkElement) Children() []Node {
	return self.element.children
}

func (self MarkElement) String() string {
	return self.element.String()
}

func (self MarkElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self MarkElement) Bytes() []byte {
	return []byte(self.String())
}

func (self MarkElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self MarkElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self MarkElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
