package reflect

import (
	"fmt"
)

type NullValue struct{}

func Null() NullValue {
	return NullValue{}
}

func (self NullValue) Type() Type {
	return VoidType{}
}

func (self NullValue) Any() any {
	return nil
}

func (self NullValue) String() string {
	return "null"
}

func (self NullValue) Equals(value Value) bool {
	switch value.(type) {
	case NullValue:
		return true
	default:
		return false
	}
}

func (self NullValue) Convert(t Type) Value {
	panic(fmt.Sprintf("value of type '%s' is not convertable to type '%s'", t.Name(), self.Type().Name()))
}
