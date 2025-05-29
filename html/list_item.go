package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/li
type ListItemElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/li
func Li(children ...any) *ListItemElement {
	return &ListItemElement{Elem("li").Add(children...)}
}

func (self *ListItemElement) Id(value string) *ListItemElement {
	self.element.Id(value)
	return self
}

func (self *ListItemElement) Style(styles ...core.KeyValue[string, string]) *ListItemElement {
	self.element.Style(styles...)
	return self
}

func (self *ListItemElement) Class(classes ...string) *ListItemElement {
	self.element.Class(classes...)
	return self
}

func (self *ListItemElement) Attr(name string, value string) *ListItemElement {
	self.element.Attr(name, value)
	return self
}

func (self *ListItemElement) Add(children ...any) *ListItemElement {
	self.element.Add(children...)
	return self
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
	return []byte(self.PrettyBytes(indent))
}
