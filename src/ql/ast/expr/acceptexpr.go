package expr

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

// acceptLhsAndRhs provides a convenient way to run Accept on both left-hand and right-hand sides of binary operands
func acceptLhsAndRhs(binaryExpr interfaces.BinaryOperatorExpr, visitor interfaces.Visitor, context interface{}) {
	binaryExpr.Lhs().Accept(visitor, context)
	binaryExpr.Rhs().Accept(visitor, context)
}

/* binary expressions */

func (this Add) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept Add")
	visitor.VisitAdd(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this And) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept And")
	visitor.VisitAnd(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this Div) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept Div")
	visitor.VisitDiv(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this Eq) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept Eq")
	visitor.VisitEq(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this GEq) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept GEq")
	visitor.VisitEq(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this GT) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept GT")
	visitor.VisitGT(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this LEq) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept LEq")
	visitor.VisitLEq(this, context)

	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this LT) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept LT")
	visitor.VisitLT(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this Mul) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept Mul")
	visitor.VisitMul(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this NEq) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept NEq")
	visitor.VisitNEq(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this Or) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept Or")
	visitor.VisitOr(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

func (this Sub) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept Sub")
	visitor.VisitSub(this, context)
	acceptLhsAndRhs(this, visitor, context)

	return nil
}

/* literals */

func (this BoolLit) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept BoolLit")
	visitor.VisitBoolLit(this, context)

	return nil
}

func (this IntLit) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept IntLit")
	visitor.VisitIntLit(this, context)

	return nil
}

func (this StrLit) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept StrLit")
	visitor.VisitStrLit(this, context)

	return nil
}

/* value types */

func (this IntType) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	visitor.VisitIntType(this, context)

	return nil
}

func (this BoolType) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	visitor.VisitBoolType(this, context)

	return nil
}

func (this StringType) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	visitor.VisitStringType(this, context)

	return nil
}

/* unary expressions */

func (this Not) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept Not")
	visitor.VisitNot(this, context)
	this.Value().Accept(visitor, context)

	return nil
}

func (this Pos) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept Pos")
	visitor.VisitPos(this, context)
	this.Value().Accept(visitor, context)

	return nil
}

func (this VarExpr) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept VarExpr")
	visitor.VisitVarExpr(this, context)
	this.Identifier().Accept(visitor, context)

	return nil
}

/* expressions */

// Accept is only called on Expr if the struct embedding it did not implement Accept, which is erroneous.
// An panic is known to accommodate for this
func (this Expr) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	panic("Expr Accept method not overridden")

	return nil
}
