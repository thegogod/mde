package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/label
type LabelElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/label
func Label(children ...any) *LabelElement {
	return &LabelElement{Elem("label").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/for
func (self *LabelElement) For(value string) *LabelElement {
	return self.Attr("for", value)
}

func (self *LabelElement) Id(value string) *LabelElement {
	self.element.Id(value)
	return self
}

func (self *LabelElement) Style(styles ...maps.KeyValue[string, string]) *LabelElement {
	self.element.Style(styles...)
	return self
}

func (self LabelElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *LabelElement) Class(classes ...string) *LabelElement {
	self.element.Class(classes...)
	return self
}

func (self LabelElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self LabelElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *LabelElement) Attr(name string, value string) *LabelElement {
	self.element.Attr(name, value)
	return self
}

func (self LabelElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *LabelElement) DelAttr(name string) *LabelElement {
	self.element.DelAttr(name)
	return self
}

func (self *LabelElement) Push(children ...any) *LabelElement {
	self.element.Push(children...)
	return self
}

func (self *LabelElement) Pop() *LabelElement {
	self.element.Pop()
	return self
}

func (self LabelElement) Children() []Node {
	return self.element.children
}

func (self LabelElement) String() string {
	return self.element.String()
}

func (self LabelElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self LabelElement) Bytes() []byte {
	return []byte(self.String())
}

func (self LabelElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self LabelElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self LabelElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
