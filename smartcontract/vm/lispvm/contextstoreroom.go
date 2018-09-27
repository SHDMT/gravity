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
	"bytes"

	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp"
	"github.com/pkg/errors"
)

const (
        authorIndexError="the author index is not int"
        paramLenError="paramLen is too long"
	    inputParamIntError="input param is not int"
	    contentIntError="content is not int"
	    contentStringError="content is not string"
	    inputParamStringError="input paramKey is not string"
	    inputParamError="input paramKey is empty"
)
//sigCount returns the number of signatures
func (lispvm *LispVM) sigCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	sigCount := int64(len(lispvm.context.TxUnit.Authors))
	return lisp.Token{Kind:lisp.Int, Text:sigCount}, nil
}

//getPK	returns public key
func (lispvm *LispVM) getPK(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(authorIndexError)
	}

	pk := string(lispvm.context.TxUnit.Authors[x.Text.(int64)].Definition)
	return lisp.Token{Kind:lisp.String, Text:pk}, nil
}

//getPKByAddr returns the public key based on the address
func (lispvm *LispVM) getPKByAddr(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}

	if x.Kind != lisp.String {
		return lisp.None, errors.New("get pk by address is not string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("get pk by address is empty")
	}

	for _, author := range lispvm.context.TxUnit.Authors {

		if bytes.Equal(author.Address, []byte(x.Text.(string))){
			definition:=string(author.Definition)
			return lisp.Token{Kind:lisp.String, Text:definition}, nil
		}
	}

	return  lisp.None, errors.New("getPKByAddr failed")

}

//getSig returns signature based on public key
func (lispvm *LispVM) getSig(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.String {
		return lisp.None, errors.New("getSig is not string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("getSig is empty")
	}

	for _, author := range lispvm.context.TxUnit.Authors {
		if bytes.Equal(author.Definition, []byte(x.Text.(string))) {
			authentifiers:=string(author.Authentifiers)
			return lisp.Token{Kind:lisp.String, Text:authentifiers}, nil
		}
	}

	return lisp.None, errors.New("not find the PK in the unit")
}

//getCurUnitHash returns the hash of the Unit
func (lispvm *LispVM) getCurUnitHash(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	CurUnitHash := string(lispvm.context.TxUnit.Hash())
	return lisp.Token{Kind:lisp.String, Text:CurUnitHash}, nil
}

//GetCurUnitHashToSig returns the hash used for signing according to Unit
func (lispvm *LispVM) getCurUnitHashToSign(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	hashToSign := string(lispvm.context.TxUnit.GetHashToSign())
	return lisp.Token{Kind:lisp.String, Text:hashToSign}, nil
}

//getCurMsgHash returns the hash of the message
func (lispvm *LispVM) getCurMsgHash(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	CurMsgHash := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].PayloadHash())
	return lisp.Token{Kind:lisp.String, Text:CurMsgHash}, nil
}

//globalParamCount returns the number of public parameters
func (lispvm *LispVM) globalParamCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	globalParamCount:= int64(len(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).GlobalParamKey))
	return  lisp.Token{Kind:lisp.Int,Text:globalParamCount},nil
}

//getGlobalParam returns public parameters
func (lispvm *LispVM) getGlobalParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New("getGlobalParam is not int")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("getGlobalParam is empty")
	}

	name := x.Text.(string)
	globalParam := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).GetParam(name))
	return lisp.Token{Kind:lisp.String, Text:globalParam}, nil
}
//getGlobalParam returns to the public parameter list
func (lispvm *LispVM) getGlobalParamList(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New("getGlobalParam is not int")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("getGlobalParam is empty")
	}

	name := x.Text.(string)
	globalParam := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).GetParam(name))
	var paramList []lisp.Token
	paramlistLen:=len(globalParam)
	startIdx:=0
	for startIdx< paramlistLen {
		paramLen:=int(globalParam[startIdx])
		if paramLen >=paramlistLen-startIdx {
			return 	lisp.None,errors.New(paramLenError)
		}
		paramV:=string(globalParam[startIdx+1:startIdx+paramLen+1])
		params:=lisp.Token{Kind:lisp.String,Text:paramV}
		paramList=append(paramList,params)

		startIdx=startIdx+paramLen+1
	}
	return lisp.Token{Kind:lisp.List, Text:paramList}, nil
}

//inputCount returns the number of Inputs
func (lispvm *LispVM) inputCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	inputcount:= int64(len(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs))
	return  lisp.Token{Kind:lisp.Int,Text:inputcount},nil
}

//getInputUnit returns unit hash according to Input
func (lispvm *LispVM) getInputUnit(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.Int {
		return lisp.None, errors.New("getInputUnit is not int")
	}

	index := x.Text.(int64)
	getInputUnit := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[index].SourceUnit)
	return lisp.Token{Kind:lisp.String, Text:getInputUnit}, nil
}

