package core

import "slices"

type Iterator struct {
	BlockQuoteDepth int
	ListDepth       int

	_scanner *Scanner
	_prev    *Token
	_curr    *Token
	_saves   []Iterator
}

func Iter(scanner *Scanner) *Iterator {
	return &Iterator{
		_scanner: scanner,
		_saves:   []Iterator{},
	}
}

func (self Iterator) Pointer() *Pointer {
	return self._scanner.Pointer()
}

func (self Iterator) Curr() *Token {
	return self._curr
}

func (self Iterator) Prev() *Token {
	return self._prev
}

func (self *Iterator) Next() bool {
	self._prev = self._curr
	token, err := self._scanner.Scan()

	if err != nil {
		return self.Next()
	}

	self._curr = token

	if token.Kind() == 0 {
		return false
	}

	return true
}

func (self *Iterator) NextWhile(kind ...rune) int {
	i := 0

	for slices.Contains(kind, self._curr.Kind()) {
		i++
		self.Next()
	}

	return i
}

func (self *Iterator) Match(kind ...rune) bool {
	self.Save()

	for _, k := range kind {
		if self._curr.Kind() != k {
			self.RevertAndPop()
			return false
		}

		self.Next()
	}

	self.Pop()
	return true
}

func (self *Iterator) MatchCount(kind rune, count int) bool {
	self.Save()

	for range count {
		if !self.Match(kind) {
			self.RevertAndPop()
			return false
		}
	}

	if self.Match(kind) {
		self.RevertAndPop()
		return false
	}

	self.Pop()
	return true
}

func (self *Iterator) MatchBytes(value ...byte) bool {
	self.Save()
	i := 0

	for i < len(value) {
		for _, b := range self._curr.value {
			if i >= len(value) {
				break
			}

			if b != value[i] {
				self.RevertAndPop()
				return false
			}

			i++
		}

		if !self.Next() {
			self.RevertAndPop()
			return false
		}
	}

	self.Pop()
	return true
}

func (self *Iterator) Consume(kind rune, message string) (*Token, error) {
	if self._curr.Kind() == kind {
		self.Next()
		return self._prev, nil
	}

	return nil, self._curr.Error(message)
}

func (self *Iterator) Save() {
	if self._saves == nil {
		self._saves = []Iterator{}
	}

	self._scanner.Save()
	self._saves = append(self._saves, Iterator{
		BlockQuoteDepth: self.BlockQuoteDepth,
		ListDepth:       self.ListDepth,
		_prev:           self._prev,
		_curr:           self._curr,
		_scanner:        self._scanner,
	})
}

func (self *Iterator) Pop() {
	self._saves = self._saves[:len(self._saves)-1]
	self._scanner.Pop()
}

func (self *Iterator) Revert() {
	if self._saves == nil || len(self._saves) == 0 {
		return
	}

	i := len(self._saves) - 1
	self.BlockQuoteDepth = self._saves[i].BlockQuoteDepth
	self.ListDepth = self._saves[i].ListDepth
	self._prev = self._saves[i]._prev
	self._curr = self._saves[i]._curr
	self._scanner.Revert()
}

func (self *Iterator) RevertAndPop() {
	self.Revert()
	self.Pop()
}
