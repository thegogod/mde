package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/slot
type SlotElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/slot
func Slot(children ...any) *SlotElement {
	return &SlotElement{Elem("slot").Push(children...)}
}

func (self SlotElement) GetTag() string {
	return self.element.GetTag()
}

func (self *SlotElement) WithAttr(name string, value string) *SlotElement {
	self.element.WithAttr(name, value)
	return self
}

func (self SlotElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self SlotElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *SlotElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *SlotElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *SlotElement) WithId(value string) *SlotElement {
	self.element.WithId(value)
	return self
}

func (self SlotElement) HasId() bool {
	return self.element.HasId()
}

func (self SlotElement) GetId() string {
	return self.element.GetId()
}

func (self *SlotElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *SlotElement) DelId() {
	self.element.DelId()
}

func (self *SlotElement) WithClass(classes ...string) *SlotElement {
	self.element.WithClass(classes...)
	return self
}

func (self SlotElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self SlotElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *SlotElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *SlotElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *SlotElement) WithStyles(styles ...maps.KeyValue[string, string]) *SlotElement {
	self.element.WithStyles(styles...)
	return self
}

func (self SlotElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *SlotElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *SlotElement) WithStyle(name string, value string) *SlotElement {
	self.element.WithStyle(name, value)
	return self
}

func (self SlotElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self SlotElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *SlotElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *SlotElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self SlotElement) Count() int {
	return self.element.Count()
}

func (self SlotElement) Children() []Node {
	return self.element.Children()
}

func (self *SlotElement) Push(children ...any) *SlotElement {
	self.element.Push(children...)
	return self
}

func (self *SlotElement) Pop() *SlotElement {
	self.element.Pop()
	return self
}

func (self SlotElement) String() string {
	return self.element.String()
}

func (self SlotElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self SlotElement) Bytes() []byte {
	return []byte(self.String())
}

func (self SlotElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self *SlotElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *SlotElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
