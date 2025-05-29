package html

import (
	"fmt"
	"strings"

	"github.com/thegogod/mde/core"
)

type Element struct {
	kind        string
	selfClosing bool
	attributes  core.OMap[string, string]
	children    []Node
}

func Elem(kind string) *Element {
	return &Element{
		kind:       kind,
		attributes: core.OMap[string, string]{},
		children:   []Node{},
	}
}

func (self *Element) Id(value string) *Element {
	return self.Attr("id", value)
}

func (self *Element) Style(styles ...core.KeyValue[string, string]) *Element {
	values := []string{}

	for _, style := range styles {
		values = append(values, fmt.Sprintf("%s: %s;", style.Key, style.Value))
	}

	return self.Attr("style", strings.Join(values, ""))
}

func (self *Element) Class(classes ...string) *Element {
	return self.Attr("class", strings.Join(classes, " "))
}

func (self *Element) SelfClosing() *Element {
	self.selfClosing = true
	return self
}

func (self *Element) Attr(name string, value string) *Element {
	if self.attributes == nil {
		self.attributes = core.OMap[string, string]{}
	}

	self.attributes.Set(name, value)
	return self
}

func (self *Element) Add(children ...any) *Element {
	if self.children == nil {
		self.children = []Node{}
	}

	for _, child := range children {
		switch v := child.(type) {
		case Node:
			self.children = append(self.children, v)
			break
		case string:
			self.children = append(self.children, Raw(v))
			break
		case *string:
			self.children = append(self.children, Raw(*v))
			break
		default:
			panic("invalid type")
		}
	}

	return self
}

func (self Element) String() string {
	html := "<" + self.kind

	for _, attr := range self.attributes {
		html += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Value)
	}

	if self.selfClosing {
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

	if self.selfClosing {
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
	return []byte(self.PrettyBytes(indent))
}
