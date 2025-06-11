package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/pre
type PreElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/pre
func Pre(children ...any) *PreElement {
	return &PreElement{Elem("pre").Push(children...)}
}

func (self *PreElement) Id(value string) *PreElement {
	self.element.Id(value)
	return self
}

func (self *PreElement) Style(styles ...maps.KeyValue[string, string]) *PreElement {
	self.element.Style(styles...)
	return self
}

func (self PreElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *PreElement) Class(classes ...string) *PreElement {
	self.element.Class(classes...)
	return self
}

func (self PreElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self PreElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *PreElement) Attr(name string, value string) *PreElement {
	self.element.Attr(name, value)
	return self
}

func (self PreElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *PreElement) DelAttr(name string) *PreElement {
	self.element.DelAttr(name)
	return self
}

func (self *PreElement) Push(children ...any) *PreElement {
	self.element.Push(children...)
	return self
}

func (self *PreElement) Pop() *PreElement {
	self.element.Pop()
	return self
}

func (self PreElement) Children() []core.Node {
	return self.element.children
}

func (self PreElement) String() string {
	return self.element.String()
}

func (self PreElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self PreElement) Bytes() []byte {
	return []byte(self.String())
}

func (self PreElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self PreElement) GetById(id string) core.Node {
	return self.element.GetById(id)
}

func (self PreElement) GetByClass(classes ...string) []core.Node {
	return self.element.GetByClass(classes...)
}
