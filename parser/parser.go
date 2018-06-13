package parser

import (
	"fmt"
	"github.com/chuanchan1116/mini-lisp/token"
	"os"
	"strconv"
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
}

type parser struct {
	t      chan token.Token
	symbol map[string]token.Token
}

func Run(t chan token.Token) (ret chan token.Token) {
	var p parser
	p.t = t
	for i := range p.state() {
		ret <- i
	}
	if i, ok := <-p.t; ok {
		fmt.Printf("Semantic error: Expecting END, got `%s'.\n", typeString[i.Type])
		os.Exit(1)
	}
	close(ret)
	return
}

func (p *parser) state() (ret chan token.Token) {
	for t := range p.t {
		ret <- p.value(t)
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
	default:
		fmt.Printf("Semantic error: Invalid token %s\n", op.Data)
		os.Exit(1)
	}
	return
}

func (p *parser) modState() (ret token.Token) {
	ret.Type = token.NUM
	a := p.value(<-p.t)
	if a.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[a.Type])
		os.Exit(1)
	} else if a.Type == token.RPARAM {
		fmt.Printf("mod: Need 2 arguments but got 0.\n")
		os.Exit(1)
	}

	b := p.value(<-p.t)
	if b.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[b.Type])
		os.Exit(1)
	} else if b.Type == token.RPARAM {
		fmt.Printf("mod: Need 2 arguments but got 1.\n")
		os.Exit(1)
	}

	av, _ := strconv.Atoi(a.Data)
	bv, _ := strconv.Atoi(b.Data)
	ret.Data = strconv.Itoa(av % bv)

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		if rp.Type == token.NUM {
			fmt.Printf("mod: Needs 2 arguments but got more.\n")
			os.Exit(1)
		} else {
			fmt.Printf("Semantic error: Expecting `)' but got `%s'\n", typeString[rp.Type])
			os.Exit(1)
		}
	}
	return
}

func (p *parser) divState() (ret token.Token) {
	ret.Type = token.NUM
	a := p.value(<-p.t)
	if a.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'\n", typeString[a.Type])
		os.Exit(1)
	} else if a.Type == token.RPARAM {
		fmt.Printf("/: Need 2 arguments but got 0.\n")
		os.Exit(1)
	}

	b := p.value(<-p.t)
	if b.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[b.Type])
		os.Exit(1)
	} else if b.Type == token.RPARAM {
		fmt.Printf("/: Need 2 arguments but got 1.\n")
		os.Exit(1)
	}

	av, _ := strconv.Atoi(a.Data)
	bv, _ := strconv.Atoi(b.Data)
	ret.Data = strconv.Itoa(av / bv)

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		if rp.Type == token.NUM {
			fmt.Printf("/: Needs 2 arguments but got more.\n")
			os.Exit(1)
		} else {
			fmt.Printf("Semantic error: Expecting `)' but got `%s'.\n", typeString[rp.Type])
			os.Exit(1)
		}
	}
	return
}

func (p *parser) mulState() (ret token.Token) {
	argc := 0
	val := 1
	ret.Type = token.NUM
	for i := range p.t {
		t := p.value(i)
		if t.Type == token.RPARAM {
			break
		} else if t.Type == token.NUM {
			argc++
			v, _ := strconv.Atoi(t.Data)
			val *= v
		} else {
			fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[t.Type])
			os.Exit(1)
		}
	}
	if argc < 2 {
		fmt.Printf("*: Need at least 2 arguments, but got %d.\n", argc)
		os.Exit(1)
	}
	ret.Data = strconv.Itoa(val)
	return
}

func (p *parser) minusState() (ret token.Token) {
	ret.Type = token.NUM
	a := p.value(<-p.t)
	if a.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[a.Type])
		os.Exit(1)
	} else if a.Type == token.RPARAM {
		fmt.Printf("-: Need 2 arguments but got 0.\n")
		os.Exit(1)
	}

	b := p.value(<-p.t)
	if b.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[b.Type])
		os.Exit(1)
	} else if b.Type == token.RPARAM {
		fmt.Printf("-: Need 2 arguments but got 1.\n")
		os.Exit(1)
	}

	av, _ := strconv.Atoi(a.Data)
	bv, _ := strconv.Atoi(b.Data)
	ret.Data = strconv.Itoa(av - bv)

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		if rp.Type == token.NUM {
			fmt.Printf("-: Needs 2 arguments but got more.\n")
			os.Exit(1)
		} else {
			fmt.Printf("Semantic error: Expecting `)' but got `%s'.\n", typeString[rp.Type])
			os.Exit(1)
		}
	}
	return
}

func (p *parser) plusState() (ret token.Token) {
	argc := 0
	val := 0
	ret.Type = token.NUM
	for i := range p.t {
		t := p.value(i)
		if t.Type == token.RPARAM {
			break
		} else if t.Type == token.NUM {
			argc++
			v, _ := strconv.Atoi(t.Data)
			val += v
		} else {
			fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[t.Type])
			os.Exit(1)
		}
	}
	if argc < 2 {
		fmt.Printf("+: Need at least 2 arguments, but got %d.\n", argc)
		os.Exit(1)
	}
	ret.Data = strconv.Itoa(val)
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
