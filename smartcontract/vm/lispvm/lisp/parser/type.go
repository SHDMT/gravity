package parser

import "fmt"

//Pattern defines all the lexical rules
//rule describes the rule by an actual function
//str defines the key word directly
type Pattern struct {
	rule []func([]byte) (interface{}, int)
	str  []string
}

//Scanner defines a lexical analyzer
//ptn is the lexical rule
//tkn is the raw code to be analyze
//skp decides whether to skip space
type Scanner struct {
	ptn *Pattern
	tkn []byte
	skp bool
}

//Add adds a lexical rule to a Scanner
func (p *Pattern) Add(f func([]byte) (interface{}, int)) {
	p.rule = append(p.rule, f)
}

//AddString adds a key word to a Scanner
func (p *Pattern) AddString(s string) {
	p.str = append(p.str, s)
}

//NewScanner generates a specific Scanner to input code by Pattern(lexical rules)
func (p *Pattern) NewScanner(s []byte, t bool) *Scanner {
	return &Scanner{ptn: p, tkn: s, skp: t}
}

//Skip ignores all the coming space
func (s *Scanner) Skip() {
	i, l := 0, len(s.tkn)
	if l == 0 {
		return
	}
	for i < l && IsSpace(s.tkn[i]) {
		i++
	}
	s.tkn = s.tkn[i:]
}

//Scan try to analyze ONE lexical unit on raw data and prune the analyzed raw data
func (s *Scanner) Scan() (interface{}, int, error) {
	if s.skp {
		s.Skip()
	}
	if len(s.tkn) == 0 {
		return nil, 0, fmt.Errorf("empty string")
	}
	for i, t := range s.ptn.str {
		l := len(t)
		if len(s.tkn) >= l && t == string(s.tkn[:l]) {
			s.tkn = s.tkn[l:]
			return t, -(i + 1), nil
		}
	}
	for i, f := range s.ptn.rule {
		a, l := f(s.tkn)
		if l > 0 {
			s.tkn = s.tkn[l:]
			return a, +(i + 1), nil
		}
	}
	return nil, 0, fmt.Errorf("unrecognised")
}

//Over tells whether there is data which is not analyzed yet
func (s *Scanner) Over() bool {
	return len(s.tkn) == 0
}
