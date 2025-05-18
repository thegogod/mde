package ast

import (
	"github.com/thegogod/mde/reflect"
)

type VarStmt struct {
	Keyword *Token
	Name    *Token
	Type    reflect.Type
	Nilable *Token
	Init    Expr
}

func NewVarStmt(
	keyword *Token,
	name *Token,
	_type reflect.Type,
	nilable *Token,
	init Expr,
) *VarStmt {
	return &VarStmt{
		Keyword: keyword,
		Name:    name,
		Type:    _type,
		Nilable: nilable,
		Init:    init,
	}
}

func (self *VarStmt) Accept(v StmtVisitor) (*reflect.Value, error) {
	return v.VisitVarStmt(self)
}
