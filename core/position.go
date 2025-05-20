package mde

import "fmt"

type Position struct {
	Ln    int
	Start int
	End   int
}

func (self Position) IsValid() bool {
	return self.Ln > -1 && self.Start <= self.End
}

func (self Position) String() string {
	return fmt.Sprintf(
		"%d:%d-%d",
		self.Ln+1,
		self.Start+1,
		self.End+1,
	)
}
