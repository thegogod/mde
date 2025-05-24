package markdown

import "github.com/thegogod/mde/core"

type Token struct {
	Kind     TokenKind
	Position core.Position
	Value    []byte
}

func NewToken(kind TokenKind, position core.Position, value []byte) *Token {
	return &Token{
		Kind:     kind,
		Position: position,
		Value:    value,
	}
}

func (self Token) GetKind() uint8 {
	return uint8(self.Kind)
}

func (self Token) GetPosition() core.Position {
	return self.Position
}

func (self Token) GetBytes() []byte {
	return self.Value
}

func (self Token) IsWhitespace() bool {
	value := self.String()
	return value == " " || value == "\n"
}

func (self Token) String() string {
	return string(self.Value)
}
