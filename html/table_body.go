package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tbody
type TableBodyElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tbody
func TBody(children ...*TableRowElement) *TableBodyElement {
	self := &TableBodyElement{Elem("tbody")}
	self.Push(children...)
	return self
}

func (self *TableBodyElement) Id(value string) *TableBodyElement {
	self.element.Id(value)
	return self
}

func (self *TableBodyElement) Style(styles ...maps.KeyValue[string, string]) *TableBodyElement {
	self.element.Style(styles...)
	return self
}

func (self TableBodyElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableBodyElement) Class(classes ...string) *TableBodyElement {
	self.element.Class(classes...)
	return self
}

func (self TableBodyElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self TableBodyElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *TableBodyElement) Attr(name string, value string) *TableBodyElement {
	self.element.Attr(name, value)
	return self
}

func (self TableBodyElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *TableBodyElement) DelAttr(name string) *TableBodyElement {
	self.element.DelAttr(name)
	return self
}

func (self *TableBodyElement) Push(children ...*TableRowElement) *TableBodyElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *TableBodyElement) Pop() *TableBodyElement {
	self.element.Pop()
	return self
}

func (self TableBodyElement) Children() []Node {
	return self.element.children
}

func (self TableBodyElement) String() string {
	return self.element.String()
}

func (self TableBodyElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TableBodyElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TableBodyElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self TableBodyElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self TableBodyElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
