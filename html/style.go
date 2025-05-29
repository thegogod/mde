package html

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/style
type StyleElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/style
func Style(children ...any) *StyleElement {
	return &StyleElement{Elem("style").Attr("type", "text/css").Add(children...)}
}

func (self *StyleElement) Attr(name string, value string) *StyleElement {
	self.element.Attr(name, value)
	return self
}

func (self *StyleElement) Add(children ...any) *StyleElement {
	self.element.Add(children...)
	return self
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
	return []byte(self.PrettyBytes(indent))
}
