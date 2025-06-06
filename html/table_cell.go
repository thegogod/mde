package html

import (
	"strconv"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/td
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/th
type TableCellElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/td
func Td(children ...any) *TableCellElement {
	return &TableCellElement{Elem("td").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/th
func Th(children ...any) *TableCellElement {
	return &TableCellElement{Elem("th").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/td#colspan
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/th#colspan
func (self *TableCellElement) ColSpan(value int) *TableCellElement {
	return self.Attr("colspan", strconv.Itoa(value))
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/td#rowspan
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/th#colspan
func (self *TableCellElement) RowSpan(value int) *TableCellElement {
	return self.Attr("rowspan", strconv.Itoa(value))
}

func (self *TableCellElement) Id(value string) *TableCellElement {
	self.element.Id(value)
	return self
}

func (self *TableCellElement) Style(styles ...maps.KeyValue[string, string]) *TableCellElement {
	self.element.Style(styles...)
	return self
}

func (self *TableCellElement) Class(classes ...string) *TableCellElement {
	self.element.Class(classes...)
	return self
}

func (self *TableCellElement) Attr(name string, value string) *TableCellElement {
	self.element.Attr(name, value)
	return self
}

func (self TableCellElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *TableCellElement) DelAttr(name string) *TableCellElement {
	self.element.DelAttr(name)
	return self
}

func (self *TableCellElement) Push(children ...any) *TableCellElement {
	self.element.Push(children...)
	return self
}

func (self *TableCellElement) Pop() *TableCellElement {
	self.element.Pop()
	return self
}

func (self TableCellElement) Children() []core.Node {
	return self.element.children
}

func (self TableCellElement) String() string {
	return self.element.String()
}

func (self TableCellElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TableCellElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TableCellElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
