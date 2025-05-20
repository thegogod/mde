package markdown

import mde "github.com/thegogod/mde/core"

type Token struct {
	Kind     Kind
	Position mde.Position
	Value    []byte
}

func NewToken(kind Kind, position mde.Position, value []byte) *Token {
	return &Token{
		Kind:     kind,
		Position: position,
		Value:    value,
	}
}

func (self Token) GetKind() uint8 {
	return uint8(self.Kind)
}

func (self Token) GetPosition() mde.Position {
	return self.Position
}

func (self Token) GetBytes() []byte {
	return self.Value
}

func (self Token) String() string {
	return string(self.Value)
}
