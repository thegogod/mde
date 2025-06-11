package core

type Scanner struct {
	_ptr   *Pointer
	_saves []Position
	_types []Tokenizer
}

func NewScanner(src []byte) *Scanner {
	return &Scanner{
		_ptr:   Ptr(src),
		_saves: []Position{},
		_types: []Tokenizer{},
	}
}

func (self *Scanner) Extend(extension Extension) *Scanner {
	self._types = append(self._types, extension.Tokenizers()...)
	return self
}

func (self Scanner) Pointer() *Pointer {
	return self._ptr
}

func (self *Scanner) Save() {
	if self._saves == nil {
		self._saves = []Position{}
	}

	self._saves = append(self._saves, self._ptr.End)
}

func (self *Scanner) Pop() {
	if len(self._saves) == 0 {
		return
	}

	self._saves = self._saves[:len(self._saves)-1]
}

func (self *Scanner) Revert() {
	if self._saves == nil || len(self._saves) == 0 {
		return
	}

	self._ptr.End = self._saves[len(self._saves)-1]
}

func (self *Scanner) RevertAndPop() {
	self.Revert()
	self.Pop()
}

func (self *Scanner) Scan() (*Token, error) {
	self._ptr.Start = self._ptr.End
	self._ptr.Next()
	self.Save()

	for _, tokenType := range self._types {
		token, err := tokenType.Scan(self._ptr)

		if token != nil || err != nil {
			self.Pop()
			return token, err
		}

		self.Revert()
	}

	self.Pop()
	return nil, nil
}
