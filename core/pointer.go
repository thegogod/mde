package core

type Pointer struct {
	Src   []byte
	Start Position
	End   Position
}

func Ptr(src []byte) *Pointer {
	return &Pointer{
		Src:   src,
		Start: Position{},
		End:   Position{},
	}
}

func (self Pointer) Peek() byte {
	if self.End.Index >= len(self.Src) {
		return byte(0)
	}

	return self.Src[self.End.Index]
}

func (self *Pointer) Next() byte {
	self.End.Index++
	self.End.Col++

	if self.Peek() == '\n' {
		self.End.Ln++
		self.End.Col = 0
	}

	return self.Peek()
}

func (self Pointer) Error(message string) error {
	return Err(self.Start, self.End, message)
}

func (self Pointer) Bytes() []byte {
	return self.Src[self.Start.Index:self.End.Index]
}
