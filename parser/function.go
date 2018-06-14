package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chuanchan1116/mini-lisp/token"
)

type funcExp struct {
	Args []token.Token
	Body []token.Token
}

func (p *parser) funcState() (ret token.Token) {
	op := <-p.t
	if op.Type == token.FUNC {
		ret = p.funcDefineState()
	} else if op.Type == token.LPARAM {
		ret = p.lparamState()
	} else {
		fmt.Printf("Semantic error: Expecting `fun' or FUN-CALL but got `%s'.\n", typeString[op.Type])
		os.Exit(1)
	}

	return
}

func (p *parser) funcDefineState() (ret token.Token) {
	ret.Type = token.FUNCEXP
	lp := <-p.t
	if lp.Type != token.LPARAM {
		fmt.Printf("Semantic error: Expecting FUN-IDs but got `%s'.\n", typeString[lp.Type])
		os.Exit(1)
	} else {
		var f funcExp
		for v := range p.t {
			if v.Type == token.RPARAM {
				break
			} else if v.Type == token.ID {
				f.Args = append(f.Args, v)
			} else {
				fmt.Printf("Semantic error: Expecting ID but got `%s'.\n", typeString[lp.Type])
				os.Exit(1)
			}
		}
		funBody := <-p.t
		f.Body = append(f.Body, funBody)
		if funBody.Type == token.LPARAM {
			level := 1
			for v := range p.t {
				f.Body = append(f.Body, v)
				if v.Type == token.RPARAM {
					level--
				} else if v.Type == token.LPARAM {
					level++
				}
				if level == -1 {
					break
				}
			}
		}

		ret.Data = funEncode(f)
	}
	return
}

func (p *parser) funcCallState(in string) (ret token.Token) {
	var subp parser
	subp.symbol = make(map[string]token.Token)
	for k, v := range p.symbol {
		subp.symbol[k] = v
	}
	f := funDecode(in)
	for i, v := range f.Args {
		arg := p.defineValue(<-p.t)
		if arg.Type == token.RPARAM {
			ret.Type = token.FUNCEXP
			for j, v := range f.Body {
				if v.Type == token.ID {
					if k, ok := subp.symbol[v.Data]; ok {
						f.Body[j] = k
					}
				}
			}
			f.Args = f.Args[i:]
			ret.Data = funEncode(f)
			return
		}
		subp.symbol[v.Data] = arg
	}

	rp := p.value(<-p.t)
	if rp.Type != token.RPARAM {
		fmt.Printf("Semantic error: Expecting `)' but got `%s'.\n", typeString[rp.Type])
		os.Exit(1)
	}

	t := make(chan token.Token, 10)
	go func() {
		for _, v := range f.Body {
			t <- v
		}
		close(t)
	}()
	subp.t = t
	ret = subp.state()
	return
}

func funEncode(t funcExp) string {
	ret, _ := json.Marshal(t)
	return string(ret)
}

func funDecode(s string) (ret funcExp) {
	json.Unmarshal([]byte(s), &ret)
	return
}
