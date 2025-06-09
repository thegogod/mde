package core

import (
	"fmt"
	"strconv"
)

type Error struct {
	_start   Position
	_end     Position
	_message string
}

func Err(start Position, end Position, message string) *Error {
	return &Error{
		_start:   start,
		_end:     end,
		_message: message,
	}
}

func (self Error) Error() string {
	return self.String()
}

func (self Error) String() string {
	line := strconv.Itoa(self._start.Ln)

	if self._end.Ln != self._start.Ln {
		line = fmt.Sprintf("%d-%d", self._start.Ln, self._end.Ln)
	}

	column := strconv.Itoa(self._start.Col)

	if self._end.Col != self._start.Col {
		column = fmt.Sprintf("%d-%d", self._start.Col, self._end.Col)
	}

	return fmt.Sprintf(
		"[%s:%s] => %s",
		line,
		column,
		self._message,
	)
}
