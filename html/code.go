package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/code
type CodeElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/code
func Code(children ...any) *CodeElement {
	return &CodeElement{Elem("code").Push(children...)}
}

func (self *CodeElement) Id(value string) *CodeElement {
	self.element.Id(value)
	return self
}

func (self *CodeElement) Style(styles ...maps.KeyValue[string, string]) *CodeElement {
	self.element.Style(styles...)
	return self
}

func (self *CodeElement) Class(classes ...string) *CodeElement {
	self.element.Class(classes...)
	return self
}

func (self *CodeElement) Attr(name string, value string) *CodeElement {
	self.element.Attr(name, value)
	return self
}

func (self CodeElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *CodeElement) DelAttr(name string) *CodeElement {
	self.element.DelAttr(name)
	return self
}

func (self *CodeElement) Push(children ...any) *CodeElement {
	self.element.Push(children...)
	return self
}

func (self *CodeElement) Pop() *CodeElement {
	self.element.Pop()
	return self
}

func (self CodeElement) Children() []core.Node {
	return self.element.children
}

func (self CodeElement) String() string {
	return self.element.String()
}

func (self CodeElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self CodeElement) Bytes() []byte {
	return []byte(self.String())
}

func (self CodeElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
