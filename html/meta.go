package html

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta
type MetaElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta
func Meta() *MetaElement {
	return &MetaElement{Elem("meta").Void()}
}

func (self *MetaElement) Charset(value string) *MetaElement {
	return self.Attr("charset", value)
}

func (self *MetaElement) Name(value string) *MetaElement {
	return self.Attr("name", value)
}

func (self *MetaElement) Content(value string) *MetaElement {
	return self.Attr("content", value)
}

func (self *MetaElement) Attr(name string, value string) *MetaElement {
	self.element.Attr(name, value)
	return self
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
	return []byte(self.PrettyBytes(indent))
}
