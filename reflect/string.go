package reflect

import "fmt"

type StringValue struct {
	value string
}

func String(value string) StringValue {
	return StringValue{value}
}

func (self StringValue) Type() Type {
	return StringType{}
}

func (self StringValue) Any() any {
	return self.value
}

func (self StringValue) String() string {
	return self.value
}

func (self StringValue) Equals(value Value) bool {
	switch v := value.(type) {
	case StringValue:
		return self.value == v.value
	default:
		return false
	}
}

func (self StringValue) Convert(t Type) Value {
	switch t.(type) {
	case StringType:
		return self
	default:
		panic(fmt.Sprintf("value of type '%s' is not convertable to type '%s'", t.Name(), self.Type().Name()))
	}
}
