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

	"encoding/hex"
	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/smartcontract/vm"
)

func TestSigCount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	Definition := []byte("test")
	author := &structure.Author{
		Address:    Definition,
		Definition: Definition,
	}

	auth := make([]*structure.Author, 1)
	auth[0] = author

	unit := structure.Unit{
		Authors: auth,
	}
	c := vm.Context{
		TxUnit: unit,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (println (sigCount))
    `))
	if err != nil {
		t.Errorf("sigCount failed,err= %v\n", err)
	}
	//Simulation passed in a parameter
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (sigCount index))
	`))

	if err1 == nil {
		t.Errorf("sigCount failed,err= %v\n", err1)
	}
}
func TestGetPK(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	Definition := []byte("test")
	author := &structure.Author{
		Definition: Definition,
	}

	auth := make([]*structure.Author, 1)
	auth[0] = author
	unit := structure.Unit{
		Authors: auth,
	}
	c := vm.Context{
		TxUnit: unit,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getPK index))
    `))
	if err != nil {
		t.Errorf("getPK failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getPK))
	`))

	if err1 == nil {
		t.Errorf("getPK failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err3 := lvm.vm.Eval([]byte(`
      (println (getPK (index)))
    `))
	if err3 == nil {
		t.Errorf("getPK failed,err= %v\n", err3)
	}
	//The parameters passed in the simulation are not int
	_, err2 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (getPK index))
	`))

	if err2 == nil {
		t.Errorf("getPK failed,err= %v\n", err2)
	}

}
func TestGetPKByAddr(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	Definition := []byte("test")
	author := &structure.Author{
		Address:    Definition,
		Definition: Definition,
	}

	auth := make([]*structure.Author, 1)
	auth[0] = author
	unit := structure.Unit{
		Authors: auth,
	}
	c := vm.Context{
		TxUnit: unit,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (define index "test")
      (println (getPKByAddr index))
    `))
	if err != nil {
		t.Errorf("getPKByAddr failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (getPKByAddr))
    `))
	if err1 == nil {
		t.Errorf("getPKByAddr failed,err= %v\n", err1)
	}
	//Parameter execution error
	_, err2 := lvm.vm.Eval([]byte(`
      (println (getPKByAddr (index)))
    `))
	if err2 == nil {
		t.Errorf("getPKByAddr failed,err= %v\n", err2)
	}
	//The mock parameter is not a string
	_, err3 := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getPKByAddr index))
    `))
	if err3 == nil {
		t.Errorf("getPKByAddr failed,err= %v\n", err3)
	}
	//The simulation parameter is empty
	_, err5 := lvm.vm.Eval([]byte(`
      (define index "")
      (println (getPKByAddr index))
    `))
	if err5 == nil {
		t.Errorf("getPKByAddr failed,err= %v\n", err5)
	}
	//The arguments passed in by the simulation are not equal to the desired address
	_, err4 := lvm.vm.Eval([]byte(`
      (define index "error")
      (println (getPKByAddr index))
    `))
	if err4 == nil {
		t.Errorf("getPKByAddr failed,err= %v\n", err4)
	}
}
func TestGetSig(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	Definition := []byte("pk")
	prvksig := []byte("preksig")
	author := &structure.Author{
		Address:       Definition,
		Authentifiers: prvksig,
		Definition:    Definition,
	}

	auth := make([]*structure.Author, 1)
	auth[0] = author
	unit := structure.Unit{
		Authors: auth,
	}
	c := vm.Context{
		TxUnit: unit,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (define index "pk")
      (println (getSig index))
    `))
	if err != nil {
		t.Errorf("getSig failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (getSig))
    `))
	if err1 == nil {
		t.Errorf("getSig failed,err= %v\n", err1)
	}
	//Parameter execution error
	_, err2 := lvm.vm.Eval([]byte(`
      (println (getSig (index)))
    `))
	if err2 == nil {
		t.Errorf("getSig failed,err= %v\n", err2)
	}
	//The mock parameter is not a string
	_, err3 := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getSig index))
    `))
	if err3 == nil {
		t.Errorf("getSig failed,err= %v\n", err3)
	}
	//The simulation parameter is empty
	_, err5 := lvm.vm.Eval([]byte(`
      (define index "")
      (println (getSig index))
    `))
	if err5 == nil {
		t.Errorf("getSig failed,err= %v\n", err5)
	}
	//The parameters passed in by the simulation are not equal to the required public key
	_, err4 := lvm.vm.Eval([]byte(`
      (define index "error")
      (println (getSig index))
    `))
	if err4 == nil {
		t.Errorf("getSig failed,err= %v\n", err4)
	}
}
func TestLispVMGetCurUnitHash(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
      (println (getCurUnitHash))
    `))
	if err != nil {
		t.Error("getCurUnitHash failed,err=", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getCurUnitHash index))
    `))
	if err1 == nil {
		t.Error("getCurUnitHash failed,err=", err1)
	}
}
func TestGetCurUnitHashToSign(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
      (println (getCurUnitHashToSign))
    `))
	if err != nil {
		t.Error("GetCurUnitHashToSign failed,err=", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getCurUnitHashToSign index))
    `))
	if err1 == nil {
		t.Error("getCurUnitHashToSign failed,err=", err1)
	}
}
func TestGetCurMsgHash(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	payloadhash := hash.Sum256([]byte("test"))
	msgheader := structure.MessageHeader{
		PayloadHash: payloadhash,
	}
	header := structure.InvokeMessage{
		Header: &msgheader,
	}
	invMsg := make([]structure.Message, 1)
	invMsg[0] = &header
	//invMsg[1] = &msg1
	unit := structure.Unit{
		Messages: invMsg,
	}

	c := vm.Context{
		TxUnit:     unit,
		TxMsgIndex: 0,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (println (getCurMsgHash))
    `))
	if err != nil {
		t.Error("GetCurMsgHash failed,err=", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getCurMsgHash index))
    `))
	if err1 == nil {
		t.Error("getCurMsgHash failed,err=", err1)
	}
}
func Test_globalParamCount(t *testing.T) {
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(println (globalParamCount))
	`))

	if err != nil {
		t.Error("globalParamCount error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (globalParamCount index))
    `))
	if err1 == nil {
		t.Error("globalParamCount failed,err=", err1)
	}
}
func Test_getGlobalParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
         (define key "gravity")
		(println (getGlobalParam key))
	`))

	if err != nil {
		t.Error("getGlobalParam error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getGlobalParam ))
	`))

	if err1 == nil {
		t.Error("getGlobalParam error:", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
		(println (getGlobalParam (key)))
	`))

	if err2 == nil {
		t.Error("getGlobalParam error:", err2)
	}
	//The mock parameter is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
         (define key 0)
		(println (getGlobalParam key))
	`))

	if err3 == nil {
		t.Error("getGlobalParam error:", err3)
	}
	//The simulation parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
         (define key "")
		(println (getGlobalParam key))
	`))

	if err4 == nil {
		t.Error("getGlobalParam error:", err4)
	}
}
func Test_getGlobalParamList(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
         (define key "gravity")
		(println (getGlobalParamList key))
	`))

	if err != nil {
		t.Error("getGlobalParamList error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getGlobalParamList ))
	`))

	if err1 == nil {
		t.Error("getGlobalParamList error:", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
		(println (getGlobalParamList (key)))
	`))

	if err2 == nil {
		t.Error("getGlobalParamList error:", err2)
	}
	//The mock parameter is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
         (define key 0)
		(println (getGlobalParamList key))
	`))

	if err3 == nil {
		t.Error("getGlobalParamList error:", err3)
	}
	//The simulation parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
         (define key "")
		(println (getGlobalParamList key))
	`))

	if err4 == nil {
		t.Error("getGlobalParamList error:", err4)
	}
	//The value data obtained is too long
	lvm.context.TxUnit.Messages[lvm.context.TxMsgIndex].(*structure.InvokeMessage).GlobalParamValue[0] = []byte("test")
	_, err5 := lvm.vm.Eval([]byte(`
         (define key "gravity")
		(println (getGlobalParamList key))
	`))

	if err5 == nil {
		t.Error("getGlobalParamList error:", err5)
	}
}
func TestInputCount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	contracInput := make([]*structure.ContractInput, 1)
	contracInput[0] = &structure.ContractInput{
		SourceOutput: uint32(1),
	}

	msg := structure.InvokeMessage{
		Inputs: contracInput,
	}

	invMsg := make([]structure.Message, 1)
	invMsg[0] = &msg
	//invMsg[1] = &msg1
	unit := structure.Unit{
		Messages: invMsg,
	}

	c := vm.Context{
		TxUnit:     unit,
		TxMsgIndex: 0,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (println (inputCount))
    `))
	if err != nil {
		t.Errorf("inputCount failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (inputCount index))
    `))
	if err1 == nil {
		t.Errorf("inputCount failed,err= %v\n", err1)
	}
}
func TestGetInputUnit(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	contracInput := make([]*structure.ContractInput, 1)
	SourceUnit := hash.Sum256([]byte("test"))
	contracInput[0] = &structure.ContractInput{
		SourceUnit: SourceUnit,
	}
	//contracInput[1] = &structure.ContractInput{
	//	SourceOutput:  uint32(1),
	//}

	msg := structure.InvokeMessage{
		Inputs: contracInput,
	}
	//msg1 := structure.InvokeMessage{
	//	Inputs:            contracInput,
	//}

	invMsg := make([]structure.Message, 1)
	invMsg[0] = &msg
	//invMsg[1] = &msg1
	unit := structure.Unit{
		Messages: invMsg,
	}

	c := vm.Context{
		TxUnit:     unit,
		TxMsgIndex: 0,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getInputUnit index))
    `))
	if err != nil {
		t.Errorf("getInputUnit failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (getInputUnit))
    `))
	if err1 == nil {
		t.Errorf("getInputUnit failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (getInputUnit (index)))
    `))
	if err2 == nil {
		t.Errorf("getInputUnit failed,err= %v\n", err2)
	}
	//The simulated parameter is not an int
	_, err3 := lvm.vm.Eval([]byte(`
      (define index "error")
      (println (getInputUnit index))
    `))
	if err3 == nil {
		t.Errorf("getInputUnit failed,err= %v\n", err3)
	}
}
func Test_getInputMsg(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define index 1)
		(println(getInputMsg index))
	`))

	if err != nil {
		t.Error("getInputMsg error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println(getInputMsg))
	`))

	if err1 == nil {
		t.Error("getInputMsg error:", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
		(println(getInputMsg (index)))
	`))

	if err2 == nil {
		t.Error("getInputMsg error:", err2)
	}
	//The simulated parameter is not an int
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println(getInputMsg index))
	`))

	if err3 == nil {
		t.Error("getInputMsg error:", err3)
	}
}
func Test_getInputParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 1)
		(define name "gravity")
		(define x (getInputParam index name))
		(println x)
	`))

	if err != nil {
		t.Error("exist error the vm execute", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 1)
		(println (getInputParam index))
	`))

	if err1 == nil {
		t.Error("getInputParam error:", err1)
	}
	//Error simulating the first parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(println (getInputParam (index) name))
	`))

	if err5 == nil {
		t.Error("getInputParam error:", err5)
	}
	//The first parameter to simulate is not an int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define name "gravity")
		(println (getInputParam index name))
	`))

	if err2 == nil {
		t.Error("getInputParam error:", err2)
	}
	//Error simulating second parameter execution
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 1)
		(println (getInputParam index (name)))
	`))

	if err6 == nil {
		t.Error("getInputParam error:", err6)
	}
	//The second argument of the simulation is not a string
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 1)
		(define name 0)
		(println (getInputParam index name))
	`))

	if err3 == nil {
		t.Error("getInputParam error:", err3)
	}
	//The second parameter of the simulation is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 1)
		(define name "")
		(println (getInputParam index name))
	`))

	if err4 == nil {
		t.Error("getInputParam error:", err4)
	}
}
func Test_getInputParamList(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 1)
		(define name "gravity")
		(define x (getInputParamList index name))
		(println x)
	`))

	if err != nil {
		t.Error("getInputParamList error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 1)
		(println (getInputParamList index))
	`))

	if err1 == nil {
		t.Error("getInputParamList error:", err1)
	}
	//Error simulating the first parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(println (getInputParamList (index) name))
	`))

	if err5 == nil {
		t.Error("getInputParamList error:", err5)
	}
	//The first parameter to simulate is not an int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define name "gravity")
		(println (getInputParamList index name))
	`))

	if err2 == nil {
		t.Error("getInputParamList error:", err2)
	}
	//Error simulating second parameter execution
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 1)
		(println (getInputParamList index (name)))
	`))

	if err6 == nil {
		t.Error("getInputParamList error:", err6)
	}
	//The second argument of the simulation is not a string
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 1)
		(define name 0)
		(println (getInputParamList index name))
	`))

	if err3 == nil {
		t.Error("getInputParamList error:", err3)
	}
	//The second parameter of the simulation is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 1)
		(define name "")
		(println (getInputParamList index name))
	`))

	if err4 == nil {
		t.Error("getInputParamList error:", err4)
	}
	//The value data obtained is too long
	lvm.context.TxUnit.Messages[lvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[0].InputParamsValue[0] = []byte("test")
	_, err7 := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "gravity")
		(define x (getInputParamList index name))
		(println x)
	`))

	if err7 == nil {
		t.Error("getInputParamList error:", err7)
	}
}
func Test_getInputPreOut(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 1)
		(println (getInputPreOut index))
	`))

	if err != nil {
		t.Error("get input pre out failed:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getInputPreOut))
	`))

	if err1 == nil {
		t.Error("get input pre out failed:", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
		(println (getInputPreOut (index)))
	`))

	if err2 == nil {
		t.Error("get input pre out failed:", err2)
	}
	//The simulated parameter is not of type int
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (getInputPreOut index))
	`))

	if err3 == nil {
		t.Error("get input pre out failed:", err3)
	}
}
func Test_getPrevOutAmount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (getPrevOutAmount index))
	`))

	if err != nil {
		t.Error("get pre out amount failed:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println (getPrevOutAmount))
	`))

	if err1 == nil {
		t.Error("get pre out amount failed:", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
		(println (getPrevOutAmount (index)))
	`))

	if err2 == nil {
		t.Error("get pre out amount failed:", err2)
	}
	//The simulated parameter is not of type int
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println (getPrevOutAmount index))
	`))

	if err3 == nil {
		t.Error("get pre out amount failed:", err3)
	}
}
func Test_getPrevOutParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 1)
		(define name "gravity")
		(define x (getPrevOutParam index name))
		(println x)
	`))

	if err != nil {
		t.Error("getPrevOutParam error::", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 1)
		(println (getPrevOutParam index))
	`))

	if err1 == nil {
		t.Error("getPrevOutParam error:", err1)
	}
	//Error simulating the first parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(println (getPrevOutParam (index) name))
	`))

	if err5 == nil {
		t.Error("getPrevOutParam error:", err5)
	}
	//The first parameter to simulate is not an int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define name "gravity")
		(println (getPrevOutParam index name))
	`))

	if err2 == nil {
		t.Error("getPrevOutParam error:", err2)
	}
	//Error simulating second parameter execution
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 1)
		(println (getPrevOutParam index (name)))
	`))

	if err6 == nil {
		t.Error("getPrevOutParam error:", err6)
	}
	//The second argument of the simulation is not a string
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 1)
		(define name 0)
		(println (getPrevOutParam index name))
	`))

	if err3 == nil {
		t.Error("getPrevOutParam error:", err3)
	}
	//The second parameter of the simulation is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 1)
		(define name "")
		(println (getPrevOutParam index name))
	`))

	if err4 == nil {
		t.Error("getPrevOutParam error:", err4)
	}

}
func Test_getPreOutExtends(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (getPreOutExtends index))
	`))

	if err != nil {
		t.Error("getPreOutExtends error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println(getPreOutExtends))
	`))

	if err1 == nil {
		t.Error("getPreOutExtends error:", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
		(println(getPreOutExtends (index)))
	`))

	if err2 == nil {
		t.Error("getPreOutExtends error:", err2)
	}
	//The simulated parameter is not an int
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println(getPreOutExtends index))
	`))

	if err3 == nil {
		t.Error("getPreOutExtends error:", err3)
	}
}
func Test_getOutputAmount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (getOutputAmount index))
	`))

	if err != nil {
		t.Error("getOutputAmount error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println(getOutputAmount))
	`))

	if err1 == nil {
		t.Error("getOutputAmount error:", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
		(println(getOutputAmount (index)))
	`))

	if err2 == nil {
		t.Error("getOutputAmount error:", err2)
	}
	//The simulated parameter is not an int
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println(getOutputAmount index))
	`))

	if err3 == nil {
		t.Error("getOutputAmount error:", err3)
	}
}
func Test_getOutputParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "gravity")
		(println (getOutputParam index name))
	`))

	if err != nil {
		t.Error("getOutputParam error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getOutputParam index))
	`))

	if err1 == nil {
		t.Error("getOutputParam error:", err1)
	}
	//Error simulating the first parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(println (getOutputParam (index) name))
	`))

	if err5 == nil {
		t.Error("getOutputParam error:", err5)
	}
	//The first parameter to simulate is not an int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define name "gravity")
		(println (getOutputParam index name))
	`))

	if err2 == nil {
		t.Error("getOutputParam error:", err2)
	}
	//Error simulating second parameter execution
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getOutputParam index (name)))
	`))

	if err6 == nil {
		t.Error("getOutputParam error:", err6)
	}
	//The second argument of the simulation is not a string
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define name 0)
		(println (getOutputParam index name))
	`))

	if err3 == nil {
		t.Error("getOutputParam error:", err3)
	}
	//The second parameter of the simulation is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define name "")
		(println (getOutputParam index name))
	`))

	if err4 == nil {
		t.Error("getOutputParam error:", err4)
	}
}
func Test_getOutputParamList(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "gravity")
		(println (getOutputParamList index name))
	`))

	if err != nil {
		t.Error("getOutputParamList error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getOutputParamList index))
	`))

	if err1 == nil {
		t.Error("getOutputParamList error:", err1)
	}
	//Error simulating the first parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(println (getOutputParamList (index) name))
	`))

	if err5 == nil {
		t.Error("getOutputParamList error:", err5)
	}
	//The first parameter to simulate is not an int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define name "gravity")
		(println (getOutputParamList index name))
	`))

	if err2 == nil {
		t.Error("getOutputParamList error:", err2)
	}
	//Error simulating second parameter execution
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getOutputParamList index (name)))
	`))

	if err6 == nil {
		t.Error("getOutputParamList error:", err6)
	}
	//The second argument of the simulation is not a string
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define name 0)
		(println (getOutputParamList index name))
	`))

	if err3 == nil {
		t.Error("getOutputParamList error:", err3)
	}
	//The second parameter of the simulation is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define name "")
		(println (getOutputParamList index name))
	`))

	if err4 == nil {
		t.Error("getOutputParamList error:", err4)
	}
	//The value data obtained is too long
	lvm.context.TxUnit.Messages[lvm.context.TxMsgIndex].(*structure.InvokeMessage).Outputs[0].OutputParamsValue[0] = []byte("test")
	_, err7 := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "gravity")
		(println (getOutputParamList index name))
	`))

	if err7 == nil {
		t.Error("getOutputParamList error:", err7)
	}
}
func Test_getOutputExtends(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(println (getOutputExtends index))
	`))

	if err != nil {
		t.Error("getOutputExtends error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
		(println(getOutputExtends))
	`))

	if err1 == nil {
		t.Error("getOutputExtends error:", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
		(println(getOutputExtends (index)))
	`))

	if err2 == nil {
		t.Error("getOutputExtends error:", err2)
	}
	//The simulated parameter is not an int
	_, err3 := lvm.vm.Eval([]byte(`
		(define index "error")
		(println(getOutputExtends index))
	`))

	if err3 == nil {
		t.Error("getOutputExtends error:", err3)
	}
}

func Test_getCurMCI(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(println (getCurrentMCI))
	`))

	if err != nil {
		t.Error("getCurrentMCI error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (getCurrentMCI index))
	`))

	if err1 == nil {
		t.Error("getCurrentMCI error:", err1)
	}

}
func Test_hasPrevOutParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "gravity")
		(define x (hasPrevOutParam index name))
		(println x)
	`))

	if err != nil {
		t.Error(" hasPrevOutParam eror:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (hasPrevOutParam index))
	`))

	if err1 == nil {
		t.Error("hasPrevOutParam error:", err1)
	}
	//Error simulating the first parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(println (hasPrevOutParam (index) name))
	`))

	if err5 == nil {
		t.Error("hasPrevOutParam error:", err5)
	}
	//The first parameter to simulate is not an int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define name "gravity")
		(println (hasPrevOutParam index name))
	`))

	if err2 == nil {
		t.Error("hasPrevOutParam error:", err2)
	}
	//Error simulating second parameter execution
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (hasPrevOutParam index (name)))
	`))

	if err6 == nil {
		t.Error("hasPrevOutParam error:", err6)
	}
	//The second argument of the simulation is not a string
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define name 0)
		(println (hasPrevOutParam index name))
	`))

	if err3 == nil {
		t.Error("hasPrevOutParam error:", err3)
	}
	//The second parameter of the simulation is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define name "")
		(println (hasPrevOutParam index name))
	`))

	if err4 == nil {
		t.Error("hasPrevOutParam error:", err4)
	}
	//The simulation parameter is different from the required value
	_, err7 := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "error")
		(define x (hasPrevOutParam index name))
		(println x)
	`))

	if err7 != nil {
		t.Error(" hasPrevOutParam eror:", err7)
	}

}
func Test_hasPKByAddr(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	Definition := []byte("test")
	author := &structure.Author{
		Address:    Definition,
		Definition: Definition,
	}

	auth := make([]*structure.Author, 1)
	auth[0] = author
	unit := structure.Unit{
		Authors: auth,
	}
	c := vm.Context{
		TxUnit: unit,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (define index "test")
      (println (hasPKByAddr index))
    `))
	if err != nil {
		t.Errorf("hasPKByAddr failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (hasPKByAddr))
    `))
	if err1 == nil {
		t.Errorf("hasPKByAddr failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (hasPKByAddr (index)))
    `))
	if err2 == nil {
		t.Errorf("hasPKByAddr failed,err= %v\n", err2)
	}
	//The mock parameter type is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (hasPKByAddr index))
    `))
	if err3 == nil {
		t.Errorf("hasPKByAddr failed,err= %v\n", err3)
	}
	//The simulation parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
      (define index "")
      (println (hasPKByAddr index))
    `))
	if err4 == nil {
		t.Errorf("hasPKByAddr failed,err= %v\n", err4)
	}
	//The simulation parameter is different from the required value
	_, err5 := lvm.vm.Eval([]byte(`
      (define index "error")
      (println (hasPKByAddr index))
    `))
	if err5 != nil {
		t.Errorf("hasPKByAddr failed,err= %v\n", err5)
	}
}
func Test_getAuthorSig(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	Definition := []byte("pk")
	prvksig := []byte("preksig")
	author := &structure.Author{
		Address:       Definition,
		Authentifiers: prvksig,
		Definition:    Definition,
	}

	auth := make([]*structure.Author, 1)
	auth[0] = author
	unit := structure.Unit{
		Authors: auth,
	}
	c := vm.Context{
		TxUnit: unit,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getAuthorSig index))
    `))
	if err != nil {
		t.Errorf("getSig failed,err= %v\n", err)
	}
	//The number of simulated parameters is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (getAuthorSig))
    `))
	if err1 == nil {
		t.Errorf("getSig failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (getAuthorSig (index)))
    `))
	if err2 == nil {
		t.Errorf("getSig failed,err= %v\n", err2)
	}
	//The simulated parameter is not an int
	_, err3 := lvm.vm.Eval([]byte(`
      (define index "error")
      (println (getAuthorSig index))
    `))
	if err3 == nil {
		t.Errorf("getSig failed,err= %v\n", err3)
	}
}
func Test_getAuthorAddr(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	Definition := []byte("test")
	author := &structure.Author{
		Address:    Definition,
		Definition: Definition,
	}

	auth := make([]*structure.Author, 1)
	auth[0] = author
	unit := structure.Unit{
		Authors: auth,
	}
	c := vm.Context{
		TxUnit: unit,
	}
	lvm.context = c
	_, err := lvm.vm.Eval([]byte(`
      (define index 0)
      (println (getAuthorAddr index))
    `))
	if err != nil {
		t.Errorf("getAuthorAddr failed,err= %v\n", err)
	}
	//The number of simulated parameters is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (getAuthorAddr))
    `))
	if err1 == nil {
		t.Errorf("getAuthorAddr failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (getAuthorAddr (index)))
    `))
	if err2 == nil {
		t.Errorf("getAuthorAddr failed,err= %v\n", err2)
	}
	//The simulated parameter is not an int
	_, err3 := lvm.vm.Eval([]byte(`
      (define index "error")
      (println (getAuthorAddr index))
    `))
	if err3 == nil {
		t.Errorf("getAuthorAddr failed,err= %v\n", err3)
	}
}
func Test_hasInputParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "gravity")
		(define x (hasInputParam index name))
		(println x)
	`))

	if err != nil {
		t.Error("hasInputParam error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (hasInputParam index))
	`))

	if err1 == nil {
		t.Error("hasInputParam error:", err1)
	}
	//Error simulating the first parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(println (hasInputParam (index) name))
	`))

	if err5 == nil {
		t.Error("hasInputParam error:", err5)
	}
	//The first parameter to simulate is not an int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define name "gravity")
		(println (hasInputParam index name))
	`))

	if err2 == nil {
		t.Error("hasInputParam error:", err2)
	}
	//Error simulating second parameter execution
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 0)
		(println (hasInputParam index (name)))
	`))

	if err6 == nil {
		t.Error("hasInputParam error:", err6)
	}
	//The second argument of the simulation is not a string
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define name 0)
		(println (hasInputParam index name))
	`))

	if err3 == nil {
		t.Error("hasInputParam error:", err3)
	}
	//The second parameter of the simulation is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define name "")
		(println (hasInputParam index name))
	`))

	if err4 == nil {
		t.Error("hasInputParam error:", err4)
	}
	//The simulation parameter is different from the required value
	_, err7 := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "error")
		(define x (hasInputParam index name))
		(println x)
	`))

	if err7 != nil {
		t.Error(" hasInputParam eror:", err7)
	}

}
func Test_hasCurPrevOutParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(define x (hasCurPrevOutParam name))
		(println x)
	`))

	if err != nil {
		t.Errorf("hasCurPrevOutParam failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (hasCurPrevOutParam))
    `))
	if err1 == nil {
		t.Errorf("hasCurPrevOutParam failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (hasCurPrevOutParam (name)))
    `))
	if err2 == nil {
		t.Errorf("hasCurPrevOutParam failed,err= %v\n", err2)
	}
	//The mock parameter type is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
      (define name 0)
      (println (hasCurPrevOutParam name))
    `))
	if err3 == nil {
		t.Errorf("hasCurPrevOutParam failed,err= %v\n", err3)
	}
	//The simulation parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
      (define name "")
      (println (hasCurPrevOutParam name))
    `))
	if err4 == nil {
		t.Errorf("hasCurPrevOutParam failed,err= %v\n", err4)
	}
	//The simulation parameter is different from the required value
	_, err5 := lvm.vm.Eval([]byte(`
      (define name "error")
      (println (hasCurPrevOutParam name))
    `))
	if err5 != nil {
		t.Errorf("hasCurPrevOutParam failed,err= %v\n", err5)
	}
}
func Test_getCurPrevOutParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(define x (getCurPrevOutParam name))
		(println x)
	`))

	if err != nil {
		t.Errorf("getCurPrevOutParam failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (getCurPrevOutParam))
    `))
	if err1 == nil {
		t.Errorf("getCurPrevOutParam failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (getCurPrevOutParam (name)))
    `))
	if err2 == nil {
		t.Errorf("getCurPrevOutParam failed,err= %v\n", err2)
	}
	//The mock parameter type is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
      (define name 0)
      (println (getCurPrevOutParam name))
    `))
	if err3 == nil {
		t.Errorf("getCurPrevOutParam failed,err= %v\n", err3)
	}
	//The simulation parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
      (define name "")
      (println (getCurPrevOutParam name))
    `))
	if err4 == nil {
		t.Errorf("getCurPrevOutParam failed,err= %v\n", err4)
	}
	//The simulation parameter is different from the required value
	_, err5 := lvm.vm.Eval([]byte(`
      (define name "error")
      (println (getCurPrevOutParam name))
    `))
	if err5 == nil {
		t.Errorf("getCurPrevOutParam failed,err= %v\n", err5)
	}
}
func Test_getCurPrevOutParamList(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(define x (getCurPrevOutParamList name))
		(println x)
	`))

	if err != nil {
		t.Errorf("getCurPrevOutParamList failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (getCurPrevOutParamList))
    `))
	if err1 == nil {
		t.Errorf("getCurPrevOutParamList failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (getCurPrevOutParamList (name)))
    `))
	if err2 == nil {
		t.Errorf("getCurPrevOutParamList failed,err= %v\n", err2)
	}
	//The mock parameter type is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
      (define name 0)
      (println (getCurPrevOutParamList name))
    `))
	if err3 == nil {
		t.Errorf("getCurPrevOutParamList failed,err= %v\n", err3)
	}
	//The simulation parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
      (define name "")
      (println (getCurPrevOutParamList name))
    `))
	if err4 == nil {
		t.Errorf("getCurPrevOutParamList failed,err= %v\n", err4)
	}
	//The simulation parameter is different from the required value
	_, err5 := lvm.vm.Eval([]byte(`
      (define name "error")
      (println (getCurPrevOutParamList name))
    `))
	if err5 == nil {
		t.Errorf("getCurPrevOutParamList failed,err= %v\n", err5)
	}
	//The value data obtained is too long
	lvm.context.PrevOut.OutputParamsValue[0] = []byte("test")
	_, err6 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(define x (getCurPrevOutParamList name))
		(println x)
	`))

	if err6 == nil {
		t.Errorf("getCurPrevOutParamList failed,err= %v\n", err6)
	}
}
func Test_getCurPrevOutAmount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define x (getCurPrevOutAmount))
		(println x)
	`))

	if err != nil {
		t.Error("getCurPrevOutAmount error:", err)
	}
	//The number of simulated parameters is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define x (getCurPrevOutAmount index))
		(println x)
	`))

	if err1 == nil {
		t.Error("getCurPrevOutAmount error:", err1)
	}
}
func Test_getCurPrevOutExtends(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define x (getCurPrevOutExtends))
		(println x)
	`))

	if err != nil {
		t.Error("getCurPrevOutExtends error:", err)
	}
	//The number of simulated parameters is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define x (getCurPrevOutExtends index))
		(println x)
	`))

	if err1 == nil {
		t.Error("getCurPrevOutExtends error:", err1)
	}
}
func Test_getCurInputParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(define x (getCurInputParam name))
		(println x)
	`))

	if err != nil {
		t.Errorf("getCurInputParam failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (getCurInputParam))
    `))
	if err1 == nil {
		t.Errorf("getCurInputParam failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (getCurInputParam (name)))
    `))
	if err2 == nil {
		t.Errorf("getCurInputParam failed,err= %v\n", err2)
	}
	//The mock parameter type is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
      (define name 0)
      (println (getCurInputParam name))
    `))
	if err3 == nil {
		t.Errorf("getCurInputParam failed,err= %v\n", err3)
	}
	//The simulation parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
      (define name "")
      (println (getCurInputParam name))
    `))
	if err4 == nil {
		t.Errorf("getCurInputParam failed,err= %v\n", err4)
	}
	//The simulation parameter is different from the required value
	_, err5 := lvm.vm.Eval([]byte(`
      (define name "error")
      (println (getCurInputParam name))
    `))
	if err5 == nil {
		t.Errorf("getCurInputParam failed,err= %v\n", err5)
	}

}
func Test_hasCurInputParam(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(define x (hasCurInputParam name))
		(println x)
	`))

	if err != nil {
		t.Errorf("hasCurInputParam failed,err= %v\n", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
      (println (hasCurInputParam))
    `))
	if err1 == nil {
		t.Errorf("hasCurInputParam failed,err= %v\n", err1)
	}
	//Error in simulated parameter execution
	_, err2 := lvm.vm.Eval([]byte(`
      (println (hasCurInputParam (name)))
    `))
	if err2 == nil {
		t.Errorf("hasCurInputParam failed,err= %v\n", err2)
	}
	//The mock parameter type is not a string type
	_, err3 := lvm.vm.Eval([]byte(`
      (define name 0)
      (println (hasCurInputParam name))
    `))
	if err3 == nil {
		t.Errorf("hasCurInputParam failed,err= %v\n", err3)
	}
	//The simulation parameter is empty
	_, err4 := lvm.vm.Eval([]byte(`
      (define name "")
      (println (hasCurInputParam name))
    `))
	if err4 == nil {
		t.Errorf("hasCurInputParam failed,err= %v\n", err4)
	}
	//The simulation parameter is different from the required value
	_, err5 := lvm.vm.Eval([]byte(`
      (define name "error")
      (println (hasCurInputParam name))
    `))
	if err5 != nil {
		t.Errorf("hasCurInputParam failed,err= %v\n", err5)
	}

}
func Test_getCurInputParamsCount(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define x (getCurInputParamsCount))
		(println x)
	`))

	if err != nil {
		t.Error("getCurInputParamsCount error:", err)
	}
	//The number of simulated parameters is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
         (define index 0)
		(define x (getCurInputParamsCount index))
		(println x)
	`))

	if err1 == nil {
		t.Error("getCurInputParamsCount error:", err1)
	}
}
func Test_getCurInputUnit(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define x (getCurInputUnit))
		(println x)
	`))

	if err != nil {
		t.Error("getCurInputUnit error:", err)
	}
	//The number of simulated parameters is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define x (getCurInputUnit index))
		(println x)
	`))

	if err1 == nil {
		t.Error("getCurInputUnit error:", err1)
	}
}
func Test_getCurInputMsg(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define x (getCurInputMsg))
		(println x)
	`))

	if err != nil {
		t.Error("getCurInputMsg error:", err)
	}
	//The number of simulated parameters is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define x (getCurInputMsg index))
		(println x)
	`))

	if err1 == nil {
		t.Error("getCurInputMsg error:", err1)
	}

}
func Test_getCurInputOutput(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	setLvm(lvm)
	_, err := lvm.vm.Eval([]byte(`
		(define x (getCurInputOutput))
		(println x)
	`))

	if err != nil {
		t.Error("getCurInputOutput error:", err)
	}
	//The number of simulated parameters is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 0)
		(define x (getCurInputOutput index))
		(println x)
	`))

	if err1 == nil {
		t.Error("getCurInputOutput error:", err1)
	}
}
func Test_getPrevOutParamList(t *testing.T) {
	//Run successfully
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	setLvm(lvm)

	_, err := lvm.vm.Eval([]byte(`
		(define index 1)
		(define name "gravity")
		(define x (getPrevOutParamList index name))
		(println x)
	`))

	if err != nil {
		t.Error("getPrevOutParamList error:", err)
	}
	//The number of parameters passed in the simulation is incorrect
	_, err1 := lvm.vm.Eval([]byte(`
        (define index 1)
		(println (getPrevOutParamList index))
	`))

	if err1 == nil {
		t.Error("getPrevOutParamList error:", err1)
	}
	//Error simulating the first parameter execution
	_, err5 := lvm.vm.Eval([]byte(`
		(define name "gravity")
		(println (getPrevOutParamList (index) name))
	`))

	if err5 == nil {
		t.Error("getPrevOutParamList error:", err5)
	}
	//The first parameter to simulate is not an int
	_, err2 := lvm.vm.Eval([]byte(`
        (define index "error")
		(define name "gravity")
		(println (getPrevOutParamList index name))
	`))

	if err2 == nil {
		t.Error("getPrevOutParamList error:", err2)
	}
	//Error simulating second parameter execution
	_, err6 := lvm.vm.Eval([]byte(`
        (define index 1)
		(println (getPrevOutParamList index (name)))
	`))

	if err6 == nil {
		t.Error("getPrevOutParamList error:", err6)
	}
	//The second argument of the simulation is not a string
	_, err3 := lvm.vm.Eval([]byte(`
        (define index 1)
		(define name 0)
		(println (getPrevOutParamList index name))
	`))

	if err3 == nil {
		t.Error("getPrevOutParamList error:", err3)
	}
	//The second parameter of the simulation is empty
	_, err4 := lvm.vm.Eval([]byte(`
        (define index 1)
		(define name "")
		(println (getPrevOutParamList index name))
	`))

	if err4 == nil {
		t.Error("getPrevOutParamList error:", err4)
	}
	//The value data obtained is too long
	lvm.context.FetchPrevOut = GetContractOutput1
	_, err7 := lvm.vm.Eval([]byte(`
		(define index 0)
		(define name "gravity")
		(define x (getPrevOutParamList index name))
		(println x)

	`))

	if err7 == nil {
		t.Error("getPrevOutParamList error:", err7)
	}
}

func setLvm(lvm *LispVM) {
	contraInput := make([]*structure.ContractInput, 4)
	byt := []byte("1")
	param := make([]string, 4)
	paramName := make([][]byte, 4)
	res := make([]hash.HashType, 4)
	for i := 0; i < len(param); i++ {
		param[i] = "gravity"
		paramName[i] = append(paramName[i], 10)
		paramName[i] = append(paramName[i], []byte("gravitygra")...)
		res[i] = hash.Sum256([]byte("test"))
	}
	for i := 0; i < len(contraInput); i++ {
		contraInput[i] = generateInputs(hash.Sum256([]byte("test")), 12, 34, param, paramName)
	}

	contraOutPut := make([]*structure.ContractOutput, 4)
	for i := 0; i < len(contraOutPut); i++ {
		contraOutPut[i] = generateOutPuts(uint64(2100), byt, param, paramName, res)
	}

	interMsg := make([]structure.Message, 4)
	for i := 0; i < len(interMsg); i++ {
		interMsg[i] = generateInvokeMsg(&structure.MessageHeader{}, defaultBytes, 1, param, paramName, contraOutPut, contraInput)
	}
	utxoheader := structure.UtxoHeader{
		Amount:    uint64(3),
		Address:   defaultBytes,
		SpentMci:  uint64(0),
		SpentHash: res,
	}
	contractdef := &structure.ContractDef{
		Address:     defaultBytes,
		ParamsKey:   param,
		ParamsValue: paramName,
	}
	lvm.context = vm.Context{
		TxUnit: structure.Unit{
			Messages: interMsg,
		},
		TxMsgIndex:   1,
		FetchPrevOut: GetContractOutput,
		MCI:          10,
		PrevOut:      generateTxUtxo(utxoheader, defaultBytes, 0, 0, defaultBytes, byt, param, paramName, res),
		Input:        generateInputs(hash.Sum256([]byte("test")), 12, 34, param, paramName),
		ContractDef:  contractdef,
	}
}

func generateOutPuts(amount uint64, extends []byte, key []string,
	value [][]byte, res []hash.HashType) *structure.ContractOutput {
	output := &structure.ContractOutput{
		Amount:  amount,
		Extends: extends,

		OutputParamsKey:   key,
		OutputParamsValue: value,
		Restricts:         res,
	}

	return output
}

func generateInputs(sourceUnit hash.HashType, sourceMsg, sourceOut uint32,
	inputParam []string, value [][]byte) *structure.ContractInput {
	input := &structure.ContractInput{
		SourceUnit:    sourceUnit,
		SourceMessage: sourceMsg,
		SourceOutput:  sourceOut,

		InputParamsKey:   inputParam,
		InputParamsValue: value,
	}

	return input
}

func generateInvokeMsg(header *structure.MessageHeader, asset hash.HashType, contractIndex uint32,
	stringParam []string, param [][]byte, outputs []*structure.ContractOutput,
	inputs []*structure.ContractInput) *structure.InvokeMessage {
	invokeMsg := &structure.InvokeMessage{
		Header: header,
		Asset:  asset,

		ContractAddr:     hash.Sum256(defaultBytes),
		GlobalParamKey:   stringParam,
		GlobalParamValue: param,

		Outputs: outputs,
		Inputs:  inputs,
	}

	return invokeMsg
}
func generateTxUtxo(UtxoHeader structure.UtxoHeader, Unit hash.HashType, Message uint32, Output uint32, Asset hash.HashType, Extends []byte, OutputParamsKey []string, OutputParamsValue [][]byte, Restricts []hash.HashType) *structure.TxUtxo {
	return &structure.TxUtxo{
		UtxoHeader: UtxoHeader,
		Unit:       Unit,
		Message:    Message,
		Output:     Output,
		Asset:      Asset,
		Extends:    Extends,

		OutputParamsKey:   OutputParamsKey,
		OutputParamsValue: OutputParamsValue,

		Restricts: Restricts,
	}
}
func GetContractOutput1(input *structure.ContractInput) *structure.ContractOutput {
	contractOutput := structure.NewContractOutput()
	contractOutput.Amount = 3
	contractOutput.Extends = []byte("Hello World!")

	//for test
	pubKByte, err := hex.DecodeString("0421af2c7f64c10a34fbf4891ac5862a71f0c5805e0e4b50ce03244943a7859b7dd26e5a9ffd9ac0bccb5aadc81552b5961cc7f361a244df703d248aa5c13fb013")
	if err != nil {
		return nil
	}
	paraValue2 := make([][]byte, 2)
	paraValue2[0] = append(paraValue2[0], []byte("test")...)
	contractOutput.AddParam("addr", hash.Sum256([]byte(pubKByte)))
	contractOutput.AddParam("gravity", paraValue2[0])

	contractOutput.AddRestrict(hash.Sum256([]byte(pubKByte)))
	contractOutput.AddRestrict(hash.Sum256([]byte(pubKByte)))

	return contractOutput
}

var defaultBytes = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}
