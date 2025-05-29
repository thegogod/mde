package html

import (
	"github.com/thegogod/mde/core"
)

type InputElement struct {
	element *Element
}

func Input() *InputElement {
	return &InputElement{Elem("input").SelfClosing()}
}

func (self *InputElement) Id(value string) *InputElement {
	self.element.Id(value)
	return self
}

func (self *InputElement) Style(styles ...core.KeyValue[string, string]) *InputElement {
	self.element.Style(styles...)
	return self
}

func (self *InputElement) Class(classes ...string) *InputElement {
	self.element.Class(classes...)
	return self
}

func (self *InputElement) Attr(name string, value string) *InputElement {
	self.element.Attr(name, value)
	return self
}

func (self InputElement) String() string {
	return self.element.String()
}

func (self InputElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self InputElement) Bytes() []byte {
	return []byte(self.String())
}

func (self InputElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
