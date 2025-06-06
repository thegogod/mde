package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/li
type ListItemElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/li
func Li(children ...any) *ListItemElement {
	return &ListItemElement{Elem("li").Push(children...)}
}

func (self *ListItemElement) Id(value string) *ListItemElement {
	self.element.Id(value)
	return self
}

func (self *ListItemElement) Style(styles ...maps.KeyValue[string, string]) *ListItemElement {
	self.element.Style(styles...)
	return self
}

func (self ListItemElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *ListItemElement) Class(classes ...string) *ListItemElement {
	self.element.Class(classes...)
	return self
}

func (self *ListItemElement) Attr(name string, value string) *ListItemElement {
	self.element.Attr(name, value)
	return self
}

func (self ListItemElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *ListItemElement) DelAttr(name string) *ListItemElement {
	self.element.DelAttr(name)
	return self
}

func (self *ListItemElement) Push(children ...any) *ListItemElement {
	self.element.Push(children...)
	return self
}

func (self *ListItemElement) Pop() *ListItemElement {
	self.element.Pop()
	return self
}

func (self ListItemElement) Children() []core.Node {
	return self.element.children
}

func (self ListItemElement) String() string {
	return self.element.String()
}

func (self ListItemElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self ListItemElement) Bytes() []byte {
	return []byte(self.String())
}

func (self ListItemElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
