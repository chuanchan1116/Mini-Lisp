package main

import (
	"fmt"
	"github.com/chuanchan1116/mini-lisp/lexer"
	"github.com/chuanchan1116/mini-lisp/parser"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
	} else {
		f, err := ioutil.ReadFile(args[1])
		if err != nil {
			panic(err)
		}
		l := lexer.Lex(string(f))
		parser.Run(l.Tokens)
	}
}
