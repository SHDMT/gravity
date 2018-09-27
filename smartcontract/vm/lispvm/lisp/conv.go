package lisp

import (
	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp/parser"
)

func init() {
	Add("Int", convInt)
	Add("Float", convFloat)
	Add("Str2List", str2List)
	Add("List2Str", list2String)
}

//convInt converts input to integer
//input with type int, float or string is valid
func convInt(t []Token, p *Lisp) (Token, error) {
	if len(t) != 1 {
		return None, ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return None, err
	}
	switch u.Kind {
	case Int:
		return u, nil
	case Float:
		return Token{Kind: Int, Text: int64(u.Text.(float64))}, nil
	case String:

		a, b := parser.ParseInt([]byte(u.Text.(string)))
		if b == 0 {
			return None, ErrNotConv
		}
		return Token{Kind: Int, Text: a}, nil
	}
	return None, ErrFitType
}

//convFloat convers input to float
//input with type int, float or string is valid
func convFloat(t []Token, p *Lisp) (Token, error) {
	if len(t) != 1 {
		return None, ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return None, err
	}
	switch u.Kind {
	case Int:
		return Token{Kind: Float, Text: float64(u.Text.(int64))}, nil
	case Float:
		return u, nil
	case String:
		a, b := parser.ParseFloat([]byte(u.Text.(string)))
		if b == 0 {
			return None, ErrNotConv
		}
		return Token{Kind: Float, Text: a}, nil
	}
	return None, ErrFitType
}

//str2List converts string input to list
func str2List(t []Token, p *Lisp) (Token, error) {
	if len(t) != 1 {
		return None, ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return None, err
	}
	if u.Kind != String {
		return None, ErrFitType
	}
	s := u.Text.(string)
	x := make([]Token, 0, len(s))
	for _, c := range s {
		x = append(x, Token{Kind: Int, Text: int64(c)})
	}
	return Token{Kind: List, Text: x}, nil
}

//list2String converts list input to string
func list2String(t []Token, p *Lisp) (Token, error) {
	if len(t) != 1 {
		return None, ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return None, err
	}
	if u.Kind != List {
		return None, ErrFitType
	}
	s := u.Text.([]Token)
	x := make([]rune, 0, len(s))
	for _, c := range s {
		if c.Kind != Int {
			return None, ErrFitType
		}
		x = append(x, rune(c.Text.(int64)))
	}
	return Token{Kind: String, Text: string(x)}, nil
}
