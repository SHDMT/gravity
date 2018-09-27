package lisp

func init() {
	//implementation of the system function "builtin" used to judge if the parameter is a system function
	//if it is, an address of RAM will be returned
	Add("builtin", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		ans, ok := Global.env[t[0].Text.(Name)]
		if !ok {
			return None, ErrNotFind
		}
		return ans, nil
	})

	//implementation of the system function "define" used to define a variable or a function
	Add("define", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) < 2 {
			return None, ErrParaNum
		}
		a := t[0]
		switch a.Kind {
		case Label:
			if len(t) != 2 {
				return None, ErrParaNum
			}

			ans, err = p.Exec(t[1])
			if err == nil {
				p.env[a.Text.(Name)] = ans
			}
			return ans, err
		case List:
			paras := a.Text.([]Token)
			if len(paras) <= 0 {
				return None, ErrParaNum
			}
			x := make([]Name, len(paras))
			for i, c := range paras {
				if c.Kind != Label {
					return None, ErrNotName
				}
				x[i] = c.Text.(Name)
			}
			ans = Token{Kind: Front, Text: &Lfac{Para: x[1:], Text: t[1:], Make: p, FuncName: x[0]}}
			p.env[x[0]] = ans
			return ans, nil
		}
		return None, ErrFitType
	})

	//implementation of the system function "update" used to update data of a variable
	//if the variable is not defined, error will be returned
	Add("update", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		a, b := t[0], t[1]
		var n Name
		switch a.Kind {
		case Label:
			n = a.Text.(Name)
		case List:
			if b.Kind != List {
				return None, ErrFitType
			}
			t = a.Text.([]Token)
			if len(t) <= 0 {
				return None, ErrParaNum
			}
			n = t[0].Text.(Name)
		default:
			return None, ErrFitType
		}
		for v := p; p != Global; p = p.parent {
			_, ok := p.env[n]
			if ok {
				if a.Kind == Label {
					ans, err = p.Exec(b)
					if err == nil {
						p.env[n] = ans
					}
					return ans, err
				}
				x := make([]Name, len(t)-1)
				for i, c := range t[1:] {
					if c.Kind != Label {
						return None, ErrNotName
					}
					x[i] = c.Text.(Name)
				}
				ans = Token{Kind: Front, Text: &Lfac{Para: x, Text: b.Text.([]Token), Make: p, FuncName: n}}
				v.env[n] = ans
				return ans, nil
			}
		}
		_, ok := p.env[n]
		if !ok {
			return None, ErrNotFind
		}
		return None, ErrRefused
	})

	//implementation of the system function "remove" used to remove a function or variable definition
	//system function definition can also be removed
	Add("remove", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		n := t[0].Text.(Name)
		for ; p != Global; p = p.parent {
			_, ok := p.env[n]
			if ok {
				delete(p.env, n)
				return None, nil
			}
		}
		_, ok := p.env[n]
		if !ok {
			return None, ErrNotFind
		}
		return None, ErrRefused
	})

	//implementation of the system function "present" used to list all the user defined function and variable names
	Add("present", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		x := make([]Token, 0, len(p.env))
		for i := range p.env {
			x = append(x, Token{Kind: Label, Text: i})
		}
		return Token{Kind: List, Text: x}, nil
	})

	//implementation of the system function "context" used to list all the function and variable names
	//(both system functions and user defined functions and variables)
	Add("context", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		x := make([]Token, 0, 128)
		for v := p; v != nil; v = v.parent {
			for i := range v.env {
				x = append(x, Token{Kind: Label, Text: i})
			}
		}
		return Token{Kind: List, Text: x}, nil
	})

	//implementation of the system function "clear" used to clear all the user defined functions and variables
	Add("clear", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		p.env = map[Name]Token{}
		return None, nil
	})
}
