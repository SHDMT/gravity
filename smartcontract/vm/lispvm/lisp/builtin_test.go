package lisp

import (
	"testing"
)

func TestLisp_builtin(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(atom 1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(atom "haha")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(atom (+ 2.5e5 1))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(atom ())`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(atom '(+ 1 1))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(atom 3 "haha")`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters must be 1")
	}

	r, err = l.Eval([]byte(`(atom (1 2))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should return an error")
	}

	r, err = l.Eval([]byte(`(eq 3 (+ 1 2))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(eq "hoho" "haha")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be True, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(eq "hehe")`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should be 2\n")
	}

	r, err = l.Eval([]byte(`(eq (ng) 1.5)`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should return an error\n")
	}

	r, err = l.Eval([]byte(`(eq 1.5 (ng))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should return an error\n")
	}

	r, err = l.Eval([]byte(`(eq println println)`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should return an error\n")
	}

	r, err = l.Eval([]byte(`(quote (a b))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: List, Text: []Token{{Kind: Label, Text: Name("a")}, {Kind: Label, Text: Name("b")}}}) {
		t.Errorf("Result is not right\n")
	}

	r, err = l.Eval([]byte(`(quote a)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Label, Text: Name("a")}) {
		t.Errorf("Result is not right\n")
	}

	r, err = l.Eval([]byte(`(quote 2 2)`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters number should be 1\n")
	}

	r, err = l.Eval([]byte(`(define a 3)
	(eval (quote a))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(3)}) {
		t.Errorf("The result should be 3, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(eval 2 2)`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters number should be 1\n")
	}

	r, err = l.Eval([]byte(`(eval (ng))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters number should be 1\n")
	}

	r, err = l.Eval([]byte(`(car '(1 2 3))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(1)}) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(cdr '(1 2 3))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: List, Text: []Token{{Kind: Int, Text: int64(2)}, {Kind: Int, Text: int64(3)}}}) {
		t.Errorf("The result is not right\n")
	}

	r, err = l.Eval([]byte(`(cons 1 '(2 3))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: List, Text: []Token{{Kind: Int, Text: int64(1)}, {Kind: Int, Text: int64(2)}, {Kind: Int,
		Text: int64(3)}}}) {
		t.Errorf("The result is not right\n")
	}

	r, err = l.Eval([]byte(`(car '(a) '(b))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters number should be 1\n")
	}

	r, err = l.Eval([]byte(`(car (ng))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameter should return an error 1\n")
	}

	r, err = l.Eval([]byte(`(car (+ 1 1))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should be a list\n")
	}

	r, err = l.Eval([]byte(`(car '())`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should be a list\n")
	}

	r, err = l.Eval([]byte(`(cdr '(a) '(b))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters number should be 1\n")
	}

	r, err = l.Eval([]byte(`(cdr (ng))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameter should return an error 1\n")
	}

	r, err = l.Eval([]byte(`(cdr (+ 1 1))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should be a list\n")
	}

	r, err = l.Eval([]byte(`(cdr '())`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should be a list\n")
	}

	r, err = l.Eval([]byte(`(cons 1)`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters number should be 2\n")
	}

	r, err = l.Eval([]byte(`(cons (ng) '(a))`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameter should return an error\n")
	}

	r, err = l.Eval([]byte(`(cons '(a) (ng) )`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters should return an error\n")
	}

	r, err = l.Eval([]byte(`(cons 1 2)`))
	if err == nil {
		t.Errorf("Error should occur but not checked: parameters 2 should be a list\n")
	}
}
