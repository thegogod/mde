package html

import "strings"

type Raw []byte

func (self Raw) String() string {
	return string(self)
}

func (self Raw) PrettyString(indent string) string {
	lines := strings.Split(string(self), "\n")
	return strings.Join(lines, "\n"+indent)
}

func (self Raw) Bytes() []byte {
	return self
}

func (self Raw) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}
