package reflect

import (
	"strings"
)

type UnionType []Type

func (self UnionType) Name() string {
	return self.String()
}

func (self UnionType) String() string {
	names := []string{}

	for _, t := range self {
		names = append(names, t.Name())
	}

	return strings.Join(names, " | ")
}

func (self UnionType) AssignableTo(t Type) bool {
	for _, ut := range self {
		if ut.AssignableTo(t) {
			return true
		}
	}

	return false
}

func (self UnionType) ConvertableTo(t Type) bool {
	for _, ut := range self {
		if ut.ConvertableTo(t) {
			return true
		}
	}

	return false
}
