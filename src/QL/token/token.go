
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
		"ident",
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
		"lbrace",
		"rbrace",
		"assign",
		"semicol",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"form": 2,
		"ident": 3,
		"integer": 4,
		"boolean": 5,
		"string": 6,
		"money": 7,
		"date": 8,
		"integer_lit": 9,
		"money_lit": 10,
		"bool_lit": 11,
		"str_lit": 12,
		"mulop": 13,
		"divop": 14,
		"true": 15,
		"false": 16,
		"col": 17,
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
		"lbrace": 33,
		"rbrace": 34,
		"assign": 35,
		"semicol": 36,
	},
}

