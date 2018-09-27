package lisp

import (
	"testing"
)

var t1 = Token{Kind: Int, Text: int64(1)}
var t2 = Token{Kind: Int, Text: int64(2)}
var t3 = Token{Kind: Int, Text: int64(3)}
var t4 = Token{Kind: Int, Text: int64(4)}

func Test_add(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(+ 5 3)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(8)}) {
		t.Errorf("The result should be 8, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(+ 5 3 4)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(12)}) {
		t.Errorf("The result should be 12, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(+ 5 1 1.2 1.7 2)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(10.9)}) {
		t.Errorf("The result should be 10.9, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(+ "haha" "hoho")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: String, Text: "hahahoho"}) {
		t.Errorf("The result should be hahahoho, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(+ '(1 2) '(3 4))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: List, Text: []Token{t1, t2, t3, t4}}) {
		t.Errorf("The result should be (1 2 3 4), but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(+ (1 2) (3 4))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: List, Text: []Token{t1, t2, t3, t4}}) {
		t.Errorf("The result should be (1 2 3 4), but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(+ 3 ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(+ ng 3)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(+ + 3)`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

}

func Test_minus(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(- 5 3)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(2)}) {
		t.Errorf("The result should be 2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(- 5 3 4)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(-2)}) {
		t.Errorf("The result should be -2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(- 5 3 4 -1.1 -1.2 3)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(-2.7)}) {
		t.Errorf("The result should be -2.7, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(- 3 "hoho")`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

	r, err = l.Eval([]byte(`(- 3 ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(- ng 3)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(- + 3)`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}
}

func Test_multiply(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(* 5 3)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(15)}) {
		t.Errorf("The result should be 15, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(* 5 3 1.2 1.5 2)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(54)}) {
		t.Errorf("The result should be 54, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(* 3 4 ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(* ng 3)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(* + 3)`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}
}

func Test_divide(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(/ 5 3)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(1)}) {
		t.Errorf("The result should be 1, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(/ 12 1 1.5 2.0 2)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(2)}) {
		t.Errorf("The result should be 2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(/ 3 4 ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(/ ng 3)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(/ + 3)`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

	r, err = l.Eval([]byte(`(/ 12 3 0)`))
	if err == nil {
		t.Errorf("Error not checked, dividor is zero\n")
	}

	r, err = l.Eval([]byte(`(/ 2.5 0)`))
	if err == nil {
		t.Errorf("Error not checked, dividor is zero\n")
	}

	r, err = l.Eval([]byte(`(/ 12 (+ 2.5 -2.5))`))
	if err == nil {
		t.Errorf("Error not checked, dividor is zero\n")
	}

	r, err = l.Eval([]byte(`(/ 12.5 2.5 (+ 2.5 -2.5))`))
	if err == nil {
		t.Errorf("Error not checked, dividor is zero\n")
	}
}

func Test_mod(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(% 19 3)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(1)}) {
		t.Errorf("The result should be 1, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(mod 19 13 5)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(1)}) {
		t.Errorf("The result should be 1, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(% 3.3 2)`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

	r, err = l.Eval([]byte(`(% ng 2)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(% 5 2 ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(% 5 0)`))
	if err == nil {
		t.Errorf("Error not checked, parameter 2 is zero\n")
	}
}
