package reflect

type StringType struct{}

func (self StringType) Name() string {
	return "string"
}

func (self StringType) String() string {
	return "string"
}

func (self StringType) AssignableTo(t Type) bool {
	switch t.(type) {
	case StringType:
		return true
	default:
		return false
	}
}

func (self StringType) ConvertableTo(t Type) bool {
	switch t.(type) {
	case StringType:
		return true
	default:
		return false
	}
}
