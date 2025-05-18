package tokens

import (
	"github.com/thegogod/mde/strings"
)

type Error struct {
	Position Position
	Message  string
}

func NewError(position Position, message string) *Error {
	return &Error{
		Position: position,
		Message:  message,
	}
}

func (self Error) Error() string {
	return self.String()
}

func (self Error) String() string {
	builder := strings.Builder().
		AddLine(self.Position.String()).
		AddLine(self.Message)

	return builder.String()
}
