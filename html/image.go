package html

import (
	"github.com/thegogod/mde/core"
)

type ImageElement struct {
	element *Element
}

func Img() *ImageElement {
	return &ImageElement{Elem("img").SelfClosing()}
}

func (self *ImageElement) Id(value string) *ImageElement {
	self.element.Id(value)
	return self
}

func (self *ImageElement) Style(styles ...core.KeyValue[string, string]) *ImageElement {
	self.element.Style(styles...)
	return self
}

func (self *ImageElement) Class(classes ...string) *ImageElement {
	self.element.Class(classes...)
	return self
}

func (self *ImageElement) Attr(name string, value string) *ImageElement {
	self.element.Attr(name, value)
	return self
}

func (self ImageElement) String() string {
	return self.element.String()
}

func (self ImageElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self ImageElement) Bytes() []byte {
	return []byte(self.String())
}

func (self ImageElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
