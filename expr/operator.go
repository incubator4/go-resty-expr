package expr

var (
	Equal                   = OperatorExpr("==")
	NotEqual                = OperatorExpr("~=")
	RegularMatch            = OperatorExpr("~~")
	InsensitiveRegularMatch = OperatorExpr("~*")
	Not                     = OperatorExpr("!")
	Has                     = OperatorExpr("has")
	In                      = OperatorExpr("in")
	GreaterThan             = OperatorExpr(">")
	LessThan                = OperatorExpr("<")
	IPMatch                 = OperatorExpr("ipmatch")
	AND                     = OperatorExpr("AND")
	OR                      = OperatorExpr("OR")
)

func (e Expr) Equals(expr Expr) Expr {
	return e.AppendExpr(Equal, expr)
}

func (e Expr) Not() Expr {
	return e.AppendExpr(Not)
}

func (e Expr) NotEquals(expr Expr) Expr {
	return e.AppendExpr(NotEqual, expr)
}

func (e Expr) In(expr Expr) Expr {
	return e.AppendExpr(In, expr)
}

func (e Expr) NotIn(expr Expr) Expr {
	return e.Not().In(expr)
}

func (e Expr) GreaterThan(expr Expr) Expr {
	return e.AppendExpr(GreaterThan, expr)
}

func (e Expr) LessThan(expr Expr) Expr {
	return e.AppendExpr(LessThan, expr)
}

func And(exprArray ...Expr) Expr {
	return AND.CreateArrayExpr(exprArray...)
}

func Or(exprArray ...Expr) Expr {
	return OR.CreateArrayExpr(exprArray...)
}
