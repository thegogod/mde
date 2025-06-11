package syntax

import (
	"bytes"
	"strings"

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
	fragment := html.Fragment()

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return fragment, err
	}

	for {
		line, err := parser.ParseTextUntil(core.NewLine, parser, iter)

		if line == nil {
			return fragment, iter.Curr().Error("expected newline at end of key value pair")
		}

		if err != nil {
			return fragment, err
		}

		parts := bytes.SplitN(line, []byte{':'}, 2)

		if len(parts) != 2 {
			return fragment, iter.Curr().Error("expected keyvalue pair separated by ':'")
		}

		meta := html.Meta()
		meta = meta.Name(strings.TrimSpace(string(parts[0])))
		meta = meta.Content(strings.TrimSpace(string(parts[1])))
		fragment.Push(meta)

		if iter.MatchCount(core.Dash, 3) {
			break
		}
	}

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return fragment, err
	}

	if _, err := iter.Consume(core.NewLine, "expected newline"); err != nil {
		return fragment, err
	}

	return fragment, nil
}
