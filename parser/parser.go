package parser

import (
	"fmt"
	"os"

	"github.com/chuanchan1116/mini-lisp/token"
)

var typeString = map[token.TokenType]string{
	token.NUM:       "number",
	token.BOOL:      "boolean",
	token.ID:        "ID",
	token.DEFINE:    "define",
	token.FUNC:      "fun",
	token.IF:        "if",
	token.PLUS:      "+",
	token.MINUS:     "-",
	token.MUL:       "*",
	token.DIV:       "/",
	token.MOD:       "mod",
	token.GREATER:   ">",
	token.SMALLER:   "<",
	token.EQU:       "=",
	token.AND:       "and",
	token.OR:        "or",
	token.NOT:       "not",
	token.LPARAM:    "(",
	token.RPARAM:    ")",
	token.PRINTNUM:  "print-num",
	token.PRINTBOOL: "print-bool",
	token.NULL:      "NULL",
}

type parser struct {
	t      chan token.Token
	symbol map[string]token.Token
}

func Run(t chan token.Token) {
	var p parser
	p.t = t
	p.state()
	if i, ok := <-p.t; ok {
		fmt.Printf("Semantic error: Expecting END, got `%s'.\n", typeString[i.Type])
		os.Exit(1)
	}
	return
}

func (p *parser) state() {
	for t := range p.t {
		p.value(t)
	}
	return
}

func (p *parser) lparamState() (ret token.Token) {
	op := <-p.t
	switch op.Type {
	case token.PLUS:
		ret = p.plusState()
	case token.MINUS:
		ret = p.minusState()
	case token.MUL:
		ret = p.mulState()
	case token.DIV:
		ret = p.divState()
	case token.MOD:
		ret = p.modState()
	case token.GREATER, token.SMALLER:
		ret = p.compareState(op.Data)
	case token.EQU:
		ret = p.equState()
	case token.AND:
		ret = p.andState()
	case token.OR:
		ret = p.orState()
	case token.NOT:
		ret = p.notState()
	case token.DEFINE:
		ret = p.defineState()
	case token.PRINTNUM:
		ret = p.printNum()
	case token.PRINTBOOL:
		ret = p.printBool()
	default:
		fmt.Printf("Runtime error: Unimplemented token %s\n", op.Data)
		os.Exit(1)
	}
	return
}

func (p *parser) printBool() (ret token.Token) {
	ret.Type = token.NULL
	val := p.value(<-p.t)
	if val.Type != token.BOOL {
		fmt.Printf("Type Error: Expecting `boolean' but got `%s'.\n", typeString[val.Type])
		os.Exit(1)
	}
	fmt.Println(val.Data)
	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		fmt.Printf("Semantic error: Expecting `)' but got `%s'.\n", typeString[rp.Type])
		os.Exit(1)
	}
	return
}

func (p *parser) printNum() (ret token.Token) {
	ret.Type = token.NULL
	val := p.value(<-p.t)
	if val.Type != token.NUM {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[val.Type])
		os.Exit(1)
	}
	fmt.Println(val.Data)
	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		fmt.Printf("Semantic error: Expecting `)' but got `%s'.\n", typeString[rp.Type])
		os.Exit(1)
	}
	return
}

func (p *parser) value(t token.Token) (ret token.Token) {
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
