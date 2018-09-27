package lisp

import "testing"

func Test_Int(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(Int 2.5)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(2)}) {
		t.Errorf("The result should be 2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Int -3.2)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(-3)}) {
		t.Errorf("The result should be -3, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Int 99)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(99)}) {
		t.Errorf("The result should be 99, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Int "12")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(12)}) {
		t.Errorf("The result should be 12, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Int "15.62")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Int, Text: int64(15)}) {
		t.Errorf("The result should be 15, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Int '(a b))`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

	r, err = l.Eval([]byte(`(Int ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(Int 3 5)`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter number\n")
	}

	r, err = l.Eval([]byte(`(Int "haha")`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter\n")
	}

}

func Test_Float(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(Float 2)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(2)}) {
		t.Errorf("The result should be 2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Float -3.2)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(-3.2)}) {
		t.Errorf("The result should be -3.2, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Float 2.4e3)`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(2400)}) {
		t.Errorf("The result should be 2400, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Float "12")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(12)}) {
		t.Errorf("The result should be 12.0, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Float "15.62")`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: Float, Text: float64(15.62)}) {
		t.Errorf("The result should be 15.62, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(Float '(a b))`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

	r, err = l.Eval([]byte(`(Float ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(Float 3 5)`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter number\n")
	}

	r, err = l.Eval([]byte(`(Float "haha")`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter\n")
	}

}

func Test_Str2List(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(Str2List "go")`))
	if err != nil {
		t.Fatal(err)
	}
	if r.Kind != List {
		t.Errorf("The result should be a list, but got %v\n", r.Kind)
	}
	if !r.Text.([]Token)[0].Eq(&Token{Kind: Int, Text: int64('g')}) {
		t.Errorf("The first element of list should be h\n")
	}
	if !r.Text.([]Token)[1].Eq(&Token{Kind: Int, Text: int64('o')}) {
		t.Errorf("The first element of list should be h\n")
	}

	r, err = l.Eval([]byte(`(Str2List -3.2)`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

	r, err = l.Eval([]byte(`(Str2List ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(Str2List 3 5)`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter number\n")
	}

}

func Test_List2Str(t *testing.T) {
	l := NewLisp()
	r, err := l.Eval([]byte(`(List2Str '(104 97 104 97))`))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Eq(&Token{Kind: String, Text: "haha"}) {
		t.Errorf("The result should be haha, but got %v\n", r)
	}

	r, err = l.Eval([]byte(`(List2Str -3.2)`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

	r, err = l.Eval([]byte(`(List2Str '(100.5 2))`))
	if err == nil {
		t.Errorf("Error not checked, unsupported type\n")
	}

	r, err = l.Eval([]byte(`(List2Str ng)`))
	if err == nil {
		t.Errorf("Error not checked, error returned\n")
	}

	r, err = l.Eval([]byte(`(List2Str 3 5)`))
	if err == nil {
		t.Errorf("Error not checked, wrong parameter number\n")
	}

}
