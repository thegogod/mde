package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tr
type TableRowElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tr
func Tr(children ...*TableCellElement) *TableRowElement {
	self := &TableRowElement{Elem("tr")}
	self.Push(children...)
	return self
}

func (self *TableRowElement) Id(value string) *TableRowElement {
	self.element.Id(value)
	return self
}

func (self *TableRowElement) Style(styles ...maps.KeyValue[string, string]) *TableRowElement {
	self.element.Style(styles...)
	return self
}

func (self TableRowElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableRowElement) Class(classes ...string) *TableRowElement {
	self.element.Class(classes...)
	return self
}

func (self TableRowElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self TableRowElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *TableRowElement) Attr(name string, value string) *TableRowElement {
	self.element.Attr(name, value)
	return self
}

func (self TableRowElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *TableRowElement) DelAttr(name string) *TableRowElement {
	self.element.DelAttr(name)
	return self
}

func (self *TableRowElement) Push(children ...*TableCellElement) *TableRowElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *TableRowElement) Pop() *TableRowElement {
	self.element.Pop()
	return self
}

func (self TableRowElement) Children() []Node {
	return self.element.children
}

func (self TableRowElement) String() string {
	return self.element.String()
}

func (self TableRowElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TableRowElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TableRowElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self TableRowElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self TableRowElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
