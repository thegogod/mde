package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/html
type HtmlElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/html
func Html(children ...any) *HtmlElement {
	return &HtmlElement{Elem("html").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/html#accessibility
func (self *HtmlElement) Lang(value string) *HtmlElement {
	return self.Attr("lang", value)
}

func (self *HtmlElement) Style(styles ...maps.KeyValue[string, string]) *HtmlElement {
	self.element.Style(styles...)
	return self
}

func (self *HtmlElement) Class(classes ...string) *HtmlElement {
	self.element.Class(classes...)
	return self
}

func (self *HtmlElement) Attr(name string, value string) *HtmlElement {
	self.element.Attr(name, value)
	return self
}

func (self HtmlElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *HtmlElement) DelAttr(name string) *HtmlElement {
	self.element.DelAttr(name)
	return self
}

func (self *HtmlElement) Push(children ...any) *HtmlElement {
	self.element.Push(children...)
	return self
}

func (self *HtmlElement) Pop() *HtmlElement {
	self.element.Pop()
	return self
}

func (self HtmlElement) Children() []core.Node {
	return self.element.children
}

func (self HtmlElement) String() string {
	return self.element.String()
}

func (self HtmlElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self HtmlElement) Bytes() []byte {
	return []byte(self.String())
}

func (self HtmlElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
