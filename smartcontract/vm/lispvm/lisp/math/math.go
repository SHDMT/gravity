package math

import (
	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp"
	"math"
)
//These define math function in go-lisp
var (
	Sin, Sinh, Asin, Asinh lisp.Gfac
	Cos, Cosh, Acos, Acosh lisp.Gfac
	Tan, Tanh, Atan, Atanh lisp.Gfac
	Exp, Log               lisp.Gfac
	Sqrt                   lisp.Gfac
	Pow                    lisp.Gfac
)

//Wrap1 is a function conversion that converts a math function with form func(float64) float64 to Gfac
func Wrap1(f func(float64) float64) func([]lisp.Token, *lisp.Lisp) (lisp.Token, error) {
	return func(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
		if len(t) != 1 {
			return lisp.None, lisp.ErrParaNum
		}
		u, err := p.Exec(t[0])
		if err != nil {
			return lisp.None, err
		}
		if u.Kind != lisp.Float {
			if u.Kind == lisp.Int {
				return lisp.Token{Kind: lisp.Float, Text: f(float64(u.Text.(int64)))}, nil
			}
			return lisp.None, lisp.ErrFitType
		}
		return lisp.Token{Kind: lisp.Float, Text: f(u.Text.(float64))}, nil
	}
}

//Wrap2 is a function conversion that converts a math function with form func(float64, float64) float64 to Gfac
func Wrap2(f func(float64, float64) float64) func([]lisp.Token, *lisp.Lisp) (lisp.Token, error) {
	return func(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
		if len(t) != 2 {
			return lisp.None, lisp.ErrParaNum
		}
		u, err := p.Exec(t[0])
		if err != nil {
			return lisp.None, err
		}
		if u.Kind != lisp.Float {
			return lisp.None, lisp.ErrFitType
		}
		v, err := p.Exec(t[1])
		if err != nil {
			return lisp.None, err
		}
		if v.Kind != lisp.Float {
			return lisp.None, lisp.ErrFitType
		}
		return lisp.Token{Kind: lisp.Float, Text: f(u.Text.(float64), v.Text.(float64))}, nil
	}
}

func init() {
	Sin, Sinh, Asin, Asinh = Wrap1(math.Sin), Wrap1(math.Sinh), Wrap1(math.Asin), Wrap1(math.Asinh)
	Cos, Cosh, Acos, Acosh = Wrap1(math.Cos), Wrap1(math.Cosh), Wrap1(math.Acos), Wrap1(math.Acosh)
	Tan, Tanh, Atan, Atanh = Wrap1(math.Tan), Wrap1(math.Tanh), Wrap1(math.Atan), Wrap1(math.Atanh)
	Exp, Log = Wrap1(math.Exp), Wrap1(math.Log)
	Sqrt = Wrap1(math.Sqrt)
	Pow = Wrap2(math.Pow)
}
