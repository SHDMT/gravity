package lisp

func init() {
	//implementation of the system function "progn" used to execute every parameters sequentially
	//the return value of last parameter will the return value of "progn"
	Add("progn", func(t []Token, p *Lisp) (ans Token, err error) {
		ans = False
		for _, token := range t {
			ans, err = p.Exec(token)
			if err != nil {
				break
			}
		}
		return ans, err
	})

	//implementation of the system function "list" used to execute every parameters sequentially
	//the return value is the list of the return value of every parameter
	Add("list", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) == 0 {
			return False, nil
		}
		elements := make([]Token, 0, len(t))

		for _, token := range t {
			var result Token
			result, err = p.Exec(token)
			if err != nil {
				break
			}
			elements = append(elements, result)
		}
		ans = Token{Kind: List, Text: elements}
		return ans, err
	})

	//implementation of the system function "length" used to return the length of parameter
	//valid input type includes list and string
	Add("length", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		r, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if r.Kind == List {
			ans = Token{Kind: Int, Text: int64(len(r.Text.([]Token)))}
			return ans, err
		} else if r.Kind == String {
			ans = Token{Kind: Int, Text: int64(len(r.Text.(string)))}
			return ans, err
		}

		return None, ErrFitType
	})

	//implementation of the system function "setq" used to set data to a variable
	Add("setq", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		symbol := t[0]
		if symbol.Kind != Label {
			return None, ErrFitType
		}
		ans, err = p.Exec(t[1])
		if err == nil {
			scope := p
			for {
				if scope.parent == Global || scope.env[symbol.Text.(Name)] != None {
					scope.env[symbol.Text.(Name)] = ans
					break
				}
				scope = scope.parent
			}
		}
		return ans, err
	})

	//implementation of the system function "defun" used for function definition
	//difference from "define" is that "defun" follows the syntax of "defun" in common-lisp
	Add("defun", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) < 2 {
			return None, ErrParaNum
		}
		symbol := t[0]
		if symbol.Kind != Label {
			return None, ErrFitType
		}
		funcName := symbol.Text.(Name)

		paraList := t[1]
		if paraList.Kind != List {
			return None, ErrFitType
		}
		paras := paraList.Text.([]Token)
		paraNames := make([]Name, len(paras))
		for i, para := range paras {
			if para.Kind != Label {
				return None, ErrFitType
			}
			paraNames[i] = para.Text.(Name)
		}
		funcBody := make([]Token, len(t)-2)
		//funcBody[0] = Token{Kind: Label, Text: Name("progn")}
		copy(funcBody[:], t[2:])
		f := &Lfac{
			FuncName: funcName,
			Para:     paraNames,
			Text:     funcBody,
			Make:     p,
		}
		ans = Token{Kind: Front, Text: f}

		scope := p
		for {
			if scope.parent == Global || scope.env[funcName] != None {
				scope.env[funcName] = ans
				break
			}
			scope = scope.parent
		}
		return ans, nil
	})

	//implementation of the system function "return-from" used to return the return value directly
	// in function or macro or block
	Add("return-from", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 && len(t) != 2 {
			return None, ErrParaNum
		}
		found := false
		dest := t[0]
		if dest.Kind != Label {
			return None, ErrFitType
		}
		if len(t) == 2 {
			ans, err = p.Exec(t[1])
		} else {
			ans, err = False, nil
		}

		scope := p
		path := make([]*Lisp, 0)
		for {
			if scope == Global {
				break
			}
			//scope.returnValue = ans
			path = append(path, scope)
			if scope.scopeName == dest.Text.(Name) {
				found = true
				break
			}
			scope = scope.parent
		}
		if !found {
			return None, ErrNotFind
		}
		for _, l := range path {
			l.returnValue = ans
		}
		return ans, err
	})

	//implementation of the system function "return" used to return the parameter directly in a loop
	Add("return", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 && len(t) != 1 {
			return None, ErrParaNum
		}
		found := false
		if len(t) == 1 {
			ans, err = p.Exec(t[0])
		} else {
			ans, err = False, nil
		}

		scope := p
		path := make([]*Lisp, 0)
		for {
			if scope == Global {
				break
			}
			path = append(path, scope)
			if scope.scopeName == "0" {
				found = true
				break
			}
			scope = scope.parent
		}
		if !found {
			return None, ErrNotFind
		}
		for _, l := range path {
			l.returnValue = ans
		}
		return ans, err
	})

	//implementation of the system function "defmacro" used for macro definition
	Add("defmacro", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) < 2 {
			return None, ErrParaNum
		}
		symbol := t[0]
		if symbol.Kind != Label {
			return None, ErrFitType
		}
		funcName := symbol.Text.(Name)

		paraList := t[1]
		if paraList.Kind != List {
			return None, ErrFitType
		}
		paras := paraList.Text.([]Token)
		paraNames := make([]Name, len(paras))
		for i, para := range paras {
			if para.Kind != Label {
				return None, ErrFitType
			}
			paraNames[i] = para.Text.(Name)
		}
		funcBody := make([]Token, len(t)-2)
		copy(funcBody[:], t[2:])
		f := &Lfac{
			FuncName: funcName,
			Para:     paraNames,
			Text:     funcBody,
			Make:     p,
		}
		ans = Token{Kind: Macro, Text: f}
		scope := p
		for {
			if scope.parent == Global || scope.env[funcName] != None {
				scope.env[funcName] = ans
				break
			}
			scope = scope.parent
		}
		return ans, nil
	})
}
