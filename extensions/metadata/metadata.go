package metadata

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/extensions/metadata/syntax"
)

type MetaData struct{}

func New() MetaData {
	return MetaData{}
}

func (self MetaData) Name() string {
	return "metadata"
}

func (self MetaData) Tokenizers() []core.Tokenizer {
	return []core.Tokenizer{}
}

func (self MetaData) Syntax() []core.Syntax {
	return []core.Syntax{syntax.Header{}}
}
