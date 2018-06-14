package parser

import (
	"fmt"
	"os"

	"github.com/chuanchan1116/mini-lisp/token"
)

func (p *parser) defineState() (ret token.Token) {
	ret.Type = token.NULL
	id := <-p.t
	if id.Type != token.ID {
		fmt.Printf("Type Error: Expecting ID but got `%s'.\n", typeString[id.Type])
		os.Exit(1)
	}
	exp := p.defineValue(<-p.t)
	if exp.Type != token.NUM && exp.Type != token.BOOL && exp.Type != token.FUNCEXP {
		fmt.Printf("Type Error: Expecting EXP but got `%s'.\n", typeString[exp.Type])
		os.Exit(1)
	}
	p.symbol[id.Data] = exp
	rp := p.defineValue(<-p.t)
	if rp.Type != token.RPARAM {
		fmt.Printf("Semantic error: Expecting `)' but got `%s'.\n", typeString[rp.Type])
		os.Exit(1)
	}
	return
}

func (p *parser) defineValue(t token.Token) (ret token.Token) {
	switch t.Type {
	case token.NUM, token.BOOL, token.RPARAM:
		ret = t
	case token.ID:
		if v, ok := p.symbol[t.Data]; ok {
			ret = v
		} else {
			fmt.Printf("Undeclaired variable: `%s'.\n", t.Data)
			os.Exit(1)
		}
	case token.LPARAM:
		ret = p.lparamState()
	default:
		fmt.Printf("Semantic error: Expecting `number', `boolean', `variable' or EXP, got `%s'.\n", typeString[t.Type])
		os.Exit(1)
	}
	return
}
