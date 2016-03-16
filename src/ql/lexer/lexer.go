
package lexer

import (
	
	// "fmt"
	// "ql/util"
	
	"io/ioutil"
	"unicode/utf8"
	"ql/token"
)

const(
	NoState = -1
	NumStates = 69
	NumSymbols = 87
) 

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {
	
	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
	
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

	
		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end
	

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '"'
1: '"'
2: 't'
3: 'r'
4: 'u'
5: 'e'
6: 'f'
7: 'a'
8: 'l'
9: 's'
10: 'e'
11: '+'
12: '*'
13: '-'
14: '/'
15: '&'
16: '&'
17: '|'
18: '|'
19: '!'
20: '<'
21: '<'
22: '='
23: '>'
24: '>'
25: '='
26: '!'
27: '='
28: '='
29: '='
30: '{'
31: '}'
32: '('
33: ')'
34: ':'
35: '='
36: '_'
37: 'f'
38: 'o'
39: 'r'
40: 'm'
41: 'i'
42: 'n'
43: 't'
44: 'e'
45: 'g'
46: 'e'
47: 'r'
48: 'b'
49: 'o'
50: 'o'
51: 'l'
52: 'e'
53: 'a'
54: 'n'
55: 's'
56: 't'
57: 'r'
58: 'i'
59: 'n'
60: 'g'
61: 'i'
62: 'f'
63: 'e'
64: 'l'
65: 's'
66: 'e'
67: '/'
68: '/'
69: '\n'
70: '/'
71: '*'
72: '*'
73: '*'
74: '/'
75: ' '
76: '\t'
77: '\n'
78: '\r'
79: 'a'-'z'
80: 'A'-'Z'
81: 'a'-'z'
82: 'A'-'Z'
83: '0'-'9'
84: '0'-'9'
85: '0'-'9'
86: .

*/
