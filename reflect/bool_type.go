package reflect

type BoolType struct{}

func (self BoolType) Name() string {
	return "bool"
}

func (self BoolType) String() string {
	return "bool"
}

func (self BoolType) Assignable(t Type) bool {
	switch t.(type) {
	case BoolType:
		return true
	default:
		return false
	}
}

func (self BoolType) Convertable(t Type) bool {
	switch t.(type) {
	case BoolType:
		return true
	case StringType:
		return true
	default:
		return false
	}
}
