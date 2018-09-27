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
	"errors"

	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/infrastructure/log"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp"
)

const (
	inputSmartContractIndexError = "input smart contract index is not integer"
)

//cap returns total asset amount
func (lispvm *LispVM) cap(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	cap := lispvm.context.AssetMsg.Cap
	return lisp.Token{Kind: lisp.Int, Text: cap}, nil
}

//isDenominations returns if the asset is fixed denomination
func (lispvm *LispVM) isDenominations(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	isDenomination := lispvm.context.AssetMsg.FixedDenominations
	if isDenomination {
		return lisp.True, nil
	}
	return lisp.False, nil
}

//getDenominationCount returns count of fixed denomination
func (lispvm *LispVM) getDenominationCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	denominationCount := int64(len(lispvm.context.AssetMsg.Denominations))

	return lisp.Token{Kind: lisp.Int, Text: denominationCount}, nil

}

//getDenomination returns denomination value according to the index
func (lispvm *LispVM) getDenomination(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New("input asset index is not a Integer")
	}

	assetAmount := int64(lispvm.context.AssetMsg.Denominations[x.Text.(int64)])

	return lisp.Token{Kind: lisp.Int, Text: assetAmount}, nil
}

//getAssetContractCount returns asset contracts count
func (lispvm *LispVM) getAssetContractCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	contractCount := int64(len(lispvm.context.AssetMsg.Contracts))

	return lisp.Token{Kind: lisp.Int, Text: contractCount}, nil
}

//getAssetContract returns contract address(hash) according to the index
func (lispvm *LispVM) getAssetContract(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(inputSmartContractIndexError)
	}

	contractAddress := string(lispvm.context.AssetMsg.Contracts[x.Text.(int64)].Address)

	return lisp.Token{Kind: lisp.String, Text: contractAddress}, nil
}

//getAllocationsCount returns asset initial allocation address count
func (lispvm *LispVM) getAllocationsCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	count := int64(len(lispvm.context.AssetMsg.AllocationAmount))

	return lisp.Token{Kind: lisp.Int, Text: count}, nil
}

//getAllocationsAddr returns allocation address according to the index
func (lispvm *LispVM) getAllocationsAddr(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New("input AllocationAddr index is not a Integer")
	}

	address := string(lispvm.context.AssetMsg.AllocationAddr[x.Text.(int64)])
	return lisp.Token{Kind: lisp.String, Text: address}, nil
}

//getAllocationsAmount returns allocation amount according to the index
func (lispvm *LispVM) getAllocationsAmount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.String {
		return lisp.None, errors.New("get allocation amount: input param(allocation address) is not string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("the address is nil")
	}

	amount := int64(lispvm.context.AssetMsg.GetAllocation(hash.HashType(x.Text.(string))))
	return lisp.Token{Kind: lisp.Int, Text: amount}, nil
}

//getAssetExtends returns asset notes
func (lispvm *LispVM) getAssetExtends(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	msg := string(lispvm.context.AssetMsg.Note)

	return lisp.Token{Kind: lisp.String, Text: msg}, nil
}

//getPublisherAddr returns publisher address
func (lispvm *LispVM) getPublisherAddr(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	creatorAddress := string(lispvm.context.AssetMsg.PublisherAddress)
	return lisp.Token{Kind: lisp.String, Text: creatorAddress}, nil
}

//getPublishUnitMCI returns unit MCI when asset published
func (lispvm *LispVM) getPublishUnitMCI(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	unitMci := int64(lispvm.context.MCI)
	return lisp.Token{Kind: lisp.Int, Text: unitMci}, nil
}

//getContractParamCount returns contract parameter count according to the index
func (lispvm *LispVM) getContractParamCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New(inputSmartContractIndexError)
	}

	paramCount := int64(len(lispvm.context.AssetMsg.Contracts[x.Text.(int64)].ParamsKey))
	return lisp.Token{Kind: lisp.Int, Text: paramCount}, nil
}

//getCurContractDefParamCount returns contract definition count
func (lispvm *LispVM) getCurContractDefParamCount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}

	paramCount := int64(len(lispvm.context.ContractDef.ParamsKey))

	return lisp.Token{Kind: lisp.Int, Text: paramCount}, nil
}

//getCurContractDefParamName returns contract definition name according to the index
func (lispvm *LispVM) getCurContractDefParamName(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, lisp.ErrParaNum
	}

	if x.Kind != lisp.Int {
		return lisp.None, errors.New("get asset parameter name the index is not int")
	}

	paramName := string(lispvm.context.ContractDef.ParamsKey[x.Text.(int64)])
	return lisp.Token{Kind: lisp.String, Text: paramName}, nil
}

//getCurContractDefParam returns contract definition value according to the parameter name
func (lispvm *LispVM) getCurContractDefParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		log.Errorf("running getCurContractDefParam with err", err)
		return lisp.None, err
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New("input smart contract index is not a string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("the name is nil")
	}

	paramValue := string(lispvm.context.ContractDef.GetParam(x.Text.(string)))

	return lisp.Token{Kind: lisp.String, Text: paramValue}, nil
}

