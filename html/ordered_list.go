package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ol
type OrderedListElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ol
func Ol(children ...*ListItemElement) *OrderedListElement {
	el := Elem("ol")

	for _, child := range children {
		el.Push(child)
	}

	return &OrderedListElement{el}
}

func (self *OrderedListElement) Id(value string) *OrderedListElement {
	self.element.Id(value)
	return self
}

func (self *OrderedListElement) Style(styles ...maps.KeyValue[string, string]) *OrderedListElement {
	self.element.Style(styles...)
	return self
}

func (self *OrderedListElement) Class(classes ...string) *OrderedListElement {
	self.element.Class(classes...)
	return self
}

func (self *OrderedListElement) Attr(name string, value string) *OrderedListElement {
	self.element.Attr(name, value)
	return self
}

func (self OrderedListElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *OrderedListElement) DelAttr(name string) *OrderedListElement {
	self.element.DelAttr(name)
	return self
}

func (self *OrderedListElement) Push(children ...*ListItemElement) *OrderedListElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *OrderedListElement) Pop() *OrderedListElement {
	self.element.Pop()
	return self
}

func (self OrderedListElement) Children() []core.Node {
	return self.element.children
}

func (self OrderedListElement) String() string {
	return self.element.String()
}

func (self OrderedListElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self OrderedListElement) Bytes() []byte {
	return []byte(self.String())
}

func (self OrderedListElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
