package lisp

func init() {
	Add("logand", computeBitAnd) //按位与
	Add("logior", computeBitOr)  //按位或
	Add("logxor", computeBitXor) //按位异或
	Add("lognor", computeBitNor) //按位或非
	Add("logeqv", computeBitEqv) //按位同或
	Add("lognot", computeBitNot) //按位非
}

//computeBitAnd computes the result of multiple inputs by operation "and on bits"
func computeBitAnd(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = bitAnd(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeBitOr computes the result of multiple inputs by operation "or on bits"
func computeBitOr(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = bitOr(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeBitXor computes the result of multiple inputs by operation "exclusive or on bits"
func computeBitXor(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = bitXor(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeBitNor computes the result of multiple inputs by operation "nor on bits"
func computeBitNor(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = bitNor(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeBitEqv computes the result of multiple inputs by operation "exclusive nor on bits"
func computeBitEqv(t []Token, p *Lisp) (Token, error) {
	var fpara Token
	var err error
	for i, para := range t {
		if i == 0 {
			fpara = para
		} else {
			fpara, err = bitEqv(fpara, para, p)
			if err != nil {
				return None, err
			}
		}
	}
	return fpara, nil
}

//computeBitNot computes the result inputs by operation "not on bits"
func computeBitNot(t []Token, p *Lisp) (Token, error) {
	if len(t) != 1 {
		return None, ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return None, err
	}
	if x.Kind == Int {
		return Token{Kind: Int, Text: ^x.Text.(int64)}, nil
	}
	return None, ErrFitType
}

//bitAnd computes the result of two inputs by operation "and on bits"
func bitAnd(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	if x.Kind == Int && y.Kind == Int {
		return Token{Kind: Int, Text: x.Text.(int64) & y.Text.(int64)}, nil
	}
	return None, ErrFitType
}

//bitAnd computes the result of two inputs by operation "or on bits"
func bitOr(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	if x.Kind == Int && y.Kind == Int {
		return Token{Kind: Int, Text: x.Text.(int64) | y.Text.(int64)}, nil
	}
	return None, ErrFitType
}

//bitAnd computes the result of two inputs by operation "exclusive or on bits"
func bitXor(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	if x.Kind == Int && y.Kind == Int {
		return Token{Kind: Int, Text: x.Text.(int64) ^ y.Text.(int64)}, nil
	}
	return None, ErrFitType
}

//bitAnd computes the result of two inputs by operation "nor on bits"
func bitNor(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	if x.Kind == Int && y.Kind == Int {
		return Token{Kind: Int, Text: ^(x.Text.(int64) | y.Text.(int64))}, nil
	}
	return None, ErrFitType
}

//bitAnd computes the result of two inputs by operation "exclusive-nor on bits"
func bitEqv(t1 Token, t2 Token, p *Lisp) (Token, error) {
	x, err := p.Exec(t1)
	if err != nil {
		return None, err
	}
	y, err := p.Exec(t2)
	if err != nil {
		return None, err
	}
	if x.Kind == Int && y.Kind == Int {
		return Token{Kind: Int, Text: ^(x.Text.(int64) ^ y.Text.(int64))}, nil
	}
	return None, ErrFitType
}
