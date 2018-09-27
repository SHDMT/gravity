package lisp

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

//Lisp struct describes a scope
//there is an unique global lisp scope which is the root scope
//all system functions are recorded in the global scope
//each running program will create a lisp instance as a child of global lisp
//each function/macro/block/loop will create a new lisp instance as a child of the program lisp
type Lisp struct {
	parent      *Lisp
	scopeName   Name
	env         map[Name]Token
	returnValue Token
}

//NewLisp returns a Lisp instance for a new running program
func NewLisp() *Lisp {
	x := new(Lisp)
	x.env = map[Name]Token{}
	x.parent = Global
	return x
}

//Add adds the system function implementation to the global Lisp instance
func Add(s string, f func([]Token, *Lisp) (Token, error)) {
	Global.env[Name(s)] = Token{Back, Gfac(f)}
}

//Exec complete a lisp interpretation
//input is a root of syntax tree given by a parser
//output is the return value
func (l *Lisp) Exec(f Token) (ans Token, err error) {
	if l.returnValue != None {
		return l.returnValue, nil
	}
	var (
		ls []Token
		ct Token
		ok bool
	)
	switch f.Kind {
	case Fold:
		return Token{List, f.Text.([]Token)}, nil
	case Label:
		nm := f.Text.(Name)
		for ; l != nil; l = l.parent {
			ct, ok = l.env[nm]
			if ok {
				return ct, nil
			}
		}
		return None, ErrNotFind
	case List:
		ls = f.Text.([]Token)
		if len(ls) == 0 {
			return False, nil
		}
		ct = ls[0]
		switch ct.Kind {
		case Label:
			nm := ct.Text.(Name)
			for v := l; v != nil; v = v.parent {
				ct, ok = v.env[nm]
				if ok {
					break
				}
			}
			if !ok {
				return None, ErrNotFind
			}
		case List:
			ct, err = l.Exec(ct)
			if err != nil {
				return None, err
			}
		}
		switch ct.Kind {
		case Back:
			return ct.Text.(Gfac)(ls[1:], l)
		case Macro:
			lp := ct.Text.(*Lfac)
			if len(ls) != len(lp.Para)+1 {
				return None, ErrParaNum
			}
			q := &Lisp{parent: lp.Make, env: map[Name]Token{}, scopeName: lp.FuncName}
			q.env[Name("self")] = ct
			for i, t := range ls[1:] {
				q.env[lp.Para[i]] = t
			}
			var v Token
			for _, body := range lp.Text {
				v, err = q.Exec(body)
				if err != nil {
					return None, err
				}
			}
			return l.Exec(v)
		case Front:
			lp := ct.Text.(*Lfac)
			if len(ls) != len(lp.Para)+1 {
				return None, ErrParaNum
			}
			q := &Lisp{parent: lp.Make, env: map[Name]Token{}, scopeName: lp.FuncName}
			q.env[Name("self")] = ct
			for i, t := range ls[1:] {
				q.env[lp.Para[i]], err = l.Exec(t)
				if err != nil {
					return None, err
				}
			}
			var v Token
			for _, body := range lp.Text {
				v, err = q.Exec(body)
				if err != nil {
					return None, err
				}
			}
			return v, nil

		default:
			return None, ErrNotFunc
		}
	default:
		return f, nil
	}
}

//Eval interprets the raw code completely, including lexical analysis, parsing and executing
func (l *Lisp) Eval(s []byte) (Token, error) {
	var (
		a, b []Token
		c, d Token
		e    error
	)
	a, e = Scan(s)
	if e != nil {
		return None, e
	}
	b, e = Tree(a)
	if e != nil {
		return None, e
	}
	for _, c = range b {
		d, e = l.Exec(c)
		if e != nil {
			return None, e
		}
	}
	return d, nil
}

//Load reads the raw code from a file and call Eval to interpret the code
func (l *Lisp) Load(s string) (Token, error) {
	var file *os.File
	var data []byte
	var err error
	file, err = os.Open(s)
	if err != nil {
		file, err = os.Open(s + ".lsp")
		if err != nil {
			return None, err
		}
	}
	defer file.Close()
	data, err = ioutil.ReadAll(file)
	if err != nil {
		return None, err
	}
	buf := bytes.NewBuffer(data)
	one := section{}
	for {
		data, err := buf.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				return None, err
			}
			err = one.feed(data)
			break
		}
		err = one.feed(data)
		if err != nil {
			return None, err
		}
	}
	if !one.over() {
		return None, ErrUnquote
	}
	return l.Eval([]byte(one.total))
}
