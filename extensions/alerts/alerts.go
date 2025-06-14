package alerts

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/extensions/alerts/syntax"
)

type Alerts struct{}

func New() Alerts {
	return Alerts{}
}

func (self Alerts) Name() string {
	return "alerts"
}

func (self Alerts) Tokenizers() []core.Tokenizer {
	return []core.Tokenizer{}
}

func (self Alerts) Syntax() []core.Syntax {
	return []core.Syntax{syntax.Alert{}}
}
