package lisp

import "testing"

func Test_progn(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(progn
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

	r, err = l.Eval([]byte(`(progn
	(setq a 5)
	(setq b 2.7 2.8)
	(setq c (+ a b))
	(+ c 10)
	)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}
}

func Test_list(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(list
	(setq a 1)
	(setq b 2)
	(setq c (+ a b))
	4
	)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: List, Text: []Token{t1, t2, t3, t4}}) {
		t.Errorf("The result is not right\n")
	}

	r, err = l.Eval([]byte(`(list
	(setq a 5)
	(setq b 2.7 2.8)
	(setq c (+ a b))
	(+ c 10)
	)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(list)`))
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

}

func Test_length(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(length "abcde")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(5)}) {
		t.Errorf("The result should be 5, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(length '(a b c d e))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(5)}) {
		t.Errorf("The result should be 5, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(length)`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter length\n")
	}

	r, err = l.Eval([]byte(`(length ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(length 2.5)`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

}

func Test_setq(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(setq a 10)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(10)}) {
		t.Errorf("The result should be 10, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(setq a "haha")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: String, Text: "haha"}) {
		t.Errorf("The result should be haha, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(setq 2 2.5)`))
	if err == nil {
		t.Errorf("Error not checked, wrong type\n")
	}

}

func Test_defun(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(defun add (a b) (setq c 0) (+ a b c))
	(add 12.5 13.7)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(26.2)}) {
		t.Errorf("The result should be 26.2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(defun add (a b) 
	(defun inc (a) (+ a 1)) (+ (inc a) (inc b)))
	(add 12.5 13.7)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(28.2)}) {
		t.Errorf("The result should be 28.2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(loop (setq n 0)(< n 1)(progn
	(setq n (+ n 1))
	(define (add a) (defun add (c) (+ c 1)) (+ (add  a) 1))
	(add 3)
	))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(5)}) {
		t.Errorf("The result should be 5, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(defun add)`))
	if err == nil {
		t.Errorf("Error not checked, too less parameter\n")
	}

	r, err = l.Eval([]byte(`(defun "add" (a b) (+ a b))`))
	if err == nil {
		t.Errorf("Error not checked, wrong type\n")
	}

	r, err = l.Eval([]byte(`(defun add ("a" b) (+ a b))`))
	if err == nil {
		t.Errorf("Error not checked, wrong type\n")
	}

	r, err = l.Eval([]byte(`(defun add a (+ a 2))`))
	if err == nil {
		t.Errorf("Error not checked, wrong type\n")
	}

}

func Test_defmacro(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(defmacro add (a b) (setq c 0) (+ a b c))
	(add 12.5 13.7)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(26.2)}) {
		t.Errorf("The result should be 26.2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(defmacro add (a b) 
	(defun inc (a) (+ a 1)) (+ (inc a) (inc b)))
	(add 12.5 13.7)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(28.2)}) {
		t.Errorf("The result should be 28.2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(loop (setq n 0)(< n 1)(progn
	(setq n (+ n 1))
	(define (add a) (defmacro add (c) (+ (eval c) 1)) (+ (add a) 1))
	(add 3)
	))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(5)}) {
		t.Errorf("The result should be 5, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(defmacro add)`))
	if err == nil {
		t.Errorf("Error not checked, too less parameter\n")
	}

	r, err = l.Eval([]byte(`(defmacro "add" (a b) (+ a b))`))
	if err == nil {
		t.Errorf("Error not checked, wrong type\n")
	}

	r, err = l.Eval([]byte(`(defmacro add ("a" b) (+ a b))`))
	if err == nil {
		t.Errorf("Error not checked, wrong type\n")
	}

	r, err = l.Eval([]byte(`(defmacro add a (+ a 2))`))
	if err == nil {
		t.Errorf("Error not checked, wrong type\n")
	}

}

func Test_returnfrom(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(defun add () 3 (return-from add 4) 5)
	(add)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(4)}) {
		t.Errorf("The result should be 4, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(defun add (n) 
	(loop (setq x 1)(+ 1 1)(progn
	(setq x (+ x 1))
	(if (>= x 5)(return-from add x)))
	))
	(add 0)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(5)}) {
		t.Errorf("The result should be 5, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(defun add (n) 
	(loop (setq x 1)(+ 1 1)(progn
	(setq x (+ x 1))
	(if (>= x 5)(return-from add)))
	))
	(add 0)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(defun add()(setq x 1)(return-from add x 2))(add)`))
	if err == nil {
		t.Errorf("Error not checked, too many parameters\n")
	}

	r, err = l.Eval([]byte(`(defun add()(setq x 1)(return-from 1 2))(add)`))
	if err == nil {
		t.Errorf("Error not checked, return label is not a label\n")
	}

	r, err = l.Eval([]byte(`(defun add()(setq x 1)(return-from minus 2))(add)`))
	if err == nil {
		t.Errorf("Error not checked, label not found\n")
	}

}

func Test_return(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(loop(setq n 1)(< n 10)(return 3))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(3)}) {
		t.Errorf("The result should be 3, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(loop(setq n 1)(< n 10)(progn (
	loop (setq x 1)(< x 10)
	(return)

	)(setq n 11)))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(11)}) {
		t.Errorf("The result should be 11, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(loop (setq n 1)(< n 10)(return l 1))`))
	if err == nil {
		t.Errorf("Error not checked, too many parameters\n")
	}

	r, err = l.Eval([]byte(`(defun add()(return 3))(add)`))
	if err == nil {
		t.Errorf("Error not checked, loop not found\n")
	}

}
