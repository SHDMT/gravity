package lisp

import "testing"

func Test_uneql(t *testing.T) {
	l := NewLisp()
	t1 := Token{Kind: Int, Text: int64(1)}
	t2 := Token{Kind: Int, Text: int64(3)}
	t3 := Token{Kind: Int, Text: int64(1)}

	result, err := uneql(t1, t2, l)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Eq(&True) {
		t.Errorf("The result should be true, bug got %v\n", result)
	}

	result, err = uneql(t1, t3, l)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Eq(&False) {
		t.Errorf("The result should be false, bug got %v\n", result)
	}

	t4 := Token{Kind: Label, Text: Name("a")}
	result, err = uneql(t1, t4, l)
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	result, err = uneql(t4, t1, l)
	if err == nil {
		t.Errorf("Error not checked: error returned")
	}

	t5 := Token{Kind: String, Text: "haha"}
	result, err = uneql(t1, t5, l)
	if err == nil {
		t.Errorf("Error not checked: different type")
	}

	t6 := Token{Kind: Fold, Text: []Token{t2, t3}}
	result, err = uneql(t1, t6, l)
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}

	result, err = uneql(t6, t1, l)
	if err == nil {
		t.Errorf("Error not checked: unsupported type")
	}
}
