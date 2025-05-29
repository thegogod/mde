package html

import (
	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
type HeadingElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H1(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h1").Add(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H2(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h2").Add(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H3(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h3").Add(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H4(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h4").Add(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H5(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h5").Add(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H6(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h6").Add(children...)}
}

func (self *HeadingElement) Id(value string) *HeadingElement {
	self.element.Id(value)
	return self
}

func (self *HeadingElement) Style(styles ...core.KeyValue[string, string]) *HeadingElement {
	self.element.Style(styles...)
	return self
}

func (self *HeadingElement) Class(classes ...string) *HeadingElement {
	self.element.Class(classes...)
	return self
}

func (self *HeadingElement) Attr(name string, value string) *HeadingElement {
	self.element.Attr(name, value)
	return self
}

func (self *HeadingElement) DelAttr(name string) *HeadingElement {
	self.element.DelAttr(name)
	return self
}

func (self *HeadingElement) Add(children ...any) *HeadingElement {
	self.element.Add(children...)
	return self
}

func (self HeadingElement) String() string {
	return self.element.String()
}

func (self HeadingElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self HeadingElement) Bytes() []byte {
	return []byte(self.String())
}

func (self HeadingElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
