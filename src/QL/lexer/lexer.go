
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
	NumStates = 83
	NumSymbols = 108
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
4: '&'
5: '&'
6: '|'
7: '|'
8: '!'
9: '<'
10: '<'
11: '='
12: '>'
13: '>'
14: '='
15: '!'
16: '='
17: '='
18: '='
19: '{'
20: '}'
21: '('
22: ')'
23: ';'
24: ':'
25: '='
26: 'f'
27: 'o'
28: 'r'
29: 'm'
30: 'i'
31: 'f'
32: 'e'
33: 'l'
34: 's'
35: 'e'
36: '_'
37: 't'
38: 'r'
39: 'u'
40: 'e'
41: 'f'
42: 'a'
43: 'l'
44: 's'
45: 'e'
46: '/'
47: '/'
48: '\n'
49: '/'
50: '*'
51: '*'
52: '*'
53: '/'
54: '.'
55: ','
56: '"'
57: '"'
58: 't'
59: 'r'
60: 'u'
61: 'e'
62: 'f'
63: 'a'
64: 'l'
65: 's'
66: 'e'
67: 'i'
68: 'n'
69: 't'
70: 'e'
71: 'g'
72: 'e'
73: 'r'
74: 'b'
75: 'o'
76: 'o'
77: 'l'
78: 'e'
79: 'a'
80: 'n'
81: 's'
82: 't'
83: 'r'
84: 'i'
85: 'n'
86: 'g'
87: 'm'
88: 'o'
89: 'n'
90: 'e'
91: 'y'
92: 'd'
93: 'a'
94: 't'
95: 'e'
96: ' '
97: '\t'
98: '\n'
99: '\r'
100: 'a'-'z'
101: 'A'-'Z'
102: 'a'-'z'
103: 'A'-'Z'
104: '0'-'9'
105: '0'-'9'
106: '0'-'9'
107: .

*/
