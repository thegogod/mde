package tokens

import "fmt"

type Position struct {
	Path  string
	Ln    int
	Start int
	End   int
}

func (self Position) IsValid() bool {
	return self.Ln > -1 && self.Start <= self.End
}

func (self Position) String() string {
	return fmt.Sprintf(
		"--> %s:%d:%d-%d",
		self.Path,
		self.Ln+1,
		self.Start+1,
		self.End+1,
	)
}
