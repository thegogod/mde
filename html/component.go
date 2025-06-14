package html

import "github.com/thegogod/mde/maps"

type Component struct {
	Selector string
	Template string
	Styles   []string

	element *Element
}

func (self *Component) init() {
	if self.element == nil {
		self.element = Elem(self.Selector)
	}
}

func (self Component) GetSelector() string {
	return self.Selector
}

func (self *Component) WithAttr(name string, value string) *Component {
	self.init()
	self.element.WithAttr(name, value)
	return self
}

func (self Component) HasAttr(name string) bool {
	self.init()
	return self.element.HasAttr(name)
}

func (self Component) GetAttr(name string) string {
	self.init()
	return self.element.GetAttr(name)
}

func (self *Component) SetAttr(name string, value string) {
	self.init()
	self.element.SetAttr(name, value)
}

func (self *Component) DelAttr(name string) {
	self.init()
	self.element.DelAttr(name)
}

func (self *Component) WithId(value string) *Component {
	self.init()
	self.element.WithId(value)
	return self
}

func (self Component) HasId() bool {
	self.init()
	return self.element.HasId()
}

func (self Component) GetId() string {
	self.init()
	return self.element.GetId()
}

func (self *Component) SetId(id string) {
	self.init()
	self.element.SetId(id)
}

func (self *Component) DelId() {
	self.init()
	self.element.DelId()
}

func (self *Component) WithClass(classes ...string) *Component {
	self.init()
	self.element.WithClass(classes...)
	return self
}

func (self Component) HasClass(classes ...string) bool {
	self.init()
	return self.element.HasClass(classes...)
}

func (self Component) GetClass() []string {
	self.init()
	return self.element.GetClass()
}

func (self *Component) AddClass(name ...string) {
	self.init()
	self.element.AddClass(name...)
}

func (self *Component) DelClass(name ...string) {
	self.init()
	self.element.DelClass(name...)
}

func (self *Component) WithStyles(styles ...maps.KeyValue[string, string]) *Component {
	self.init()
	self.element.WithStyles(styles...)
	return self
}

func (self Component) GetStyles() maps.OMap[string, string] {
	self.init()
	return self.element.GetStyles()
}

func (self *Component) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.init()
	self.element.SetStyles(styles...)
}

func (self *Component) WithStyle(name string, value string) *Component {
	self.init()
	self.element.WithStyle(name, value)
	return self
}

func (self Component) HasStyle(name ...string) bool {
	self.init()
	return self.element.HasStyle(name...)
}

func (self Component) GetStyle(name string) string {
	self.init()
	return self.element.GetStyle(name)
}

func (self *Component) SetStyle(name string, value string) {
	self.init()
	self.element.SetStyle(name, value)
}

func (self *Component) DelStyle(name ...string) {
	self.init()
	self.element.DelStyle(name...)
}

func (self Component) Count() int {
	self.init()
	return self.element.Count()
}

func (self Component) Children() []Node {
	self.init()
	return self.element.Children()
}

func (self Component) String() string {
	self.init()
	return self.element.String()
}

func (self Component) PrettyString(indent string) string {
	self.init()
	return self.element.PrettyString(indent)
}

func (self Component) Bytes() []byte {
	return []byte(self.String())
}

func (self Component) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self *Component) GetById(id string) Node {
	self.init()
	return self.element.GetById(id)
}

func (self *Component) Select(query ...any) []Node {
	self.init()
	return self.element.Select(query...)
}
