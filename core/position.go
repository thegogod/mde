package core

import "fmt"

type Position struct {
	Ln    int
	Col   int
	Start int
	End   int

	save *Position
}

func (self Position) IsValid() bool {
	return self.Ln > -1 && self.Start <= self.End
}

func (self *Position) Save() {
	self.save = &Position{
		Ln:    self.Ln,
		Col:   self.Col,
		Start: self.Start,
		End:   self.End,
	}
}

func (self *Position) Revert() {
	if self.save == nil {
		return
	}

	self.Ln = self.save.Ln
	self.Col = self.save.Col
	self.Start = self.save.Start
	self.End = self.save.End
	self.save = nil
}

func (self Position) String() string {
	return fmt.Sprintf(
		"%d:%d",
		self.Ln+1,
		self.Col+1,
	)
}
