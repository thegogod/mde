package core

import "fmt"

type Position struct {
	Ln    int
	Col   int
	Start int
	End   int
	saves []Position
}

func (self Position) IsValid() bool {
	return self.Ln > -1 && self.Start <= self.End
}

func (self *Position) Save() {
	if self.saves == nil {
		self.saves = []Position{}
	}

	self.saves = append(self.saves, Position{
		Ln:    self.Ln,
		Col:   self.Col,
		Start: self.Start,
		End:   self.End,
	})
}

func (self *Position) Revert() {
	if self.saves == nil || len(self.saves) == 0 {
		return
	}

	self.Ln = self.saves[len(self.saves)-1].Ln
	self.Col = self.saves[len(self.saves)-1].Col
	self.Start = self.saves[len(self.saves)-1].Start
	self.End = self.saves[len(self.saves)-1].End
}

func (self *Position) Pop() {
	self.saves = self.saves[:len(self.saves)-1]
}

func (self Position) String() string {
	return fmt.Sprintf(
		"%d:%d",
		self.Ln+1,
		self.Col+1,
	)
}
