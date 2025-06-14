package reflect

type NumberType struct{}

func (self NumberType) Name() string {
	return "number"
}

func (self NumberType) String() string {
	return "number"
}

func (self NumberType) AssignableTo(t Type) bool {
	switch t.(type) {
	case NumberType:
		return true
	default:
		return false
	}
}

func (self NumberType) ConvertableTo(t Type) bool {
	switch t.(type) {
	case NumberType:
		return true
	case StringType:
		return true
	default:
		return false
	}
}
