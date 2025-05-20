package tokens

import "strconv"

type Token struct {
	Kind     Kind
	Position Position
	Value    []byte
}

func NewToken(kind Kind, position Position, value []byte) *Token {
	return &Token{
		Kind:     kind,
		Position: position,
		Value:    value,
	}
}

func (self Token) Byte() byte {
	return self.Value[0]
}

func (self Token) String() string {
	return string(self.Value)
}

func (self Token) Int() (int, error) {
	return strconv.Atoi(string(self.Value))
}

func (self Token) Float() (float64, error) {
	return strconv.ParseFloat(string(self.Value), 64)
}

func (self Token) Bool() (bool, error) {
	return strconv.ParseBool(string(self.Value))
}
