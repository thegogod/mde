package html

import (
	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a
type AnchorElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a
func A(children ...any) *AnchorElement {
	return &AnchorElement{Elem("a").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#href
func (self *AnchorElement) WithHref(value string) *AnchorElement {
	return self.WithAttr("href", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#target
func (self *AnchorElement) WithTarget(value string) *AnchorElement {
	return self.WithAttr("target", value)
}

func (self AnchorElement) GetTag() string {
	return self.element.GetTag()
}

func (self *AnchorElement) WithAttr(name string, value string) *AnchorElement {
	self.element.WithAttr(name, value)
	return self
}

func (self AnchorElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self AnchorElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *AnchorElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *AnchorElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *AnchorElement) WithId(value string) *AnchorElement {
	self.element.WithId(value)
	return self
}

func (self AnchorElement) HasId() bool {
	return self.element.HasId()
}

func (self AnchorElement) GetId() string {
	return self.element.GetId()
}

func (self *AnchorElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *AnchorElement) DelId() {
	self.element.DelId()
}

func (self *AnchorElement) WithClass(classes ...string) *AnchorElement {
	self.element.WithClass(classes...)
	return self
}

func (self AnchorElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self AnchorElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *AnchorElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *AnchorElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *AnchorElement) WithStyles(styles ...maps.KeyValue[string, string]) *AnchorElement {
	self.element.WithStyles(styles...)
	return self
}

func (self AnchorElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *AnchorElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *AnchorElement) WithStyle(name string, value string) *AnchorElement {
	self.element.WithStyle(name, value)
	return self
}

func (self AnchorElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self AnchorElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *AnchorElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *AnchorElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self AnchorElement) Count() int {
	return self.element.Count()
}

func (self AnchorElement) Children() []Node {
	return self.element.Children()
}

func (self *AnchorElement) Push(children ...any) *AnchorElement {
	self.element.Push(children...)
	return self
}

func (self *AnchorElement) Pop() *AnchorElement {
	self.element.Pop()
	return self
}

func (self AnchorElement) String() string {
	return self.element.String()
}

func (self AnchorElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self AnchorElement) Bytes() []byte {
	return []byte(self.String())
}

func (self AnchorElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self *AnchorElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *AnchorElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
