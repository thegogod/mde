package reflect

import (
	"fmt"
	"strconv"
)

type BoolValue struct {
	value bool
}

func Bool(value bool) BoolValue {
	return BoolValue{value}
}

func (self BoolValue) Type() Type {
	return BoolType{}
}

func (self BoolValue) Any() any {
	return self.value
}

func (self BoolValue) String() string {
	return strconv.FormatBool(self.value)
}

func (self BoolValue) Bool() bool {
	return self.value
}

func (self BoolValue) Equals(value Value) bool {
	switch v := value.(type) {
	case BoolValue:
		return self.value == v.value
	default:
		return false
	}
}

func (self BoolValue) Convert(t Type) Value {
	switch t.(type) {
	case BoolType:
		return self
	case StringType:
		return String(self.String())
	default:
		panic(fmt.Sprintf("value of type '%s' is not convertable to type '%s'", t.Name(), self.Type().Name()))
	}
}
