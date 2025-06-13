package syntax

import (
	"bytes"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Table struct{}

func (self Table) IsBlock() bool {
	return true
}

func (self Table) IsInline() bool {
	return false
}

func (self Table) Name() string {
	return "table"
}

func (self Table) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(core.Pipe, 1)
}

func (self Table) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	table := html.Table()
	columns := []*html.TableCellElement{}

	for !iter.Match(core.NewLine) {
		th := html.Th()

		for {
			if !iter.Match(core.Space) && !iter.Match(core.Tab) {
				break
			}
		}

		buff := []byte{}

		for !iter.Match(core.Pipe) {
			node, err := parser.ParseInline(parser, iter)

			if node == nil {
				return table, iter.Curr().Error("invalid table header")
			}

			if err != nil {
				return table, err
			}

			if _, ok := node.(*html.BreakLineElement); ok {
				continue
			}

			if raw, ok := node.(html.Raw); ok && (raw.String() == " " || raw.String() == "\t") {
				buff = append(buff, raw.Bytes()...)
				continue
			}

			if len(buff) > 0 {
				th.Push(buff)
				buff = []byte{}
			}

			th.Push(node)
		}

		columns = append(columns, th)
	}

	if _, err := iter.Consume(core.Pipe, "expected opening '|'"); err != nil {
		return table, err
	}

	for i := range len(columns) {
		node, err := parser.ParseTextUntil(core.Pipe, parser, iter)

		if node == nil {
			return table, iter.Curr().Error("expected closing '|'")
		}

		if err != nil {
			return table, err
		}

		node = bytes.TrimSpace(node)
		dashes := 0

		for _, b := range node {
			if b == '-' {
				dashes++
			}

			if b != ':' && b != '-' {
				return table, iter.Curr().Error("invalid header separator")
			}
		}

		if dashes < 3 {
			return table, iter.Curr().Error("invalid header seperator")
		}

		leftAlign := bytes.HasPrefix(node, []byte{':'})
		rightAlign := bytes.HasSuffix(node, []byte{':'})

		if leftAlign && rightAlign {
			columns[i].SetStyle("text-align", "center")
		} else if leftAlign {
			columns[i].SetStyle("text-align", "left")
		} else if rightAlign {
			columns[i].SetStyle("text-align", "right")
		}
	}

	if _, err := iter.Consume(core.NewLine, "expected new line"); err != nil {
		return table, err
	}

	table.Push(html.THead(html.Tr(columns...)))
	rows := []*html.TableCellElement{}

	for iter.Match(core.Pipe) {
		for i := range len(columns) {
			td := html.Td()

			for {
				if !iter.Match(core.Space) && !iter.Match(core.Tab) {
					break
				}
			}

			buff := []byte{}

			for !iter.Match(core.Pipe) {
				node, err := parser.ParseInline(parser, iter)

				if node == nil {
					return table, iter.Curr().Error("expected closing '|'")
				}

				if err != nil {
					return table, err
				}

				if _, ok := node.(*html.BreakLineElement); ok {
					continue
				}

				if raw, ok := node.(html.Raw); ok && (raw.String() == " " || raw.String() == "\t") {
					buff = append(buff, raw.Bytes()...)
					continue
				}

				if len(buff) > 0 {
					td.Push(buff)
					buff = []byte{}
				}

				td.Push(node)
			}

			rows = append(rows, td.WithStyles(columns[i].GetStyles()...))
		}

		if iter.Curr().Kind() == core.Eof {
			break
		}

		if _, err := iter.Consume(core.NewLine, "expected new line"); err != nil {
			return table, err
		}
	}

	table.Push(html.TBody(html.Tr(rows...)))
	return table, nil
}
