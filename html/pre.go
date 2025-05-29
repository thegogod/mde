package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/pre
type PreElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/pre
func Pre(children ...any) *PreElement {
	return &PreElement{Elem("pre").Add(children...)}
}

func (self *PreElement) Id(value string) *PreElement {
	self.element.Id(value)
	return self
}

func (self *PreElement) Style(styles ...core.KeyValue[string, string]) *PreElement {
	self.element.Style(styles...)
	return self
}

func (self *PreElement) Class(classes ...string) *PreElement {
	self.element.Class(classes...)
	return self
}

func (self *PreElement) Attr(name string, value string) *PreElement {
	self.element.Attr(name, value)
	return self
}

func (self *PreElement) Add(children ...any) *PreElement {
	self.element.Add(children...)
	return self
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
	return []byte(self.PrettyBytes(indent))
}
