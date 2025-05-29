package html

import (
	"strings"
)

// https://react.dev/reference/react/Fragment
type FragmentElement []Node

// https://react.dev/reference/react/Fragment
func Fragment(children ...Node) *FragmentElement {
	self := &FragmentElement{}
	return self.Add(children...)
}

func (self *FragmentElement) Add(children ...Node) *FragmentElement {
	*self = append(*self, children...)
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
	return []byte(self.PrettyBytes(indent))
}
