package html

import "github.com/thegogod/mde/maps"

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta
type MetaElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta
func Meta() *MetaElement {
	return &MetaElement{Elem("meta").Void()}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#charset
func (self *MetaElement) WithCharset(value string) *MetaElement {
	return self.WithAttr("charset", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#name
func (self *MetaElement) WithName(value string) *MetaElement {
	return self.WithAttr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#content
func (self *MetaElement) WithContent(value string) *MetaElement {
	return self.WithAttr("content", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#media
func (self *MetaElement) WithMedia(value string) *MetaElement {
	return self.WithAttr("media", value)
}

func (self MetaElement) GetTag() string {
	return self.element.GetTag()
}

func (self *MetaElement) WithAttr(name string, value string) *MetaElement {
	self.element.WithAttr(name, value)
	return self
}

func (self MetaElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self MetaElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *MetaElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *MetaElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *MetaElement) WithId(value string) *MetaElement {
	self.element.WithId(value)
	return self
}

func (self MetaElement) HasId() bool {
	return self.element.HasId()
}

func (self MetaElement) GetId() string {
	return self.element.GetId()
}

func (self *MetaElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *MetaElement) DelId() {
	self.element.DelId()
}

func (self *MetaElement) WithClass(classes ...string) *MetaElement {
	self.element.WithClass(classes...)
	return self
}

func (self MetaElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self MetaElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *MetaElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *MetaElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *MetaElement) WithStyles(styles ...maps.KeyValue[string, string]) *MetaElement {
	self.element.WithStyles(styles...)
	return self
}

func (self MetaElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *MetaElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *MetaElement) WithStyle(name string, value string) *MetaElement {
	self.element.WithStyle(name, value)
	return self
}

func (self MetaElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self MetaElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *MetaElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *MetaElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
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

func (self *MetaElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *MetaElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