//getInputMsg returns Msg index according to Input
func (lispvm *LispVM) getInputMsg(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(inputParamIntError)
	}

	msgIndex := int64(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)].SourceMessage)

	return lisp.Token{Kind:lisp.Int, Text:msgIndex}, nil
}

//getInputMsg returns the parameter value according to Input
func (lispvm *LispVM) getInputParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(inputParamIntError)
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}

	if y.Kind != lisp.String {
		return lisp.None, errors.New("input name is not int")
	}
	if len(y.Text.(string))==0{
		return lisp.None, errors.New("input name is empty")
	}

	inputParam := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)].GetParam(y.Text.(string)))

	return lisp.Token{Kind:lisp.String, Text:inputParam}, nil
}
//getInputMsg returns the parameter list according to Input
func (lispvm *LispVM) getInputParamList(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(inputParamIntError)
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}

	if y.Kind != lisp.String {
		return lisp.None, errors.New("input name is not int")
	}
	if len(y.Text.(string))==0{
		return lisp.None, errors.New("input name is empty")
	}

	inputParam := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)].GetParam(y.Text.(string)))
	var paramList []lisp.Token
	paramlistLen:=len(inputParam)
	startIdx:=0
	for startIdx< paramlistLen {
		paramLen:=int(inputParam[startIdx])
		if paramLen >=paramlistLen-startIdx {
			return 	lisp.None,errors.New(paramLenError)
		}
		paramV:=string(inputParam[startIdx+1:startIdx+paramLen+1])
		params:=lisp.Token{Kind:lisp.String,Text:paramV}
		paramList=append(paramList,params)

		startIdx=startIdx+paramLen+1
	}
	return lisp.Token{Kind:lisp.List, Text:paramList}, nil

}

//getInputPreOut returns the Output index according to Input
func (lispvm *LispVM) getInputPreOut(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}

	outputIndex := int64(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)].SourceOutput)

	return lisp.Token{Kind:lisp.Int, Text:outputIndex}, nil
}

//getPrevOutAmount returns the PrevOut amount based on Input
func (lispvm *LispVM) getPrevOutAmount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}

	msg := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage)
	preOutAmount := int64(lispvm.context.FetchPrevOut(msg.Inputs[x.Text.(int64)]).Amount)

	return lisp.Token{Kind:lisp.Int, Text:preOutAmount}, nil
}

//getPrevOutParam returns the PrevOut parameter
func (lispvm *LispVM) getPrevOutParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}
	if y.Kind != lisp.String {
		return lisp.None, errors.New(contentStringError)
	}

	if len(y.Text.(string)) == 0 {
		return lisp.None, errors.New("pre index is empty")
	}

	//lispvm.context.FetchPrevOut
	input := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)]

	preOutParam := string(lispvm.context.FetchPrevOut(input).GetParam(y.Text.(string)))
	return lisp.Token{Kind:lisp.String, Text:preOutParam}, nil
}

//getPrevOutParamList returns to the PrevOut parameter list
func (lispvm *LispVM) getPrevOutParamList(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}
	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}
	if y.Kind != lisp.String {
		return lisp.None, errors.New(contentStringError)
	}
	if len(y.Text.(string)) == 0 {
		return lisp.None, errors.New("pre index is empty")
	}
	//lispvm.context.FetchPrevOut
	input := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)]
	preOutParam := string(lispvm.context.FetchPrevOut(input).GetParam(y.Text.(string)))
	var preOutParamList []lisp.Token
	paramlistLen:=len(preOutParam)
	startIdx:=0
	for startIdx< paramlistLen {
		paramLen:=int(preOutParam[startIdx])
		if paramLen >=paramlistLen-startIdx {
			return 	lisp.None,errors.New(paramLenError)
		}
		paramV:=string(preOutParam[startIdx+1:startIdx+paramLen+1])
		params:=lisp.Token{Kind:lisp.String,Text:paramV}
		preOutParamList=append(preOutParamList,params)
		startIdx=startIdx+paramLen+1
	}
	return lisp.Token{Kind:lisp.List, Text:preOutParamList}, nil
}

//getPreOutExtends returns to PrevOut Additional Information
func (lispvm *LispVM) getPreOutExtends(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	//Get parameter values
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}

	input := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)]
	outputExtends := string(lispvm.context.FetchPrevOut(input).Extends)

	return lisp.Token{Kind:lisp.String, Text:outputExtends}, nil
}