//getCurContractDefParamList returns contract definition value list according to the parameter name
func (lispvm *LispVM) getCurContractDefParamList(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New("input smart contract index is not a string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("the name is nil")
	}
	paramValues := lispvm.context.ContractDef.GetParam(x.Text.(string))
	var paramList []lisp.Token
	paramlistLen := len(paramValues)
	startIdx := 0
	for startIdx < paramlistLen {
		paramLen := int(paramValues[startIdx])
		if paramLen >= paramlistLen-startIdx {
			return lisp.None, errors.New("paramLen is too long")
		}
		paramV := string(paramValues[startIdx+1 : startIdx+paramLen+1])
		params := lisp.Token{Kind: lisp.String, Text: paramV}
		paramList = append(paramList, params)

		startIdx = startIdx + paramLen + 1
	}
	return lisp.Token{Kind: lisp.List, Text: paramList}, nil
}

//getContractParamName returns contract parameter name according to indexes
func (lispvm *LispVM) getContractParamName(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.Int {
		return lisp.None, errors.New(inputSmartContractIndexError)
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}
	if y.Kind != lisp.Int {
		return lisp.None, errors.New("input smart contract ParamsKey index is not a Integer")
	}

	param := string(lispvm.context.AssetMsg.Contracts[x.Text.(int64)].ParamsKey[y.Text.(int64)])
	return lisp.Token{Kind: lisp.String, Text: param}, nil
}

//getContractParamByIndex returns contract parameter name value according to indexes
func (lispvm *LispVM) getContractParamByIndex(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.Int {
		return lisp.None, errors.New(inputSmartContractIndexError)
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}
	if y.Kind != lisp.Int {
		return lisp.None, errors.New("input smart contract ParamsValue index is not a Integer")
	}

	param := string(lispvm.context.AssetMsg.Contracts[x.Text.(int64)].ParamsValue[y.Text.(int64)])
	return lisp.Token{Kind: lisp.String, Text: param}, nil
}

//getContractParam returns parameter values according to contract index and parameter name
func (lispvm *LispVM) getContractParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 2 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.Int {
		return lisp.None, errors.New(inputSmartContractIndexError)
	}

	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}

	if y.Kind != lisp.String {
		return lisp.None, errors.New("input smart contract ParamsKey is not a string")
	}
	if len(y.Text.(string)) == 0 {
		return lisp.None, errors.New("the ParamsKey is nil")
	}

	param := string(lispvm.context.AssetMsg.Contracts[x.Text.(int64)].GetParam(y.Text.(string)))
	return lisp.Token{Kind: lisp.String, Text: param}, nil
}

//isExistAtGlobalParam returns if key existed in Global Parameter keys
func (lispvm *LispVM) isExistAtGlobalParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New("input key index is not a string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("the key is nil")
	}
	globalparamkey := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage).GlobalParamKey
	for _, globalparam := range globalparamkey {
		if globalparam == x.Text.(string) {

			return lisp.True, nil
		}
	}
	return lisp.False, nil
}

// isExistAtOutputParam returns if key is existed in contract output parameter keys
func (lispvm *LispVM) isExistAtOutputParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.String {
		return lisp.None, errors.New("not exist ExistAtOutputParam")
	}

	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("isExistAtOutputParam the input param is not string")
	}

	invMsg := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage)
	for _, output := range invMsg.Outputs {
		for _, param := range output.OutputParamsKey {
			if x.Text.(string) == param {
				return lisp.True, nil
			}
		}
	}

	return lisp.False, nil
}

// isExistAtInputParam returns if key is existed in contract input parameters keys
func (lispvm *LispVM) isExistAtInputParam(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}

	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}

	if x.Kind != lisp.String {
		return lisp.None, errors.New("not exist ExistAtInputParam")
	}

	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("isExistAtInputParam the input param is not string")
	}

	invMsg := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage)
	for _, input := range invMsg.Inputs {
		for _, param := range input.InputParamsKey {
			if x.Text.(string) == param {
				return lisp.True, nil
			}
		}
	}
	return lisp.False, nil
}

//calcOutputAmount returns total outputs amount of invoke message
func (lispvm *LispVM) calcOutputAmount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	invMsg := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage)

	totalOutput := uint64(0)
	for _, output := range invMsg.Outputs {
		totalOutput = totalOutput + output.Amount
	}
	totalOutputInt64 := int64(totalOutput)
	return lisp.Token{Kind: lisp.Int, Text: totalOutputInt64}, nil
}

//calcInputAmount returns total inputs amount of invoke message
func (lispvm *LispVM) calcInputAmount(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	invMsg := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage)
	totalInput := uint64(0)
	for _, input := range invMsg.Inputs {
		totalInput = totalInput + lispvm.context.FetchPrevOut(input).Amount
	}
	totalinput := int64(totalInput)
	return lisp.Token{Kind: lisp.Int, Text: totalinput}, nil
}

//calcBalance returns difference between inputs amount and outputs amount in invoke message
func (lispvm *LispVM) calcBalance(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 0 {
		return lisp.None, lisp.ErrParaNum
	}
	invMsg := lispvm.context.TxUnit.Messages[lispvm.context.TxMsgIndex].(*structure.InvokeMessage)
	totalInput := uint64(0)
	for _, input := range invMsg.Inputs {
		totalInput = totalInput + lispvm.context.FetchPrevOut(input).Amount
	}

	totalOutput := uint64(0)
	for _, output := range invMsg.Outputs {
		totalOutput = totalOutput + output.Amount
	}
	balance := int64(totalInput - totalOutput)
	return lisp.Token{Kind: lisp.Int, Text: balance}, nil
}
