package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/extensions/markdown/syntax"
)

type Markdown struct{}

func New() Markdown {
	return Markdown{}
}

func (self Markdown) Name() string {
	return "markdown"
}

func (self Markdown) Tokenizers() []core.Tokenizer {
	return []core.Tokenizer{
		{
			Id:   core.Eof,
			Name: "end-of-file",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != 0 {
					return nil, nil
				}

				return ptr.Done(core.Eof), nil
			},
		},
		{
			Id:   core.Space,
			Name: "space",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != ' ' {
					return nil, nil
				}

				return ptr.Done(core.Space), nil
			},
		},
		{
			Id:   core.NewLine,
			Name: "newline",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '\n' {
					return nil, nil
				}

				return ptr.Done(core.NewLine), nil
			},
		},
		{
			Id:   core.Tab,
			Name: "tab",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '\t' {
					return nil, nil
				}

				return ptr.Done(core.Tab), nil
			},
		},
		{
			Id:   core.Colon,
			Name: "colon",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != ':' {
					return nil, nil
				}

				return ptr.Done(core.Colon), nil
			},
		},
		{
			Id:   core.NotEquals,
			Name: "not-equals",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '!' || ptr.Peek() != '=' {
					return nil, nil
				}

				ptr.Next()
				return ptr.Done(core.NotEquals), nil
			},
		},
		{
			Id:   core.Bang,
			Name: "bang",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '!' {
					return nil, nil
				}

				return ptr.Done(core.Bang), nil
			},
		},
		{
			Id:   core.Hash,
			Name: "hash",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '#' {
					return nil, nil
				}

				return ptr.Done(core.Hash), nil
			},
		},
		{
			Id:   core.At,
			Name: "at",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '@' {
					return nil, nil
				}

				return ptr.Done(core.At), nil
			},
		},
		{
			Id:   core.LeftBracket,
			Name: "left-bracket",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '[' {
					return nil, nil
				}

				return ptr.Done(core.LeftBracket), nil
			},
		},
		{
			Id:   core.RightBracket,
			Name: "right-bracket",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != ']' {
					return nil, nil
				}

				return ptr.Done(core.RightBracket), nil
			},
		},
		{
			Id:   core.LeftParen,
			Name: "left-paren",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '(' {
					return nil, nil
				}

				return ptr.Done(core.LeftParen), nil
			},
		},
		{
			Id:   core.RightParen,
			Name: "right-paren",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != ')' {
					return nil, nil
				}

				return ptr.Done(core.RightParen), nil
			},
		},
		{
			Id:   core.LeftBrace,
			Name: "left-brace",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '{' {
					return nil, nil
				}

				return ptr.Done(core.LeftBrace), nil
			},
		},
		{
			Id:   core.RightBrace,
			Name: "right-brace",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '}' {
					return nil, nil
				}

				return ptr.Done(core.RightBrace), nil
			},
		},
		{
			Id:   core.Asterisk,
			Name: "asterisk",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '*' {
					return nil, nil
				}

				return ptr.Done(core.Asterisk), nil
			},
		},
		{
			Id:   core.Plus,
			Name: "plus",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '+' {
					return nil, nil
				}

				return ptr.Done(core.Plus), nil
			},
		},
		{
			Id:   core.Percent,
			Name: "percent",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '%' {
					return nil, nil
				}

				return ptr.Done(core.Percent), nil
			},
		},

		{
			Id:   core.Dash,
			Name: "dash",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '-' {
					return nil, nil
				}

				return ptr.Done(core.Dash), nil
			},
		},
		{
			Id:   core.Underscore,
			Name: "underscore",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '_' {
					return nil, nil
				}

				return ptr.Done(core.Underscore), nil
			},
		},
		{
			Id:   core.Tilde,
			Name: "tilde",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '~' {
					return nil, nil
				}

				return ptr.Done(core.Tilde), nil
			},
		},
		{
			Id:   core.EqualsEquals,
			Name: "equals-equals",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '=' || ptr.Peek() != '=' {
					return nil, nil
				}

				ptr.Next()
				return ptr.Done(core.EqualsEquals), nil
			},
		},
		{
			Id:   core.Equals,
			Name: "equals",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '=' {
					return nil, nil
				}

				return ptr.Done(core.Equals), nil
			},
		},
		{
			Id:   core.GreaterThanEquals,
			Name: "greater-than-equals",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '>' || ptr.Peek() != '=' {
					return nil, nil
				}

				ptr.Next()
				return ptr.Done(core.GreaterThanEquals), nil
			},
		},
		{
			Id:   core.GreaterThan,
			Name: "greater-than",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '>' {
					return nil, nil
				}

				if ptr.Peek() == ' ' {
					ptr.Next()
				}

				return ptr.Done(core.GreaterThan), nil
			},
		},
		{
			Id:   core.LessThanEquals,
			Name: "less-than-equals",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '<' || ptr.Peek() != '=' {
					return nil, nil
				}

				ptr.Next()
				return ptr.Done(core.LessThanEquals), nil
			},
		},
		{
			Id:   core.LessThan,
			Name: "less-than",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '<' {
					return nil, nil
				}

				if ptr.Peek() == ' ' {
					ptr.Next()
				}

				return ptr.Done(core.LessThan), nil
			},
		},
		{
			Id:   core.Quote,
			Name: "quote",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '\'' {
					return nil, nil
				}

				return ptr.Done(core.Quote), nil
			},
		},
		{
			Id:   core.DoubleQuote,
			Name: "double-quote",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '"' {
					return nil, nil
				}

				return ptr.Done(core.DoubleQuote), nil
			},
		},
		{
			Id:   core.BackQuote,
			Name: "back-quote",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '`' {
					return nil, nil
				}

				return ptr.Done(core.BackQuote), nil
			},
		},
		{
			Id:   core.Period,
			Name: "period",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '.' {
					return nil, nil
				}

				return ptr.Done(core.Period), nil
			},
		},
		{
			Id:   core.Or,
			Name: "or",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '|' || ptr.Peek() != '|' {
					return nil, nil
				}

				ptr.Next()
				return ptr.Done(core.Or), nil
			},
		},
		{
			Id:   core.Pipe,
			Name: "pipe",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '|' {
					return nil, nil
				}

				return ptr.Done(core.Pipe), nil
			},
		},
		{
			Id:   core.And,
			Name: "and",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '&' || ptr.Peek() != '&' {
					return nil, nil
				}

				ptr.Next()
				return ptr.Done(core.And), nil
			},
		},
		{
			Id:   core.Ampersand,
			Name: "ampersand",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '&' {
					return nil, nil
				}

				return ptr.Done(core.Ampersand), nil
			},
		},
		{
			Id:   core.Slash,
			Name: "slash",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '/' {
					return nil, nil
				}

				return ptr.Done(core.Slash), nil
			},
		},
		{
			Id:   core.BackSlash,
			Name: "back-slash",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() != '\\' {
					return nil, nil
				}

				return ptr.Done(core.BackSlash), nil
			},
		},
		{
			Id:   core.Decimal,
			Name: "decimal",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() < '0' || ptr.Curr() > '9' {
					return nil, nil
				}

				for ptr.Peek() >= '0' && ptr.Peek() <= '9' {
					ptr.Next()
				}

				if ptr.Peek() != '.' {
					return nil, nil
				}

				if ptr.Peek() < '0' || ptr.Peek() > '9' {
					return nil, nil
				}

				for ptr.Peek() >= '0' && ptr.Peek() <= '9' {
					ptr.Next()
				}

				return ptr.Done(core.Decimal), nil
			},
		},
		{
			Id:   core.Integer,
			Name: "integer",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				if ptr.Curr() < '0' || ptr.Curr() > '9' {
					return nil, nil
				}

				for ptr.Peek() >= '0' && ptr.Peek() <= '9' {
					ptr.Next()
				}

				return ptr.Done(core.Integer), nil
			},
		},
		{
			Id:   core.Text,
			Name: "text",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				for (ptr.Peek() >= 'a' && ptr.Peek() <= 'z') || (ptr.Peek() >= 'A' && ptr.Peek() <= 'Z') {
					ptr.Next()
				}

				return ptr.Done(core.Text), nil
			},
		},
	}
}

func (self Markdown) Syntax() []core.Syntax {
	return []core.Syntax{
		&syntax.Html{},
		&syntax.H1{},
		&syntax.H2{},
		&syntax.H3{},
		&syntax.H4{},
		&syntax.H5{},
		&syntax.H6{},
		&syntax.HorizontalRule{},
		&syntax.CodeBlock{},
		&syntax.BlockQuote{},
		&syntax.UnOrderedList{},
		&syntax.OrderedList{},
		&syntax.Table{},
		&syntax.Paragraph{},

		&syntax.Bold{},
		&syntax.BoldAlt{},
		&syntax.Italic{},
		&syntax.ItalicAlt{},
		&syntax.Strike{},
		&syntax.StrikeAlt{},
		&syntax.Highlight{},
		&syntax.Code{},
		&syntax.Emoji{},
		&syntax.Link{},
		&syntax.Url{},
		&syntax.Image{},
		&syntax.Break{},
		&syntax.ListItem{},
		&syntax.NewLine{},
	}
}
