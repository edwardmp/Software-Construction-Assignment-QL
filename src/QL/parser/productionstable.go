
package parser

import(
	"ql/ast"
	"ql/token"
	"ql/util"
    "ql/ast/vari"
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
		String: `Form : "form" ident Block	<< ast.NewForm(X[1], X[2]) >>`,
		Id: "Form",
		NTType: 1,
		Index: 1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewForm(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Type : "integer"	<< vari.INT, nil >>`,
		Id: "Type",
		NTType: 2,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return vari.INT, nil
		},
	},
	ProdTabEntry{
		String: `Type : "boolean"	<< vari.BOOLEAN, nil >>`,
		Id: "Type",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return vari.BOOLEAN, nil
		},
	},
	ProdTabEntry{
		String: `Type : "string"	<< vari.STRING, nil >>`,
		Id: "Type",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return vari.STRING, nil
		},
	},
	ProdTabEntry{
		String: `Type : "money"	<< vari.MONEY, nil >>`,
		Id: "Type",
		NTType: 2,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return vari.MONEY, nil
		},
	},
	ProdTabEntry{
		String: `Type : "date"	<< vari.DATE, nil >>`,
		Id: "Type",
		NTType: 2,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return vari.DATE, nil
		},
	},
	ProdTabEntry{
		String: `Literal : integer_lit	<<  >>`,
		Id: "Literal",
		NTType: 3,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Literal : money_lit	<<  >>`,
		Id: "Literal",
		NTType: 3,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Literal : bool_lit	<<  >>`,
		Id: "Literal",
		NTType: 3,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Literal : str_lit	<<  >>`,
		Id: "Literal",
		NTType: 3,
		Index: 10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NumLiteral : integer_lit	<< ast.NewIntLit((util.IntValue(X[0].(*token.Token).Lit))) >>`,
		Id: "NumLiteral",
		NTType: 4,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntLit((util.IntValue(X[0].(*token.Token).Lit)))
		},
	},
	ProdTabEntry{
		String: `StrLiteral : str_lit	<< ast.NewStrLit(X[0]) >>`,
		Id: "StrLiteral",
		NTType: 5,
		Index: 12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStrLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Term : Term mulop Factor	<< ast.NewMul(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 6,
		Index: 13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMul(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Term divop Factor	<< ast.NewDiv(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 6,
		Index: 14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDiv(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id: "Term",
		NTType: 6,
		Index: 15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : NumLiteral	<<  >>`,
		Id: "Factor",
		NTType: 7,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< ast.NewBoolLit(ast.TRUE) >>`,
		Id: "Bool",
		NTType: 8,
		Index: 17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolLit(ast.TRUE)
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< ast.NewBoolLit(ast.FALSE) >>`,
		Id: "Bool",
		NTType: 8,
		Index: 18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolLit(ast.FALSE)
		},
	},
	ProdTabEntry{
		String: `VarDecl : VarId col Type	<< ast.NewVarDecl(X[0], X[2]) >>`,
		Id: "VarDecl",
		NTType: 9,
		Index: 19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarDecl(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `VarId : ident	<< ast.NewVarId(X[0]) >>`,
		Id: "VarId",
		NTType: 10,
		Index: 20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarId(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : addop Expr	<< ast.NewPos(X[1]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 21,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPos(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : subop Expr	<< ast.NewNeg(X[1]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 22,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNeg(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : notop Expr	<< ast.NewNot(X[1]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 23,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNot(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr addop Term	<< ast.NewAdd(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAdd(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr subop Term	<< ast.NewSub(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSub(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr eqop Expr	<< ast.NewEq(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr neqop Expr	<< ast.NewNEq(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr gtop Expr	<< ast.NewGT(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGT(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr ltop Expr	<< ast.NewLT(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLT(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr geqop Expr	<< ast.NewGEq(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 30,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr leqop Expr	<< ast.NewLEq(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr andop Expr	<< ast.NewAnd(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnd(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr orop Expr	<< ast.NewOr(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 11,
		Index: 33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : lpar Expr rpar	<< X[1], nil >>`,
		Id: "Expr",
		NTType: 11,
		Index: 34,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Term	<<  >>`,
		Id: "Expr",
		NTType: 11,
		Index: 35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Bool	<<  >>`,
		Id: "Expr",
		NTType: 11,
		Index: 36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : VarId	<<  >>`,
		Id: "Expr",
		NTType: 11,
		Index: 37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IfStmt : "if" lpar Expr rpar Block	<< ast.NewIf(X[2], X[4]) >>`,
		Id: "IfStmt",
		NTType: 12,
		Index: 38,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIf(X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `IfStmt : "if" lpar Expr rpar Block "else" Block	<< ast.NewIfElse(X[2], X[4], X[5]) >>`,
		Id: "IfStmt",
		NTType: 12,
		Index: 39,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfElse(X[2], X[4], X[5])
		},
	},
	ProdTabEntry{
		String: `Question : StrLiteral VarDecl	<< ast.NewQuestion(X[0], X[1]) >>`,
		Id: "Question",
		NTType: 13,
		Index: 40,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewQuestion(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Block : lbrace StmtList rbrace	<< X[1], nil >>`,
		Id: "Block",
		NTType: 14,
		Index: 41,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Block : lbrace rbrace	<< ast.NewEmptyStmtList() >>`,
		Id: "Block",
		NTType: 14,
		Index: 42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEmptyStmtList()
		},
	},
	ProdTabEntry{
		String: `StmtList : Stmt	<< ast.NewStmtList(X[0]) >>`,
		Id: "StmtList",
		NTType: 15,
		Index: 43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtList(X[0])
		},
	},
	ProdTabEntry{
		String: `StmtList : StmtList Stmt	<< ast.AppendStmt(X[0], X[1]) >>`,
		Id: "StmtList",
		NTType: 15,
		Index: 44,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStmt(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Stmt : Question	<<  >>`,
		Id: "Stmt",
		NTType: 16,
		Index: 45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : IfStmt	<<  >>`,
		Id: "Stmt",
		NTType: 16,
		Index: 46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	
}
