package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tfoot
type TableFootElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tfoot
func TFoot(children ...*TableRowElement) *TableFootElement {
	self := &TableFootElement{Elem("tfoot")}
	self.Push(children...)
	return self
}

func (self *TableFootElement) Id(value string) *TableFootElement {
	self.element.Id(value)
	return self
}

func (self *TableFootElement) Style(styles ...maps.KeyValue[string, string]) *TableFootElement {
	self.element.Style(styles...)
	return self
}

func (self TableFootElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableFootElement) Class(classes ...string) *TableFootElement {
	self.element.Class(classes...)
	return self
}

func (self TableFootElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self TableFootElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *TableFootElement) Attr(name string, value string) *TableFootElement {
	self.element.Attr(name, value)
	return self
}

func (self TableFootElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *TableFootElement) DelAttr(name string) *TableFootElement {
	self.element.DelAttr(name)
	return self
}

func (self *TableFootElement) Push(children ...*TableRowElement) *TableFootElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *TableFootElement) Pop() *TableFootElement {
	self.element.Pop()
	return self
}

func (self TableFootElement) Children() []Node {
	return self.element.children
}

func (self TableFootElement) String() string {
	return self.element.String()
}

func (self TableFootElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TableFootElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TableFootElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self TableFootElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self TableFootElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
