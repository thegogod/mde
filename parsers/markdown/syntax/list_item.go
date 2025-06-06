package syntax

import (
	"github.com/google/uuid"
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type ListItem struct{}

func (self ListItem) IsBlock() bool {
	return false
}

func (self ListItem) IsInline() bool {
	return false
}

func (self ListItem) Name() string {
	return "list_item"
}

func (self ListItem) Select(parser core.Parser, iter core.Iterator) bool {
	return false
}

func (self ListItem) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	it := iter.(*tokens.Iterator)
	li := html.Li()
	iter.Save()
	node, err := self.parseTask(parser, iter)

	if err == nil && node != nil {
		li.Push(node)
		iter.Pop()
		return li, nil
	}

	iter.Revert()

	for iter.Curr().IsInline() {
		node, err := parser.ParseInline(iter)

		if err != nil {
			iter.Revert()
			iter.Pop()
			return li, err
		}

		if node == nil {
			break
		}

		if node.String() == "\n" {
			if !iter.MatchCount(tokens.Tab, it.ListDepth) {
				break
			}

			node, err = nil, nil
			iter.Save()

			if iter.Match(tokens.Integer) && iter.Match(tokens.Period) && iter.Match(tokens.Space) {
				node, err = parser.ParseSyntax("ordered_list", iter)
			} else if iter.Match(tokens.Dash) && iter.Match(tokens.Space) {
				node, err = parser.ParseSyntax("unordered_list", iter)
			}

			if node != nil && err == nil {
				li.Push(node)
			} else {
				iter.Revert()
			}

			iter.Pop()
			break
		}

		li.Push(node)
	}

	iter.Pop()
	return li, nil
}

func (self ListItem) parseTask(parser core.Parser, iter core.Iterator) (core.Node, error) {
	id := uuid.NewString()
	label := html.Label().For(id)
	input := html.CheckBoxInput().Id(id)

	if _, err := iter.Consume(tokens.LeftBracket, "expected '['"); err != nil {
		return label, err
	}

	checked, err := iter.Consume(tokens.Space, "expected ' ' or 'x'")

	if err != nil {
		checked, err = iter.Consume(tokens.Text, "expected ' ' or 'x'")

		if err != nil {
			return label, err
		}
	}

	if checked.String() != " " && checked.String() != "x" {
		return label, iter.Curr().Error("expected ' ' or 'x'")
	}

	if checked.String() == "x" {
		input.Checked(true)
	}

	if _, err = iter.Consume(tokens.RightBracket, "expected ']'"); err != nil {
		return label, err
	}

	if _, err = iter.Consume(tokens.Space, "expected ' '"); err != nil {
		return label, err
	}

	text := ""

	for !iter.Match(tokens.NewLine) {
		node, err := parser.ParseText(iter)

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
