package html

import (
	"fmt"
	"slices"
	"strings"
)

type Element struct {
	kind        string
	selfClosing bool
	attributes  []Attribute
	children    []Node
}

func Elem(kind string) *Element {
	return &Element{
		kind:       kind,
		attributes: []Attribute{},
		children:   []Node{},
	}
}

func (self *Element) SelfClosing() *Element {
	self.selfClosing = true
	return self
}

func (self *Element) Attr(name string, value string) *Element {
	if self.attributes == nil {
		self.attributes = []Attribute{}
	}

	i := slices.IndexFunc(self.attributes, func(attr Attribute) bool {
		return attr.Name == name
	})

	if i > -1 {
		self.attributes[i].Value = value
	} else {
		self.attributes = append(self.attributes, Attribute{
			Name:  name,
			Value: value,
		})
	}

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
		html += fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value)
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
		html += fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value)
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
