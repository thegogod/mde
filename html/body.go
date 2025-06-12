package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/body
type BodyElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/body
func Body(children ...any) *BodyElement {
	return &BodyElement{Elem("body").Push(children...)}
}

func (self *BodyElement) Id(value string) *BodyElement {
	self.element.Id(value)
	return self
}

func (self *BodyElement) Style(styles ...maps.KeyValue[string, string]) *BodyElement {
	self.element.Style(styles...)
	return self
}

func (self BodyElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *BodyElement) Class(classes ...string) *BodyElement {
	self.element.Class(classes...)
	return self
}

func (self BodyElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self BodyElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *BodyElement) Attr(name string, value string) *BodyElement {
	self.element.Attr(name, value)
	return self
}

func (self BodyElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *BodyElement) DelAttr(name string) *BodyElement {
	self.element.DelAttr(name)
	return self
}

func (self *BodyElement) Push(children ...any) *BodyElement {
	self.element.Push(children...)
	return self
}

func (self *BodyElement) Pop() *BodyElement {
	self.element.Pop()
	return self
}

func (self BodyElement) Children() []Node {
	return self.element.children
}

func (self BodyElement) String() string {
	return self.element.String()
}

func (self BodyElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self BodyElement) Bytes() []byte {
	return []byte(self.String())
}

func (self BodyElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self BodyElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self BodyElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
