package html

import (
	"strconv"

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
func (self *TableCellElement) WithColSpan(value int) *TableCellElement {
	return self.WithAttr("colspan", strconv.Itoa(value))
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/td#rowspan
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/th#colspan
func (self *TableCellElement) WithRowSpan(value int) *TableCellElement {
	return self.WithAttr("rowspan", strconv.Itoa(value))
}

func (self TableCellElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TableCellElement) WithAttr(name string, value string) *TableCellElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TableCellElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TableCellElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TableCellElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TableCellElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TableCellElement) WithId(value string) *TableCellElement {
	self.element.WithId(value)
	return self
}

func (self TableCellElement) HasId() bool {
	return self.element.HasId()
}

func (self TableCellElement) GetId() string {
	return self.element.GetId()
}

func (self *TableCellElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TableCellElement) DelId() {
	self.element.DelId()
}

func (self *TableCellElement) WithClass(classes ...string) *TableCellElement {
	self.element.WithClass(classes...)
	return self
}

func (self TableCellElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TableCellElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TableCellElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TableCellElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TableCellElement) WithStyles(styles ...maps.KeyValue[string, string]) *TableCellElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TableCellElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableCellElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TableCellElement) WithStyle(name string, value string) *TableCellElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TableCellElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TableCellElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TableCellElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TableCellElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TableCellElement) Count() int {
	return self.element.Count()
}

func (self TableCellElement) Children() []Node {
	return self.element.Children()
}

func (self *TableCellElement) Push(children ...any) *TableCellElement {
	self.element.Push(children...)
	return self
}

func (self *TableCellElement) Pop() *TableCellElement {
	self.element.Pop()
	return self
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

func (self *TableCellElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TableCellElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
