package html

import (
	"fmt"
	"slices"
	"strings"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

type Element struct {
	Kind       string
	void       bool // https://developer.mozilla.org/en-US/docs/Glossary/Void_element
	attributes maps.OMap[string, string]
	children   []core.Node
}

func Elem(kind string) *Element {
	return &Element{
		Kind:       kind,
		attributes: maps.OMap[string, string]{},
		children:   []core.Node{},
	}
}

func (self *Element) Id(value string) *Element {
	return self.Attr("id", value)
}

func (self *Element) Style(styles ...maps.KeyValue[string, string]) *Element {
	value := []string{}
	set := self.GetStyles()
	set.Merge(styles)

	if len(set) == 0 {
		return self.DelAttr("style")
	}

	for _, pair := range set {
		if pair.Key == "" || pair.Value == "" {
			continue
		}

		value = append(value, fmt.Sprintf("%s: %s", pair.Key, pair.Value))
	}

	return self.Attr("style", strings.Join(append(value, ""), ";"))
}

func (self Element) GetStyles() maps.OMap[string, string] {
	styles := maps.OMap[string, string]{}
	existing := self.attributes.GetOrDefault("style")
	lines := strings.SplitSeq(existing, ";")

	for line := range lines {
		parts := strings.Split(line, ": ")

		if len(parts) != 2 {
			continue
		}

		styles.Set(parts[0], parts[1])
	}

	return styles
}

func (self *Element) Class(classes ...string) *Element {
	return self.Attr("class", strings.Join(classes, " "))
}

func (self Element) GetClasses() []string {
	return strings.Split(self.GetAttr("class"), " ")
}

func (self Element) HasClass(classes ...string) bool {
	existing := self.GetClasses()

	for _, cls := range classes {
		if !slices.Contains(existing, cls) {
			return false
		}
	}

	return true
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
		case Host:
			for key, value := range v {
				self.Attr(key, fmt.Sprintf("%v", value))
			}

			break
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
	html := "<" + self.Kind

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

	html += fmt.Sprintf("</%s>", self.Kind)
	return html
}

func (self Element) PrettyString(indent string) string {
	html := "<" + self.Kind

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

	html += fmt.Sprintf("\n</%s>", self.Kind)
	return html
}

func (self Element) Bytes() []byte {
	return []byte(self.String())
}

func (self Element) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self Element) GetById(id string) core.Node {
	if self.GetAttr("id") == id {
		return self
	}

	for _, child := range self.children {
		if node := child.GetById(id); node != nil {
			return node
		}
	}

	return nil
}

func (self Element) GetByClass(classes ...string) []core.Node {
	nodes := []core.Node{}

	if self.HasClass(classes...) {
		nodes = append(nodes, self)
	}

	for _, child := range self.children {
		n := child.GetByClass(classes...)
		nodes = append(nodes, n...)
	}

	return nodes
}
