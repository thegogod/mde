package reflect

import (
	"fmt"
	"strconv"
)

type NumberValue struct {
	value float64
}

func Number(value float64) NumberValue {
	return NumberValue{value}
}

func (self NumberValue) Type() Type {
	return NumberType{}
}

func (self NumberValue) Any() any {
	return self.value
}

func (self NumberValue) String() string {
	return strconv.FormatFloat(self.value, 'f', -1, 64)
}

func (self NumberValue) Int() int64 {
	return int64(self.value)
}

func (self NumberValue) Float() float64 {
	return self.value
}

func (self NumberValue) Equals(value Value) bool {
	switch v := value.(type) {
	case NumberValue:
		return self.value == v.value
	default:
		return false
	}
}

func (self NumberValue) Convert(t Type) Value {
	switch t.(type) {
	case NumberType:
		return self
	case StringType:
		return String(self.String())
	default:
		panic(fmt.Sprintf("value of type '%s' is not convertable to type '%s'", t.Name(), self.Type().Name()))
	}
}
