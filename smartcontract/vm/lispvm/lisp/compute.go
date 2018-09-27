package lisp

func init() {

	Add("+", computeAdd)
	Add("-", computeSub)
	Add("*", computeMul)
	Add("/", computeDiv)

	Add("%", computeMod)
	Add("mod", computeMod)
}

//computeAdd computes the result of multiple inputs by operation +
func computeAdd(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = addf(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeSub computes the result of multiple inputs by operation -
func computeSub(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = subf(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeMul computes the result of multiple inputs by operation *
func computeMul(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = mulf(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeDiv computes the result of multiple inputs by operation /
func computeDiv(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = divf(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeMod computes the result of multiple inputs by operation %
func computeMod(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = modf(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//addf computes the result of two inputs by operation +
func addf(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		if t1.Kind == List {
			x = t1
		} else {
			return None, err
		}
	}
	y, err := p.Exec(t2)
	if err != nil {
		if t2.Kind == List {
			y = t2
		} else {
			return None, err
		}
	}
	switch x.Kind {
	case Int:
		switch y.Kind {
		case Int:
			return Token{Kind: Int, Text: x.Text.(int64) + y.Text.(int64)}, nil
		case Float:
			return Token{Kind: Float, Text: float64(x.Text.(int64)) + y.Text.(float64)}, nil
		}
	case Float:
		switch y.Kind {
		case Int:
			return Token{Kind: Float, Text: x.Text.(float64) + float64(y.Text.(int64))}, nil
		case Float:
			return Token{Kind: Float, Text: x.Text.(float64) + y.Text.(float64)}, nil
		}
	case String:
		switch y.Kind {
		case String:
			return Token{Kind: String, Text: x.Text.(string) + y.Text.(string)}, nil
		}
	case List:
		switch y.Kind {
		case List:
			a, b := x.Text.([]Token), y.Text.([]Token)
			c := make([]Token, len(a)+len(b))
			copy(c, a)
			copy(c[len(a):], b)
			return Token{Kind: List, Text: c}, nil
		}

	}

	return None, ErrFitType
}

//subf computes the result of two inputs by operation -
func subf(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int:
		switch y.Kind {
		case Int:
			return Token{Kind: Int, Text: x.Text.(int64) - y.Text.(int64)}, nil
		case Float:
			return Token{Kind: Float, Text: float64(x.Text.(int64)) - y.Text.(float64)}, nil
		}
	case Float:
		switch y.Kind {
		case Int:
			return Token{Kind: Float, Text: x.Text.(float64) - float64(y.Text.(int64))}, nil
		case Float:
			return Token{Kind: Float, Text: x.Text.(float64) - y.Text.(float64)}, nil
		}
	}
	return None, ErrFitType
}

//mulf computes the result of two inputs by operation *
func mulf(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int:
		switch y.Kind {
		case Int:
			return Token{Kind: Int, Text: x.Text.(int64) * y.Text.(int64)}, nil
		case Float:
			return Token{Kind: Float, Text: float64(x.Text.(int64)) * y.Text.(float64)}, nil
		}
	case Float:
		switch y.Kind {
		case Int:
			return Token{Kind: Float, Text: x.Text.(float64) * float64(y.Text.(int64))}, nil
		case Float:
			return Token{Kind: Float, Text: x.Text.(float64) * y.Text.(float64)}, nil
		}
	}
	return None, ErrFitType
}

//divf computes the result of two inputs by operation /
func divf(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	switch x.Kind {
	case Int:
		switch y.Kind {
		case Int:
			if y.Text.(int64) == 0 {
				return None, ErrDivZero
			}
			return Token{Kind: Int, Text: x.Text.(int64) / y.Text.(int64)}, nil
		case Float:
			if y.Text.(float64) == 0 {
				return None, ErrDivZero
			}
			return Token{Kind: Float, Text: float64(x.Text.(int64)) / y.Text.(float64)}, nil
		}
	case Float:
		switch y.Kind {
		case Int:
			if y.Text.(int64) == 0 {
				return None, ErrDivZero
			}
			return Token{Kind: Float, Text: x.Text.(float64) / float64(y.Text.(int64))}, nil
		case Float:
			if y.Text.(float64) == 0 {
				return None, ErrDivZero
			}
			return Token{Kind: Float, Text: x.Text.(float64) / y.Text.(float64)}, nil
		}
	}
	return None, ErrFitType
}

//modf computes the result of two inputs by operation %
func modf(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	if x.Kind == Int && y.Kind == Int {
		if y.Text.(int64) == 0 {
			return None, ErrModZero
		}
		return Token{Kind: Int, Text: x.Text.(int64) % y.Text.(int64)}, nil
	}
	return None, ErrFitType
}
