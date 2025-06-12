package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/thead
type TableHeadElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/thead
func THead(children ...*TableRowElement) *TableHeadElement {
	self := &TableHeadElement{Elem("thead")}
	self.Push(children...)
	return self
}

func (self *TableHeadElement) Id(value string) *TableHeadElement {
	self.element.Id(value)
	return self
}

func (self *TableHeadElement) Style(styles ...maps.KeyValue[string, string]) *TableHeadElement {
	self.element.Style(styles...)
	return self
}

func (self TableHeadElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableHeadElement) Class(classes ...string) *TableHeadElement {
	self.element.Class(classes...)
	return self
}

func (self TableHeadElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self TableHeadElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *TableHeadElement) Attr(name string, value string) *TableHeadElement {
	self.element.Attr(name, value)
	return self
}

func (self TableHeadElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *TableHeadElement) DelAttr(name string) *TableHeadElement {
	self.element.DelAttr(name)
	return self
}

func (self *TableHeadElement) Push(children ...*TableRowElement) *TableHeadElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *TableHeadElement) Pop() *TableHeadElement {
	self.element.Pop()
	return self
}

func (self TableHeadElement) Children() []Node {
	return self.element.children
}

func (self TableHeadElement) String() string {
	return self.element.String()
}

func (self TableHeadElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TableHeadElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TableHeadElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self TableHeadElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self TableHeadElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
