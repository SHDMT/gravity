package lisp

//greater computes the result of two inputs by operation >
func greater(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int, Float, String:
		switch y.Kind {
		case Int, Float, String:
			if z, err := x.Cmp(&y); err != nil {
				return None, err
			} else if z > 0 {
				return True, nil
			} else {
				return False, nil
			}
		default:
			return None, ErrFitType
		}
	default:
		return None, ErrFitType
	}
}

//geq computes the result of two inputs by operation >=
func geq(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int, Float, String:
		switch y.Kind {
		case Int, Float, String:
			if z, err := x.Cmp(&y); err != nil {
				return None, err
			} else if z >= 0 {
				return True, nil
			} else {
				return False, nil
			}
		default:
			return None, ErrFitType
		}
	default:
		return None, ErrFitType
	}
}

//less computes the result of two inputs by operation <
func less(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int, Float, String:
		switch y.Kind {
		case Int, Float, String:
			if z, err := x.Cmp(&y); err != nil {
				return None, err
			} else if z < 0 {
				return True, nil
			} else {
				return False, nil
			}
		default:
			return None, ErrFitType
		}
	default:
		return None, ErrFitType
	}
}

//leq computes the result of two inputs by operation <=
func leq(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int, Float, String:
		switch y.Kind {
		case Int, Float, String:
			if z, err := x.Cmp(&y); err != nil {
				return None, err
			} else if z <= 0 {
				return True, nil
			} else {
				return False, nil
			}
		default:
			return None, ErrFitType
		}
	default:
		return None, ErrFitType
	}
}

//eql computes the result of two inputs by operation =
func eql(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int, Float, String:
		switch y.Kind {
		case Int, Float, String:
			if z, err := x.Cmp(&y); err != nil {
				return None, err
			} else if z == 0 {
				return True, nil
			} else {
				return False, nil
			}
		default:
			return None, ErrFitType
		}
	default:
		return None, ErrFitType
	}
}

//uneql computes the result of two inputs by operation !=
func uneql(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int, Float, String:
		switch y.Kind {
		case Int, Float, String:
			if z, err := x.Cmp(&y); err != nil {
				return None, err
			} else if z != 0 {
				return True, nil
			} else {
				return False, nil
			}
		default:
			return None, ErrFitType
		}
	default:
		return None, ErrFitType
	}
}
