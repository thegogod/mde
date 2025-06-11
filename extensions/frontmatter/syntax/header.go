package syntax

import (
	"github.com/goccy/go-yaml"
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Header struct{}

func (self Header) IsBlock() bool {
	return true
}

func (self Header) IsInline() bool {
	return false
}

func (self Header) Name() string {
	return "header"
}

func (self Header) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Pointer().Sof() && iter.MatchCount(core.Dash, 3)
}

func (self Header) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	host := html.Host{}

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return host, err
	}

	data := []byte{}

	for {
		line, err := parser.ParseTextUntil(core.NewLine, parser, iter)

		if line == nil {
			return host, iter.Curr().Error("expected newline at end of key value pair")
		}

		if err != nil {
			return host, err
		}

		data = append(data, line...)

		if iter.MatchCount(core.Dash, 3) {
			break
		}

		data = append(data, '\n')
	}

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return host, err
	}

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return host, err
	}

	if err := yaml.Unmarshal(data, &host); err != nil {
		return host, err
	}

	return host, nil
}
