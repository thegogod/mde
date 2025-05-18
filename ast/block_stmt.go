package ast

import (
	"github.com/thegogod/mde/reflect"
)

type BlockStmt struct {
	Stmts []Stmt
}

func NewBlockStmt(stmts []Stmt) *BlockStmt {
	return &BlockStmt{
		Stmts: stmts,
	}
}

func (self *BlockStmt) Accept(v StmtVisitor) (*reflect.Value, error) {
	return v.VisitBlockStmt(self)
}
