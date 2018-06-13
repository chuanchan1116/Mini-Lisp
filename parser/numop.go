package parser

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chuanchan1116/mini-lisp/token"
)

func (p *parser) equState() (ret token.Token) {
	ret.Type = token.BOOL
	a := p.value(<-p.t)
	var equ bool
	if a.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[a.Type])
		os.Exit(1)
	} else if a.Type == token.RPARAM {
		fmt.Printf("=: Need at least 2 arguments but got 0.\n")
		os.Exit(1)
	}

	argc := 1
	for i := range p.t {
		t := p.value(i)
		if t.Type == token.RPARAM {
			break
		} else if t.Type == token.NUM {
			argc++
			equ = (a == t)
		} else {
			fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[t.Type])
			os.Exit(1)
		}
	}
	if argc < 2 {
		fmt.Printf("*: Need at least 2 arguments, but got %d.\n", argc)
		os.Exit(1)
	}
	if equ {
		ret.Data = "#t"
	} else {
		ret.Data = "#f"
	}
	return
}

func (p *parser) compareState(op string) (ret token.Token) {
	ret.Type = token.BOOL
	ret.Data = "#f"
	a := p.value(<-p.t)
	if a.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[a.Type])
		os.Exit(1)
	} else if a.Type == token.RPARAM {
		fmt.Printf("%s: Need 2 arguments but got 0.\n", op)
		os.Exit(1)
	}

	b := p.value(<-p.t)
	if b.Type == token.BOOL {
		fmt.Printf("Type Error: Expecting `number' but got `%s'.\n", typeString[b.Type])
		os.Exit(1)
	} else if b.Type == token.RPARAM {
		fmt.Printf("%s: Need 2 arguments but got 1.\n", op)
		os.Exit(1)
	}

	av, _ := strconv.Atoi(a.Data)
	bv, _ := strconv.Atoi(b.Data)
	switch op {
	case ">":
		if av > bv {
			ret.Data = "#t"
		}
	case "<":
		if av < bv {
			ret.Data = "#t"
		}
	}

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		if rp.Type == token.NUM {
			fmt.Printf("%s: Needs 2 arguments but got more.\n", op)
			os.Exit(1)
		} else {
			fmt.Printf("Semantic error: Expecting `)' but got `%s'\n", typeString[rp.Type])
			os.Exit(1)
		}
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
