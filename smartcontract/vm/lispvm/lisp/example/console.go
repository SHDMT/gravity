package main

import (
	"os"

	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp"
	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp/math"
)

var env *lisp.Lisp

//main provides a console program to test lisp
func main() {
	l := len(os.Args) - 1
	if l > 0 {
		for _, f := range os.Args[1:l] {
			env.Eval([]byte(`(load "` + f + `")`))
		}
		if os.Args[l] != "-" {
			env.Eval([]byte(`(println (load "` + os.Args[l] + `"))`))
			return
		}
	}
	env.Eval([]byte(`
	(loop
		()
		1
		(each
			(println "==>")
			(catch
				(error
					(println (scan))
				)
			)
		)
	)`))
}
func init() {

	lisp.Add("sin", math.Sin)
	lisp.Add("sinh", math.Sinh)
	lisp.Add("asin", math.Asin)
	lisp.Add("asinh", math.Asinh)
	lisp.Add("cos", math.Cos)
	lisp.Add("cosh", math.Cosh)
	lisp.Add("acos", math.Acos)
	lisp.Add("acosh", math.Acosh)
	lisp.Add("tan", math.Tan)
	lisp.Add("tanh", math.Tanh)
	lisp.Add("atan", math.Atan)
	lisp.Add("atanh", math.Atanh)
	lisp.Add("exp", math.Exp)
	lisp.Add("log", math.Log)
	lisp.Add("pow", math.Pow)
	lisp.Add("sqrt", math.Sqrt)

	env = lisp.NewLisp()
}
