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
	"encoding/hex"
	"testing"

	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/smartcontract/vm"
)

func Test_cap(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (cap))
	`))

	if err != nil {
		t.Error("cap error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (cap index))
	`))

	if err1 == nil {
		t.Error("cap error:", err1)
	}
}

func Test_isDenominations(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (isDenominations))
	`))

	if err != nil {
		t.Error("isDenominations error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (isDenominations index))
	`))

	if err1 == nil {
		t.Error("isDenominations error:", err1)
	}
	//Simulate FixedDenominations to false
	lvm.context.AssetMsg.FixedDenominations = false

	_, err2 := lvm.vm.Eval([]byte(`
		(println (isDenominations))
	`))

	if err2 != nil {
		t.Error("isDenominations error:", err2)
	}

}

func Test_getDenominationCount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (getDenominationCount))
	`))

	if err != nil {
		t.Error("getDenominationCount error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
         (define index 0)
		(println (getDenominationCount index))
	`))

	if err1 == nil {
		t.Error("getDenominationCount error:", err1)
	}
}
func Test_getDenomination(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 1)
		(println (getDenomination index))
	`))

	if err != nil {
		t.Error("getDenomination error:", err)
	}
	//The number of parameters transmitted by the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getDenomination ))
	`))

	if err1 == nil {
		t.Error("getDenomination error:", err1)
	}
	//The first parameter execution error
	_, err3 := lvm.vm.Eval([]byte(`
		(println (getDenomination (index)))
	`))

	if err3 == nil {
		t.Error("getDenomination error:", err3)
	}
	//The simulated parameter is not int
	_, err2 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (getDenomination index))
	`))

	if err2 == nil {
		t.Error("getDenomination error:", err2)
	}

}
func Test_getAssetContractCount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (getAssetContractCount))
	`))

	if err != nil {
		t.Error("getAssetContractCount error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getAssetContractCount index))
	`))

	if err1 == nil {
		t.Error("getAssetContractCount error:", err1)
	}
}
func Test_getAssetContract(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (getAssetContract index))
	`))

	if err != nil {
		t.Error("getAssetContract error:", err)
	}
	//The number of parameters passed in the simulation is incorrect.
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getAssetContract))
	`))

	if err1 == nil {
		t.Error("getAssetContract error:", err1)
	}
	//Error in simulated parameter execution
	_, err3 := lvm.vm.Eval([]byte(`
		(println (getAssetContract (index)))
	`))

	if err3 == nil {
		t.Error("getAssetContract error:", err3)
	}
	//The parameters passed in the simulation are not int
	_, err2 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (getAssetContract index))
	`))

	if err2 == nil {
		t.Error("getAssetContract error:", err2)
	}
}

func Test_getAllocationsCount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (getAllocationsCount))
	`))

	if err != nil {
		t.Error("getAllocationsCount error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getAllocationsCount index))
	`))

	if err1 == nil {
		t.Error("getAllocationsCount error:", err1)
	}

}

