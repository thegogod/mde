package html

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
type ScriptElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
func Script(children ...any) *ScriptElement {
	return &ScriptElement{Elem("script").Add(children...)}
}

func (self *ScriptElement) Src(value string) *ScriptElement {
	return self.Attr("src", value)
}

func (self *ScriptElement) Type(value string) *ScriptElement {
	return self.Attr("type", value)
}

func (self *ScriptElement) Async() *ScriptElement {
	return self.Attr("async", "")
}

func (self *ScriptElement) Defer() *ScriptElement {
	return self.Attr("defer", "")
}

func (self *ScriptElement) Attr(name string, value string) *ScriptElement {
	self.element.Attr(name, value)
	return self
}

func (self *ScriptElement) Add(children ...any) *ScriptElement {
	self.element.Add(children...)
	return self
}

func (self ScriptElement) String() string {
	return self.element.String()
}

func (self ScriptElement) PrettyString(indent string) string {
	return self.element.PrettyString(indent)
}

func (self ScriptElement) Bytes() []byte {
	return []byte(self.String())
}

func (self ScriptElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
