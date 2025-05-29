package html

import (
	"github.com/thegogod/mde/core"
)

type BreakLineElement struct {
	element *Element
}

func Br() *BreakLineElement {
	return &BreakLineElement{Elem("br").SelfClosing()}
}

func (self *BreakLineElement) Id(value string) *BreakLineElement {
	self.element.Id(value)
	return self
}

func (self *BreakLineElement) Style(styles ...core.KeyValue[string, string]) *BreakLineElement {
	self.element.Style(styles...)
	return self
}

func (self *BreakLineElement) Class(classes ...string) *BreakLineElement {
	self.element.Class(classes...)
	return self
}

func (self *BreakLineElement) Attr(name string, value string) *BreakLineElement {
	self.element.Attr(name, value)
	return self
}

func (self BreakLineElement) String() string {
	return self.element.String()
}

func (self BreakLineElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self BreakLineElement) Bytes() []byte {
	return []byte(self.String())
}

func (self BreakLineElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
