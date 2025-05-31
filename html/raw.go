package html

import "strings"

type Raw []byte

func (self Raw) String() string {
	value := string(self)

	for strings.HasSuffix(value, "\n") {
		value = strings.TrimSuffix(value, "\n")
	}

	return value
}

func (self Raw) PrettyString(indent string) string {
	value := string(self)

	for strings.HasSuffix(value, "\n") {
		value = strings.TrimSuffix(value, "\n")
	}

	lines := strings.Split(value, "\n")
	return strings.Join(lines, "\n"+indent)
}

func (self Raw) Bytes() []byte {
	return []byte(self.String())
}

func (self Raw) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyBytes(indent))
}
