package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link
type LinkElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link
func Link() *LinkElement {
	return &LinkElement{Elem("link").SelfClosing()}
}

func (self *LinkElement) Rel(value string) *LinkElement {
	return self.Attr("rel", value)
}

func (self *LinkElement) Href(value string) *LinkElement {
	return self.Attr("href", value)
}

func (self *LinkElement) Id(value string) *LinkElement {
	self.element.Id(value)
	return self
}

func (self *LinkElement) Style(styles ...core.KeyValue[string, string]) *LinkElement {
	self.element.Style(styles...)
	return self
}

func (self *LinkElement) Class(classes ...string) *LinkElement {
	self.element.Class(classes...)
	return self
}

func (self *LinkElement) Attr(name string, value string) *LinkElement {
	self.element.Attr(name, value)
	return self
}

func (self LinkElement) String() string {
	return self.element.String()
}

func (self LinkElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self LinkElement) Bytes() []byte {
	return []byte(self.String())
}

func (self LinkElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
