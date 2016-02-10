
package parser

import(
	"QL/ast"
	"QL/token"
	"QL/util"
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
		String: `S' : Expr	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : addop Expr	<< ast.NewPos(X[1]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPos(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : subop Expr	<< ast.NewNeg(X[1]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNeg(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : notop Expr	<< ast.NewNot(X[1]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNot(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr addop Term	<< ast.NewAdd(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAdd(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr subop Expr	<< ast.NewSub(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSub(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Term	<<  >>`,
		Id: "Expr",
		NTType: 1,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Expr eqop Expr	<< ast.NewEq(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr neqop Expr	<< ast.NewNEq(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr gtop Expr	<< ast.NewGT(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGT(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr ltop Expr	<< ast.NewLT(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLT(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr geqop Expr	<< ast.NewGEq(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr leqop Expr	<< ast.NewLEq(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr andop Expr	<< ast.NewAnd(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnd(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr orop Expr	<< ast.NewOr(X[0], X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : lpar Expr rpar	<< X[1], nil >>`,
		Id: "Expr",
		NTType: 1,
		Index: 15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Bool	<<  >>`,
		Id: "Expr",
		NTType: 1,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : ident	<<  >>`,
		Id: "Expr",
		NTType: 1,
		Index: 17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term mulop Factor	<< ast.NewMul(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMul(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Term divop Factor	<< ast.NewDiv(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDiv(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id: "Term",
		NTType: 2,
		Index: 20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : int	<< ast.NewIntLit((util.IntValue(X[0].(*token.Token).Lit))) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntLit((util.IntValue(X[0].(*token.Token).Lit)))
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< ast.NewBoolLit(ast.TRUE) >>`,
		Id: "Bool",
		NTType: 4,
		Index: 22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolLit(ast.TRUE)
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< ast.NewBoolLit(ast.FALSE) >>`,
		Id: "Bool",
		NTType: 4,
		Index: 23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolLit(ast.FALSE)
		},
	},
	ProdTabEntry{
		String: `Form : ident lbrace Body rbrace	<<  >>`,
		Id: "Form",
		NTType: 5,
		Index: 24,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Question : ident type	<<  >>`,
		Id: "Question",
		NTType: 6,
		Index: 25,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Ifstmt : if Expr Body Elsestmt	<<  >>`,
		Id: "Ifstmt",
		NTType: 7,
		Index: 26,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Elsestmt : else Body	<<  >>`,
		Id: "Elsestmt",
		NTType: 8,
		Index: 27,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Body : Ifstmt	<<  >>`,
		Id: "Body",
		NTType: 9,
		Index: 28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Body : Question	<<  >>`,
		Id: "Body",
		NTType: 9,
		Index: 29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	
}
