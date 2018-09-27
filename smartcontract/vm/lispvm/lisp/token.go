package lisp

import "fmt"

//Token is an important structure in lisp interpretation
//Token can describe a lexical unit
//Token can also describe a node in syntax tree
type Token struct {
	Kind
	Text interface{}
}

//Bool tells whether a Token is considered a boolean true or boolean false
func (t *Token) Bool() bool {
	switch t.Kind {
	case Null:
		return false
	case Int:
		return t.Text.(int64) != 0
	case Float:
		return t.Text.(float64) != 0
	case String:
		return t.Text.(string) != ""
	case List:
		return len(t.Text.([]Token)) != 0
	}
	return true
}

//Eq tells whether two tokens are same
func (t *Token) Eq(p *Token) bool {
	var (
		a, b []Token
		c, d []Name
	)
	if t.Kind != p.Kind {
		return false
	}
	switch t.Kind {
	case Null:
		return true
	case Int:
		return t.Text.(int64) == p.Text.(int64)
	case Float:
		return t.Text.(float64) == p.Text.(float64)
	case String:
		return t.Text.(string) == p.Text.(string)
	case Back:
		return false
	case Label:
		return t.Text.(Name) == p.Text.(Name)
	case Front:
		if t.Text.(*Lfac).Make != p.Text.(*Lfac).Make {
			return false
		}
		a, b = t.Text.(*Lfac).Text, p.Text.(*Lfac).Text
		c, d = t.Text.(*Lfac).Para, p.Text.(*Lfac).Para
	case Macro:
		a, b = t.Text.(*Lfac).Text, p.Text.(*Lfac).Text
		c, d = t.Text.(*Lfac).Para, p.Text.(*Lfac).Para
	case Fold, List:
		a, b = t.Text.([]Token), p.Text.([]Token)
		c, d = nil, nil
	}
	m, n := len(a), len(b)
	if m != n {
		return false
	}
	for i := 0; i < m; i++ {
		if !a[i].Eq(&b[i]) {
			return false
		}
	}
	if c != nil {
		m, n := len(c), len(d)
		if m != n {
			return false
		}
		for i := 0; i < m; i++ {
			if c[i] != d[i] {
				return false
			}
		}
	}
	return true
}

//Cmp compares two tokens
//return value 1 indicates that caller is considered greater than parameter
//return value -1 indicates that caller is considered less than parameter
//return value 0 indicate that caller is equal to parameter
func (t *Token) Cmp(p *Token) (int, error) {
	var a, b bool
	switch t.Kind {
	case Int:
		switch p.Kind {
		case Int:
			a = t.Text.(int64) > p.Text.(int64)
			b = t.Text.(int64) < p.Text.(int64)
			//c = t.text.(int64) == p.Text.(int64)
		case Float:
			a = float64(t.Text.(int64)) > p.Text.(float64)
			b = float64(t.Text.(int64)) < p.Text.(float64)
			//c = float64(t.Text.(int64)) == p.Text.(float64)
		default:
			return 0, ErrFitType
		}
	case Float:
		switch p.Kind {
		case Int:
			a = t.Text.(float64) > float64(p.Text.(int64))
			b = t.Text.(float64) < float64(p.Text.(int64))
			//c = t.Text.(float64) == float64(p.Text.(int64))
		case Float:
			a = t.Text.(float64) > p.Text.(float64)
			b = t.Text.(float64) < p.Text.(float64)
			//c = t.Text.(float64) == p.Text.(float64)
		default:
			return 0, ErrFitType
		}
	case String:
		switch p.Kind {
		// case Int, Float:
		// 	return 0, 1
		case String:
			a = t.Text.(string) > p.Text.(string)
			b = t.Text.(string) < p.Text.(string)
			//c = t.Text.(string) == p.Text.(string)
		default:
			return 0, ErrFitType
		}
		// case List:
		// 	switch p.Kind {
		// 	case Int, Float, String:
		// 		return 0, 1
		// 	case List:
		// 		x, y := t.Text.([]Token), p.Text.([]Token)
		// 		m, n := len(x), len(y)
		// 		for i := 0; i < m && i < n; i++ {
		// 			j := x[i].Cmp(&y[i])
		// 			if j != 0 {
		// 				return 0, j
		// 			}
		// 		}
		// 		a = m > n
		// 		b = m < n
		// 	default:
		// 		return 0, 0
		// 	}
		// default:
		// 	return 0, 0
	}
	if a {
		return +1, nil
	}
	if b {
		return -1, nil
	}
	return 0, nil
}

//String implements the Stringer interface and give the string expression of token
func (t Token) String() string {
	switch t.Kind {
	case Null:
		return ""
	default:
		return fmt.Sprint(t.Text)
	}
}
