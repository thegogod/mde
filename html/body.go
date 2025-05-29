package html

import "github.com/thegogod/mde/maps"

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/body
type BodyElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/body
func Body(children ...any) *BodyElement {
	return &BodyElement{Elem("body").Add(children...)}
}

func (self *BodyElement) Id(value string) *BodyElement {
	self.element.Id(value)
	return self
}

func (self *BodyElement) Style(styles ...maps.KeyValue[string, string]) *BodyElement {
	self.element.Style(styles...)
	return self
}

func (self *BodyElement) Class(classes ...string) *BodyElement {
	self.element.Class(classes...)
	return self
}

func (self *BodyElement) Attr(name string, value string) *BodyElement {
	self.element.Attr(name, value)
	return self
}

func (self *BodyElement) DelAttr(name string) *BodyElement {
	self.element.DelAttr(name)
	return self
}

func (self *BodyElement) Add(children ...any) *BodyElement {
	self.element.Add(children...)
	return self
}

func (self BodyElement) String() string {
	return self.element.String()
}

func (self BodyElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self BodyElement) Bytes() []byte {
	return []byte(self.String())
}

func (self BodyElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
