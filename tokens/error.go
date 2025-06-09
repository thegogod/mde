package tokens

import (
	"fmt"
	"strconv"

	"github.com/thegogod/mde/core"
)

type Error struct {
	Start   core.Position
	End     core.Position
	Message string
}

func NewError(start core.Position, end core.Position, message string) *Error {
	return &Error{
		Start:   start,
		End:     end,
		Message: message,
	}
}

func (self Error) Error() string {
	return self.String()
}

func (self Error) String() string {
	line := strconv.Itoa(self.Start.Ln)

	if self.End.Ln != self.Start.Ln {
		line = fmt.Sprintf("%d-%d", self.Start.Ln, self.End.Ln)
	}

	column := strconv.Itoa(self.Start.Col)

	if self.End.Col != self.Start.Col {
		column = fmt.Sprintf("%d-%d", self.Start.Col, self.End.Col)
	}

	return fmt.Sprintf(
		"[%s:%s]",
		line,
		column,
	)
}