//getOutputAmount returns to the Output amount
func (lispvm *LispVM) getOutputAmount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	//Get parameter values
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}

	amount := int64(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Outputs[x.Text.(int64)].Amount)

	return lisp.Token{Kind:lisp.Int, Text:amount}, nil
}

//getOutputParam  returns Output parameter
func (lispvm *LispVM) getOutputParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}

	if y.Kind != lisp.String {
		return lisp.None, errors.New(contentStringError)
	}

	if len(y.Text.(string)) == 0 {
		return lisp.None, errors.New("pre index is nil")
	}

	outputParam := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Outputs[x.Text.(int64)].GetParam(y.Text.(string)))

	return lisp.Token{Kind:lisp.String, Text:outputParam}, nil
}
//getOutputParam returns to the Output parameter list
func (lispvm *LispVM) getOutputParamList(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}

	if y.Kind != lisp.String {
		return lisp.None, errors.New(contentStringError)
	}

	if len(y.Text.(string)) == 0 {
		return lisp.None, errors.New("pre index is nil")
	}

	outputParam := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Outputs[x.Text.(int64)].GetParam(y.Text.(string)))

	var paramList []lisp.Token
	paramlistLen:=len(outputParam)
	startIdx:=0
	for startIdx< paramlistLen {
		paramLen:=int(outputParam[startIdx])
		if paramLen >=paramlistLen-startIdx {
			return 	lisp.None,errors.New(paramLenError)
		}
		paramV:=string(outputParam[startIdx+1:startIdx+paramLen+1])
		params:=lisp.Token{Kind:lisp.String,Text:paramV}
		paramList=append(paramList,params)

		startIdx=startIdx+paramLen+1
	}
	return lisp.Token{Kind:lisp.List, Text:paramList}, nil
}

//getOutputExtends returns to Output Additional Information
func (lispvm *LispVM) getOutputExtends(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(contentIntError)
	}

	outputExtend := string(lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Outputs[x.Text.(int64)].Extends)

	return lisp.Token{Kind:lisp.String, Text:outputExtend}, nil
}

//getCurMCI returns to current MCI
func (lispvm *LispVM) getCurrentMCI(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	mci:=int64(lispvm.context.MCI)
	return lisp.Token{Kind:lisp.Int,Text:mci},nil

}

//hasPrevOutParam is to determine if the PrevOut parameter already exists
func (lispvm *LispVM) hasPrevOutParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.Int {
		return lisp.None, errors.New("param is not int")
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}
	if y.Kind != lisp.String {
		return lisp.None, errors.New("param is not string")
	}


	if len(y.Text.(string)) == 0{
		return lisp.None,errors.New("param is empty")

	}

	//lispvm.context.FetchPrevOut
	input := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)]

	preoutparam := lispvm.context.FetchPrevOut(input).FindParam(y.Text.(string))
	if preoutparam >= 0 {
		return lisp.True, nil
	}

	return lisp.False, nil
}

//hasPKByAddr is to determine whether the Author public key exists
func (lispvm *LispVM) hasPKByAddr(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}

	if x.Kind != lisp.String {
		return lisp.None, errors.New("pk by address is not string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("get pk by address empty")
	}
	for _, author := range lispvm.context.TxUnit.Authors {

		if bytes.Equal(author.Address, []byte(x.Text.(string))) {
			return lisp.True, nil
		}
	}
	return lisp.False, nil
}

//getAuthorSig returns the Author signature
func (lispvm *LispVM) getAuthorSig(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(authorIndexError)
	}
	sig:=string(lispvm.context.TxUnit.Authors[x.Text.(int64)].Authentifiers)
	return lisp.Token{Kind:lisp.String,Text:sig}, nil
}

//getAuthorAddr returns the Author address
func (lispvm *LispVM) getAuthorAddr(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(authorIndexError)
	}
	addr:=string(lispvm.context.TxUnit.Authors[x.Text.(int64)].Address)
	return lisp.Token{Kind:lisp.String,Text:addr}, nil
}

//hasInputParam is to determine whether the current inputParam exists
func (lispvm *LispVM) hasInputParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {

	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.Int {
		return lisp.None, errors.New("param is not int")
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}
	if y.Kind != lisp.String {
		return lisp.None, errors.New("param is not string")
	}
	if len(y.Text.(string)) == 0{
		return lisp.None,errors.New("param is empty")
	}
	input := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).Inputs[x.Text.(int64)]
	inputparam := input.FindParam(y.Text.(string))
	if inputparam >= 0 {
		return lisp.True, nil
	}

	return lisp.False, nil

}

