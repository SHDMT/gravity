package lisp

func init() {
	//implementation of the system function "each" used to execute the parameter sequentially
	//same as "progn"
	Add("each", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) == 0 {
			return None, ErrParaNum
		}
		for _, i := range t {
			ans, err = p.Exec(i)
			if err != nil {
				break
			}
		}
		return ans, err
	})

	//implementation of the system function "block" used create a named scope
	Add("block", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) < 2 {
			return None, ErrParaNum
		}
		a := t[0]
		if a.Kind != Label {
			return None, ErrNotName
		}

		q := &Lisp{parent: p, env: map[Name]Token{}, scopeName: a.Text.(Name)}
		for _, i := range t[1:] {
			ans, err = q.Exec(i)
			if err != nil {
				break
			}
		}
		return ans, err
	})

	//implementation of the system function "if" used for condition flow control
	Add("if", func(t []Token, p *Lisp) (Token, error) {
		if len(t) < 2 || len(t) > 3 {
			return None, ErrParaNum
		}
		ans, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if ans.Bool() {
			return p.Exec(t[1])
		}
		if len(t) == 3 {
			return p.Exec(t[2])
		}
		return None, nil
	})

	//implementation of the system function "cond" used for condition flow control
	Add("cond", func(t []Token, p *Lisp) (Token, error) {
		if len(t) == 0 {
			return None, ErrParaNum
		}
		for _, i := range t {
			if i.Kind != List {
				return None, ErrFitType
			}
			t := i.Text.([]Token)
			if len(t) != 2 {
				return None, ErrParaNum
			}
			ans, err := p.Exec(t[0])
			if err != nil {
				return None, err
			}
			if ans.Bool() {
				return p.Exec(t[1])
			}
		}
		return None, nil
	})

	//implementation of the system function "while" used for condition flow control
	Add("while", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		q := &Lisp{parent: p, env: map[Name]Token{}, scopeName: "0"}
		var rv Token
		for {
			a, err := q.Exec(t[0])
			if err != nil {
				return None, err
			}
			if q.returnValue != None || !a.Bool() {
				break
			}
			rv, err = q.Exec(t[1])
			if err != nil {
				return None, err
			}
		}
		if q.returnValue == None {
			return rv, nil
		}
		return q.returnValue, nil
	})

	//implementation of the system function "until" used for condition flow control
	Add("until", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		q := &Lisp{parent: p, env: map[Name]Token{}, scopeName: "0"}
		var rv Token
		for {
			a, err := q.Exec(t[0])
			if err != nil {
				return None, err
			}
			if q.returnValue != None || a.Bool() {
				break
			}
			rv, err = q.Exec(t[1])
			if err != nil {
				return None, err
			}
		}
		if q.returnValue == None {
			return rv, nil
		}
		return q.returnValue, nil
	})

	//implementation of the system function "loop" used for condition flow control
	Add("loop", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 3 {
			return None, ErrParaNum
		}
		q := &Lisp{parent: p, env: map[Name]Token{}, scopeName: "0"}
		_, err := q.Exec(t[0])
		if err != nil {
			return None, err
		}
		var rv Token
		for {
			a, err := q.Exec(t[1])
			if err != nil {
				return None, err
			}
			if q.returnValue != None || !a.Bool() {
				break
			}
			rv, err = q.Exec(t[2])
			if err != nil {
				return None, err
			}
		}
		if q.returnValue == None {
			return rv, nil
		}
		return q.returnValue, nil
	})

	//implementation of the system function "for" used for condition flow control
	Add("for", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 3 {
			return None, ErrParaNum
		}
		q := &Lisp{parent: p, env: map[Name]Token{}, scopeName: "0"}
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		iter, err := q.Exec(t[1])
		if err != nil {
			return None, err
		}
		if iter.Kind != List {
			return None, ErrFitType
		}
		n := t[0].Text.(Name)
		var rv Token
		for _, m := range iter.Text.([]Token) {
			p.env[n] = m
			rv, err = q.Exec(t[2])
			if err != nil {
				return None, err
			}
		}
		if q.returnValue == None {
			return rv, nil
		}
		return q.returnValue, nil
	})
}
