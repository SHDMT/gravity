package lisp

func init() {
	Add(">", compareG)
	Add(">=", compareGeq)
	Add("<", compareLess)
	Add("<=", compareLeq)
	Add("==", compareEql)
	Add("=", compareEql)
	Add("!=", compareUneql)
	Add("/=", compareUneql)
}

//compareG returns the result of multiple inputs by operation >
//the result will be true if and only if every parameter is greater than next
func compareG(t []Token, p *Lisp) (Token, error) {
	l := len(t)
	if l < 2 {
		return None, ErrParaNum
	}
	var res Token
	var err error

	for i := 0; i < l-1; i++ {

		if res, err = greater(t[i], t[i+1], p); err != nil {
			return None, err
		}
		if !res.Bool() {
			return False, nil
		}
	}
	return True, nil
}

//compareGeq returns the result of multiple inputs by operation >=
//the result will be true if and only if every parameter is greater or equal than next
func compareGeq(t []Token, p *Lisp) (Token, error) {
	l := len(t)
	if l < 2 {
		return None, ErrParaNum
	}
	var res Token
	var err error

	for i := 0; i < l-1; i++ {

		if res, err = geq(t[i], t[i+1], p); err != nil {
			return None, err
		}
		if !res.Bool() {
			return False, nil
		}
	}
	return True, nil
}

//compareLess returns the result of multiple inputs by operation <
//the result will be true if and only if every parameter is less than next
func compareLess(t []Token, p *Lisp) (Token, error) {
	l := len(t)
	if l < 2 {
		return None, ErrParaNum
	}
	var res Token
	var err error

	for i := 0; i < l-1; i++ {

		if res, err = less(t[i], t[i+1], p); err != nil {
			return None, err
		}
		if !res.Bool() {
			return False, nil
		}
	}
	return True, nil
}

//compareGeq returns the result of multiple inputs by operation <=
//the result will be true if and only if every parameter is less or equal than next
func compareLeq(t []Token, p *Lisp) (Token, error) {
	l := len(t)
	if l < 2 {
		return None, ErrParaNum
	}
	var res Token
	var err error

	for i := 0; i < l-1; i++ {

		if res, err = leq(t[i], t[i+1], p); err != nil {
			return None, err
		}
		if !res.Bool() {
			return False, nil
		}
	}
	return True, nil
}

//compareGeq returns the result of multiple inputs by operation =
//the result will be true if and only if every parameter is equal
func compareEql(t []Token, p *Lisp) (Token, error) {
	l := len(t)
	if l < 2 {
		return None, ErrParaNum
	}
	var res Token
	var err error

	for i := 0; i < l-1; i++ {

		if res, err = eql(t[i], t[i+1], p); err != nil {
			return None, err
		}
		if !res.Bool() {
			return False, nil
		}
	}
	return True, nil
}

//compareGeq returns the result of multiple inputs by operation !=
//the result will be true if and only if every parameter is unequal
func compareUneql(t []Token, p *Lisp) (Token, error) {
	l := len(t)
	if l < 2 {
		return None, ErrParaNum
	}
	var x Token
	var err error
	nMap := make(map[float64]bool)
	sMap := make(map[string]bool)
	for i := 0; i < l; i++ {
		if x, err = p.Exec(t[i]); err != nil {
			return None, err
		}
		switch x.Kind {
		case Int:
			n := x.Text.(int64)
			if ok := nMap[float64(n)]; ok {
				return False, nil
			}
			nMap[float64(n)] = true
		case Float:
			n := x.Text.(float64)
			if ok := nMap[n]; ok {
				return False, nil
			}
			nMap[n] = true
		case String:
			n := x.Text.(string)
			if ok := sMap[n]; ok {
				return False, nil
			}
			sMap[n] = true
		default:
			return None, ErrFitType
		}
	}
	return True, nil
}
