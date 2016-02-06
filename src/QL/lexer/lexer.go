
package lexer

import (
	
	// "fmt"
	// "QL/util"
	
	"io/ioutil"
	"unicode/utf8"
	"QL/token"
)

const(
	NoState = -1
	NumStates = 97
	NumSymbols = 123
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
0: '+'
1: '*'
2: '-'
3: '/'
4: '{'
5: '}'
6: '('
7: ')'
8: ';'
9: ':'
10: '='
11: 'f'
12: 'o'
13: 'r'
14: 'm'
15: 'i'
16: 'f'
17: 'e'
18: 'l'
19: 's'
20: 'e'
21: 't'
22: 'r'
23: 'u'
24: 'e'
25: 'f'
26: 'a'
27: 'l'
28: 's'
29: 'e'
30: '/'
31: '/'
32: '\n'
33: '/'
34: '*'
35: '*'
36: '*'
37: '/'
38: '.'
39: ','
40: 't'
41: 'r'
42: 'u'
43: 'e'
44: 'f'
45: 'a'
46: 'l'
47: 's'
48: 'e'
49: '"'
50: '"'
51: '&'
52: '&'
53: '|'
54: '|'
55: '!'
56: '<'
57: '<'
58: '='
59: '>'
60: '>'
61: '='
62: '!'
63: '='
64: '='
65: '='
66: 'd'
67: 'e'
68: 'c'
69: 'i'
70: 'm'
71: 'a'
72: 'l'
73: 'i'
74: 'n'
75: 't'
76: 'e'
77: 'g'
78: 'e'
79: 'r'
80: 'b'
81: 'o'
82: 'o'
83: 'l'
84: 'e'
85: 'a'
86: 'n'
87: 's'
88: 't'
89: 'r'
90: 'i'
91: 'n'
92: 'g'
93: 'm'
94: 'o'
95: 'n'
96: 'e'
97: 'y'
98: 'c'
99: 'u'
100: 'r'
101: 'r'
102: 'e'
103: 'n'
104: 'c'
105: 'y'
106: 'd'
107: 'a'
108: 't'
109: 'e'
110: '_'
111: ' '
112: '\t'
113: '\n'
114: '\r'
115: '0'-'9'
116: '0'-'9'
117: 'a'-'z'
118: 'A'-'Z'
119: 'a'-'z'
120: 'A'-'Z'
121: '0'-'9'
122: .

*/
