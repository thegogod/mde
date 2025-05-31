package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/strong
type StrongElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/strong
func Strong(children ...any) *StrongElement {
	return &StrongElement{Elem("strong").Push(children...)}
}

func (self *StrongElement) Id(value string) *StrongElement {
	self.element.Id(value)
	return self
}

func (self *StrongElement) Style(styles ...maps.KeyValue[string, string]) *StrongElement {
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

func (self StrongElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *StrongElement) DelAttr(name string) *StrongElement {
	self.element.DelAttr(name)
	return self
}

func (self *StrongElement) Push(children ...any) *StrongElement {
	self.element.Push(children...)
	return self
}

func (self *StrongElement) Pop() *StrongElement {
	self.element.Pop()
	return self
}

func (self StrongElement) Children() []core.Node {
	return self.element.children
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
