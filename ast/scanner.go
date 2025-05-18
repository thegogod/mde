package ast

import (
	"slices"
)

type Scanner struct {
	path  string
	src   []byte
	ln    int
	left  int
	right int
}

func NewScanner(path string, src []byte) *Scanner {
	return &Scanner{
		path:  path,
		src:   src,
		ln:    0,
		left:  0,
		right: 0,
	}
}

func (self *Scanner) Next() (*Token, error) {
	if self.right >= len(self.src) {
		return self.create(Eof), nil
	}

	self.left = self.right
	b := self.src[self.right]
	self.right++

	switch b {
	case ' ':
	case '\r':
	case '\t':
		// ignore whitespace
		break
	case '\n':
		self.ln++
		break
	case '(':
		return self.create(LeftParen), nil
	case ')':
		return self.create(RightParen), nil
	case '{':
		return self.create(LeftBrace), nil
	case '}':
		return self.create(RightBrace), nil
	case '[':
		return self.create(LeftBracket), nil
	case ']':
		return self.create(RightBracket), nil
	case ',':
		return self.create(Comma), nil
	case '.':
		return self.create(Dot), nil
	case ':':
		if self.peek() == ':' {
			self.right++
			return self.create(DoubleColon), nil
		}

		return self.create(Colon), nil
	case ';':
		return self.create(SemiColon), nil
	case '?':
		return self.create(QuestionMark), nil
	case '@':
		return self.create(AtSymbol), nil
	case '#':
		return self.create(HashSymbol), nil
	case '_':
		return self.create(Underscore), nil
	case '|':
		if self.peek() != '|' {
			return nil, self.error("expected '|'")
		}

		self.right++
		return self.create(Or), nil
	case '&':
		if self.peek() != '&' {
			return nil, self.error("expected '&'")
		}

		self.right++
		return self.create(And), nil
	case '+':
		if self.peek() == '=' {
			self.right++
			return self.create(PlusEq), nil
		}

		return self.create(Plus), nil
	case '-':
		if self.peek() == '=' {
			self.right++
			return self.create(MinusEq), nil
		} else if self.isInt(self.peek()) {
			self.right++
			return self.onNumeric()
		}

		return self.create(Minus), nil
	case '*':
		if self.peek() == '=' {
			self.right++
			return self.create(StarEq), nil
		}

		return self.create(Star), nil
	case '/':
		if self.peek() == '/' {
			return self.onComment()
		} else if self.peek() == '=' {
			self.right++
			return self.create(SlashEq), nil
		}

		return self.create(Slash), nil
	case '!':
		if self.peek() == '=' {
			self.right++
			return self.create(NotEq), nil
		}

		return self.create(Not), nil
	case '=':
		if self.peek() == '=' {
			self.right++
			return self.create(EqEq), nil
		}

		return self.create(Eq), nil
	case '>':
		if self.peek() == '=' {
			self.right++
			return self.create(GtEq), nil
		}

		return self.create(Gt), nil
	case '<':
		if self.peek() == '=' {
			self.right++
			return self.create(LtEq), nil
		}

		return self.create(Lt), nil
	case '\'':
		return self.onByte()
	case '"':
		return self.onString()
	default:
		if self.isInt(b) {
			return self.onNumeric()
		} else if self.isAlpha(b) {
			return self.onIdentifier()
		}

		return nil, self.error("unexpected character")
	}

	return self.Next()
}

func (self *Scanner) onComment() (*Token, error) {
	for self.peek() != '\n' && self.peek() != 0 {
		self.right++
	}

	self.ln++
	self.right++
	return self.Next()
}

func (self *Scanner) onByte() (*Token, error) {
	self.right++

	if self.peek() != '\'' {
		return nil, self.error("unterminated byte")
	}

	self.left++
	token := self.create(LByte)
	self.right++
	return token, nil
}

func (self *Scanner) onString() (*Token, error) {
	for self.peek() != '"' && self.peek() != 0 {
		if self.peek() == '\n' {
			self.ln++
		} else if self.peek() == '\\' {
			err := self.onEscape()

			if err != nil {
				return nil, err
			}
		}

		self.right++
	}

	if self.right == len(self.src) {
		return nil, self.error("unterminated string")
	}

	self.left++
	token := self.create(LString)
	self.right++
	return token, nil
}

func (self *Scanner) onEscape() error {
	self.right++

	defer func() {
		self.right--
	}()

	switch self.peek() {
	case 'a': // bell
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\a')
	case 'b': // backspace
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\b')
	case 'f': // form feed
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\f')
	case 'n': // new line
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\n')
	case 'r': // carriage return
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\r')
	case 't': // horizontal tab
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\t')
	case 'v': // verical tab
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\v')
	case '\'': // single quote
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\'')
	case '"': // double quote
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '"')
	case '\\': // back slash
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\\')
	default:
		return self.error("unknown escape sequence")
	}

	return nil
}

func (self *Scanner) onNumeric() (*Token, error) {
	kind := LInt

	for self.isInt(self.peek()) {
		self.right++
	}

	if self.peek() == '.' {
		kind = LFloat
		self.right++

		for self.isInt(self.peek()) {
			self.right++
		}
	}

	return self.create(kind), nil
}

func (self *Scanner) onIdentifier() (*Token, error) {
	for self.isAlpha(self.peek()) || self.isInt(self.peek()) {
		self.right++
	}

	name := self.src[self.left:self.right]

	if kind, ok := Keywords[string(name)]; ok {
		return self.create(kind), nil
	}

	return self.create(Identifier), nil
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

func (self Scanner) create(kind TokenKind) *Token {
	return NewToken(
		kind,
		self.path,
		self.ln,
		self.left,
		self.right,
		self.src[self.left:self.right],
	)
}

func (self Scanner) error(message string) error {
	return NewError(
		self.path,
		self.ln,
		self.left,
		self.right,
		message,
	)
}
