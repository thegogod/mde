package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a
type AnchorElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a
func A(children ...any) *AnchorElement {
	return &AnchorElement{Elem("a").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#href
func (self *AnchorElement) Href(value string) *AnchorElement {
	return self.Attr("href", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#target
func (self *AnchorElement) Target(value string) *AnchorElement {
	return self.Attr("target", value)
}

func (self *AnchorElement) Id(value string) *AnchorElement {
	self.element.Id(value)
	return self
}

func (self *AnchorElement) Style(styles ...maps.KeyValue[string, string]) *AnchorElement {
	self.element.Style(styles...)
	return self
}

func (self AnchorElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *AnchorElement) Class(classes ...string) *AnchorElement {
	self.element.Class(classes...)
	return self
}

func (self AnchorElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self AnchorElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *AnchorElement) Attr(name string, value string) *AnchorElement {
	self.element.Attr(name, value)
	return self
}

func (self AnchorElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *AnchorElement) DelAttr(name string) *AnchorElement {
	self.element.DelAttr(name)
	return self
}

func (self *AnchorElement) Push(children ...any) *AnchorElement {
	self.element.Push(children...)
	return self
}

func (self *AnchorElement) Pop() *AnchorElement {
	self.element.Pop()
	return self
}

func (self AnchorElement) Children() []core.Node {
	return self.element.children
}

func (self AnchorElement) String() string {
	return self.element.String()
}

func (self AnchorElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self AnchorElement) Bytes() []byte {
	return []byte(self.String())
}

func (self AnchorElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self AnchorElement) GetById(id string) core.Node {
	return self.element.GetById(id)
}

func (self AnchorElement) GetByClass(classes ...string) []core.Node {
	return self.element.GetByClass(classes...)
}
