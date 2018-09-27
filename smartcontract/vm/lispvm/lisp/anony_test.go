package lisp

import (
	"testing"
)

func TestLisp_lambda(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`((lambda () (println 1) 2))`))
	if err != nil {
		t.Fatal(err)
	}
	two := Token{Kind: Int, Text: int64(2)}
	if !r.Eq(&two) {
		t.Errorf("The return value should be 2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`((lambda (a aa) (setq a (+ a 1)) (setq aa (- aa 1))(* a aa)) 5 9)`))
	if err != nil {
		t.Fatal(err)
	}

	if !r.Eq(&Token{Kind: Int, Text: int64(48)}) {
		t.Errorf("The return value should be 48, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`
	((lambda (a) (if (> a 10) a (self (+ a 1)))) 0)

	`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(11)}) {
		t.Errorf("The return value should be 48, but got %v\n", r)
	}

	_, err = l.Eval([]byte(`((lambda ()))`))
	if err == nil {
		t.Errorf("lambda should have at least 2 error")
	}

	_, err = l.Eval([]byte(`((lambda a (println a)))`))
	if err == nil {
		t.Errorf("lambda para list should be a list")
	}

	_, err = l.Eval([]byte(`((lambda ((println a) a) (println 1)) 2 4)`))
	if err == nil {
		t.Errorf("lambda para should be a label")
	}
}
