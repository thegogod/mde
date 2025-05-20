package tokens

import "unicode"

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
		if self.peek() == ' ' {
			return self.create(Br), nil
		}

		break
	case '\n':
		return self.create(Space), nil
	case '#':
		token, err := self.onHeading()

		if err == nil {
			return token, nil
		}

		break
	case '-':
		if unicode.IsSpace(rune(self.peek())) {
			self.right++
			return self.create(Ul), nil
		}

		break
	case '_':
		if self.peek() == '_' {
			self.right++
			return self.create(Bold), nil
		}

		return self.create(Italic), nil
	case '*':
		if self.peek() == '*' {
			self.right++
			return self.create(Bold), nil
		}

		return self.create(Italic), nil
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
	case '!':
		return self.create(Bang), nil
	case '>':
		return self.create(Gt), nil
	case '<':
		return self.create(Lt), nil
	default:
		if unicode.IsNumber(rune(b)) && self.peek() == '.' {
			self.right++

			if unicode.IsSpace(rune(self.peek())) {
				return self.create(Ol), nil
			}

			self.right--
		}
	}

	return self.create(Text), nil
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

func (self Scanner) peek() byte {
	if self.right >= len(self.src) {
		return byte(Eof)
	}

	return self.src[self.right]
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
