
package parser

import(
	"ql/ast"
    "ql/token"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Form	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Form : "form" VarId Block	<< ast.NewForm(X[1], X[2], X[0].(*token.Token).Pos) >>`,
		Id: "Form",
		NTType: 1,
		Index: 1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewForm(X[1], X[2], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Type : "integer"	<< ast.NewIntType(X[0]) >>`,
		Id: "Type",
		NTType: 2,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntType(X[0])
		},
	},
	ProdTabEntry{
		String: `Type : "boolean"	<< ast.NewBoolType(X[0]) >>`,
		Id: "Type",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolType(X[0])
		},
	},
	ProdTabEntry{
		String: `Type : "string"	<< ast.NewStringType(X[0]) >>`,
		Id: "Type",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringType(X[0])
		},
	},
	ProdTabEntry{
		String: `Term : Term mulop Factor	<< ast.NewMul(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Term",
		NTType: 3,
		Index: 5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMul(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Term : Term divop Factor	<< ast.NewDiv(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Term",
		NTType: 3,
		Index: 6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDiv(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id: "Term",
		NTType: 3,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : NumLiteral	<<  >>`,
		Id: "Factor",
		NTType: 4,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : BoolLiteral	<<  >>`,
		Id: "Factor",
		NTType: 4,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : StrLiteral	<<  >>`,
		Id: "Factor",
		NTType: 4,
		Index: 10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : VarId	<< ast.NewVarExpr(X[0]) >>`,
		Id: "Factor",
		NTType: 4,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarExpr(X[0])
		},
	},
	ProdTabEntry{
		String: `NumLiteral : integer_lit	<< ast.NewIntLit(X[0]) >>`,
		Id: "NumLiteral",
		NTType: 5,
		Index: 12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntLit(X[0])
		},
	},
	ProdTabEntry{
		String: `StrLiteral : str_lit	<< ast.NewStrLit(X[0]) >>`,
		Id: "StrLiteral",
		NTType: 6,
		Index: 13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStrLit(X[0])
		},
	},
	ProdTabEntry{
		String: `BoolLiteral : booltrue_lit	<< ast.NewBoolLit(ast.TRUE, X[0].(*token.Token).Pos) >>`,
		Id: "BoolLiteral",
		NTType: 7,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolLit(ast.TRUE, X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `BoolLiteral : boolfalse_lit	<< ast.NewBoolLit(ast.FALSE, X[0].(*token.Token).Pos) >>`,
		Id: "BoolLiteral",
		NTType: 7,
		Index: 15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolLit(ast.FALSE, X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `VarDecl : VarId col Type	<< ast.NewVarDecl(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "VarDecl",
		NTType: 8,
		Index: 16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarDecl(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `VarId : ident	<< ast.NewVarId(X[0]) >>`,
		Id: "VarId",
		NTType: 9,
		Index: 17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarId(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : addop Expr	<< ast.NewPos(X[1], X[0].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 18,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPos(X[1], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : subop Expr	<< ast.NewNeg(X[1], X[0].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 19,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNeg(X[1], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : notop Expr	<< ast.NewNot(X[1], X[0].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 20,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNot(X[1], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr addop Term	<< ast.NewAdd(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAdd(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr subop Term	<< ast.NewSub(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSub(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr eqop Expr	<< ast.NewEq(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEq(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr neqop Expr	<< ast.NewNEq(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNEq(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr gtop Expr	<< ast.NewGT(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGT(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr ltop Expr	<< ast.NewLT(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLT(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr geqop Expr	<< ast.NewGEq(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGEq(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr leqop Expr	<< ast.NewLEq(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLEq(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr andop Expr	<< ast.NewAnd(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnd(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr orop Expr	<< ast.NewOr(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 10,
		Index: 30,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOr(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : lpar Expr rpar	<< X[1], nil >>`,
		Id: "Expr",
		NTType: 10,
		Index: 31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Term	<<  >>`,
		Id: "Expr",
		NTType: 10,
		Index: 32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : BoolLiteral	<<  >>`,
		Id: "Expr",
		NTType: 10,
		Index: 33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : StrLiteral	<<  >>`,
		Id: "Expr",
		NTType: 10,
		Index: 34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IfStmt : "if" lpar Expr rpar Block	<< ast.NewIf(X[2], X[4], X[0].(*token.Token).Pos) >>`,
		Id: "IfStmt",
		NTType: 11,
		Index: 35,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIf(X[2], X[4], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `IfStmt : "if" lpar Expr rpar Block "else" Block	<< ast.NewIfElse(X[2], X[4], X[6], X[0].(*token.Token).Pos) >>`,
		Id: "IfStmt",
		NTType: 11,
		Index: 36,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfElse(X[2], X[4], X[6], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Question : StrLiteral VarDecl	<< ast.NewInputQuestion(X[0], X[1]) >>`,
		Id: "Question",
		NTType: 12,
		Index: 37,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInputQuestion(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Question : StrLiteral VarDecl assign Expr	<< ast.NewComputedQuestion(X[0], X[1], X[3], X[2].(*token.Token).Pos) >>`,
		Id: "Question",
		NTType: 12,
		Index: 38,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComputedQuestion(X[0], X[1], X[3], X[2].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Block : lbrace StmtList rbrace	<< X[1], nil >>`,
		Id: "Block",
		NTType: 13,
		Index: 39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Block : lbrace rbrace	<< ast.NewEmptyStmtList(X[0].(*token.Token).Pos) >>`,
		Id: "Block",
		NTType: 13,
		Index: 40,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEmptyStmtList(X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `StmtList : Stmt	<< ast.NewStmtList(X[0]) >>`,
		Id: "StmtList",
		NTType: 14,
		Index: 41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtList(X[0])
		},
	},
	ProdTabEntry{
		String: `StmtList : StmtList Stmt	<< ast.AppendStmt(X[0], X[1]) >>`,
		Id: "StmtList",
		NTType: 14,
		Index: 42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStmt(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Stmt : Question	<<  >>`,
		Id: "Stmt",
		NTType: 15,
		Index: 43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : IfStmt	<<  >>`,
		Id: "Stmt",
		NTType: 15,
		Index: 44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	
}
