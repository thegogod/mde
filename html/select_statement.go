package html

type SelectStatement interface {
	Eval(node Node) bool
}

type Selector []SelectStatement

func Select(query ...SelectStatement) Selector {
	stmt := Selector{}

	for _, q := range query {
		stmt = append(stmt, q)
	}

	return stmt
}

func (self *Selector) And(query ...SelectStatement) *Selector {
	arr := *self
	arr = append(arr, And(query...))
	*self = arr
	return self
}

func (self *Selector) Or(query ...SelectStatement) *Selector {
	arr := *self
	arr = append(arr, Or(query...))
	*self = arr
	return self
}

func (self Selector) Eval(node Node) bool {
	for _, query := range self {
		if !query.Eval(node) {
			return false
		}
	}

	return true
}

type AndStatement []SelectStatement

func And(query ...SelectStatement) AndStatement {
	and := AndStatement{}

	for _, q := range query {
		and = append(and, q)
	}

	return and
}

func (self AndStatement) Eval(node Node) bool {
	for _, query := range self {
		if !query.Eval(node) {
			return false
		}
	}

	return true
}

type OrStatement []SelectStatement

func Or(query ...SelectStatement) OrStatement {
	or := OrStatement{}

	for _, q := range query {
		or = append(or, q)
	}

	return or
}

func (self OrStatement) Eval(node Node) bool {
	for _, query := range self {
		if query.Eval(node) {
			return true
		}
	}

	return false
}

type ClassExpression struct {
	classes []string
}

func WithClass(classes ...string) ClassExpression {
	return ClassExpression{classes}
}

func (self ClassExpression) Eval(node Node) bool {
	return false
}
