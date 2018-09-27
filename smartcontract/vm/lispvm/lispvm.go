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
	"github.com/SHDMT/gravity/infrastructure/log"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/smartcontract/vm"
	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp"
)

const (
	//LispVMVersion is current VM version
	LispVMVersion = 1
)

//LispVM is Lisp virtual machine
type LispVM struct {
	id          uint32
	config      vm.Config
	context     vm.Context
	curContract structure.Contract
	vm          *lisp.Lisp
}

//ID returns unique identification of LispVM
func (lispvm *LispVM) ID() uint32 {
	return 0
}

//Version returns LispVM version
func (lispvm *LispVM) Version() uint16 {
	return LispVMVersion
}

//ScriptCode returns  Lisp language code
func (lispvm *LispVM) ScriptCode() byte {
	return vm.LispScriptCode
}

//Config returns
func (lispvm *LispVM) Config() vm.Config {
	return lispvm.config
}

//Context returns context of LispVM
func (lispvm *LispVM) Context() vm.Context {
	return lispvm.context
}

//Exec returns contract execute result
func (lispvm *LispVM) Exec(contract *structure.Contract) (ret bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Lisp vm error, recovered from %v", r)
			ret = false
		}
	}()

	result, err := lispvm.vm.Eval(contract.Code)
	if err != nil {
		log.Error("execute the contract failed:", err)
		return false
	}

	return result.Bool()
}

//SetEnv setup LispVM environment
func (lispvm *LispVM) SetEnv(context vm.Context, config vm.Config) {
	lispvm.context = context
	lispvm.config = config
}

//NewLispVM creates a new LispVM object
func NewLispVM(context vm.Context, config vm.Config) *LispVM {
	lispvm := new(LispVM)
	lispvm.context = context
	lispvm.config = config
	lispvm.vm = lisp.NewLisp()

	lisp.Add("verify", lispvm.verify)
	lisp.Add("hash", lispvm.hash)
	lisp.Add("verifyMultiSign", lispvm.verifyMultiSign)

	lisp.Add("sigCount", lispvm.sigCount)
	lisp.Add("getPK", lispvm.getPK)
	lisp.Add("getPKByAddr", lispvm.getPKByAddr)
	lisp.Add("getSig", lispvm.getSig)
	lisp.Add("hasPKByAddr", lispvm.hasPKByAddr)
	lisp.Add("getAuthorSig", lispvm.getAuthorSig)
	lisp.Add("getAuthorAddr", lispvm.getAuthorAddr)

	lisp.Add("getCurUnitHash", lispvm.getCurUnitHash)
	lisp.Add("getCurUnitHashToSign", lispvm.getCurUnitHashToSign)
	lisp.Add("getCurMsgHash", lispvm.getCurMsgHash)

	lisp.Add("getCurrentMCI", lispvm.getCurrentMCI)
	lisp.Add("countBytes", lispvm.countBytes)

	if config.Mode == vm.VMModeContract {
		lisp.Add("inputCount", lispvm.inputCount)
		lisp.Add("getInputUnit", lispvm.getInputUnit)
		lisp.Add("getInputMsg", lispvm.getInputMsg)
		lisp.Add("getInputParam", lispvm.getInputParam)
		lisp.Add("getInputPreOut", lispvm.getInputPreOut)
		lisp.Add("getPrevOutAmount", lispvm.getPrevOutAmount)
		lisp.Add("getPrevOutParam", lispvm.getPrevOutParam)
		//add on 20180817
		lisp.Add("getPrevOutParamList", lispvm.getPrevOutParamList)
		//add on 20180817
		lisp.Add("getPreOutExtends", lispvm.getPreOutExtends)
		lisp.Add("getOutputAmount", lispvm.getOutputAmount)
		lisp.Add("getOutputParam", lispvm.getOutputParam)
		lisp.Add("getOutputExtends", lispvm.getOutputExtends)

		lisp.Add("cap", lispvm.cap)
		lisp.Add("globalParamCount", lispvm.globalParamCount)
		lisp.Add("getGlobalParam", lispvm.getGlobalParam)
		lisp.Add("isDenominations", lispvm.isDenominations)
		lisp.Add("getDenominationCount", lispvm.getDenominationCount)
		lisp.Add("getDenomination", lispvm.getDenomination)
		lisp.Add("getAssetContractCount", lispvm.getAssetContractCount)
		lisp.Add("getAssetContract", lispvm.getAssetContract)
		lisp.Add("getAllocationsCount", lispvm.getAllocationsCount)
		lisp.Add("getAllocationsAddr", lispvm.getAllocationsAddr)
		lisp.Add("getAllocationsAmount", lispvm.getAllocationsAmount)
		lisp.Add("getAssetExtends", lispvm.getAssetExtends)
		lisp.Add("getPublisherAddr", lispvm.getPublisherAddr)
		lisp.Add("getPublisherUnitMCI", lispvm.getPublishUnitMCI)
		lisp.Add("getContractParamCount", lispvm.getContractParamCount)
		lisp.Add("getCurContractDefParamCount", lispvm.getCurContractDefParamCount)
		lisp.Add("getCurContractDefParamName", lispvm.getCurContractDefParamName)
		lisp.Add("getCurContractDefParam", lispvm.getCurContractDefParam)
		//add on 20180817
		lisp.Add("getCurContractDefParamList", lispvm.getCurContractDefParamList)
		//add on 20180817
		lisp.Add("getContractParamName", lispvm.getContractParamName)

		lisp.Add("getContractParamByIndex", lispvm.getContractParamByIndex)
		lisp.Add("getContractParam", lispvm.getContractParam)
		lisp.Add("isExistAtGlobalParam", lispvm.isExistAtGlobalParam)
		lisp.Add("isExistAtOutputParam", lispvm.isExistAtOutputParam)
		lisp.Add("isExistAtInputParam", lispvm.isExistAtInputParam)
		lisp.Add("calcOutputAmount", lispvm.calcOutputAmount)
		lisp.Add("calcInputAmount", lispvm.calcInputAmount)
		lisp.Add("calcBalance", lispvm.calcBalance)

		lisp.Add("hasPrevOutParam", lispvm.hasPrevOutParam)

		lisp.Add("hasInputParam", lispvm.hasInputParam)

		lisp.Add("getInputParamList", lispvm.getInputParamList)
		lisp.Add("getOutputParamList", lispvm.getOutputParamList)
		lisp.Add("getGlobalParamList", lispvm.getGlobalParamList)
	} else if config.Mode == vm.VMModeRestrict {
		lisp.Add("hasCurPrevOutParam", lispvm.hasCurPrevOutParam)
		lisp.Add("getCurPrevOutParam", lispvm.getCurPrevOutParam)
		//2018/8/24
		lisp.Add("getCurPrevOutParamList", lispvm.getCurPrevOutParamList)
		//2018/8/24
		lisp.Add("getCurPrevOutAmount", lispvm.getCurPrevOutAmount)
		lisp.Add("getCurPrevOutExtends", lispvm.getCurPrevOutExtends)

		lisp.Add("getCurInputParam", lispvm.getCurInputParam)
		lisp.Add("hasCurInputParam", lispvm.hasCurInputParam)
		lisp.Add("getCurInputParamsCount", lispvm.getCurInputParamsCount)
		lisp.Add("getCurInputUnit", lispvm.getCurInputUnit)
		lisp.Add("getCurInputMsg", lispvm.getCurInputMsg)
		lisp.Add("getCurInputOutput", lispvm.getCurInputOutput)
	}
	return lispvm
}
