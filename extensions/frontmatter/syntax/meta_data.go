package syntax

import (
	"github.com/goccy/go-yaml"
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type MetaData struct{}

func (self MetaData) IsBlock() bool {
	return true
}

func (self MetaData) IsInline() bool {
	return false
}

func (self MetaData) Name() string {
	return "metadata"
}

func (self MetaData) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Pointer().Sof() && iter.MatchCount(core.Dash, 3)
}

func (self MetaData) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	el := html.MetaData{}

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return el, err
	}

	data := []byte{}

	for {
		line, err := parser.ParseTextUntil(core.NewLine, parser, iter)

		if line == nil {
			return el, iter.Curr().Error("expected newline at end of key value pair")
		}

		if err != nil {
			return el, err
		}

		data = append(data, line...)

		if iter.MatchCount(core.Dash, 3) {
			break
		}

		data = append(data, '\n')
	}

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return el, err
	}

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return el, err
	}

	if err := yaml.Unmarshal(data, &el); err != nil {
		return el, err
	}

	return el, nil
}
