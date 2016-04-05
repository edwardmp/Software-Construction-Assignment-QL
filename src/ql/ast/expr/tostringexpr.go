package expr

import (
	"fmt"
	"ql/interfaces"
)

/* binary expressions */

func (this Add) String() string {
	return fmt.Sprintf("%s + %s", this.Lhs().String(), this.Rhs().String())
}

func (this And) String() string {
	return fmt.Sprintf("%s && %s", this.Lhs().String(), this.Rhs().String())
}

func (this Div) String() string {
	return fmt.Sprintf("%s / %s", this.Lhs().String(), this.Rhs().String())
}

func (this Eq) String() string {
	return fmt.Sprintf("%s == %s", this.Lhs().String(), this.Rhs().String())
}

func (this GEq) String() string {
	return fmt.Sprintf("%s >= %s", this.Lhs().String(), this.Rhs().String())
}

func (this GT) String() string {
	return fmt.Sprintf("%s > %s", this.Lhs().String(), this.Rhs().String())
}

func (this LEq) String() string {
	return fmt.Sprintf("%s <= %s", this.Lhs().String(), this.Rhs().String())
}

func (this LT) String() string {
	return fmt.Sprintf("%s < %s", this.Lhs().String(), this.Rhs().String())
}

func (this Mul) String() string {
	return fmt.Sprintf("%s () %s", this.Lhs().String(), this.Rhs().String())
}

func (this NEq) String() string {
	return fmt.Sprintf("%s != %s", this.Lhs().String(), this.Rhs().String())
}

func (this Or) String() string {
	return fmt.Sprintf("%s || %s", this.Lhs().String(), this.Rhs().String())
}

func (this Sub) String() string {
	return fmt.Sprintf("%s - %s", this.Lhs().String(), this.Rhs().String())
}

/* literals */

func (this BoolLit) String() string {
	return fmt.Sprintf("%t", this.Value())
}

func (this IntLit) String() string {
	return fmt.Sprintf("%d", this.Value())
}

func (this StrLit) String() string {
	return this.Value().(interfaces.StringValue).PrimitiveValueString()
}

/* unary expressions */

func (this Not) String() string {
	return fmt.Sprintf("%s", this.Value())
}

func (this Pos) String() string {
	return fmt.Sprintf("%s", this.Value())
}

func (this VarExpr) String() string {
	return fmt.Sprintf("%s", this.Identifier())
}

/* value types */

func (this ValueType) String() string {
	return this.typeString
}

/* expressions */

// String is only called on Expr if the struct embedding it did not implement String, which is erroneous
func (this Expr) String() string {
	panic("Expr String method not overridden")
	return ""
}
