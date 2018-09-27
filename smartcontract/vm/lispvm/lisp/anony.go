package lisp

func init() {
	//implementation of the system function "lambda" used for anonymous function definition
	Add("lambda", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) < 2 {
			return None, ErrParaNum
		}
		a := t[0]
		if a.Kind != List {
			return None, ErrFitType
		}
		tokens := a.Text.([]Token)
		x := make([]Name, 0, len(tokens))
		for _, token := range tokens {
			if token.Kind != Label {
				return None, ErrNotName
			}
			x = append(x, token.Text.(Name))
		}
		ans = Token{Kind: Front, Text: &Lfac{Para: x, Text: t[1:], Make: p, FuncName: "1"}}
		return ans, nil
	})
}