//hasCurPrevOutParam is to determine whether the paramValue in the current PrevOut exists
func (lispvm *LispVM) hasCurPrevOutParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New(inputParamStringError)
	}
	if len(x.Text.(string))==0 {
		return lisp.None, errors.New(inputParamError)
	}
	for _, v := range lispvm.context.PrevOut.OutputParamsKey{
		if v==x.Text.(string){
			return lisp.True, nil
		}
	}
	return lisp.False, nil
}

//getCurPrevOutParam returns the paramValue in the current PrevOut
func (lispvm *LispVM) getCurPrevOutParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New(inputParamStringError)
	}
	if len(x.Text.(string))==0 {
		return lisp.None, errors.New(inputParamError)
	}
	for i, v := range lispvm.context.PrevOut.OutputParamsKey{
		if v==x.Text.(string){
			paramsvalue:=string(lispvm.context.PrevOut.OutputParamsValue[i])
			return lisp.Token{Kind:lisp.String,Text:paramsvalue}, nil
		}
	}
	return lisp.None,errors.New("getCurPrevOutParam failed ")
}
//getCurPrevOutParamList returns the list of paramValues in the current PrevOut
func (lispvm *LispVM) getCurPrevOutParamList(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New(inputParamStringError)
	}
	if len(x.Text.(string))==0 {
		return lisp.None, errors.New(inputParamError)
	}
	for i, v := range lispvm.context.PrevOut.OutputParamsKey{
		if v==x.Text.(string){
			paramsvalue:=string(lispvm.context.PrevOut.OutputParamsValue[i])
			var paramList []lisp.Token
			paramlistLen:=len(paramsvalue)
			startIdx:=0
			for startIdx< paramlistLen {
				paramLen:=int(paramsvalue[startIdx])
				if paramLen >=paramlistLen-startIdx {
					return 	lisp.None,errors.New(paramLenError)
				}
				paramV:=string(paramsvalue[startIdx+1:startIdx+paramLen+1])
				params:=lisp.Token{Kind:lisp.String,Text:paramV}
				paramList=append(paramList,params)

				startIdx=startIdx+paramLen+1
			}
			return lisp.Token{Kind:lisp.List, Text:paramList}, nil
		}
	}
	return lisp.None,errors.New("getCurPrevOutParam failed ")
}
//getCurPrevOutAmount returns the amount in current PrevOut
func (lispvm *LispVM) getCurPrevOutAmount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	amount:=int64(lispvm.context.PrevOut.UtxoHeader.Amount)
	return lisp.Token{Kind:lisp.Int,Text:amount},nil
}

//getCurPrevOutExtends returns to Extends in current PrevOnt
func (lispvm *LispVM) getCurPrevOutExtends(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	extends := string(lispvm.context.PrevOut.Extends)
	return lisp.Token{Kind:lisp.String, Text:extends}, nil
}

//getCurInputParam returns the ParamValue in the current Input
func (lispvm *LispVM) getCurInputParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New(inputParamStringError)
	}
	if len(x.Text.(string))==0 {
		return lisp.None, errors.New(inputParamError)
	}
	for i, v := range lispvm.context.Input.InputParamsKey{
		if v==x.Text.(string){
			paramsvale:=string(lispvm.context.Input.InputParamsValue[i])
			return lisp.Token{Kind:lisp.String,Text:paramsvale}, nil
		}
	}
	return lisp.None,errors.New("getCurInputParam failed")
}

//hasCurInputParam is to determine whether the current Input ParamValue exists
func (lispvm *LispVM) hasCurInputParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New(inputParamStringError)
	}
	if len(x.Text.(string))==0 {
		return lisp.None, errors.New(inputParamError)
	}
	for _, v := range lispvm.context.Input.InputParamsKey{
		if v==x.Text.(string){
			return lisp.True, nil
		}
	}
	return lisp.False, nil
}

//getCurInputParamsCount returns the number of Params of the current Input
func (lispvm *LispVM) getCurInputParamsCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	paraCount := int64(len(lispvm.context.Input.InputParamsKey))
	return lisp.Token{Kind:lisp.Int, Text:paraCount}, nil
}

//getCurInputUnit returns the sourceUnit of the current Input
func (lispvm *LispVM) getCurInputUnit(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	sUnit := string(lispvm.context.Input.SourceUnit)
	return lisp.Token{Kind:lisp.String, Text:sUnit}, nil
}

//getCurInputMsg returns the Msg number of the current Input
func (lispvm *LispVM) getCurInputMsg(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	sMsg := int64(lispvm.context.Input.SourceMessage)
	return lisp.Token{Kind:lisp.Int, Text:sMsg}, nil
}

//getCurInputOutput returns the current Input's Output number
func (lispvm *LispVM) getCurInputOutput(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	sOutput := int64(lispvm.context.Input.SourceOutput)
	return lisp.Token{Kind:lisp.Int, Text:sOutput}, nil
}
