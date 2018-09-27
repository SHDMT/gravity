// This file is part of the Dazzle Gravity library.
//
// The Dazzle Gravity library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Dazzle Gravity library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Dazzle Gravity library. If not, see <e <http://www.gnu.org/licenses/>./>.
package lispvm

import (
	"testing"

	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/smartcontract/vm"
)

var code1 = `
	(getGlobalParam "hello")
`
var code2 = `
	(/ 1 0)
`

func TestLispVMPanic(t *testing.T) {
	contract1 := structure.Contract{
		Version:    1,
		Name:       "error program1",
		ScriptCode: vm.LispScriptCode,
		Code:       []byte(code1),
	}
	contract2 := structure.Contract{
		Version:    1,
		Name:       "error program2",
		ScriptCode: vm.LispScriptCode,
		Code:       []byte(code2),
	}

	vm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})

	ret := vm.Exec(&contract1)

	if ret {
		t.Errorf("Code should detected no pointer err.")
	}

	ret = vm.Exec(&contract2)

	if ret {
		t.Errorf("Code should detected divide zero err.")
	}
}
