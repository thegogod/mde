package reflect

type BoolType struct{}

func (self BoolType) Name() string {
	return "bool"
}

func (self BoolType) String() string {
	return "bool"
}

func (self BoolType) AssignableTo(t Type) bool {
	switch t.(type) {
	case BoolType:
		return true
	default:
		return false
	}
}

func (self BoolType) ConvertableTo(t Type) bool {
	switch t.(type) {
	case BoolType:
		return true
	case StringType:
		return true
	default:
		return false
	}
}
