package lisp

//section maintains the state of feed function
type section struct {
	quote bool
	count int
	total string
}

//feed removes the comments in program following a '#'
func (b *section) feed(s []byte) error {
	single := false
	for i, l := 0, len(s); i < l; i++ {
		if b.quote {
			switch s[i] {
			case '"':
				b.quote = false
			case '\\':
				i++
			}
		} else if single {
			switch s[i] {
			case '\'':
				single = false
			case '\\':
				i++
			}
		} else {
			switch s[i] {
			case '(':
				b.count++
			case ')':
				b.count--
			case '\'':
				if i+1 < len(s) {
					if s[i+1] == '(' && (i+2 >= len(s) || s[i+2] != '\'') {
						b.count++
						i++
					} else {
						single = true
					}
				}
			case '"':
				b.quote = true
			case '#':
				s, l = s[:i], i
			}
		}
	}
	if single || b.count < 0 {
		return ErrUnquote
	}
	b.total += string(s)
	return nil
}

//over tells whether there is
func (b *section) over() bool {
	return b.count == 0 && !b.quote
}
