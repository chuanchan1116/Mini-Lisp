package lexer

import (
	"github.com/chuanchan1116/mini-lisp/token"
)

type lexer struct {
	input  string
	start  int
	pos    int
	Tokens chan token.Token
}

var separator = map[byte]bool{
	' ':  true,
	'\r': true,
	'\t': true,
	'\n': true,
}

func Lex(s string) *lexer {
	l := &lexer{
		input:  s,
		start:  0,
		pos:    0,
		Tokens: make(chan token.Token, 20),
	}
	go l.run()
	return l
}

func (l *lexer) skipSeparator() {
	for separator[l.input[l.start]] && l.start < len(l.input) {
		l.start++
	}
}

func (l *lexer) syntaxErr() {
	for separator[l.input[l.pos-1]] && l.input[l.pos-1] != ')' {
		l.pos++
	}
	panic("Syntax error: Unknown token `" + l.input[l.start:l.pos] + "'.")
}

func (l *lexer) run() {
	for l.start != len(l.input) {
		tok := *new(token.Token)
		l.skipSeparator()

		l.pos = l.start + 1
		c := l.input[l.start]
		switch {
		case c == '(':
			tok.Type = token.LPARAM
		case c == ')':
			tok.Type = token.RPARAM
		case c == '+':
			tok.Type = token.PLUS
		case c == '-':
			if l.input[l.pos] >= '1' && l.input[l.pos] <= '9' {
				l.pos++
				for isDigit(l.input[l.pos]) {
					l.pos++
				}
				tok.Type = token.NUM
			} else {
				tok.Type = token.MINUS
			}
		case c == '*':
			tok.Type = token.MUL
		case c == '/':
			tok.Type = token.DIV
		case c == '>':
			tok.Type = token.GREATER
		case c == '<':
			tok.Type = token.SMALLER
		case c == '=':
			tok.Type = token.EQU
		case c == '#':
			if l.pos < len(l.input) && (l.input[l.pos] == 't' || l.input[l.pos] == 'f') {
				l.pos++
				tok.Type = token.BOOL
			} else {
				l.syntaxErr()
			}
		case c == '0':
			tok.Type = token.NUM
		case '1' <= c && c <= '9':
			for isDigit(l.input[l.pos]) {
				l.pos++
			}
			tok.Type = token.NUM
		case isLower(c):
			for l.pos < len(l.input) && (isLower(l.input[l.pos]) || isDigit(l.input[l.pos]) || l.input[l.pos] == '-') {
				l.pos++
			}
			switch l.input[l.start:l.pos] {
			case "and":
				tok.Type = token.AND
			case "or":
				tok.Type = token.OR
			case "not":
				tok.Type = token.NOT
			case "mod":
				tok.Type = token.MOD
			case "define":
				tok.Type = token.DEFINE
			case "fun":
				tok.Type = token.FUNC
			case "if":
				tok.Type = token.IF
			case "print-num":
				tok.Type = token.PRINTNUM
			case "print-bool":
				tok.Type = token.PRINTBOOL
			default:
				tok.Type = token.ID
			}
		default:
			l.syntaxErr()
		}

		tok.Data = l.input[l.start:l.pos]
		l.start = l.pos
		l.Tokens <- tok
	}
	close(l.Tokens)
	return
}

func isLower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
