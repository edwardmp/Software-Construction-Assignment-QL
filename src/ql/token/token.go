
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
		"mulop",
		"divop",
		"integer_lit",
		"str_lit",
		"booltrue_lit",
		"boolfalse_lit",
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
		"mulop": 6,
		"divop": 7,
		"integer_lit": 8,
		"str_lit": 9,
		"booltrue_lit": 10,
		"boolfalse_lit": 11,
		"col": 12,
		"ident": 13,
		"addop": 14,
		"subop": 15,
		"notop": 16,
		"eqop": 17,
		"neqop": 18,
		"gtop": 19,
		"ltop": 20,
		"geqop": 21,
		"leqop": 22,
		"andop": 23,
		"orop": 24,
		"lpar": 25,
		"rpar": 26,
		"if": 27,
		"else": 28,
		"assign": 29,
		"lbrace": 30,
		"rbrace": 31,
	},
}

