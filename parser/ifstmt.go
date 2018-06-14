package parser

import (
	"fmt"

	"github.com/chuanchan1116/mini-lisp/token"
)

func (p *parser) ifState() (ret token.Token) {
	testExp := p.value(<-p.t)
	if testExp.Type != token.BOOL {
		fmt.Printf("Type Error: Expecting `boolean' but got `%s'.\n", typeString[testExp.Type])
	}
	thanExp := p.value(<-p.t)
	elseExp := p.value(<-p.t)

	if testExp.Data == "#t" {
		ret = thanExp
	} else {
		ret = elseExp
	}
	return
}
