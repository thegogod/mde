package html

import (
	"bytes"
	"unicode"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
type HeadingElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func Heading(size int, children ...any) *HeadingElement {
	switch size {
	case 1:
		return H1(children...)
	case 2:
		return H2(children...)
	case 3:
		return H3(children...)
	case 4:
		return H4(children...)
	case 5:
		return H5(children...)
	case 6:
		return H6(children...)
	default:
		panic("invalid heading size")
	}
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

func (self *HeadingElement) Style(styles ...maps.KeyValue[string, string]) *HeadingElement {
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

func (self HeadingElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *HeadingElement) DelAttr(name string) *HeadingElement {
	self.element.DelAttr(name)
	return self
}

func (self *HeadingElement) Add(children ...any) *HeadingElement {
	self.element.Add(children...)
	return self
}

func (self *HeadingElement) Pop() *HeadingElement {
	self.element.Pop()
	return self
}

func (self HeadingElement) Children() []core.Node {
	return self.element.children
}

func (self HeadingElement) String() string {
	if !self.element.attributes.Exists("id") {
		id := []byte{}

		for _, child := range self.element.children {
			value := child.Bytes()

			for _, b := range value {
				if unicode.IsSpace(rune(b)) {
					id = append(id, '-')
				} else if unicode.IsNumber(rune(b)) {
					id = append(id, b)
				} else if unicode.IsLetter(rune(b)) {
					id = append(id, bytes.ToLower([]byte{b})...)
				}
			}
		}

		self.Attr("id", string(id))
	}

	return self.element.String()
}

func (self HeadingElement) PrettyString(indent string) string {
	if !self.element.attributes.Exists("id") {
		id := []byte{}

		for _, child := range self.element.children {
			value := child.Bytes()

			for _, b := range value {
				if unicode.IsSpace(rune(b)) {
					id = append(id, '-')
				} else if unicode.IsNumber(rune(b)) {
					id = append(id, b)
				} else if unicode.IsLetter(rune(b)) {
					id = append(id, bytes.ToLower([]byte{b})...)
				}
			}
		}

		self.Attr("id", string(id))
	}

	return self.element.PrettyString(indent)
}

func (self HeadingElement) Bytes() []byte {
	return []byte(self.String())
}

func (self HeadingElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
