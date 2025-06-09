package tokens

import (
	"github.com/thegogod/mde/core"
)

type Iterator struct {
	BlockQuoteDepth int
	ListDepth       int

	scanner *Scanner
	prev    core.Token
	curr    core.Token
	saves   []Iterator
}

func Iter(scanner *Scanner) *Iterator {
	return &Iterator{
		scanner: scanner,
		saves:   []Iterator{},
	}
}

func (self Iterator) Curr() core.Token {
	return self.curr
}

func (self Iterator) Prev() core.Token {
	return self.prev
}

func (self Iterator) Position() core.Position {
	return self.scanner.pos
}

func (self *Iterator) Next() bool {
	self.prev = self.curr
	token, err := self.scanner.Scan()

	if err != nil {
		return self.Next()
	}

	self.curr = token

	if token.Kind() == Eof {
		return false
	}

	return true
}

func (self *Iterator) Match(kind ...byte) bool {
	self.Save()

	for _, k := range kind {
		if self.curr.Kind() != k {
			self.RevertAndPop()
			return false
		}

		self.Next()
	}

	self.Pop()
	return true
}

func (self *Iterator) MatchCount(kind byte, count int) bool {
	self.Save()

	for range count {
		if !self.Match(kind) {
			self.Revert()
			self.Pop()
			return false
		}
	}

	if self.Match(kind) {
		self.Revert()
		self.Pop()
		return false
	}

	self.Pop()
	return true
}

func (self *Iterator) MatchBytes(value ...byte) bool {
	self.Save()
	i := 0

	for i < len(value) {
		for _, b := range self.curr.Bytes() {
			if i >= len(value) {
				break
			}

			if b != value[i] {
				self.Revert()
				self.Pop()
				return false
			}

			i++
		}

		if !self.Next() {
			self.Revert()
			self.Pop()
			return false
		}
	}

	self.Pop()
	return true
}

func (self *Iterator) Consume(kind byte, message string) (core.Token, error) {
	if self.curr.Kind() == kind {
		self.Next()
		return self.prev, nil
	}

	return nil, self.curr.Error(message)
}

func (self *Iterator) Save() {
	if self.saves == nil {
		self.saves = []Iterator{}
	}

	self.scanner.Save()
	self.saves = append(self.saves, Iterator{
		BlockQuoteDepth: self.BlockQuoteDepth,
		ListDepth:       self.ListDepth,
		prev:            self.prev,
		curr:            self.curr,
		scanner:         self.scanner,
	})
}

func (self *Iterator) Pop() {
	self.saves = self.saves[:len(self.saves)-1]
	self.scanner.Pop()
}

func (self *Iterator) Revert() {
	if self.saves == nil || len(self.saves) == 0 {
		return
	}

	i := len(self.saves) - 1
	self.BlockQuoteDepth = self.saves[i].BlockQuoteDepth
	self.ListDepth = self.saves[i].ListDepth
	self.prev = self.saves[i].prev
	self.curr = self.saves[i].curr
	self.scanner.Revert()
}

func (self *Iterator) RevertAndPop() {
	self.Revert()
	self.Pop()
}
