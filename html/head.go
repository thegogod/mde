package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/head
type HeadElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/head
func Head() *HeadElement {
	return &HeadElement{Elem("head")}
}

func (self *HeadElement) Id(value string) *HeadElement {
	self.element.Id(value)
	return self
}

func (self *HeadElement) Style(styles ...core.KeyValue[string, string]) *HeadElement {
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

func (self *HeadElement) Add(children ...any) *HeadElement {
	self.element.Add(children...)
	return self
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
	return []byte(self.PrettyBytes(indent))
}
