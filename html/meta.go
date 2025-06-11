package html

import "github.com/thegogod/mde/core"

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta
type MetaElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta
func Meta() *MetaElement {
	return &MetaElement{Elem("meta").Void()}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#charset
func (self *MetaElement) Charset(value string) *MetaElement {
	return self.Attr("charset", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#name
func (self *MetaElement) Name(value string) *MetaElement {
	return self.Attr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#content
func (self *MetaElement) Content(value string) *MetaElement {
	return self.Attr("content", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#media
func (self *MetaElement) Media(value string) *MetaElement {
	return self.Attr("media", value)
}

func (self *MetaElement) Attr(name string, value string) *MetaElement {
	self.element.Attr(name, value)
	return self
}

func (self MetaElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *MetaElement) DelAttr(name string) *MetaElement {
	self.element.DelAttr(name)
	return self
}

func (self MetaElement) String() string {
	return self.element.String()
}

func (self MetaElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self MetaElement) Bytes() []byte {
	return []byte(self.String())
}

func (self MetaElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self MetaElement) GetById(id string) core.Node {
	return self.element.GetById(id)
}

func (self MetaElement) GetByClass(classes ...string) []core.Node {
	return self.element.GetByClass(classes...)
}
