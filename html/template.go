package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/template
type TemplateElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/template
func Template(children ...any) *TemplateElement {
	return &TemplateElement{Elem("template").Push(children...)}
}

func (self *TemplateElement) Id(value string) *TemplateElement {
	self.element.Id(value)
	return self
}

func (self *TemplateElement) Style(styles ...maps.KeyValue[string, string]) *TemplateElement {
	self.element.Style(styles...)
	return self
}

func (self TemplateElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TemplateElement) Class(classes ...string) *TemplateElement {
	self.element.Class(classes...)
	return self
}

func (self TemplateElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self TemplateElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *TemplateElement) Attr(name string, value string) *TemplateElement {
	self.element.Attr(name, value)
	return self
}

func (self TemplateElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *TemplateElement) DelAttr(name string) *TemplateElement {
	self.element.DelAttr(name)
	return self
}

func (self *TemplateElement) Push(children ...any) *TemplateElement {
	self.element.Push(children...)
	return self
}

func (self *TemplateElement) Pop() *TemplateElement {
	self.element.Pop()
	return self
}

func (self TemplateElement) Children() []core.Node {
	return self.element.children
}

func (self TemplateElement) String() string {
	return self.element.String()
}

func (self TemplateElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TemplateElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TemplateElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self TemplateElement) GetById(id string) core.Node {
	return self.element.GetById(id)
}

func (self TemplateElement) GetByClass(classes ...string) []core.Node {
	return self.element.GetByClass(classes...)
}
