package core

import "fmt"

// Represents a 2D location
type Position struct {
	Ln    int
	Col   int
	Index int

	_saves []Position
}

func (self Position) String() string {
	return fmt.Sprintf(
		"%d:%d",
		self.Ln+1,
		self.Col+1,
	)
}
