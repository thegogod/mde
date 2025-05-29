package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/label
type LabelElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/label
func Label(children ...any) *LabelElement {
	return &LabelElement{Elem("label").Add(children...)}
}

func (self *LabelElement) For(value string) *LabelElement {
	return self.Attr("for", value)
}

func (self *LabelElement) Id(value string) *LabelElement {
	self.element.Id(value)
	return self
}

func (self *LabelElement) Style(styles ...core.KeyValue[string, string]) *LabelElement {
	self.element.Style(styles...)
	return self
}

func (self *LabelElement) Class(classes ...string) *LabelElement {
	self.element.Class(classes...)
	return self
}

func (self *LabelElement) Attr(name string, value string) *LabelElement {
	self.element.Attr(name, value)
	return self
}

func (self *LabelElement) Add(children ...any) *LabelElement {
	self.element.Add(children...)
	return self
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
	return []byte(self.PrettyBytes(indent))
}
