package syntax

import (
	"fmt"
	"strings"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
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

func (self Url) Select(parser core.Parser, iter core.Iterator) bool {
	iter.Save()
	text, err := parser.ParseText(iter)
	iter.RevertAndPop()

	if text == nil || err != nil {
		return false
	}

	return strings.HasPrefix(string(text), "http")
}

func (self Url) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	link := html.A()
	protocol, _ := parser.ParseText(iter)

	if _, err := iter.Consume(tokens.Colon, "expected ':'"); err != nil {
		return link, err
	}

	if token, err := iter.Consume(tokens.Text, "expected '/'"); err != nil || token.String() != "/" {
		return link, err
	}

	if token, err := iter.Consume(tokens.Text, "expected '/'"); err != nil || token.String() != "/" {
		return link, err
	}

	path := []byte{}

	for {
		part, err := parser.ParseText(iter)

		if part == nil || err != nil {
			return link, err
		}

		path = append(path, part...)

		if iter.Curr().Kind() != tokens.Period {
			break
		}
	}

	url := fmt.Sprintf("%s://%s", protocol, path)

	link.Href(url)
	link.Push(url)
	return link, nil
}
