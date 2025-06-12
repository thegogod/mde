package html

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/style
type StyleElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/style
func Style(children ...any) *StyleElement {
	return &StyleElement{Elem("style").Attr("type", "text/css").Push(children...)}
}

func (self *StyleElement) Attr(name string, value string) *StyleElement {
	self.element.Attr(name, value)
	return self
}

func (self StyleElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *StyleElement) DelAttr(name string) *StyleElement {
	self.element.DelAttr(name)
	return self
}

func (self *StyleElement) Push(children ...any) *StyleElement {
	self.element.Push(children...)
	return self
}

func (self *StyleElement) Pop() *StyleElement {
	self.element.Pop()
	return self
}

func (self StyleElement) Children() []Node {
	return self.element.children
}

func (self StyleElement) String() string {
	return self.element.String()
}

func (self StyleElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self StyleElement) Bytes() []byte {
	return []byte(self.String())
}

func (self StyleElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self StyleElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self StyleElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
