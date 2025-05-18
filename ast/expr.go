package ast

import (
	"github.com/thegogod/mde/reflect"
)

type Expr interface {
	GetType() (reflect.Type, error)
	Accept(ExprVisitor) (*reflect.Value, error)
}

type ExprVisitor interface {
	VisitAssignExpr(*AssignExpr) (*reflect.Value, error)
	VisitBinaryExpr(*BinaryExpr) (*reflect.Value, error)
	VisitCallExpr(*CallExpr) (*reflect.Value, error)
	VisitGetExpr(*GetExpr) (*reflect.Value, error)
	VisitGroupingExpr(*GroupingExpr) (*reflect.Value, error)
	VisitLiteralExpr(*LiteralExpr) (*reflect.Value, error)
	VisitLogicalExpr(*LogicalExpr) (*reflect.Value, error)
	VisitSelfExpr(*SelfExpr) (*reflect.Value, error)
	VisitSetExpr(*SetExpr) (*reflect.Value, error)
	VisitSliceExpr(*SliceExpr) (*reflect.Value, error)
	VisitUnaryExpr(*UnaryExpr) (*reflect.Value, error)
	VisitVarExpr(*VarExpr) (*reflect.Value, error)
}
