package lisp

import (
	"errors"
	"fmt"
)

//The followings define all the error strings
var (
	ErrNotOver = errors.New("cannot scan to the end")
	ErrUnquote = errors.New("quote is unfold")
	ErrNotFind = errors.New("not find this Name")
	ErrNotFunc = errors.New("not a function")
	ErrParaNum = errors.New("wrong parament number")
	ErrFitType = errors.New("lisp type is wrong")
	ErrNotName = errors.New("this's not a Name")
	ErrIsEmpty = errors.New("fold is empty")
	ErrNotConv = errors.New("cannot translate")
	ErrRefused = errors.New("can't remove a back function")
	ErrDivZero = errors.New("cannot divide zero")
	ErrModZero = errors.New("cannot mod zero")
)

func init() {
	//implementation of the system function "raise" used to converts a string to an error info
	Add("raise", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		ans, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if ans.Kind != String {
			return None, ErrFitType
		}
		return None, fmt.Errorf(ans.Text.(string))
	})

	//implementation of the system function "catch" used to catch the error of executed parameter and make
	//the error returned as return value with string type
	Add("catch", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		_, err := p.Exec(t[0])
		if err != nil {
			return Token{Kind: String, Text: fmt.Sprint(err)}, nil
		}
		return None, nil
	})

	//implementation of the system function "error" used to clear the result if there is an error returned
	Add("error", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		ans, err := p.Exec(t[0])
		if err != nil {
			//fmt.Println(err)
			return None, err
		}
		return ans, nil
	})
}
