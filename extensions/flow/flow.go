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
	return []core.Tokenizer{}
}

func (self Flow) Syntax() []core.Syntax {
	return []core.Syntax{syntax.If{}}
}
