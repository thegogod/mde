package core

import "fmt"

// Represents a 2D location
type Position struct {
	Ln    int
	Col   int
	Start int
	End   int

	_saves []Position
}

func (self Position) IsValid() bool {
	return self.Ln > -1 && self.Start <= self.End
}

func (self *Position) Save() {
	if self._saves == nil {
		self._saves = []Position{}
	}

	self._saves = append(self._saves, Position{
		Ln:    self.Ln,
		Col:   self.Col,
		Start: self.Start,
		End:   self.End,
	})
}

func (self *Position) Pop() {
	if len(self._saves) == 0 {
		return
	}

	self._saves = self._saves[:len(self._saves)-1]
}

func (self *Position) Revert() {
	if self._saves == nil || len(self._saves) == 0 {
		return
	}

	self.Ln = self._saves[len(self._saves)-1].Ln
	self.Col = self._saves[len(self._saves)-1].Col
	self.Start = self._saves[len(self._saves)-1].Start
	self.End = self._saves[len(self._saves)-1].End
}

func (self *Position) RevertAndPop() {
	self.Revert()
	self.Pop()
}

func (self Position) String() string {
	return fmt.Sprintf(
		"%d:%d",
		self.Ln+1,
		self.Col+1,
	)
}
