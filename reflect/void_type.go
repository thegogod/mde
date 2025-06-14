package reflect

type VoidType struct{}

func (self VoidType) Name() string {
	return "void"
}

func (self VoidType) String() string {
	return "void"
}

func (self VoidType) AssignableTo(t Type) bool {
	return false
}

func (self VoidType) ConvertableTo(t Type) bool {
	return false
}
