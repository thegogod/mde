package html

import (
	"github.com/thegogod/mde/core"
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

func (self *SlotElement) Id(value string) *SlotElement {
	self.element.Id(value)
	return self
}

func (self *SlotElement) Style(styles ...maps.KeyValue[string, string]) *SlotElement {
	self.element.Style(styles...)
	return self
}

func (self SlotElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *SlotElement) Class(classes ...string) *SlotElement {
	self.element.Class(classes...)
	return self
}

func (self SlotElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self SlotElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *SlotElement) Attr(name string, value string) *SlotElement {
	self.element.Attr(name, value)
	return self
}

func (self SlotElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *SlotElement) DelAttr(name string) *SlotElement {
	self.element.DelAttr(name)
	return self
}

func (self *SlotElement) Push(children ...any) *SlotElement {
	self.element.Push(children...)
	return self
}

func (self *SlotElement) Pop() *SlotElement {
	self.element.Pop()
	return self
}

func (self SlotElement) Children() []core.Node {
	return self.element.children
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

func (self SlotElement) GetById(id string) core.Node {
	return self.element.GetById(id)
}

func (self SlotElement) GetByClass(classes ...string) []core.Node {
	return self.element.GetByClass(classes...)
}
