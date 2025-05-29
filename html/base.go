package html

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base
type BaseElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base
func Base() *BaseElement {
	return &BaseElement{Elem("base").Void()}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base#href
func (self *BaseElement) Href(value string) *BaseElement {
	return self.Attr("href", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base#target
func (self *BaseElement) Target(value string) *BaseElement {
	return self.Attr("target", value)
}

func (self *BaseElement) Attr(name string, value string) *BaseElement {
	self.element.Attr(name, value)
	return self
}

func (self BaseElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *BaseElement) DelAttr(name string) *BaseElement {
	self.element.DelAttr(name)
	return self
}

func (self BaseElement) String() string {
	return self.element.String()
}

func (self BaseElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self BaseElement) Bytes() []byte {
	return []byte(self.String())
}

func (self BaseElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
