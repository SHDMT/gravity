package lisp

import (
	"math/rand"
	"time"
)

//Scan does the lexical analysis
//input is the raw data of go-lisp code
//output is result of lexical analysis with an array of token
func Scan(s []byte) (list []Token, err error) {
	scanner := pattern.NewScanner(s, true)
	list = make([]Token, 0, 100)
	for {
		a, b, c := scanner.Scan()
		if c != nil {
			break
		}
		switch b {
		case 1:
			list = append(list, Token{Kind: Operator, Text: a})
		case 2:
			list = append(list, Token{Kind: Int, Text: a})
		case 3:
			list = append(list, Token{Kind: Float, Text: a})
		case 4:
			list = append(list, Token{Kind: Int, Text: a})
		case 5:
			list = append(list, Token{Kind: String, Text: a})
		case 6:
			list = append(list, Token{Kind: Label, Text: a})
		}
	}
	if !scanner.Over() {
		err = ErrNotOver
	}
	return
}

//Tree works as a parser
//input is the result of Scan(lexical analysis) with an array of Token
//output is an array of root of a syntax tree
func Tree(tkn []Token) ([]Token, error) {
	var f Token
	var s int
	if len(tkn) == 0 {
		return nil, nil
	}
	if tkn[0].Kind == Operator {
		var t bool
		switch tkn[0].Text.(byte) {
		case '(':
			t = true
		case '[':
			t = false
		default:
			return nil, ErrUnquote
		}
		i, j, l := 1, 1, len(tkn)
		for i < l && j > 0 {
			if tkn[i].Kind == Operator {
				switch tkn[i].Text.(byte) {
				case '(', '[':
					j++
				case ')':
					j--
				}
			}
			i++
		}
		if j <= 0 {
			fold, err := Tree(tkn[1 : i-1])
			if err != nil {
				return nil, err
			}
			if t {
				f = Token{Text: fold, Kind: List}
			} else {
				f = Token{Text: fold, Kind: Fold}
			}
			s = i
		} else {
			return nil, ErrUnquote
		}
	} else {
		f = tkn[0]
		s = 1
	}
	rest, err := Tree(tkn[s:])
	if err != nil {
		return nil, err
	}
	ans := make([]Token, 1+len(rest))
	ans[0] = f
	copy(ans[1:], rest)
	return ans, nil
}

func init() {
	rand.Seed(time.Now().Unix())
}
