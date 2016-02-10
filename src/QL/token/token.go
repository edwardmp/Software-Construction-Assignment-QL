
package token

import(
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const(
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap  []string
	idMap map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"addop",
		"subop",
		"notop",
		"eqop",
		"neqop",
		"gtop",
		"ltop",
		"geqop",
		"leqop",
		"andop",
		"orop",
		"lpar",
		"rpar",
		"ident",
		"mulop",
		"divop",
		"int",
		"true",
		"false",
		"lbrace",
		"rbrace",
		"type",
		"if",
		"else",
		"assign",
		"col",
		"form",
		"literal",
		"num_literal",
		"semicol",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"addop": 2,
		"subop": 3,
		"notop": 4,
		"eqop": 5,
		"neqop": 6,
		"gtop": 7,
		"ltop": 8,
		"geqop": 9,
		"leqop": 10,
		"andop": 11,
		"orop": 12,
		"lpar": 13,
		"rpar": 14,
		"ident": 15,
		"mulop": 16,
		"divop": 17,
		"int": 18,
		"true": 19,
		"false": 20,
		"lbrace": 21,
		"rbrace": 22,
		"type": 23,
		"if": 24,
		"else": 25,
		"assign": 26,
		"col": 27,
		"form": 28,
		"literal": 29,
		"num_literal": 30,
		"semicol": 31,
	},
}

