package lisp

func init() {
	//implementation of the system function "atom" used for atom parameter judgement
	Add("atom", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind != List || len(x.Text.([]Token)) == 0 {
			return True, nil
		}
		return False, nil

	})
	//implementation of the system function "eq" used to judge if two parameters are equal
	Add("eq", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if x.Kind == Back || y.Kind == Back {
			return None, ErrFitType
		}
		if x.Eq(&y) {
			return True, nil
		}
		return False, nil

	})
	//implementation of the system function "car" used to get the first element of a list
	Add("car", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind == List {
			if len(x.Text.([]Token)) != 0 {
				return x.Text.([]Token)[0], nil
			}
			return None, ErrIsEmpty
		}
		return None, ErrFitType
	})
	//implementation of the system function "cdr" used to get the sublist(
	// include all elements except the firs one) of a list
	Add("cdr", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind == List {
			if len(x.Text.([]Token)) != 0 {
				return Token{Kind: List, Text: x.Text.([]Token)[1:]}, nil
			}
			return None, ErrIsEmpty
		}
		return None, ErrFitType
	})
	//implementation of the system function "cons" used to combine an elemnent and a list to a new list
	Add("cons", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if y.Kind == List {
			a := y.Text.([]Token)
			b := make([]Token, len(a)+1)
			b[0] = x
			copy(b[1:], a)
			return Token{Kind: List, Text: b}, nil
		}
		return None, ErrFitType
	})
	//implementation of the system function "eval" used for execution of parameter
	Add("eval", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		ans, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		return p.Exec(ans)
	})
	//implementation of the system function "quote" used to return the parameter directly without execution
	Add("quote", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		return t[0], nil
	})
}
