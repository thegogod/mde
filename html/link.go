package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link
type LinkElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link
func Link() *LinkElement {
	return &LinkElement{Elem("link").Void()}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/rel
func (self *LinkElement) Rel(value string) *LinkElement {
	return self.Attr("rel", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link#href
func (self *LinkElement) Href(value string) *LinkElement {
	return self.Attr("href", value)
}

func (self *LinkElement) Id(value string) *LinkElement {
	self.element.Id(value)
	return self
}

func (self *LinkElement) Style(styles ...maps.KeyValue[string, string]) *LinkElement {
	self.element.Style(styles...)
	return self
}

func (self LinkElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *LinkElement) Class(classes ...string) *LinkElement {
	self.element.Class(classes...)
	return self
}

func (self LinkElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self LinkElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *LinkElement) Attr(name string, value string) *LinkElement {
	self.element.Attr(name, value)
	return self
}

func (self LinkElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *LinkElement) DelAttr(name string) *LinkElement {
	self.element.DelAttr(name)
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
	return []byte(self.PrettyString(indent))
}

func (self LinkElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self LinkElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
