
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
	NumStates = 80
	NumSymbols = 106
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
0: ','
1: '"'
2: '"'
3: 't'
4: 'r'
5: 'u'
6: 'e'
7: 'f'
8: 'a'
9: 'l'
10: 's'
11: 'e'
12: '+'
13: '*'
14: '-'
15: '/'
16: '&'
17: '&'
18: '|'
19: '|'
20: '!'
21: '<'
22: '<'
23: '='
24: '>'
25: '>'
26: '='
27: '!'
28: '='
29: '='
30: '='
31: '{'
32: '}'
33: '('
34: ')'
35: ':'
36: '='
37: '_'
38: 'f'
39: 'o'
40: 'r'
41: 'm'
42: 'i'
43: 'n'
44: 't'
45: 'e'
46: 'g'
47: 'e'
48: 'r'
49: 'b'
50: 'o'
51: 'o'
52: 'l'
53: 'e'
54: 'a'
55: 'n'
56: 's'
57: 't'
58: 'r'
59: 'i'
60: 'n'
61: 'g'
62: 'm'
63: 'o'
64: 'n'
65: 'e'
66: 'y'
67: 'd'
68: 'a'
69: 't'
70: 'e'
71: 't'
72: 'r'
73: 'u'
74: 'e'
75: 'f'
76: 'a'
77: 'l'
78: 's'
79: 'e'
80: 'i'
81: 'f'
82: 'e'
83: 'l'
84: 's'
85: 'e'
86: '/'
87: '/'
88: '\n'
89: '/'
90: '*'
91: '*'
92: '*'
93: '/'
94: ' '
95: '\t'
96: '\n'
97: '\r'
98: 'a'-'z'
99: 'A'-'Z'
100: 'a'-'z'
101: 'A'-'Z'
102: '0'-'9'
103: '0'-'9'
104: '0'-'9'
105: .

*/
