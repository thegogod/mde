package html

import (
	"github.com/thegogod/mde/core"
)

type HtmlElement struct {
	element *Element
}

func Html() *HtmlElement {
	return &HtmlElement{Elem("html")}
}

func (self *HtmlElement) Id(value string) *HtmlElement {
	self.element.Id(value)
	return self
}

func (self *HtmlElement) Style(styles ...core.KeyValue[string, string]) *HtmlElement {
	self.element.Style(styles...)
	return self
}

func (self *HtmlElement) Class(classes ...string) *HtmlElement {
	self.element.Class(classes...)
	return self
}

func (self *HtmlElement) Attr(name string, value string) *HtmlElement {
	self.element.Attr(name, value)
	return self
}

func (self *HtmlElement) Add(children ...any) *HtmlElement {
	self.element.Add(children...)
	return self
}

func (self HtmlElement) String() string {
	return self.element.String()
}

func (self HtmlElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self HtmlElement) Bytes() []byte {
	return []byte(self.String())
}

func (self HtmlElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
