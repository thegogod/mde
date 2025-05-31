package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/head
type HeadElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/head
func Head(children ...any) *HeadElement {
	return &HeadElement{Elem("head").Push(children...)}
}

func (self *HeadElement) Id(value string) *HeadElement {
	self.element.Id(value)
	return self
}

func (self *HeadElement) Style(styles ...maps.KeyValue[string, string]) *HeadElement {
	self.element.Style(styles...)
	return self
}

func (self *HeadElement) Class(classes ...string) *HeadElement {
	self.element.Class(classes...)
	return self
}

func (self *HeadElement) Attr(name string, value string) *HeadElement {
	self.element.Attr(name, value)
	return self
}

func (self HeadElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *HeadElement) DelAttr(name string) *HeadElement {
	self.element.DelAttr(name)
	return self
}

func (self *HeadElement) Push(children ...any) *HeadElement {
	self.element.Push(children...)
	return self
}

func (self *HeadElement) Pop() *HeadElement {
	self.element.Pop()
	return self
}

func (self HeadElement) Children() []core.Node {
	return self.element.children
}

func (self HeadElement) String() string {
	return self.element.String()
}

func (self HeadElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self HeadElement) Bytes() []byte {
	return []byte(self.String())
}

func (self HeadElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
