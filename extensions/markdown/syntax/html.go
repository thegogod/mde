package syntax

import (
	"bytes"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Html struct {
	_path []string
}

func (self Html) IsBlock() bool {
	return true
}

func (self Html) IsInline() bool {
	return true
}

func (self Html) Name() string {
	return "html"
}

func (self Html) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Match(core.LessThan)
}

func (self *Html) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	if self._path == nil {
		self._path = []string{}
	}

	name := []byte{}
	iter.NextWhile(core.Space, core.Tab)

	for iter.Match(core.Text) || iter.Match(core.Underscore) || iter.Match(core.Dash) {
		name = append(name, iter.Prev().Render()...)
	}

	el := html.Elem(string(name))
	self._path = append(self._path, string(name))
	depth := len(self._path)

	for iter.NextWhile(core.Space, core.Tab) > 0 && iter.Curr().Kind() == core.Text {
		attr := []byte{}
		value := []byte{}

		for iter.Match(core.Text) || iter.Match(core.Underscore) || iter.Match(core.Dash) {
			attr = append(attr, iter.Prev().Render()...)
		}

		if iter.Match(core.Equals) {
			if _, err := iter.Consume(core.DoubleQuote, "expected '\"'"); err != nil {
				return el, err
			}

			v, err := parser.ParseTextUntil(core.DoubleQuote, parser, iter)

			if v == nil {
				return el, iter.Curr().Error("expected closing '\"'")
			}

			if err != nil {
				return el, err
			}

			value = v
		}

		el.SetAttr(string(attr), string(value))
	}

	isVoid := false

	if iter.Curr().Kind() == core.Slash {
		isVoid = true
		el.Void()
		iter.Next()
	}

	if _, err := iter.Consume(core.GreaterThan, "expected closing '>'"); err != nil {
		return el, err
	}

	if !isVoid {
		for {
			iter.NextWhile(core.NewLine, core.Tab)

			if self.parseClosingTag(iter, name, depth) {
				break
			}

			content, err := parser.ParseInline(parser, iter)

			if content == nil {
				return el, iter.Curr().Error("expected closing tag")
			}

			if err != nil {
				return el, err
			}

			el.Push(content)
		}
	}

	self._path = self._path[:len(self._path)-1]
	return el, nil
}

func (self Html) parseClosingTag(iter *core.Iterator, name []byte, depth int) bool {
	if !iter.Match(core.LessThan, core.Slash) {
		return false
	}

	iter.Save()
	iter.NextWhile(core.Space, core.Tab)
	tag := []byte{}

	for iter.Match(core.Text) || iter.Match(core.Underscore) || iter.Match(core.Dash) {
		tag = append(tag, iter.Prev().Render()...)
	}

	if !bytes.Equal(tag, name) {
		iter.RevertAndPop()
		return false
	}

	iter.NextWhile(core.Space, core.Tab)

	if !iter.Match(core.GreaterThan) || depth != len(self._path) {
		iter.RevertAndPop()
		return false
	}

	iter.Pop()
	return true
}
