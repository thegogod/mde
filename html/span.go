package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/span
type SpanElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/span
func Span(children ...any) *SpanElement {
	return &SpanElement{Elem("span").Push(children...)}
}

func (self *SpanElement) Id(value string) *SpanElement {
	self.element.Id(value)
	return self
}

func (self *SpanElement) Style(styles ...maps.KeyValue[string, string]) *SpanElement {
	self.element.Style(styles...)
	return self
}

func (self SpanElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *SpanElement) Class(classes ...string) *SpanElement {
	self.element.Class(classes...)
	return self
}

func (self *SpanElement) Attr(name string, value string) *SpanElement {
	self.element.Attr(name, value)
	return self
}

func (self SpanElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *SpanElement) DelAttr(name string) *SpanElement {
	self.element.DelAttr(name)
	return self
}

func (self *SpanElement) Push(children ...any) *SpanElement {
	self.element.Push(children...)
	return self
}

func (self *SpanElement) Pop() *SpanElement {
	self.element.Pop()
	return self
}

func (self SpanElement) Children() []core.Node {
	return self.element.children
}

func (self SpanElement) String() string {
	return self.element.String()
}

func (self SpanElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self SpanElement) Bytes() []byte {
	return []byte(self.String())
}

func (self SpanElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
