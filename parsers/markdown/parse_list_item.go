package markdown

import (
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseListItem(iter *tokens.Iterator) (*html.ListItemElement, error) {
	li := html.Li()
	iter.Save()
	node, err := self.parseTask(iter)

	if err == nil && node != nil {
		li.Push(node)
		iter.Pop()
		return li, nil
	}

	iter.Revert()

	for iter.Curr().IsInline() {
		node, err := self.ParseInline(iter)

		if err != nil {
			iter.Revert()
			iter.Pop()
			return li, err
		}

		if node == nil {
			break
		}

		if node.String() == "\n" {
			if !iter.MatchCount(tokens.Tab, iter.ListDepth) {
				break
			}

			node, err = nil, nil
			iter.Save()

			if iter.Match(tokens.Ol) {
				node, err = self.parseOrderedList(iter)
			} else if iter.Match(tokens.Ul) {
				node, err = self.parseUnorderedList(iter)
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
