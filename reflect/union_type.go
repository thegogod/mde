package reflect

import (
	"strings"
)

type UnionType []Type

func Union(types ...Type) UnionType {
	return types
}

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

func (self UnionType) Assignable(t Type) bool {
	for _, ut := range self {
		if ut.Assignable(t) {
			return true
		}
	}

	return false
}

func (self UnionType) Convertable(t Type) bool {
	for _, ut := range self {
		if ut.Convertable(t) {
			return true
		}
	}

	return false
}
