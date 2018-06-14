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

	if testExp.Data == "#t" {
		thanExp := p.value(<-p.t)
		p.skipExp()
		ret = thanExp
	} else {
		p.skipExp()
		elseExp := p.value(<-p.t)
		ret = elseExp
	}
	return
}

func (p *parser) skipExp() {
	t := <-p.t
	if t.Type == token.LPARAM {
		level := 1
		for i := range p.t {
			if i.Type == token.RPARAM {
				level--
			} else if i.Type == token.LPARAM {
				level++
			}
			if level == 0 {
				break
			}
		}
	}
}
