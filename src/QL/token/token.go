
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
		"mulop",
		"divop",
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
		"int",
		"true",
		"false",
		"col",
		"compop",
		"else",
		"equals",
		"form",
		"ident",
		"if",
		"lbrace",
		"literal",
		"logop",
		"num_literal",
		"rbrace",
		"semicol",
		"type",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"addop": 2,
		"subop": 3,
		"notop": 4,
		"mulop": 5,
		"divop": 6,
		"eqop": 7,
		"neqop": 8,
		"gtop": 9,
		"ltop": 10,
		"geqop": 11,
		"leqop": 12,
		"andop": 13,
		"orop": 14,
		"lpar": 15,
		"rpar": 16,
		"int": 17,
		"true": 18,
		"false": 19,
		"col": 20,
		"compop": 21,
		"else": 22,
		"equals": 23,
		"form": 24,
		"ident": 25,
		"if": 26,
		"lbrace": 27,
		"literal": 28,
		"logop": 29,
		"num_literal": 30,
		"rbrace": 31,
		"semicol": 32,
		"type": 33,
	},
}

