package parser

import (
	"github.com/chuanchan1116/mini-lisp/token"
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
		panic("Semantic error: Expecting END, got `" + typeString[i.Type] + "'.")
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
		panic("Semantic error: Invalid token " + op.Data)
	}
	return
}

func (p *parser) modState() (ret token.Token) {
	ret.Type = token.NUM
	a := p.value(<-p.t)
	if a.Type == token.BOOL {
		panic("Type Error: Expecting `number' but got `" + typeString[a.Type] + "'.")
	} else if a.Type == token.RPARAM {
		panic("mod: Need 2 arguments but got 0.")
	}

	b := p.value(<-p.t)
	if b.Type == token.BOOL {
		panic("Type Error: Expecting `number' but got `" + typeString[b.Type] + "'.")
	} else if b.Type == token.RPARAM {
		panic("mod: Need 2 arguments but got 1.")
	}

	av, _ := strconv.Atoi(a.Data)
	bv, _ := strconv.Atoi(b.Data)
	ret.Data = strconv.Itoa(av % bv)

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		if rp.Type == token.NUM {
			panic("mod: Needs 2 arguments but got more.")
		} else {
			panic("Semantic error: Expecting `)' but got `" + typeString[rp.Type] + "'.")
		}
	}
	return
}

func (p *parser) divState() (ret token.Token) {
	ret.Type = token.NUM
	a := p.value(<-p.t)
	if a.Type == token.BOOL {
		panic("Type Error: Expecting `number' but got `" + typeString[a.Type] + "'.")
	} else if a.Type == token.RPARAM {
		panic("/: Need 2 arguments but got 0.")
	}

	b := p.value(<-p.t)
	if b.Type == token.BOOL {
		panic("Type Error: Expecting `number' but got `" + typeString[b.Type] + "'.")
	} else if b.Type == token.RPARAM {
		panic("/: Need 2 arguments but got 1.")
	}

	av, _ := strconv.Atoi(a.Data)
	bv, _ := strconv.Atoi(b.Data)
	ret.Data = strconv.Itoa(av / bv)

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		if rp.Type == token.NUM {
			panic("/: Needs 2 arguments but got more.")
		} else {
			panic("Semantic error: Expecting `)' but got `" + typeString[rp.Type] + "'.")
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
			panic("Type Error: Expecting `number' but got `" + typeString[t.Type] + "'.")
		}
	}
	if argc < 2 {
		panic("*: Need at least 2 arguments, but got " + strconv.Itoa(argc) + ".")
	}
	ret.Data = strconv.Itoa(val)
	return
}

func (p *parser) minusState() (ret token.Token) {
	ret.Type = token.NUM
	a := p.value(<-p.t)
	if a.Type == token.BOOL {
		panic("Type Error: Expecting `number' but got `" + typeString[a.Type] + "'.")
	} else if a.Type == token.RPARAM {
		panic("-: Need 2 arguments but got 0.")
	}

	b := p.value(<-p.t)
	if b.Type == token.BOOL {
		panic("Type Error: Expecting `number' but got `" + typeString[b.Type] + "'.")
	} else if b.Type == token.RPARAM {
		panic("-: Need 2 arguments but got 1.")
	}

	av, _ := strconv.Atoi(a.Data)
	bv, _ := strconv.Atoi(b.Data)
	ret.Data = strconv.Itoa(av - bv)

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		if rp.Type == token.NUM {
			panic("-: Needs 2 arguments but got more.")
		} else {
			panic("Semantic error: Expecting `)' but got `" + typeString[rp.Type] + "'.")
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
			panic("Type Error: Expecting `number' but got `" + typeString[t.Type] + "'.")
		}
	}
	if argc < 2 {
		panic("+: Need at least 2 arguments, but got " + strconv.Itoa(argc) + ".")
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
			panic("Undeclaired variable: " + t.Data + ".")
		}
	case token.LPARAM:
		ret = p.lparamState()
	default:
		panic("Semantic error: Expecting `number', `boolean', `variable' or EXP, got `" + typeString[t.Type] + "'.")
	}
	return
}
