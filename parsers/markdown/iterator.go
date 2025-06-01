package markdown

import "github.com/thegogod/mde/core"

type Iterator struct {
	Prev *Token
	Curr *Token

	scanner         *Scanner
	blockQuoteDepth int
	listDepth       int
	saves           []Iterator
}

func (self *Iterator) Reset(src []byte) {
	self.Prev = nil
	self.Curr = nil
	self.blockQuoteDepth = 0
	self.listDepth = 0
	self.scanner = NewScanner(src)
	self.saves = []Iterator{}
}

func (self *Iterator) Next() bool {
	self.Prev = self.Curr
	token, err := self.scanner.Scan()

	if err != nil {
		return self.Next()
	}

	self.Curr = token.(*Token)

	if TokenKind(token.GetKind()) == Eof {
		return false
	}

	return true
}

func (self *Iterator) Match(kind TokenKind) bool {
	if self.Curr.Kind != kind {
		return false
	}

	self.Next()
	return true
}

func (self *Iterator) MatchCount(kind TokenKind, count int) bool {
	self.Save()

	for range count {
		if !self.Match(kind) {
			self.Revert()
			self.Pop()
			return false
		}
	}

	self.Pop()
	return true
}

func (self *Iterator) Consume(kind TokenKind, message string) (core.Token, error) {
	if self.Curr.Kind == kind {
		self.Next()
		return self.Prev, nil
	}

	return nil, NewError(self.Curr.Position, message)
}

func (self *Iterator) Save() {
	if self.saves == nil {
		self.saves = []Iterator{}
	}

	self.scanner.pos.Save()
	self.saves = append(self.saves, Iterator{
		Prev:            self.Prev,
		Curr:            self.Curr,
		blockQuoteDepth: self.blockQuoteDepth,
		listDepth:       self.listDepth,
		scanner:         self.scanner,
	})
}

func (self *Iterator) Revert() {
	if self.saves == nil || len(self.saves) == 0 {
		return
	}

	i := len(self.saves) - 1
	self.Prev = self.saves[i].Prev
	self.Curr = self.saves[i].Curr
	self.blockQuoteDepth = self.saves[i].blockQuoteDepth
	self.listDepth = self.saves[i].listDepth
	self.scanner.pos.Revert()
}

func (self *Iterator) Pop() {
	self.saves = self.saves[:len(self.saves)-1]
	self.scanner.pos.Pop()
}
