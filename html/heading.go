package html

import (
	"github.com/thegogod/mde/core"
)

type HeadingElement struct {
	element *Element
}

func H1() *HeadingElement {
	return &HeadingElement{Elem("h1")}
}

func H2() *HeadingElement {
	return &HeadingElement{Elem("h2")}
}

func H3() *HeadingElement {
	return &HeadingElement{Elem("h3")}
}

func H4() *HeadingElement {
	return &HeadingElement{Elem("h4")}
}

func H5() *HeadingElement {
	return &HeadingElement{Elem("h5")}
}

func H6() *HeadingElement {
	return &HeadingElement{Elem("h6")}
}

func (self *HeadingElement) Id(value string) *HeadingElement {
	self.element.Id(value)
	return self
}

func (self *HeadingElement) Style(styles ...core.KeyValue[string, string]) *HeadingElement {
	self.element.Style(styles...)
	return self
}

func (self *HeadingElement) Class(classes ...string) *HeadingElement {
	self.element.Class(classes...)
	return self
}

func (self *HeadingElement) Attr(name string, value string) *HeadingElement {
	self.element.Attr(name, value)
	return self
}

func (self *HeadingElement) Add(children ...any) *HeadingElement {
	self.element.Add(children...)
	return self
}

func (self HeadingElement) String() string {
	return self.element.String()
}

func (self HeadingElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self HeadingElement) Bytes() []byte {
	return []byte(self.String())
}

func (self HeadingElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
