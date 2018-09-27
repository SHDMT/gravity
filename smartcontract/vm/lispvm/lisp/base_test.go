package lisp

import (
	"fmt"
	"testing"
)

func Test_base(t *testing.T) {
	answer := []string{"unknown", "int", "float", "string",
		"fold list", "list", "go", "macro", "lisp", "Name",
		"operator"}
	for i := Null; i <= Operator; i++ {
		token := Token{Kind: i, Text: nil}
		s := fmt.Sprintf("%v", token.Kind)
		if s != answer[i] {
			t.Errorf("Wrong output of %v\n", i)
		}
	}

	l := NewLisp()
	r, err := l.Eval([]byte(`(define (inc a) (+ a 'a'))`))
	if err != nil {
		t.Fatal(err)
	}
	if r.String() != "{front : [a] => [[+ a 97]]}" {
		t.Errorf("Wrong string output of lfac\n")
	}
}
