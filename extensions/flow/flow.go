package frontmatter

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/extensions/flow/syntax"
)

type Flow struct{}

func New() Flow {
	return Flow{}
}

func (self Flow) Name() string {
	return "flow"
}

func (self Flow) Tokenizers() []core.Tokenizer {
	return []core.Tokenizer{
		{
			Id:   core.True,
			Name: "true",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				for ptr.Peek() >= 'a' && ptr.Peek() <= 'z' {
					ptr.Next()
				}

				if string(ptr.Bytes()) != "true" {
					return nil, nil
				}

				return ptr.Done(core.True), nil
			},
		},
		{
			Id:   core.False,
			Name: "false",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				for ptr.Peek() >= 'a' && ptr.Peek() <= 'z' {
					ptr.Next()
				}

				if string(ptr.Bytes()) != "false" {
					return nil, nil
				}

				return ptr.Done(core.False), nil
			},
		},
		{
			Id:   core.Null,
			Name: "null",
			Scan: func(ptr *core.Pointer) (*core.Token, error) {
				for ptr.Peek() >= 'a' && ptr.Peek() <= 'z' {
					ptr.Next()
				}

				if string(ptr.Bytes()) != "null" {
					return nil, nil
				}

				return ptr.Done(core.Null), nil
			},
		},
	}
}

func (self Flow) Syntax() []core.Syntax {
	return []core.Syntax{
		syntax.IfStatement{},

		syntax.PrimaryExpression{},
	}
}
