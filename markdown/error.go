package markdown

import (
	mde "github.com/thegogod/mde/core"
	"github.com/thegogod/mde/strings"
)

type Error struct {
	Position mde.Position
	Message  string
}

func NewError(position mde.Position, message string) *Error {
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
