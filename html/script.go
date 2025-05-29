package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
type ScriptElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
func Script(children ...any) *ScriptElement {
	return &ScriptElement{Elem("script").Add(children...)}
}

func (self *ScriptElement) Src(value string) *ScriptElement {
	self.Attr("src", value)
	return self
}

func (self *ScriptElement) Async() *ScriptElement {
	self.Attr("async", "")
	return self
}

func (self *ScriptElement) Defer() *ScriptElement {
	self.Attr("defer", "")
	return self
}

func (self *ScriptElement) Id(value string) *ScriptElement {
	self.element.Id(value)
	return self
}

func (self *ScriptElement) Style(styles ...core.KeyValue[string, string]) *ScriptElement {
	self.element.Style(styles...)
	return self
}

func (self *ScriptElement) Class(classes ...string) *ScriptElement {
	self.element.Class(classes...)
	return self
}

func (self *ScriptElement) Attr(name string, value string) *ScriptElement {
	self.element.Attr(name, value)
	return self
}

func (self *ScriptElement) Add(children ...any) *ScriptElement {
	self.element.Add(children...)
	return self
}

func (self ScriptElement) String() string {
	return self.element.String()
}

func (self ScriptElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self ScriptElement) Bytes() []byte {
	return []byte(self.String())
}

func (self ScriptElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
