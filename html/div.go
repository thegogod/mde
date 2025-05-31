package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/div
type DivElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/div
func Div(children ...any) *DivElement {
	return &DivElement{Elem("div").Add(children...)}
}

func (self *DivElement) Id(value string) *DivElement {
	self.element.Id(value)
	return self
}

func (self *DivElement) Style(styles ...maps.KeyValue[string, string]) *DivElement {
	self.element.Style(styles...)
	return self
}

func (self *DivElement) Class(classes ...string) *DivElement {
	self.element.Class(classes...)
	return self
}

func (self *DivElement) Attr(name string, value string) *DivElement {
	self.element.Attr(name, value)
	return self
}

func (self DivElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *DivElement) DelAttr(name string) *DivElement {
	self.element.DelAttr(name)
	return self
}

func (self *DivElement) Add(children ...any) *DivElement {
	self.element.Add(children...)
	return self
}

func (self *DivElement) Pop() *DivElement {
	self.element.Pop()
	return self
}

func (self DivElement) Children() []core.Node {
	return self.element.children
}

func (self DivElement) String() string {
	return self.element.String()
}

func (self DivElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self DivElement) Bytes() []byte {
	return []byte(self.String())
}

func (self DivElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
