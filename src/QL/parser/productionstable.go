package parser

import "QL/ast"
import "QL/util"
import "QL/token"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Expr	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : addop Expr	<< ast.NewPos(X[1]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPos(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : subop Expr	<< ast.NewNeg(X[1]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNeg(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : notop Expr	<< ast.NewNot(X[1]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNot(X[1])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr mulop Expr	<< ast.NewMul(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMul(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr divop Expr	<< ast.NewDiv(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDiv(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr addop Expr	<< ast.NewAdd(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAdd(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr subop Expr	<< ast.NewSub(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSub(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr eqop Expr	<< ast.NewEq(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr neqop Expr	<< ast.NewNEq(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr gtop Expr	<< ast.NewGT(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGT(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr ltop Expr	<< ast.NewLT(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLT(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr geqop Expr	<< ast.NewGEq(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr leqop Expr	<< ast.NewLEq(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLEq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr andop Expr	<< ast.NewAnd(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnd(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr orop Expr	<< ast.NewOr(X[0], X[2]) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : lpar Expr rpar	<< X[1], nil >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Expr : int	<< util.IntValue(X[0].(*token.Token).Lit) >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return util.IntValue(X[0].(*token.Token).Lit)
		},
	},
	ProdTabEntry{
		String: `Expr : Val	<<  >>`,
		Id:         "Expr",
		NTType:     1,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Val : "true"	<< ast.TRUE, nil >>`,
		Id:         "Val",
		NTType:     2,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.TRUE, nil
		},
	},
	ProdTabEntry{
		String: `Val : "false"	<< ast.FALSE, nil >>`,
		Id:         "Val",
		NTType:     2,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.FALSE, nil
		},
	},
}
