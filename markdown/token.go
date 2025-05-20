package markdown

import "github.com/thegogod/mde/core"

type Token struct {
	TokenKind TokenKind
	Position  core.Position
	Value     []byte
}

func NewToken(TokenKind TokenKind, position core.Position, value []byte) *Token {
	return &Token{
		TokenKind: TokenKind,
		Position:  position,
		Value:     value,
	}
}

func (self Token) GetTokenKind() uint8 {
	return uint8(self.TokenKind)
}

func (self Token) GetPosition() core.Position {
	return self.Position
}

func (self Token) GetBytes() []byte {
	return self.Value
}

func (self Token) String() string {
	return string(self.Value)
}
