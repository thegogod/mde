package syntax

import (
	"fmt"
	"strings"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Url struct{}

func (self Url) IsBlock() bool {
	return false
}

func (self Url) IsInline() bool {
	return true
}

func (self Url) Name() string {
	return "url"
}

func (self Url) Select(parser core.Parser, iter *core.Iterator) bool {
	iter.Save()
	text, err := parser.ParseText(parser, iter)
	iter.RevertAndPop()

	if text == nil || err != nil {
		return false
	}

	return strings.HasPrefix(string(text), "http")
}

func (self Url) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	link := html.A()
	protocol, _ := parser.ParseText(parser, iter)

	if _, err := iter.Consume(core.Colon, "expected ':'"); err != nil {
		return link, err
	}

	if _, err := iter.Consume(core.Slash, "expected '/'"); err != nil {
		return link, err
	}

	if _, err := iter.Consume(core.Slash, "expected '/'"); err != nil {
		return link, err
	}

	path := []byte{}

	for {
		part, err := parser.ParseText(parser, iter)

		if part == nil || err != nil {
			return link, err
		}

		path = append(path, part...)

		if iter.Curr().Kind() != core.Period {
			break
		}
	}

	url := fmt.Sprintf("%s://%s", protocol, path)

	link.WithHref(url)
	link.Push(url)
	return link, nil
}
