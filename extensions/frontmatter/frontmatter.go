package frontmatter

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/extensions/frontmatter/syntax"
)

type FrontMatter struct{}

func New() FrontMatter {
	return FrontMatter{}
}

func (self FrontMatter) Name() string {
	return "frontmatter"
}

func (self FrontMatter) Tokenizers() []core.Tokenizer {
	return []core.Tokenizer{}
}

func (self FrontMatter) Syntax() []core.Syntax {
	return []core.Syntax{syntax.MetaData{}}
}
