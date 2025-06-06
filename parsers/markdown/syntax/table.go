package syntax

import (
	"bytes"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/maps"
	"github.com/thegogod/mde/parsers/markdown/tokens"
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

func (self Table) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.MatchCount(tokens.Pipe, 1)
}

func (self Table) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	table := html.Table()
	columns := []*html.TableCellElement{}

	for !iter.Match(tokens.NewLine) {
		th := html.Th()

		for {
			if !iter.Match(tokens.Space) && !iter.Match(tokens.Tab) {
				break
			}
		}

		buff := []byte{}

		for !iter.Match(tokens.Pipe) {
			node, err := parser.ParseInline(iter)

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

	if _, err := iter.Consume(tokens.Pipe, "expected opening '|'"); err != nil {
		return table, err
	}

	for i := range len(columns) {
		node, err := parser.ParseTextUntil(iter, tokens.Pipe)

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
			columns[i].Style(maps.KeyValue[string, string]{
				Key:   "text-align",
				Value: "center",
			})
		} else if leftAlign {
			columns[i].Style(maps.KeyValue[string, string]{
				Key:   "text-align",
				Value: "left",
			})
		} else if rightAlign {
			columns[i].Style(maps.KeyValue[string, string]{
				Key:   "text-align",
				Value: "right",
			})
		}
	}

	if _, err := iter.Consume(tokens.NewLine, "expected new line"); err != nil {
		return table, err
	}

	table.Push(html.THead(html.Tr(columns...)))
	rows := []*html.TableCellElement{}

	for iter.Match(tokens.Pipe) {
		for i := range len(columns) {
			td := html.Td()

			for {
				if !iter.Match(tokens.Space) && !iter.Match(tokens.Tab) {
					break
				}
			}

			buff := []byte{}

			for !iter.Match(tokens.Pipe) {
				node, err := parser.ParseInline(iter)

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

			rows = append(rows, td.Style(columns[i].GetStyles()...))
		}

		if iter.Curr().Kind() == tokens.Eof {
			break
		}

		if _, err := iter.Consume(tokens.NewLine, "expected new line"); err != nil {
			return table, err
		}
	}

	table.Push(html.TBody(html.Tr(rows...)))
	return table, nil
}
