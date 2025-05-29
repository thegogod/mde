package html

import (
	"strconv"

	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file
type FileInputElement struct {
	element *InputElement
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file
func FileInput() *FileInputElement {
	return &FileInputElement{Input("file")}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#name
func (self *FileInputElement) Name(value string) *FileInputElement {
	return self.Attr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *FileInputElement) Disabled() *FileInputElement {
	return self.Attr("disabled", "true")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *FileInputElement) Enabled() *FileInputElement {
	return self.Attr("disabled", "false")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#form
func (self *FileInputElement) Form(formId string) *FileInputElement {
	return self.Attr("form", formId)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file#value
func (self *FileInputElement) Value(value string) *FileInputElement {
	return self.Attr("value", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file#accept
func (self *FileInputElement) Accept(value string) *FileInputElement {
	return self.Attr("accept", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file#capture
func (self *FileInputElement) Capture(value string) *FileInputElement {
	return self.Attr("capture", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file#multiple
func (self *FileInputElement) Multiple(value bool) *FileInputElement {
	return self.Attr("multiple", strconv.FormatBool(value))
}

func (self *FileInputElement) Id(value string) *FileInputElement {
	self.element.Id(value)
	return self
}

func (self *FileInputElement) Style(styles ...maps.KeyValue[string, string]) *FileInputElement {
	self.element.Style(styles...)
	return self
}

func (self *FileInputElement) Class(classes ...string) *FileInputElement {
	self.element.Class(classes...)
	return self
}

func (self *FileInputElement) Attr(name string, value string) *FileInputElement {
	self.element.Attr(name, value)
	return self
}

func (self FileInputElement) GetAttr(name string) string {
	return self.element.element.attributes.GetOrDefault(name)
}

func (self *FileInputElement) DelAttr(name string) *FileInputElement {
	self.element.DelAttr(name)
	return self
}

func (self FileInputElement) String() string {
	return self.element.String()
}

func (self FileInputElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self FileInputElement) Bytes() []byte {
	return []byte(self.String())
}

func (self FileInputElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
