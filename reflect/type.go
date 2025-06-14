package reflect

type Type interface {
	Name() string
	String() string
	Assignable(Type) bool
	Convertable(Type) bool
}
