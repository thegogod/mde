package html

import (
	"fmt"
	"strings"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

type Element struct {
	kind       string
	void       bool // https://developer.mozilla.org/en-US/docs/Glossary/Void_element
	attributes maps.OMap[string, string]
	children   []core.Node
}

func Elem(kind string) *Element {
	return &Element{
		kind:       kind,
		attributes: maps.OMap[string, string]{},
		children:   []core.Node{},
	}
}

func (self *Element) Id(value string) *Element {
	return self.Attr("id", value)
}

func (self *Element) Style(styles ...maps.KeyValue[string, string]) *Element {
	values := []string{}

	for _, style := range styles {
		values = append(values, fmt.Sprintf("%s: %s;", style.Key, style.Value))
	}

	return self.Attr("style", strings.Join(values, ""))
}

func (self *Element) Class(classes ...string) *Element {
	return self.Attr("class", strings.Join(classes, " "))
}

func (self *Element) Void() *Element {
	self.void = true
	return self
}

func (self *Element) Attr(name string, value string) *Element {
	if self.attributes == nil {
		self.attributes = maps.OMap[string, string]{}
	}

	self.attributes.Set(name, value)
	return self
}

func (self Element) GetAttr(name string) string {
	return self.attributes.GetOrDefault(name)
}

func (self *Element) DelAttr(name string) *Element {
	self.attributes.Del(name)
	return self
}

func (self *Element) Push(children ...any) *Element {
	if self.children == nil {
		self.children = []core.Node{}
	}

	for _, child := range children {
		switch v := child.(type) {
		case core.Node:
			self.children = append(self.children, v)
			break
		case string:
			self.children = append(self.children, Raw(v))
			break
		case *string:
			self.children = append(self.children, Raw(*v))
			break
		case []byte:
			self.children = append(self.children, Raw(v))
			break
		default:
			panic("invalid type")
		}
	}

	if len(self.children) > 1 {
		curr, isCurrRaw := self.children[len(self.children)-1].(Raw)
		prev, isPrevRaw := self.children[len(self.children)-2].(Raw)

		if isCurrRaw && isPrevRaw {
			prev = append(prev, curr...)
			self.Pop()
			self.children[len(self.children)-1] = prev
		}
	}

	return self
}

func (self *Element) Pop() *Element {
	self.children = self.children[:len(self.children)-1]
	return self
}

func (self Element) Children() []core.Node {
	return self.children
}

func (self Element) String() string {
	html := "<" + self.kind

	for _, attr := range self.attributes {
		html += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Value)
	}

	if self.void {
		html += " />"
		return html
	}

	html += ">"

	for _, child := range self.children {
		html += child.String()
	}

	html += fmt.Sprintf("</%s>", self.kind)
	return html
}

func (self Element) PrettyString(indent string) string {
	html := "<" + self.kind

	for _, attr := range self.attributes {
		html += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Value)
	}

	if self.void {
		html += " />"
		return html
	}

	html += ">"

	for _, child := range self.children {
		lines := strings.Split(child.PrettyString(indent), "\n")

		if len(lines) == 3 {
			html += "\n" + indent + child.String()
			continue
		}

		html += "\n" + indent + strings.Join(lines, "\n"+indent)
	}

	html += fmt.Sprintf("\n</%s>", self.kind)
	return html
}

func (self Element) Bytes() []byte {
	return []byte(self.String())
}

func (self Element) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