func Test_getAllocationsAddr(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (getAllocationsAddr index))
	`))

	if err != nil {
		t.Error("getAllocationsAddr error:", err)
	}
	//The number of parameters passed in the simulation is incorrect.
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getAllocationsAddr))
	`))

	if err1 == nil {
		t.Error("getAllocationsAddr error:", err1)
	}
	//Error in simulated parameter execution
	_, err3 := lvm.vm.Eval([]byte(`
		(println (getAllocationsAddr (index)))
	`))

	if err3 == nil {
		t.Error("getAllocationsAddr error:", err3)
	}
	//The parameters passed in the simulation are not int
	_, err2 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (getAllocationsAddr index))
	`))

	if err2 == nil {
		t.Error("getAllocationsAddr error:", err2)
	}

}
func Test_getAllocationsAmount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index  "test")
		(println (getAllocationsAmount index))
	`))

	if err != nil {
		t.Error("getAllocationsAmount error:", err)
	}
	//The number of parameters passed in the simulation is incorrect.
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getAllocationsAmount))
	`))

	if err1 == nil {
		t.Error("getAllocationsAmount error:", err1)
	}
	//Error in simulated parameter execution
	_, err4 := lvm.vm.Eval([]byte(`
		(println (getAllocationsAmount (index)))
	`))

	if err4 == nil {
		t.Error("getAllocationsAmount error:", err4)
	}
	//The parameters passed in the simulation are not string
	_, err2 := lvm.vm.Eval([]byte(`
		(define index  0)
		(println (getAllocationsAmount index))
	`))

	if err2 == nil {
		t.Error("getAllocationsAmount error:", err2)
	}
	//Simulated incoming parameters are empty
	_, err3 := lvm.vm.Eval([]byte(`
		(define index  "")
		(println (getAllocationsAmount index))
	`))

	if err3 == nil {
		t.Error("getAllocationsAmount error:", err3)
	}

}
func Test_getAssetExtends(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (getAssetExtends))
	`))

	if err != nil {
		t.Error("getAssetExtends error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getAssetExtends index))
	`))

	if err1 == nil {
		t.Error("getAssetExtends error:", err1)
	}

}
func Test_getPublisherAddr(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (getPublisherAddr))
	`))

	if err != nil {
		t.Error("getPublisherAddr error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
         (define index 0)
		(println (getPublisherAddr index))
	`))

	if err1 == nil {
		t.Error("getPublisherAddr error:", err1)
	}
}
func Test_getPublisherUnitMCI(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (getPublisherUnitMCI))
	`))

	if err != nil {
		t.Error("getPublisherUnitMCI error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
         (define index 0)
		(println (getPublisherUnitMCI index))
	`))

	if err1 == nil {
		t.Error("getPublisherUnitMCI error:", err1)
	}
}
func Test_getContractParamCount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (getContractParamCount index))
	`))

	if err != nil {
		t.Error("getContractParamCount error:", err)
	}
	//The number of parameters passed in the simulation is incorrect.
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getContractParamCount))
	`))

	if err1 == nil {
		t.Error("getContractParamCount error:", err1)
	}
	//Error in simulated parameter execution
	_, err3 := lvm.vm.Eval([]byte(`
		(println (getContractParamCount (index)))
	`))

	if err3 == nil {
		t.Error("getContractParamCount error:", err3)
	}
	//The parameters passed in the simulation are not int
	_, err2 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (getContractParamCount index))
	`))

	if err2 == nil {
		t.Error("getContractParamCount error:", err2)
	}
}
func Test_getCurContractDefParamCount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (getCurContractDefParamCount))
	`))

	if err != nil {
		t.Error("getCurContractDefParamCount error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getCurContractDefParamCount index))
	`))

	if err1 == nil {
		t.Error("getCurContractDefParamCount error:", err1)
	}

}
func Test_getCurContractDefParamName(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (getCurContractDefParamName index))
	`))

	if err != nil {
		t.Error("getCurContractDefParamName error:", err)
	}
	//The number of parameters passed in the simulation is incorrect.
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getCurContractDefParamName))
	`))

	if err1 == nil {
		t.Error("getCurContractDefParamName error:", err1)
	}
	//Error in simulated parameter execution
	_, err3 := lvm.vm.Eval([]byte(`
		(println (getCurContractDefParamName (index)))
	`))

	if err3 == nil {
		t.Error("getCurContractDefParamName error:", err3)
	}
	//The parameters passed in the simulation are not int
	_, err2 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (getCurContractDefParamName index))
	`))

	if err2 == nil {
		t.Error("getCurContractDefParamName error:", err2)
	}
}
func Test_getCurContractDefParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define key "north")
		(println (getCurContractDefParam key))
	`))

	if err != nil {
		t.Error("getCurContractDefParam error:", err)
	}
	//The number of parameters passed in the simulation is incorrect.
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getCurContractDefParam))
	`))

	if err1 == nil {
		t.Error("getCurContractDefParam error:", err1)
	}
	//Error in simulated parameter execution
	_, err4 := lvm.vm.Eval([]byte(`
		(println (getCurContractDefParam (key)))
	`))

	if err4 == nil {
		t.Error("getCurContractDefParam error:", err4)
	}
	//The parameters passed in the simulation are not string
	_, err2 := lvm.vm.Eval([]byte(`
		(define key 0)
		(println (getCurContractDefParam key))
	`))

	if err2 == nil {
		t.Error("getCurContractDefParam error:", err2)
	}
	//Simulated incoming parameters are empty
	_, err3 := lvm.vm.Eval([]byte(`
		(define key "")
		(println (getCurContractDefParam key))
	`))

	if err3 == nil {
		t.Error("getCurContractDefParam error:", err3)
	}
}
func Test_getContractParamName(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
        (define index 0)
		(define keyindex 0)
		(println (getContractParamName index keyindex))
	`))

	if err != nil {
		t.Error("getContractParamName error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getContractParamName index))
	`))

	if err1 == nil {
		t.Error("getContractParamName error:", err1)
	}
	//Simulate the first parameter execution error
	_, err4 := lvm.vm.Eval([]byte(`
		(define keyindex 0)
		(println (getContractParamName (index) keyindex))
	`))

	if err4 == nil {
		t.Error("getContractParamName error:", err4)
	}
	//Simulating the first parameter is not int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define keyindex 0)
		(println (getContractParamName index keyindex))
	`))

	if err2 == nil {
		t.Error("getContractParamName error:", err2)
	}
	//Simulate the second parameter execution error
	_, err5 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getContractParamName index (keyindex)))
	`))

	if err5 == nil {
		t.Error("getContractParamName error:", err5)
	}
	//Simulating the second parameter is not int
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define keyindex "error")
		(println (getContractParamName index keyindex))
	`))

	if err3 == nil {
		t.Error("getContractParamName error:", err3)
	}
}
func Test_getContractParamByIndex(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
        (define index 0)
		(define keyindex 0)
		(println (getContractParamByIndex index keyindex))
	`))

	if err != nil {
		t.Error("getContractParamByIndex error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getContractParamByIndex index))
	`))

	if err1 == nil {
		t.Error("getContractParamByIndex error:", err1)
	}
	//Simulate the first parameter execution error
	_, err4 := lvm.vm.Eval([]byte(`
		(define keyindex 0)
		(println (getContractParamByIndex (index) keyindex))
	`))

	if err4 == nil {
		t.Error("getContractParamByIndex error:", err4)
	}
	//Simulating the first parameter is not int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define keyindex 0)
		(println (getContractParamByIndex index keyindex))
	`))

	if err2 == nil {
		t.Error("getContractParamByIndex error:", err2)
	}
	//Simulate the second parameter execution error
	_, err5 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getContractParamByIndex index (keyindex)))
	`))

	if err5 == nil {
		t.Error("getContractParamByIndex error:", err5)
	}
	//Simulating the second parameter is not int
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define keyindex "error")
		(println (getContractParamByIndex index keyindex))
	`))

	if err3 == nil {
		t.Error("getContractParamByIndex error:", err3)
	}

}
func Test_getContractParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
        (define index 0)
		(define key "north")
		(println (getContractParam index key))
	`))

	if err != nil {
		t.Error("getContractParam error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getContractParam index))
	`))

	if err1 == nil {
		t.Error("getContractParam error:", err1)
	}
	//Simulate the first parameter execution error
	_, err5 := lvm.vm.Eval([]byte(`
		(define key "north")
		(println (getContractParam (index) key))
	`))

	if err5 == nil {
		t.Error("getContractParam error:", err5)
	}
	//Simulating the first parameter is not int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define key "north")
		(println (getContractParam index key))
	`))

	if err2 == nil {
		t.Error("getContractParam error:", err2)
	}
	//Simulate the second parameter execution error
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getContractParam index (key)))
	`))

	if err6 == nil {
		t.Error("getContractParam error:", err6)
	}
	//Simulating the second parameter is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define key 0)
		(println (getContractParam index key))
	`))

	if err3 == nil {
		t.Error("getContractParam error:", err3)
	}
	//Simulate the second parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define key "")
		(println (getContractParam index key))
	`))

	if err4 == nil {
		t.Error("getContractParam error:", err4)
	}
}

func Test_isExistAtGlobalParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index "west")
		(println (isExistAtGlobalParam index))
	`))

	if err != nil {
		t.Error("isExistAtGlobalParam error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (isExistAtGlobalParam))
	`))

	if err1 == nil {
		t.Error("isExistAtGlobalParam error:", err1)
	}
	//Error in simulated parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(println (isExistAtGlobalParam (index)))
	`))

	if err5 == nil {
		t.Error("isExistAtGlobalParam error:", err5)
	}
	//The mock parameter type is not a string type
	_, err2 := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (isExistAtGlobalParam index))
	`))

	if err2 == nil {
		t.Error("isExistAtGlobalParam error:", err2)
	}
	//Simulated incoming parameters are empty
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "")
		(println (isExistAtGlobalParam index))
	`))

	if err3 == nil {
		t.Error("isExistAtGlobalParam error:", err3)
	}
	//Simulated incoming parameters are different from the required values
	_, err4 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (isExistAtGlobalParam index))
	`))

	if err4 != nil {
		t.Error("isExistAtGlobalParam error:", err4)
	}
}
func Test_isExistAtOutputParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index "west")
		(println (isExistAtOutputParam index))
	`))

	if err != nil {
		t.Error("isExistAtOutputParam error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (isExistAtOutputParam))
	`))

	if err1 == nil {
		t.Error("isExistAtOutputParam error:", err1)
	}
	//Error in simulated parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(println (isExistAtOutputParam (index)))
	`))

	if err5 == nil {
		t.Error("isExistAtOutputParam error:", err5)
	}
	//The mock parameter type is not a string type
	_, err2 := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (isExistAtOutputParam index))
	`))

	if err2 == nil {
		t.Error("isExistAtOutputParam error:", err2)
	}
	//Simulated incoming parameters are empty
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "")
		(println (isExistAtOutputParam index))
	`))

	if err3 == nil {
		t.Error("isExistAtOutputParam error:", err3)
	}
	//Simulated incoming parameters are different from the required values
	_, err4 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (isExistAtOutputParam index))
	`))

	if err4 != nil {
		t.Error("isExistAtOutputParam error:", err4)
	}

}
func Test_isExistAtInputParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index "west")
		(println (isExistAtInputParam index))
	`))

	if err != nil {
		t.Error("isExistAtInputParam error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (isExistAtInputParam))
	`))

	if err1 == nil {
		t.Error("isExistAtInputParam error:", err1)
	}
	//Error in simulated parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(println (isExistAtInputParam (index)))
	`))

	if err5 == nil {
		t.Error("isExistAtInputParam error:", err5)
	}
	//The mock parameter type is not a string type
	_, err2 := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (isExistAtInputParam index))
	`))

	if err2 == nil {
		t.Error("isExistAtInputParam error:", err2)
	}
	//Simulated incoming parameters are empty
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "")
		(println (isExistAtInputParam index))
	`))

	if err3 == nil {
		t.Error("isExistAtInputParam error:", err3)
	}
	//Simulated incoming parameters are different from the required values
	_, err4 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (isExistAtInputParam index))
	`))

	if err4 != nil {
		t.Error("isExistAtInputParam error:", err4)
	}
}
func Test_CalcOutputAmount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (calcOutputAmount))
	`))

	if err != nil {
		t.Error("calcOutputAmount error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (calcOutputAmount index))
	`))

	if err1 == nil {
		t.Error("calcOutputAmount error:", err1)
	}

}
func Test_CalcInputAmount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (calcInputAmount))
	`))

	if err != nil {
		t.Error("calcInputAmount error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (calcInputAmount index))
	`))

	if err1 == nil {
		t.Error("calcInputAmount error:", err1)
	}
}
func Test_CalcBalance(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (calcBalance))
	`))

	if err != nil {
		t.Error("calcBalance error:", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (calcBalance index))
	`))

	if err1 == nil {
		t.Error("calcBalance error:", err1)
	}
}

func Test_getCurContractDefParamList(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setIssueLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define key "north")
		(define x (getCurContractDefParamList key))
		(println x)
	`))

	if err != nil {
		t.Error("getCurContractDefParamList error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(define x (getCurContractDefParamList))
		(println x)
	`))

	if err1 == nil {
		t.Error("getCurContractDefParamList error:", err1)
	}
	//The parameters passed in the simulation are not string
	_, err2 := lvm.vm.Eval([]byte(`
		(define key 0)
		(define x (getCurContractDefParamList key))
		(println x)
	`))

	if err2 == nil {
		t.Error("getCurContractDefParamList error:", err2)
	}
	//Simulated incoming parameters are empty
	_, err3 := lvm.vm.Eval([]byte(`
		(define key "")
		(define x (getCurContractDefParamList key))
		(println x)
	`))

	if err3 == nil {
		t.Error("getCurContractDefParamList error:", err3)
	}
	//The first parameter execution error
	_, err4 := lvm.vm.Eval([]byte(`
		(define x (getCurContractDefParamList (key)))
		(println x)
	`))

	if err4 == nil {
		t.Error("getCurContractDefParamList error:", err4)
	}
	//The data of the obtained value is too long
	lvm.context.ContractDef.ParamsValue[0] = []byte("north")
	_, err5 := lvm.vm.Eval([]byte(`
		(define key "north")
		(define x (getCurContractDefParamList key))
		(println x)
	`))

	if err5 == nil {
		t.Error("getCurContractDefParamList error:", err5)
	}
}

func setIssueLvm(lvm *LispVM) {

	alloAddr := make([]hash.HashType, 4)
	alloAmount := make([]int64, 4)
	alloAddr[0] = hash.HashType("test")
	alloAmount[0] = 128
	res := make([]uint32, 4)
	paraKey := make([]string, 4)
	paraValue := make([][]byte, 4)
	paraValue1 := make([][]byte, 4)
	conDef := make([]*structure.ContractDef, 4)
	paraKey[0] = "north"
	paraKey[1] = "south"
	paraKey[2] = "west"
	paraKey[3] = "east"
	paraValue1[0] = append(paraValue1[0], 10)
	paraValue1[0] = append(paraValue1[0], []byte("abcdefghij")...)
	paraValue1[1] = []byte("south")
	paraValue1[2] = []byte("west")
	paraValue1[3] = []byte("east")
	conDef[0] = &structure.ContractDef{
		Address:     hash.Sum256(defaultBytes),
		ParamsKey:   paraKey,
		ParamsValue: paraValue1,
	}
	invokeMsg := make([]structure.Message, 4)
	for i := 0; i < len(alloAddr); i++ {

		res[i] = 24
		paraValue[i] = hash.Sum256(defaultBytes)

		contractinput := make([]*structure.ContractInput, 1)
		contractinput[0] = &structure.ContractInput{
			InputParamsKey:   paraKey,
			InputParamsValue: paraValue,
			SourceOutput:     uint32(1),
		}
		contractouput := make([]*structure.ContractOutput, 2)
		contractouput[0] = &structure.ContractOutput{
			Amount:            uint64(1),
			OutputParamsKey:   paraKey,
			OutputParamsValue: paraValue,
		}
		contractouput[1] = &structure.ContractOutput{
			Amount:            uint64(2),
			OutputParamsKey:   paraKey,
			OutputParamsValue: paraValue,
		}
		invokeMsg[i] = &structure.InvokeMessage{
			ContractAddr:     hash.Sum256(defaultBytes),
			GlobalParamKey:   paraKey,
			GlobalParamValue: paraValue,
			Outputs:          contractouput,
			Inputs:           contractinput,
		}
	}

	issueMsg := generateIssueMsg(&structure.MessageHeader{}, "IssueMessage", 21000000, true, res, conDef, alloAddr, alloAmount, hash.Sum256(defaultBytes), hash.Sum256(defaultBytes))

	lvm.context = vm.Context{
		TxUnit: structure.Unit{
			Messages: invokeMsg,
		},
		TxMsgIndex:   1,
		FetchPrevOut: GetContractOutput,

		AssetMsg:    *issueMsg,
		ContractDef: conDef[0],

		MCI: 0,
	}
}

func generateIssueMsg(header *structure.MessageHeader, name string, cap int64,
	fixed bool, demo []uint32, contracts []*structure.ContractDef, alloAddr []hash.HashType,
	alloAmount []int64, address hash.HashType, note []byte) *structure.IssueMessage {
	issueMsg := &structure.IssueMessage{
		Header:             header,
		Name:               name,
		Cap:                cap,
		FixedDenominations: fixed,
		Denominations:      demo,
		Contracts:          contracts,

		AllocationAddr:   alloAddr,
		AllocationAmount: alloAmount,

		PublisherAddress: address,
		Note:             note,
	}

	return issueMsg
}

func GetContractOutput(input *structure.ContractInput) *structure.ContractOutput {
	contractOutput := structure.NewContractOutput()
	contractOutput.Amount = 3
	contractOutput.Extends = []byte("Hello World!")

	//for test
	pubKByte, err := hex.DecodeString("0421af2c7f64c10a34fbf4891ac5862a71f0c5805e0e4b50ce03244943a7859b7dd26e5a9ffd9ac0bccb5aadc81552b5961cc7f361a244df703d248aa5c13fb013")
	if err != nil {
		return nil
	}
	paraValue2 := make([][]byte, 2)
	paraValue2[0] = append(paraValue2[0], 10)
	paraValue2[0] = append(paraValue2[0], []byte("gravitygra")...)
	contractOutput.AddParam("addr", hash.Sum256([]byte(pubKByte)))
	contractOutput.AddParam("gravity", paraValue2[0])

	contractOutput.AddRestrict(hash.Sum256([]byte(pubKByte)))
	contractOutput.AddRestrict(hash.Sum256([]byte(pubKByte)))

	return contractOutput
}
