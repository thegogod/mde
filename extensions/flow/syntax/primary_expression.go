package syntax

import (
	"strconv"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/extensions/flow/ast"
	"github.com/thegogod/mde/reflect"
)

type PrimaryExpression struct{}

func (self PrimaryExpression) IsBlock() bool {
	return false
}

func (self PrimaryExpression) IsInline() bool {
	return true
}

func (self PrimaryExpression) Name() string {
	return "primary"
}

func (self PrimaryExpression) Select(parser core.Parser, iter *core.Iterator) bool {
	return false
}

func (self PrimaryExpression) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	if iter.Match(core.True) {
		return ast.Html(ast.Literal(reflect.NewBool(true))), nil
	} else if iter.Match(core.False) {
		return ast.Html(ast.Literal(reflect.NewBool(false))), nil
	} else if iter.Match(core.Integer) {
		i, err := strconv.ParseInt(string(iter.Prev().Bytes()), 0, 8)
		return ast.Html(ast.Literal(reflect.NewInt(int(i)))), err
	} else if iter.Match(core.Decimal) {
		i, err := strconv.ParseFloat(string(iter.Prev().Bytes()), 8)
		return ast.Html(ast.Literal(reflect.NewFloat(i))), err
	} else if iter.Match(core.Null) {
		return ast.Html(ast.Literal(reflect.NewNil())), nil
	}

	return nil, nil
}
