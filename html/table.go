package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/table
type TableElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/table
func Table(children ...any) *TableElement {
	return &TableElement{Elem("table").Push(children...)}
}

func (self *TableElement) Id(value string) *TableElement {
	self.element.Id(value)
	return self
}

func (self *TableElement) Style(styles ...maps.KeyValue[string, string]) *TableElement {
	self.element.Style(styles...)
	return self
}

func (self TableElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableElement) Class(classes ...string) *TableElement {
	self.element.Class(classes...)
	return self
}

func (self *TableElement) Attr(name string, value string) *TableElement {
	self.element.Attr(name, value)
	return self
}

func (self TableElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *TableElement) DelAttr(name string) *TableElement {
	self.element.DelAttr(name)
	return self
}

func (self *TableElement) Push(children ...any) *TableElement {
	self.element.Push(children...)
	return self
}

func (self *TableElement) Pop() *TableElement {
	self.element.Pop()
	return self
}

func (self TableElement) Children() []core.Node {
	return self.element.children
}

func (self TableElement) String() string {
	return self.element.String()
}

func (self TableElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TableElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TableElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
