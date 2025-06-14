package reflect

type VoidType struct{}

func (self VoidType) Name() string {
	return "void"
}

func (self VoidType) String() string {
	return "void"
}

func (self VoidType) Assignable(t Type) bool {
	return false
}

func (self VoidType) Convertable(t Type) bool {
	return false
}
