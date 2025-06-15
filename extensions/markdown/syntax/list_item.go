package syntax

import (
	"github.com/google/uuid"
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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

func (self ListItem) Select(parser core.Parser, iter *core.Iterator) bool {
	return false
}

func (self ListItem) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	li := html.Li()
	iter.Save()
	node, err := self.parseTask(parser, iter)

	if err == nil && node != nil {
		li.Push(node)
		iter.Pop()
		return li, nil
	}

	iter.Revert()

	for iter.Curr().Kind() != core.Eof {
		node, err := parser.ParseInline(parser, iter)

		if err != nil {
			iter.Revert()
			iter.Pop()
			return li, err
		}

		if node == nil {
			break
		}

		if raw, ok := node.(html.Raw); ok && string(raw) == "\n" {
			if !iter.MatchCount(core.Tab, iter.ListDepth) {
				break
			}

			node, err = nil, nil
			iter.Save()

			if iter.Match(core.Integer) && iter.Match(core.Period) && iter.Match(core.Space) {
				node, err = parser.ParseSyntax("markdown", "ordered_list", parser, iter)
			} else if iter.Match(core.Dash) && iter.Match(core.Space) {
				node, err = parser.ParseSyntax("markdown", "unordered_list", parser, iter)
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

func (self ListItem) parseTask(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	id := uuid.NewString()
	label := html.Label().WithFor(id)
	input := html.CheckBoxInput().WithId(id)

	if _, err := iter.Consume(core.LeftBracket, "expected '['"); err != nil {
		return label, err
	}

	checked, err := iter.Consume(core.Space, "expected ' ' or 'x'")

	if err != nil {
		checked, err = iter.Consume(core.Text, "expected ' ' or 'x'")

		if err != nil {
			return label, err
		}
	}

	if checked.String() != " " && checked.String() != "x" {
		return label, iter.Curr().Error("expected ' ' or 'x'")
	}

	if checked.String() == "x" {
		input.WithChecked(true)
	}

	if _, err = iter.Consume(core.RightBracket, "expected ']'"); err != nil {
		return label, err
	}

	if _, err = iter.Consume(core.Space, "expected ' '"); err != nil {
		return label, err
	}

	text := ""

	for !iter.Match(core.NewLine) {
		node, err := parser.ParseText(parser, iter)

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
