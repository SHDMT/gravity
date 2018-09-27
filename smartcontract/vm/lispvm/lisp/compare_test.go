package lisp

import "testing"

func TestLisp_compare(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(> 5 3)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(> 5 5)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(< -1 -0.5)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(< 10 1e1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(= -1 -0.5)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(== 10 1e1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(!= -1 -0.5)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(/= 10 1e1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(>= 2 2)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(>= 0 0.1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(<= -1 (* -1 1))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(<= (+ 2 5) 6)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(> 5 3 1)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(< 1 3 4)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(>= 5 3 4)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(<= 3 2 4)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(== 1 1.0 1e0)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(!= 3 3 4)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(== "wait" "wait" "wait")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&True) {
		t.Errorf("The result should be true, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(/= "wait" "wait" "wait2")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&False) {
		t.Errorf("The result should be false, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(> 3 "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(> (ng) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(> "wait2" (ng))`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(> '(1 2) 0)`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(> 3 0 '(5 6))`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(> 3)`))
	if err == nil {
		t.Errorf("Error not checked: too less parameter")
	}

	r, err = l.Eval([]byte(`(< 3 "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(< (ng) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(< 1 3 (ng))`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(< '(5 6) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(< 1 3 '(0))`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(< 3)`))
	if err == nil {
		t.Errorf("Error not checked: too less parameter")
	}

	r, err = l.Eval([]byte(`(= 3 "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(== 3)`))
	if err == nil {
		t.Errorf("Error not checked: too less parameter")
	}

	r, err = l.Eval([]byte(`(== (ng) 3 3)`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(== 3 (ng) 3)`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(== '(1 2) 3)`))
	if err == nil {
		t.Errorf("Error not checked: unsuppored type")
	}

	r, err = l.Eval([]byte(`(== 3 3 '(1))`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(!= (ng) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(!= '(2 3 3) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(/= 3)`))
	if err == nil {
		t.Errorf("Error not checked: too less parameter")
	}

	r, err = l.Eval([]byte(`(>= 3 "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(>= 3)`))
	if err == nil {
		t.Errorf("Error not checked: too less parameter")
	}

	r, err = l.Eval([]byte(`(>= (ng) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(>= 3 1 (ng))`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(>= '(5 6) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(>= 3 1 '(0))`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(<= 3 "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(<= 3)`))
	if err == nil {
		t.Errorf("Error not checked: too less parameter")
	}

	r, err = l.Eval([]byte(`(<= (ng) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(<= 1 3 (ng))`))
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	r, err = l.Eval([]byte(`(<= '(5 6) "wait2")`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	r, err = l.Eval([]byte(`(<= 1 3 '(0))`))
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

}
