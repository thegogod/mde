package tokens

type Scanner struct {
	src   []byte
	ln    int
	left  int
	right int
}

func NewScanner(src []byte) *Scanner {
	return &Scanner{
		src:   src,
		ln:    0,
		left:  0,
		right: 0,
	}
}

func (self *Scanner) Scan() (*Token, error) {
	if self.right >= len(self.src) {
		return self.create(Eof), nil
	}

	self.left = self.right
	b := self.src[self.right]
	self.right++

	switch b {
	case ' ':
		return self.create(Space), nil
	case '\n':
		return self.create(NewLine), nil
	case '#':
		token, err := self.onHeading()

		if err == nil {
			return token, nil
		}
	case '-':
		return self.create(Minus), nil
	case '_':
		return self.create(Underscore), nil
	case '*':
		return self.create(Star), nil
	case '`':
		return self.create(BackTick), nil
	case '[':
		return self.create(LeftBracket), nil
	case ']':
		return self.create(RightBracket), nil
	case '(':
		return self.create(LeftParen), nil
	case ')':
		return self.create(RightParen), nil
	case '{':
		return self.create(LeftBrace), nil
	case '}':
		return self.create(RightBrace), nil
	case '!':
		return self.create(Not), nil
	case '>':
		return self.create(Gt), nil
	case '<':
		return self.create(Lt), nil
	case '.':
		return self.create(Dot), nil
	case '@':
		if self.isAlpha(self.peek()) {
			self.right++
			token, err := self.onKeyword()

			if err == nil {
				return token, nil
			}
		}
	default:
		if self.isAlpha(self.peek()) {
			token, err := self.onKeyword()

			if err == nil {
				return token, nil
			}
		}
	}

	return self.create(PlainText), nil
}

func (self *Scanner) onHeading() (*Token, error) {
	i := 1

	for self.peek() == '#' {
		i++
		self.right++
	}

	if self.peek() != ' ' {
		return nil, self.error("expected space")
	}

	self.right++
	self.left = self.right

	for self.right < len(self.src) && self.peek() != '\n' {
		self.right++
	}

	kind := H1

	switch i {
	case 1:
		kind = H1
		break
	case 2:
		kind = H2
		break
	case 3:
		kind = H3
		break
	case 4:
		kind = H4
		break
	case 5:
		kind = H5
		break
	case 6:
		kind = H6
		break
	default:
		return nil, self.error("max heading depth is 6")
	}

	return self.create(kind), nil
}

func (self *Scanner) onKeyword() (*Token, error) {
	for self.isAlpha(self.peek()) || self.isInt(self.peek()) {
		self.right++
	}

	name := self.src[self.left:self.right]

	if kind, ok := Keywords[string(name)]; ok {
		return self.create(kind), nil
	}

	return nil, self.error("keyword not found")
}

func (self Scanner) peek() byte {
	if self.right >= len(self.src) {
		return 0
	}

	return self.src[self.right]
}

func (self Scanner) isInt(b byte) bool {
	return b >= '0' && b <= '9'
}

func (self Scanner) isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b == '_')
}

func (self Scanner) create(kind Kind) *Token {
	return NewToken(
		kind,
		Position{
			Ln:    self.ln,
			Start: self.left,
			End:   self.right,
		},
		self.src[self.left:self.right],
	)
}

func (self Scanner) error(message string) error {
	return NewError(
		Position{
			Ln:    self.ln,
			Start: self.left,
			End:   self.right,
		},
		message,
	)
}
