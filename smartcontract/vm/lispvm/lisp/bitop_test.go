package lisp

import "testing"

func TestLisp_bitop(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(logand 1 1 1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The logand value should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(logior 1 0 0)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The logior value should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(logxor 1 0 0)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The logxor value should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(lognor 1 0 0)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The lognor value should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(logeqv 1 0 1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(0)}) {
		t.Errorf("The lognor value should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(lognot 1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(-2)}) {
		t.Errorf("The lognor value should be -2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(logior (logand 1 1 0) (logand 1 0 0 0 ) (lognot 0))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(-1)}) {
		t.Errorf("The value should be -1, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(lognot 1 0 1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: lognor should have only 1 parameter")
	}

	r, err = l.Eval([]byte(`(logand (ng) 0 1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}
	r, err = l.Eval([]byte(`(logand 1 (ng) 0)`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(logior 1 (ng) 0 1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(logior (ng) 0 1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(logxor 1 (ng) 0 1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(logxor (ng) 0 1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(lognor 1 (ng) 0 )`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(lognor (ng) 0  1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(logeqv 1 (ng) 0 )`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(logeqv (ng) 0 0)`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(lognot (ng) )`))
	if err == nil {
		t.Errorf("error should occur but not checked: one of paras should return error")
	}

	r, err = l.Eval([]byte(`(logand 1 1.1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: log computation do not support float")
	}

	r, err = l.Eval([]byte(`(logior 1 1.1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: log computation do not support float")
	}

	r, err = l.Eval([]byte(`(logxor 1 1.1)`))
	if err == nil {
		t.Errorf("error should occur but not checked: log computation do not support float")
	}

	r, err = l.Eval([]byte(`(lognor 1 "haha")`))
	if err == nil {
		t.Errorf("error should occur but not checked: log computation do not support string")
	}

	r, err = l.Eval([]byte(`(logeqv 1 "haha")`))
	if err == nil {
		t.Errorf("error should occur but not checked: log computation do not support string")
	}

	r, err = l.Eval([]byte(`(lognot  "haha")`))
	if err == nil {
		t.Errorf("error should occur but not checked: log computation do not support string")
	}
}
