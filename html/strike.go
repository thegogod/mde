package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/s
type StrikeElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/s
func S(children ...any) *StrikeElement {
	return &StrikeElement{Elem("s").Push(children...)}
}

func (self *StrikeElement) Id(value string) *StrikeElement {
	self.element.Id(value)
	return self
}

func (self *StrikeElement) Style(styles ...maps.KeyValue[string, string]) *StrikeElement {
	self.element.Style(styles...)
	return self
}

func (self StrikeElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *StrikeElement) Class(classes ...string) *StrikeElement {
	self.element.Class(classes...)
	return self
}

func (self StrikeElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self StrikeElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *StrikeElement) Attr(name string, value string) *StrikeElement {
	self.element.Attr(name, value)
	return self
}

func (self StrikeElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *StrikeElement) DelAttr(name string) *StrikeElement {
	self.element.DelAttr(name)
	return self
}

func (self *StrikeElement) Push(children ...any) *StrikeElement {
	self.element.Push(children...)
	return self
}

func (self *StrikeElement) Pop() *StrikeElement {
	self.element.Pop()
	return self
}

func (self StrikeElement) Children() []Node {
	return self.element.children
}

func (self StrikeElement) String() string {
	return self.element.String()
}

func (self StrikeElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self StrikeElement) Bytes() []byte {
	return []byte(self.String())
}

func (self StrikeElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self StrikeElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self StrikeElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
