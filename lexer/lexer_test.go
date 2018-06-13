package lexer

import (
	"testing"

	"github.com/chuanchan1116/mini-lisp/token"
)

var s = "() 123 -123 0 x xyz123 #t #f define fun if + - * / mod > < = and or not print-num print-bool"

func TestLex(t *testing.T) {
	l := Lex(s)
	expected := []token.Token{
		{token.LPARAM, "("},
		{token.RPARAM, ")"},
		{token.NUM, "123"},
		{token.NUM, "-123"},
		{token.NUM, "0"},
		{token.ID, "x"},
		{token.ID, "xyz123"},
		{token.BOOL, "#t"},
		{token.BOOL, "#f"},
		{token.DEFINE, "define"},
		{token.FUNC, "fun"},
		{token.IF, "if"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MUL, "*"},
		{token.DIV, "/"},
		{token.MOD, "mod"},
		{token.GREATER, ">"},
		{token.SMALLER, "<"},
		{token.EQU, "="},
		{token.AND, "and"},
		{token.OR, "or"},
		{token.NOT, "not"},
		{token.PRINTNUM, "print-num"},
		{token.PRINTBOOL, "print-bool"},
	}

	for _, v := range expected {
		got := <-l.Tokens
		if v != got {
			t.Errorf("Expected %v, got %v\n", v, got)
		}
	}
}
