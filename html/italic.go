package html

import "github.com/thegogod/mde/maps"

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/i
type ItalicElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/i
func I(children ...any) *ItalicElement {
	return &ItalicElement{Elem("i").Add(children...)}
}

func (self *ItalicElement) Id(value string) *ItalicElement {
	self.element.Id(value)
	return self
}

func (self *ItalicElement) Style(styles ...maps.KeyValue[string, string]) *ItalicElement {
	self.element.Style(styles...)
	return self
}

func (self *ItalicElement) Class(classes ...string) *ItalicElement {
	self.element.Class(classes...)
	return self
}

func (self *ItalicElement) Attr(name string, value string) *ItalicElement {
	self.element.Attr(name, value)
	return self
}

func (self *ItalicElement) DelAttr(name string) *ItalicElement {
	self.element.DelAttr(name)
	return self
}

func (self *ItalicElement) Add(children ...any) *ItalicElement {
	self.element.Add(children...)
	return self
}

func (self ItalicElement) String() string {
	return self.element.String()
}

func (self ItalicElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self ItalicElement) Bytes() []byte {
	return []byte(self.String())
}

func (self ItalicElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
