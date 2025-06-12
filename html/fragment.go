package html

import (
	"strings"
)

// https://react.dev/reference/react/Fragment
type FragmentElement []Node

// https://react.dev/reference/react/Fragment
func Fragment(children ...Node) *FragmentElement {
	self := &FragmentElement{}
	return self.Push(children...)
}

func (self *FragmentElement) Push(children ...Node) *FragmentElement {
	for _, child := range children {
		if _, ok := child.(Host); ok {
			continue
		}

		*self = append(*self, child)
	}

	return self
}

func (self *FragmentElement) Pop() *FragmentElement {
	arr := *self
	arr = arr[:len(arr)-1]
	*self = arr
	return self
}

func (self FragmentElement) Children() []Node {
	return self
}

func (self FragmentElement) String() string {
	content := ""

	for _, node := range self {
		content += node.String()
	}

	return content
}

func (self FragmentElement) PrettyString(indent string) string {
	content := []string{}

	for _, node := range self {
		content = append(content, node.PrettyString(indent))
	}

	return strings.Join(content, "\n")
}

func (self FragmentElement) Bytes() []byte {
	return []byte(self.String())
}

func (self FragmentElement) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self FragmentElement) GetById(id string) Node {
	for _, child := range self {
		if node := child.GetById(id); node != nil {
			return node
		}
	}

	return nil
}

func (self FragmentElement) GetByClass(classes ...string) []Node {
	nodes := []Node{}

	for _, child := range self {
		n := child.GetByClass(classes...)
		nodes = append(nodes, n...)
	}

	return nodes
}
