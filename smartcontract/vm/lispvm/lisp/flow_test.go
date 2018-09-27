package lisp

import "testing"

func Test_each(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(each
	(setq a 5)
	(setq b 2.7)
	(setq c (+ a b))
	(+ c 10)
	)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(17.7)}) {
		t.Errorf("The result should be 17.7, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(each
	(setq a 5)
	(setq b 2.7 2.8)
	(setq c (+ a b))
	(+ c 10)
	)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(each
	)`))
	if err == nil {
		t.Errorf("Error not checked, no parameter\n")
	}
}

func Test_block(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(block AA (setq n 1) n )`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(1)}) {
		t.Errorf("The result should be 1, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(block AA (defun add (a) (return-from AA a)) (add 2.5) )`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(2.5)}) {
		t.Errorf("The result should be 2.5, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(block AA)`))
	if err == nil {
		t.Errorf("Error not checked, too less parameter\n")
	}

	r, err = l.Eval([]byte(`(block (setq n 3)(+ 3 1))`))
	if err == nil {
		t.Errorf("Error not checked, label not found\n")
	}

	r, err = l.Eval([]byte(`(block AA (setq n 3)(ng))`))
	if err == nil {
		t.Errorf("Error not checked, label not found\n")
	}
}

func Test_if(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(if (> 3 1) 3.3 2.5)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(3.3)}) {
		t.Errorf("The result should be 3.3, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(if (> 3 4) 3.3 2.5)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(2.5)}) {
		t.Errorf("The result should be 2.5, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(if (> 2 1) (println 3)(println 4)(println 5))`))
	if err == nil {
		t.Errorf("Error not checked, too many parameters\n")
	}

	r, err = l.Eval([]byte(`(if (ng) (println 3)(println 4))`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}
}

func Test_cond(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(cond ((> 1 3) 7)((> 3 1) 8))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(8)}) {
		t.Errorf("The result should be 8, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(cond ((> 3 1) 7)((> 3 1) 8))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(7)}) {
		t.Errorf("The result should be 7, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(cond ((> 1 3) 7)((> 1 3) 8))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&None) {
		t.Errorf("The result should be none, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(cond println ((> 1 0)(println 1)))`))
	if err == nil {
		t.Errorf("Error not checked, parameter is not a list\n")
	}

	r, err = l.Eval([]byte(`(cond ((ng)(println 1)))`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(cond)`))
	if err == nil {
		t.Errorf("Error not checked, too less parameter\n")
	}

	r, err = l.Eval([]byte(`(cond ((> 3 1) (println 3) 7)((> 3 1) 8))`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter number\n")
	}
}

func Test_loop(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(loop (setq a 0)(<= a 10)(setq a (+ a 1)))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(11)}) {
		t.Errorf("The result should be 11, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(loop (ng)(> n 5)(+ n 1))`))
	if err == nil {
		t.Errorf("Error not checked, error returned \n")
	}

	r, err = l.Eval([]byte(`(loop (setq n 1)(> n 5)(+ n 1)(+ n 1))`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter \n")
	}

	r, err = l.Eval([]byte(`(loop (setq n 1)(ng)(+ n 1))`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(loop (setq n 1)(ng)(+ n 1))`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}
}

func Test_while(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(setq a 0)(while (<= a 10)(setq a (+ a 1)))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(11)}) {
		t.Errorf("The result should be 11, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(setq a 0)(while (<= a 10)(return 12))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(12)}) {
		t.Errorf("The result should be 12, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(while (ng)(+ 1 1))`))
	if err == nil {
		t.Errorf("Error not checked, error returned \n")
	}

	r, err = l.Eval([]byte(`(while (setq n 1)(> n 5)(+ n 1))`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter \n")
	}

	r, err = l.Eval([]byte(`(while (setq n 1)(ng))`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}
}

func Test_until(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(setq a 0)(until (> a 10)(setq a (+ a 1)))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(11)}) {
		t.Errorf("The result should be 11, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(setq a 0)(until (> a 10)(return 12))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(12)}) {
		t.Errorf("The result should be 12, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(until (ng)(+ 1 1))`))
	if err == nil {
		t.Errorf("Error not checked, error returned \n")
	}

	r, err = l.Eval([]byte(`(until (setq n 1)(< n 5)(+ n 1))`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter \n")
	}

	r, err = l.Eval([]byte(`(until (setq n 0)(ng))`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}
}

func Test_for(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(for a '(0 1 2) a)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(2)}) {
		t.Errorf("The result should be 2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(for a '(1 2 3 4 5) (if (== a 5) (return 99)))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(99)}) {
		t.Errorf("The result should be 99, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(for a '(1 2 3))`))
	if err == nil {
		t.Errorf("Error not checked, too less parameter \n")
	}

	r, err = l.Eval([]byte(`(for a (0 1 2) a)`))
	if err == nil {
		t.Errorf("Error not checked, error returned \n")
	}

	r, err = l.Eval([]byte(`(for a 3 a)`))
	if err == nil {
		t.Errorf("Error not checked, range is not a list\n")
	}

	r, err = l.Eval([]byte(`(for "string" '(1 2 3) a)`))
	if err == nil {
		t.Errorf("Error not checked, first para is not a label\n")
	}

	r, err = l.Eval([]byte(`(for a '(1 2 3) (ng))`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}
}
