package html

import "github.com/thegogod/mde/maps"

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input
type InputElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#input_types
func Input(kind string) *InputElement {
	self := &InputElement{Elem("input").Void()}
	return self.Type(kind)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#input_types
func (self *InputElement) Type(value string) *InputElement {
	return self.Attr("type", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#name
func (self *InputElement) Name(value string) *InputElement {
	return self.Attr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *InputElement) Disabled() *InputElement {
	return self.Attr("disabled", "true")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *InputElement) Enabled() *InputElement {
	return self.Attr("disabled", "false")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#form
func (self *InputElement) Form(formId string) *InputElement {
	return self.Attr("form", formId)
}

func (self *InputElement) Id(value string) *InputElement {
	self.element.Id(value)
	return self
}

func (self *InputElement) Style(styles ...maps.KeyValue[string, string]) *InputElement {
	self.element.Style(styles...)
	return self
}

func (self *InputElement) Class(classes ...string) *InputElement {
	self.element.Class(classes...)
	return self
}

func (self *InputElement) Attr(name string, value string) *InputElement {
	self.element.Attr(name, value)
	return self
}

func (self InputElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *InputElement) DelAttr(name string) *InputElement {
	self.element.DelAttr(name)
	return self
}

func (self InputElement) String() string {
	return self.element.String()
}

func (self InputElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self InputElement) Bytes() []byte {
	return []byte(self.String())
}

func (self InputElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
