package lisp

func init() {
	//implementation of the system function "and" used to perform an "and" operation
	//the result will be true if and only if every parameter is non-zero/true
	Add("and", func(t []Token, p *Lisp) (Token, error) {
		if len(t) == 0 {
			return False, nil
		}
		res := False
		for _, v := range t {
			x, err := p.Exec(v)
			if err != nil {
				return None, err
			}
			if !x.Bool() {
				return False, nil
			}
			res = x
		}
		return res, nil
	})

	//implementation of the system function "or" used to perform an "or" operation
	//the result will be false if and only if every parameter is zero/false
	Add("or", func(t []Token, p *Lisp) (Token, error) {
		if len(t) == 0 {
			return False, nil
		}
		for _, v := range t {
			x, err := p.Exec(v)
			if err != nil {
				return None, err
			}
			if x.Bool() {
				return x, nil
			}
		}
		return False, nil
	})

	//implementation of the system function "or" used to perform an "xor" operation
	Add("xor", func(t []Token, p *Lisp) (Token, error) {
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
		if x.Bool() != y.Bool() {
			return True, nil
		}
		return False, nil

	})

	//implementation of the system function "not" used to perform an "not" operation
	Add("not", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Bool() {
			return False, nil
		}
		return True, nil

	})
}
