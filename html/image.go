package html

import "github.com/thegogod/mde/maps"

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img
type ImageElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img
func Img() *ImageElement {
	return &ImageElement{Elem("img").Void()}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img#src
func (self *ImageElement) Src(value string) *ImageElement {
	return self.Attr("src", value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/alt#usage_notes
func (self *ImageElement) Alt(value string) *ImageElement {
	return self.Attr("alt", value)
}

func (self *ImageElement) Id(value string) *ImageElement {
	self.element.Id(value)
	return self
}

func (self *ImageElement) Style(styles ...maps.KeyValue[string, string]) *ImageElement {
	self.element.Style(styles...)
	return self
}

func (self ImageElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *ImageElement) Class(classes ...string) *ImageElement {
	self.element.Class(classes...)
	return self
}

func (self *ImageElement) Attr(name string, value string) *ImageElement {
	self.element.Attr(name, value)
	return self
}

func (self ImageElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *ImageElement) DelAttr(name string) *ImageElement {
	self.element.DelAttr(name)
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
	return []byte(self.PrettyString(indent))
}
