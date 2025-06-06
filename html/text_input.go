package html

import (
	"strconv"

	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text
type TextInputElement struct {
	element *InputElement
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text
func TextInput() *TextInputElement {
	return &TextInputElement{Input("text")}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#name
func (self *TextInputElement) Name(value string) *TextInputElement {
	return self.Attr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *TextInputElement) Disabled() *TextInputElement {
	return self.Attr("disabled", "true")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *TextInputElement) Enabled() *TextInputElement {
	return self.Attr("disabled", "false")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#form
func (self *TextInputElement) Form(formId string) *TextInputElement {
	return self.Attr("form", formId)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#value
func (self *TextInputElement) Value(value string) *TextInputElement {
	return self.Attr("value", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#maxlength
func (self *TextInputElement) MaxLength(value int) *TextInputElement {
	return self.Attr("maxlength", strconv.Itoa(value))
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#minlength
func (self *TextInputElement) MinLength(value int) *TextInputElement {
	return self.Attr("minlength", strconv.Itoa(value))
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#pattern
func (self *TextInputElement) Pattern(value string) *TextInputElement {
	return self.Attr("pattern", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#placeholder
func (self *TextInputElement) Placeholder(value string) *TextInputElement {
	return self.Attr("placeholder", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#readonly
func (self *TextInputElement) Readonly() *TextInputElement {
	return self.Attr("readonly", "")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#size
func (self *TextInputElement) Size(value int) *TextInputElement {
	return self.Attr("size", strconv.Itoa(value))
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#spellcheck
func (self *TextInputElement) Spellcheck(value bool) *TextInputElement {
	return self.Attr("spellcheck", strconv.FormatBool(value))
}

func (self *TextInputElement) Id(value string) *TextInputElement {
	self.element.Id(value)
	return self
}

func (self *TextInputElement) Style(styles ...maps.KeyValue[string, string]) *TextInputElement {
	self.element.Style(styles...)
	return self
}

func (self TextInputElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TextInputElement) Class(classes ...string) *TextInputElement {
	self.element.Class(classes...)
	return self
}

func (self *TextInputElement) Attr(name string, value string) *TextInputElement {
	self.element.Attr(name, value)
	return self
}

func (self TextInputElement) GetAttr(name string) string {
	return self.element.element.attributes.GetOrDefault(name)
}

func (self *TextInputElement) DelAttr(name string) *TextInputElement {
	self.element.DelAttr(name)
	return self
}

func (self TextInputElement) String() string {
	return self.element.String()
}

func (self TextInputElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self TextInputElement) Bytes() []byte {
	return []byte(self.String())
}

func (self TextInputElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
