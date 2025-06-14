package reflect

type Value interface {
	Type() Type
	Any() any
	String() string
	Equals(Value) bool
	Convert(Type) Value
}
