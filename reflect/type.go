package reflect

type Type interface {
	Name() string
	String() string
	AssignableTo(Type) bool
	ConvertableTo(Type) bool
}
