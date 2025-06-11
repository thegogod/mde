package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ul
type UnorderedListElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ul
func Ul(children ...*ListItemElement) *UnorderedListElement {
	el := Elem("ul")

	for _, child := range children {
		el.Push(child)
	}

	return &UnorderedListElement{el}
}

func (self *UnorderedListElement) Id(value string) *UnorderedListElement {
	self.element.Id(value)
	return self
}

func (self *UnorderedListElement) Style(styles ...maps.KeyValue[string, string]) *UnorderedListElement {
	self.element.Style(styles...)
	return self
}

func (self UnorderedListElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *UnorderedListElement) Class(classes ...string) *UnorderedListElement {
	self.element.Class(classes...)
	return self
}

func (self UnorderedListElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self UnorderedListElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *UnorderedListElement) Attr(name string, value string) *UnorderedListElement {
	self.element.Attr(name, value)
	return self
}

func (self UnorderedListElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *UnorderedListElement) DelAttr(name string) *UnorderedListElement {
	self.element.DelAttr(name)
	return self
}

func (self *UnorderedListElement) Push(children ...*ListItemElement) *UnorderedListElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *UnorderedListElement) Pop() *UnorderedListElement {
	self.element.Pop()
	return self
}

func (self UnorderedListElement) Children() []core.Node {
	return self.element.children
}

func (self UnorderedListElement) String() string {
	return self.element.String()
}

func (self UnorderedListElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self UnorderedListElement) Bytes() []byte {
	return []byte(self.String())
}

func (self UnorderedListElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self UnorderedListElement) GetById(id string) core.Node {
	return self.element.GetById(id)
}

func (self UnorderedListElement) GetByClass(classes ...string) []core.Node {
	return self.element.GetByClass(classes...)
}
