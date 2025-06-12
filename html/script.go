package html

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
type ScriptElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
func Script(children ...any) *ScriptElement {
	return &ScriptElement{Elem("script").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script#src
func (self *ScriptElement) Src(value string) *ScriptElement {
	return self.Attr("src", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script/type
func (self *ScriptElement) Type(value string) *ScriptElement {
	return self.Attr("type", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script#async
func (self *ScriptElement) Async() *ScriptElement {
	return self.Attr("async", "")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script#defer
func (self *ScriptElement) Defer() *ScriptElement {
	return self.Attr("defer", "")
}

func (self *ScriptElement) Attr(name string, value string) *ScriptElement {
	self.element.Attr(name, value)
	return self
}

func (self ScriptElement) GetAttr(name string) string {
	return self.element.attributes.GetOrDefault(name)
}

func (self *ScriptElement) DelAttr(name string) *ScriptElement {
	self.element.DelAttr(name)
	return self
}

func (self *ScriptElement) Push(children ...any) *ScriptElement {
	self.element.Push(children...)
	return self
}

func (self *ScriptElement) Pop() *ScriptElement {
	self.element.Pop()
	return self
}

func (self ScriptElement) Children() []Node {
	return self.element.children
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
	return []byte(self.PrettyString(indent))
}

func (self ScriptElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self ScriptElement) GetByClass(classes ...string) []Node {
	return self.element.GetByClass(classes...)
}
