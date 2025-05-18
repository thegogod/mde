package strings

import (
	"fmt"
	"strings"
)

type StringBuilder []string

func Builder() StringBuilder {
	return StringBuilder{}
}

func (self StringBuilder) Add(values ...any) StringBuilder {
	for _, value := range values {
		if v, ok := value.(string); ok {
			self = append(self, v)
			continue
		}

		if v, ok := value.(fmt.Stringer); ok {
			self = append(self, v.String())
			continue
		}

		if v, ok := value.(StringBuilder); ok {
			self = append(self, v.String())
			continue
		}

		if v, ok := value.(*StringBuilder); ok {
			self = append(self, v.String())
			continue
		}

		panic("unsupported non-stringish type")
	}

	return self
}

func (self StringBuilder) AddLine(values ...any) StringBuilder {
	for _, value := range values {
		if v, ok := value.(string); ok {
			self = append(self, v)
			continue
		}

		if v, ok := value.(fmt.Stringer); ok {
			self = append(self, v.String())
			continue
		}

		if v, ok := value.(StringBuilder); ok {
			self = append(self, v.String())
			continue
		}

		if v, ok := value.(*StringBuilder); ok {
			self = append(self, v.String())
			continue
		}

		panic("unsupported non-stringish type")
	}

	self = append(self, "\n")
	return self
}

func (self StringBuilder) AddDelim(delim string, values ...any) StringBuilder {
	for _, value := range values {
		if v, ok := value.(string); ok {
			self = append(self, v)
			continue
		}

		if v, ok := value.(fmt.Stringer); ok {
			self = append(self, v.String())
			continue
		}

		if v, ok := value.(StringBuilder); ok {
			self = append(self, v.Build(delim))
			continue
		}

		if v, ok := value.(*StringBuilder); ok {
			self = append(self, v.Build(delim))
			continue
		}

		panic("unsupported non-stringish type")
	}

	return self
}

func (self StringBuilder) Build(delim string) string {
	return strings.Join(self, delim)
}

func (self StringBuilder) String() string {
	return strings.Join(self, "")
}
