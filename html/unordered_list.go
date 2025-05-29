package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ul
type UnorderedListElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ul
func Ul(children ...*ListItemElement) *UnorderedListElement {
	el := Elem("ul")

	for _, child := range children {
		el.Add(child)
	}

	return &UnorderedListElement{el}
}

func (self *UnorderedListElement) Id(value string) *UnorderedListElement {
	self.element.Id(value)
	return self
}

func (self *UnorderedListElement) Style(styles ...core.KeyValue[string, string]) *UnorderedListElement {
	self.element.Style(styles...)
	return self
}

func (self *UnorderedListElement) Class(classes ...string) *UnorderedListElement {
	self.element.Class(classes...)
	return self
}

func (self *UnorderedListElement) Attr(name string, value string) *UnorderedListElement {
	self.element.Attr(name, value)
	return self
}

func (self *UnorderedListElement) DelAttr(name string) *UnorderedListElement {
	self.element.DelAttr(name)
	return self
}

func (self *UnorderedListElement) Add(children ...*ListItemElement) *UnorderedListElement {
	for _, child := range children {
		self.element.Add(child)
	}

	return self
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
	return []byte(self.PrettyBytes(indent))
}
