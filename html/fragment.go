package html

import (
	"strings"

	"github.com/thegogod/mde/maps"
)

// https://react.dev/reference/react/Fragment
type FragmentElement []Node

// https://react.dev/reference/react/Fragment
func Fragment(children ...Node) *FragmentElement {
	self := &FragmentElement{}
	return self.Push(children...)
}

func (self FragmentElement) GetSelector() string {
	return ""
}

func (self FragmentElement) HasAttr(name string) bool {
	return false
}

func (self FragmentElement) GetAttr(name string) string {
	return ""
}

func (self FragmentElement) SetAttr(name string, value string) {
	return
}

func (self FragmentElement) DelAttr(name string) {
	return
}

func (self FragmentElement) HasId() bool {
	return false
}

func (self FragmentElement) GetId() string {
	return ""
}

func (self FragmentElement) SetId(id string) {
	return
}

func (self FragmentElement) DelId() {
	return
}

func (self FragmentElement) HasClass(name ...string) bool {
	return false
}

func (self FragmentElement) GetClass() []string {
	return []string{}
}

func (self FragmentElement) AddClass(name ...string) {
	return
}

func (self FragmentElement) DelClass(name ...string) {
	return
}

func (self FragmentElement) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self FragmentElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self FragmentElement) HasStyle(name ...string) bool {
	return false
}

func (self FragmentElement) GetStyle(name string) string {
	return ""
}

func (self FragmentElement) SetStyle(name string, value string) {
	return
}

func (self FragmentElement) DelStyle(name ...string) {
	return
}

func (self FragmentElement) Count() int {
	return len(self)
}

func (self FragmentElement) Children() []Node {
	return self
}

func (self *FragmentElement) Push(children ...Node) *FragmentElement {
	for _, child := range children {
		*self = append(*self, child)
	}

	return self
}

func (self *FragmentElement) Pop() *FragmentElement {
	arr := *self

	if len(arr) == 0 {
		return self
	}

	arr = arr[:len(arr)-1]
	*self = arr
	return self
}

func (self FragmentElement) Render() []byte {
	content := ""

	for _, node := range self {
		if strings.HasPrefix(node.GetSelector(), ":") {
			continue
		}

		content += string(node.Render())
	}

	return []byte(content)
}

func (self FragmentElement) RenderPretty(indent string) []byte {
	content := []string{}

	for _, node := range self {
		if strings.HasPrefix(node.GetSelector(), ":") {
			continue
		}

		content = append(content, string(node.RenderPretty(indent)))
	}

	return []byte(strings.Join(content, "\n"))
}

func (self FragmentElement) GetById(id string) Node {
	for _, child := range self {
		if node := child.GetById(id); node != nil {
			return node
		}
	}

	return nil
}

func (self FragmentElement) Select(query ...any) []Node {
	nodes := []Node{}

	for _, child := range self {
		nodes = append(nodes, child.Select(query...)...)
	}

	return nodes
}
