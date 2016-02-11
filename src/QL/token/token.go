
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
		"form",
		"integer",
		"boolean",
		"string",
		"money",
		"date",
		"integer_lit",
		"money_lit",
		"bool_lit",
		"str_lit",
		"mulop",
		"divop",
		"true",
		"false",
		"col",
		"ident",
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
		"if",
		"else",
		"assign",
		"lbrace",
		"rbrace",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"form": 2,
		"integer": 3,
		"boolean": 4,
		"string": 5,
		"money": 6,
		"date": 7,
		"integer_lit": 8,
		"money_lit": 9,
		"bool_lit": 10,
		"str_lit": 11,
		"mulop": 12,
		"divop": 13,
		"true": 14,
		"false": 15,
		"col": 16,
		"ident": 17,
		"addop": 18,
		"subop": 19,
		"notop": 20,
		"eqop": 21,
		"neqop": 22,
		"gtop": 23,
		"ltop": 24,
		"geqop": 25,
		"leqop": 26,
		"andop": 27,
		"orop": 28,
		"lpar": 29,
		"rpar": 30,
		"if": 31,
		"else": 32,
		"assign": 33,
		"lbrace": 34,
		"rbrace": 35,
	},
}

