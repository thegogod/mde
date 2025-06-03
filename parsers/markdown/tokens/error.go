package tokens

import (
	"strings"

	"github.com/thegogod/mde/core"
)

type Error struct {
	Position core.Position
	Message  string
}

func NewError(position core.Position, message string) *Error {
	return &Error{
		Position: position,
		Message:  message,
	}
}

func (self Error) Error() string {
	return self.String()
}

func (self Error) String() string {
	return strings.Join([]string{self.Position.String(), self.Message}, "\n")
}
