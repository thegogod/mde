package markdown

import mde "github.com/thegogod/mde/core"

type Token struct {
	TokenKind TokenKind
	Position  mde.Position
	Value     []byte
}

func NewToken(TokenKind TokenKind, position mde.Position, value []byte) *Token {
	return &Token{
		TokenKind: TokenKind,
		Position:  position,
		Value:     value,
	}
}

func (self Token) GetTokenKind() uint8 {
	return uint8(self.TokenKind)
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
