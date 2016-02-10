
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
		"tint",
		"tbool",
		"tstring",
		"tmoney",
		"tdate",
		"integer_lit",
		"money_lit",
		"bool_lit",
		"str_lit",
		"mulop",
		"divop",
		"true",
		"false",
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
		"col",
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
		"tint": 4,
		"tbool": 5,
		"tstring": 6,
		"tmoney": 7,
		"tdate": 8,
		"integer_lit": 9,
		"money_lit": 10,
		"bool_lit": 11,
		"str_lit": 12,
		"mulop": 13,
		"divop": 14,
		"true": 15,
		"false": 16,
		"addop": 17,
		"subop": 18,
		"notop": 19,
		"eqop": 20,
		"neqop": 21,
		"gtop": 22,
		"ltop": 23,
		"geqop": 24,
		"leqop": 25,
		"andop": 26,
		"orop": 27,
		"lpar": 28,
		"rpar": 29,
		"if": 30,
		"else": 31,
		"col": 32,
		"lbrace": 33,
		"rbrace": 34,
		"assign": 35,
		"semicol": 36,
	},
}

