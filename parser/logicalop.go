package parser

import (
	"fmt"
	"os"

	"github.com/chuanchan1116/mini-lisp/token"
)

func (p *parser) notState() (ret token.Token) {
	ret.Type = token.BOOL
	a := p.value(<-p.t)
	if a.Type == token.NUM {
		fmt.Printf("Type Error: Expecting `boolean' but got `%s'.\n", typeString[a.Type])
		os.Exit(1)
	} else if a.Type == token.RPARAM {
		fmt.Printf("=: Need at least 2 arguments but got 0.\n")
		os.Exit(1)
	}
	if a.Data == "#t" {
		ret.Data = "#f"
	} else {
		ret.Data = "#t"
	}

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		if rp.Type == token.BOOL {
			fmt.Printf("not: Needs 1 arguments but got more.\n")
			os.Exit(1)
		} else {
			fmt.Printf("Semantic error: Expecting `)' but got `%s'\n", typeString[rp.Type])
			os.Exit(1)
		}
	}
	return
}

func (p *parser) orState() (ret token.Token) {
	ret.Type = token.BOOL
	argc := 0
	val := false
	for i := range p.t {
		t := p.value(i)
		if t.Type == token.RPARAM {
			break
		} else if t.Type == token.BOOL {
			argc++
			val = val || (t.Data == "#t")
		} else {
			fmt.Printf("Type Error: Expecting `boolean' but got `%s'.\n", typeString[t.Type])
			os.Exit(1)
		}
	}
	if argc < 2 {
		fmt.Printf("or: Need at least 2 arguments, but got %d.\n", argc)
		os.Exit(1)
	}
	if val {
		ret.Data = "#t"
	} else {
		ret.Data = "#f"
	}
	return
}

func (p *parser) andState() (ret token.Token) {
	ret.Type = token.BOOL
	argc := 0
	val := true
	for i := range p.t {
		t := p.value(i)
		if t.Type == token.RPARAM {
			break
		} else if t.Type == token.BOOL {
			argc++
			val = val && (t.Data == "#t")
		} else {
			fmt.Printf("Type Error: Expecting `boolean' but got `%s'.\n", typeString[t.Type])
			os.Exit(1)
		}
	}
	if argc < 2 {
		fmt.Printf("and: Need at least 2 arguments, but got %d.\n", argc)
		os.Exit(1)
	}
	if val {
		ret.Data = "#t"
	} else {
		ret.Data = "#f"
	}
	return
}
