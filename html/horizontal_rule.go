package html

import "github.com/thegogod/mde/maps"

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/hr
type HorizontalRuleElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/hr
func Hr() *HorizontalRuleElement {
	return &HorizontalRuleElement{Elem("hr").Void()}
}

func (self *HorizontalRuleElement) Id(value string) *HorizontalRuleElement {
	self.element.Id(value)
	return self
}

func (self *HorizontalRuleElement) Style(styles ...maps.KeyValue[string, string]) *HorizontalRuleElement {
	self.element.Style(styles...)
	return self
}

func (self HorizontalRuleElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *HorizontalRuleElement) Class(classes ...string) *HorizontalRuleElement {
	self.element.Class(classes...)
	return self
}

func (self *HorizontalRuleElement) Attr(name string, value string) *HorizontalRuleElement {
	self.element.Attr(name, value)
	return self
}

func (self HorizontalRuleElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *HorizontalRuleElement) DelAttr(name string) *HorizontalRuleElement {
	self.element.DelAttr(name)
	return self
}

func (self HorizontalRuleElement) String() string {
	return self.element.String()
}

func (self HorizontalRuleElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self HorizontalRuleElement) Bytes() []byte {
	return []byte(self.String())
}

func (self HorizontalRuleElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
