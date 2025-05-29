package html

import (
	"github.com/thegogod/mde/core"
)

type StrongElement struct {
	element *Element
}

func Strong() *StrongElement {
	return &StrongElement{Elem("strong")}
}

func (self *StrongElement) Id(value string) *StrongElement {
	self.element.Id(value)
	return self
}

func (self *StrongElement) Style(styles ...core.KeyValue[string, string]) *StrongElement {
	self.element.Style(styles...)
	return self
}

func (self *StrongElement) Class(classes ...string) *StrongElement {
	self.element.Class(classes...)
	return self
}

func (self *StrongElement) Attr(name string, value string) *StrongElement {
	self.element.Attr(name, value)
	return self
}

func (self *StrongElement) Add(children ...any) *StrongElement {
	self.element.Add(children...)
	return self
}

func (self StrongElement) String() string {
	return self.element.String()
}

func (self StrongElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self StrongElement) Bytes() []byte {
	return []byte(self.String())
}

func (self StrongElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
