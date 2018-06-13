package parser

import (
	"fmt"
	"os"

	"github.com/chuanchan1116/mini-lisp/token"
)

func (p *parser) defineState() (ret token.Token) {
	ret.Type = token.NULL
	id := p.value(<-p.t)
	if id.Type != token.ID {
		fmt.Printf("Type Error: Expecting ID but got `%s'.\n", typeString[id.Type])
		os.Exit(1)
	}
	exp := p.value(<-p.t)
	if exp.Type != token.NUM && exp.Type != token.BOOL {
		fmt.Printf("Type Error: Expecting EXP but got `%s'.\n", typeString[exp.Type])
		os.Exit(1)
	}
	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		fmt.Printf("Semantic error: Expecting `)' but got `%s'.\n", typeString[rp.Type])
		os.Exit(1)
	}
	return
}
