package markdown

import (
	"github.com/google/uuid"
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseTask(iter *tokens.Iterator) (core.Node, error) {
	id := uuid.NewString()
	label := html.Label().For(id)
	input := html.CheckBoxInput().Id(id)
	_, err := iter.Consume(tokens.LeftBracket, "expected '['")

	if err != nil {
		return label, err
	}

	checked, err := iter.Consume(tokens.Text, "expected ' ' or 'x'")

	if err != nil {
		return label, err
	}

	if checked.String() != " " && checked.String() != "x" {
		return label, iter.Curr().Error("expected ' ' or 'x'")
	}

	if checked.String() == "x" {
		input.Checked(true)
	}

	_, err = iter.Consume(tokens.RightBracket, "expected ']'")

	if err != nil {
		return label, err
	}

	space, err := iter.Consume(tokens.Text, "expected ' '")

	if err != nil {
		return label, err
	}

	if space.String() != " " {
		return label, iter.Curr().Error("expected ' '")
	}

	text := ""

	for !iter.Match(tokens.NewLine) {
		node, err := self.parseText(iter)

		if err != nil {
			return label, err
		}

		if node == nil {
			break
		}

		text += string(node)
	}

	label.Push(input)
	label.Push(html.Span(text))
	return label, nil
}
