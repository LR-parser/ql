package lexer

import (

	// "fmt"
	// ql/util"

	"io/ioutil"
	"ql/token"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 69
	NumSymbols = 87
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
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
2: '+'
3: '*'
4: '-'
5: '/'
6: '&'
7: '&'
8: '|'
9: '|'
10: '!'
11: '<'
12: '<'
13: '='
14: '>'
15: '>'
16: '='
17: '!'
18: '='
19: '='
20: '='
21: '{'
22: '}'
23: '('
24: ')'
25: ':'
26: '='
27: '_'
28: 'f'
29: 'o'
30: 'r'
31: 'm'
32: 'i'
33: 'n'
34: 't'
35: 'e'
36: 'g'
37: 'e'
38: 'r'
39: 'b'
40: 'o'
41: 'o'
42: 'l'
43: 'e'
44: 'a'
45: 'n'
46: 's'
47: 't'
48: 'r'
49: 'i'
50: 'n'
51: 'g'
52: 't'
53: 'r'
54: 'u'
55: 'e'
56: 'f'
57: 'a'
58: 'l'
59: 's'
60: 'e'
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
79: '0'-'9'
80: '0'-'9'
81: 'a'-'z'
82: 'A'-'Z'
83: 'a'-'z'
84: 'A'-'Z'
85: '0'-'9'
86: .

*/
