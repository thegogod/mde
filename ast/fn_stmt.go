package ast

import (
	"github.com/thegogod/mde/reflect"
)

type FnStmt struct {
	Name       *Token
	Params     []*FnStmt
	ReturnType reflect.Type
	Body       []Stmt
}

func NewFnStmt(
	name *Token,
	params []*FnStmt,
	returnType reflect.Type,
	body []Stmt,
) *FnStmt {
	return &FnStmt{
		Name:       name,
		Params:     params,
		ReturnType: returnType,
		Body:       body,
	}
}

func (self *FnStmt) Accept(v StmtVisitor) (*reflect.Value, error) {
	return v.VisitFnStmt(self)
}
