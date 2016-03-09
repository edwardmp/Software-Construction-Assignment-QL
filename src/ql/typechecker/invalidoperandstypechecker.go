package typechecker

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/ast/expr"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/litexpr"
	"ql/ast/expr/unaryoperatorexpr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/symboltable"
)

type InvalidOperandsTypeChecker struct {
	TypeChecker
}

func CheckForOperatorsWithInvalidOperands(form stmt.Form, symbolTable symboltable.SymbolTable) []error {
	log.Info("Start check for invalid operands")
	invalidOperandsTypeChecker := InvalidOperandsTypeChecker{}

	form.Accept(&invalidOperandsTypeChecker, symbolTable)
	log.WithFields(log.Fields{"ErrorsEncountered": invalidOperandsTypeChecker.ErrorsEncountered}).Info("Ended check for invalid operands")

	return invalidOperandsTypeChecker.ErrorsEncountered
}

func (v *InvalidOperandsTypeChecker) Visit(t interface{}, s interface{}) interface{} {
	switch t.(type) {
	default:
		log.WithFields(log.Fields{"Node": fmt.Sprintf("%T", t)}).Panic("Unexpected node type")
	case stmt.Form:
		log.Debug("Visit Form")
		t.(stmt.Form).Identifier.Accept(v, s)
		return t.(stmt.Form).Content.Accept(v, s)
	case vari.VarId:
		log.Debug("Visit VarId")
	case vari.VarType:
		log.Debug("Visit VarType")
	case vari.VarDecl:
		log.Debug("Visit VarDecl")
		varDecl := t.(vari.VarDecl)
		varDecl.Ident.Accept(v, s)
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for _, question := range t.(stmt.StmtList).Questions {
			question.Accept(v, s)
		}

		for _, conditional := range t.(stmt.StmtList).Conditionals {
			conditional.Accept(v, s)
		}
	case stmt.InputQuestion:
		log.Debug("Visit InputQuestion")
		t.(stmt.InputQuestion).Label.Accept(v, s)
		t.(stmt.InputQuestion).VarDecl.Accept(v, s)
	case stmt.ComputedQuestion:
		log.Debug("Visit ComputedQuestion")
		t.(stmt.ComputedQuestion).Label.Accept(v, s)
		t.(stmt.ComputedQuestion).VarDecl.Accept(v, s)
		t.(stmt.ComputedQuestion).Computation.Accept(v, s)
	case stmt.If:
		log.Debug("Visit If")
		t.(stmt.If).Cond.Accept(v, s)
		t.(stmt.If).Body.Accept(v, s)
	case stmt.IfElse:
		log.Debug("Visit IfElse")
		t.(stmt.IfElse).Cond.Accept(v, s)
		t.(stmt.IfElse).IfBody.Accept(v, s)
		t.(stmt.IfElse).ElseBody.Accept(v, s)
	case litexpr.StrLit:
		log.Debug("Visit StrLit")
	case litexpr.BoolLit:
		log.Debug("Visit BoolLit")
	case litexpr.IntLit:
		log.Debug("Visit IntLit")
	case binaryoperatorexpr.BinaryOperatorExpr:
		log.Debug("Visit BinaryOperatorExpr")

		v.handleBinaryOperatorExpr(t.(binaryoperatorexpr.BinaryOperatorExpr), s.(symboltable.SymbolTable))

		t.(binaryoperatorexpr.BinaryOperatorExpr).GetLhs().(expr.Expr).Accept(v, s)
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetRhs().(expr.Expr).Accept(v, s)
	case unaryoperatorexpr.UnaryOperatorExpr:
		log.Debug("Visit UnaryOperatorExpr")

		v.handleUnaryOperatorExpr(t.(unaryoperatorexpr.UnaryOperatorExpr), s.(symboltable.SymbolTable))

		t.(unaryoperatorexpr.UnaryOperatorExpr).GetValue().(expr.Expr).Accept(v, s)
	case unaryoperatorexpr.VarExpr:
		log.Debug("Visit VarExpr")
		t.(unaryoperatorexpr.VarExpr).GetIdentifier().Accept(v, s)
	}

	return nil
}

func (v *InvalidOperandsTypeChecker) handleBinaryOperatorExpr(binaryOperatorExpr binaryoperatorexpr.BinaryOperatorExpr, symbolTable symboltable.SymbolTable) {
	lhs := binaryOperatorExpr.GetLhs()
	rhs := binaryOperatorExpr.GetRhs()
	exprType := fmt.Sprintf("%T", binaryOperatorExpr)
	lhsTypeOfValue := fmt.Sprintf("%T", lhs.Eval(symbolTable))
	rhsTypeOfValue := fmt.Sprintf("%T", rhs.Eval(symbolTable))

	if lhsTypeOfValue != rhsTypeOfValue {
		v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Encountered BinaryOperatorExpr with operands of different types: %s and %s", lhsTypeOfValue, rhsTypeOfValue))
	} else if lhsTypeOfValue == "bool" { // only check against lhs because we know they are of same type at this point
		if exprType != "binaryoperatorexpr.And" && exprType != "binaryoperatorexpr.Or" && exprType != "binaryoperatorexpr.Eq" && exprType != "binaryoperatorexpr.NEq" {
			v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Encountered invalid operation for bool operands: %s", exprType))
		}
	} else if lhsTypeOfValue == "int" {
		if exprType != "binaryoperatorexpr.Add" && exprType != "binaryoperatorexpr.Div" && exprType != "binaryoperatorexpr.Eq" && exprType != "binaryoperatorexpr.NEq" && exprType != "binaryoperatorexpr.GEq" && exprType != "binaryoperatorexpr.GT" && exprType != "binaryoperatorexpr.GEq" && exprType != "binaryoperatorexpr.LEq" && exprType != "binaryoperatorexpr.LT" && exprType != "binaryoperatorexpr.Mul" && exprType != "binaryoperatorexpr.Sub" {
			v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Encountered invalid operation for int operands: %s", exprType))
		}
	} else if lhsTypeOfValue == "string" {
		if exprType != "binaryoperatorexpr.Eq" && exprType != "binaryoperatorexpr.NEq" {
			v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Encountered invalid operation for string operands: %s", exprType))
		}
	}
}

func (v *InvalidOperandsTypeChecker) handleUnaryOperatorExpr(unaryOperatorExpr unaryoperatorexpr.UnaryOperatorExpr, symbolTable symboltable.SymbolTable) {
	value := unaryOperatorExpr.GetValue()
	exprType := fmt.Sprintf("%T", unaryOperatorExpr)
	typeOfValue := fmt.Sprintf("%T", value.(expr.Expr).Eval(symbolTable))

	if typeOfValue == "bool" {
		if exprType != "unaryoperatorexpr.Not" {
			v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Encountered invalid operation for bool operand: %s", exprType))
		}
	} else if typeOfValue == "int" {
		if exprType != "unaryoperatorexpr.Neg" && exprType != "unaryoperatorexpr.Pos" {
			v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Encountered invalid operation for int operand: %s", exprType))
		}
	} else if typeOfValue == "string" {
		v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Encountered invalid operation for string operand: %s", exprType))
	}
}
