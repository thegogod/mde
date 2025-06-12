package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote
type BlockQuoteElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote
func BlockQuote(children ...any) *BlockQuoteElement {
	return &BlockQuoteElement{Elem("blockquote").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote#cite
func (self *BlockQuoteElement) Cite(value string) *BlockQuoteElement {
	return self.Attr("cite", value)
}

func (self *BlockQuoteElement) Id(value string) *BlockQuoteElement {
	self.element.Id(value)
	return self
}

func (self *BlockQuoteElement) Style(styles ...maps.KeyValue[string, string]) *BlockQuoteElement {
	self.element.Style(styles...)
	return self
}

func (self BlockQuoteElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *BlockQuoteElement) Class(classes ...string) *BlockQuoteElement {
	self.element.Class(classes...)
	return self
}

func (self BlockQuoteElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self BlockQuoteElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *BlockQuoteElement) Attr(name string, value string) *BlockQuoteElement {
	self.element.Attr(name, value)
	return self
}

func (self BlockQuoteElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *BlockQuoteElement) DelAttr(name string) *BlockQuoteElement {
	self.element.DelAttr(name)
	return self
}

func (self *BlockQuoteElement) Push(children ...any) *BlockQuoteElement {
	self.element.Push(children...)
	return self
}

func (self *BlockQuoteElement) Pop() *BlockQuoteElement {
	self.element.Pop()
	return self
}

func (self BlockQuoteElement) Children() []Node {
	return self.element.children
}

func (self BlockQuoteElement) String() string {
	return self.element.String()
}

func (self BlockQuoteElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self BlockQuoteElement) Bytes() []byte {
	return []byte(self.String())
}

func (self BlockQuoteElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self BlockQuoteElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self BlockQuoteElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
