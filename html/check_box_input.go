package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/checkbox
type CheckBoxInputElement struct {
	element *InputElement
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/checkbox
func CheckBoxInput() *CheckBoxInputElement {
	return &CheckBoxInputElement{Input("checkbox")}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#name
func (self *CheckBoxInputElement) Name(value string) *CheckBoxInputElement {
	return self.Attr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *CheckBoxInputElement) Disabled() *CheckBoxInputElement {
	return self.Attr("disabled", "true")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *CheckBoxInputElement) Enabled() *CheckBoxInputElement {
	return self.Attr("disabled", "false")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#form
func (self *CheckBoxInputElement) Form(formId string) *CheckBoxInputElement {
	return self.Attr("form", formId)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/checkbox#value
func (self *CheckBoxInputElement) Value(value string) *CheckBoxInputElement {
	return self.Attr("value", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/checkbox#checked
func (self *CheckBoxInputElement) Checked(value bool) *CheckBoxInputElement {
	if value {
		return self.Attr("checked", "true")
	}

	return self.DelAttr("checked")
}

func (self *CheckBoxInputElement) Id(value string) *CheckBoxInputElement {
	self.element.Id(value)
	return self
}

func (self *CheckBoxInputElement) Style(styles ...maps.KeyValue[string, string]) *CheckBoxInputElement {
	self.element.Style(styles...)
	return self
}

func (self CheckBoxInputElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *CheckBoxInputElement) Class(classes ...string) *CheckBoxInputElement {
	self.element.Class(classes...)
	return self
}

func (self CheckBoxInputElement) GetClasses() []string {
	return self.element.GetClasses()
}

func (self CheckBoxInputElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self *CheckBoxInputElement) Attr(name string, value string) *CheckBoxInputElement {
	self.element.Attr(name, value)
	return self
}

func (self CheckBoxInputElement) GetAttr(name string) string {
	return self.element.element.attributes.GetOrDefault(name)
}

func (self *CheckBoxInputElement) DelAttr(name string) *CheckBoxInputElement {
	self.element.DelAttr(name)
	return self
}

func (self CheckBoxInputElement) String() string {
	return self.element.String()
}

func (self CheckBoxInputElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self CheckBoxInputElement) Bytes() []byte {
	return []byte(self.String())
}

func (self CheckBoxInputElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self CheckBoxInputElement) GetById(id string) core.Node {
	return self.element.GetById(id)
}

func (self CheckBoxInputElement) GetByClass(classes ...string) []core.Node {
	return self.element.GetByClass(classes...)
}
