package lisp

import (
	"fmt"

	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp/parser"
)

//Kind is used to be the type of Token's type
type Kind int

//Name is used to be the type of name of variables and functions
type Name string

//Gfac defines the function struct used for system go-lisp function
type Gfac func([]Token, *Lisp) (Token, error)

//Lfac defines the struct used for user defined function and macros
type Lfac struct {
	Para     []Name
	Text     []Token
	Make     *Lisp
	FuncName Name
}

//The followings define all kinds of Token
const (
	Null Kind = iota
	Int
	Float
	String
	Fold
	List
	Back
	Macro
	Front
	Label
	Operator
)

var (
	pattern = &parser.Pattern{}

	//True defines the general description of boolean true with an integer of data 1
	True = Token{Int, int64(1)}

	//False defines the general description of boolean false with an empty list
	False = Token{List, []Token(nil)}

	//None defines a Token with no data
	None = Token{}

	//Global is the root of lisp instance list, all system function are recorded in the env map of Global
	Global = &Lisp{env: map[Name]Token{}, parent: nil}
)

//String returns the description string of Kind
func (t Kind) String() string {
	switch t {
	case Int:
		return "int"
	case Float:
		return "float"
	case String:
		return "string"
	case Fold:
		return "fold list"
	case List:
		return "list"
	case Back:
		return "go"
	case Macro:
		return "macro"
	case Front:
		return "lisp"
	case Label:
		return "Name"
	case Operator:
		return "operator"
	}
	return "unknown"
}

//String returns the description string of a Lfac(user defined function)
func (l Lfac) String() string {
	return fmt.Sprintf("{front : %v => %v}", l.Para, l.Text)
}
