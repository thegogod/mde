package html

import "github.com/thegogod/mde/core"

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/title
type TitleElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/title
func Title(children ...any) *TitleElement {
	return &TitleElement{Elem("title").Push(children...)}
}

func (self *TitleElement) Attr(name string, value string) *TitleElement {
	self.element.Attr(name, value)
	return self
}

func (self TitleElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *TitleElement) DelAttr(name string) *TitleElement {
	self.element.DelAttr(name)
	return self
}

func (self *TitleElement) Push(children ...any) *TitleElement {
	self.element.Push(children...)
	return self
}

func (self *TitleElement) Pop() *TitleElement {
	self.element.Pop()
	return self
}

func (self TitleElement) Children() []core.Node {
	return self.element.children
}

func (self TitleElement) String() string {
	return self.element.String()
}

func (self TitleElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TitleElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TitleElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
