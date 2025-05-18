package ast

import (
	"github.com/thegogod/mde/reflect"
)

type Stmt interface {
	Accept(StmtVisitor) (*reflect.Value, error)
}

type StmtVisitor interface {
	VisitBlockStmt(*BlockStmt) (*reflect.Value, error)
	VisitExprStmt(*ExprStmt) (*reflect.Value, error)
	VisitForStmt(*ForStmt) (*reflect.Value, error)
	VisitFnStmt(*FnStmt) (*reflect.Value, error)
	VisitIfStmt(*IfStmt) (*reflect.Value, error)
	VisitVarStmt(*VarStmt) (*reflect.Value, error)
	VisitUseStmt(*UseStmt) (*reflect.Value, error)
}
