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
package vm

import (
	"github.com/SHDMT/gravity/platform/consensus/structure"
)

const (
	//VMModeContract is VM contract execute mode
	VMModeContract = 0

	//VMModeRestrict is asset restrict mode
	VMModeRestrict = 1
)

//Context is context of VM
type Context struct {
	//TxUnit is unit information that is used when execute contract
	TxUnit structure.Unit

	//TxMsgIndex is message index that is used when execute contract
	TxMsgIndex uint32

	//AssetUnit		structure.Unit
	//AssetMsgIndex	uint32

	//AssetMsg is issue message that is used when execute contract
	AssetMsg structure.IssueMessage

	//FetchPrevOut is a callback function to fetch previous output associated with current input
	FetchPrevOut func(input *structure.ContractInput) *structure.ContractOutput

	//MCI is Main Chain Index of DAG that is used when execute contract
	MCI uint64

	//ContractDef is contract definition information
	ContractDef *structure.ContractDef

	//For Restrict
	//Input is contract input that is used when execute contract
	Input *structure.ContractInput
	// PrevOut is previous output unit info associate with current contract input
	PrevOut *structure.TxUtxo
}

//Config is configuration of VM
type Config struct {
	//Debug marks if VM is working in debug mode
	Debug bool
	//MaxStep is the max steps of VM
	MaxStep uint32
	//Mode marks execute mode of VM
	Mode byte
}

//NewContext creates a new smart contract runtime environment object
func NewContext(unit structure.Unit, msgIndex uint32,
	assetMsg structure.IssueMessage,
	mci uint64, fetchOutput func(input *structure.ContractInput) *structure.ContractOutput) *Context {
	return &Context{
		TxUnit:       unit,
		TxMsgIndex:   msgIndex,
		AssetMsg:     assetMsg,
		MCI:          mci,
		FetchPrevOut: fetchOutput,
	}
}

//NewRestrictContext creates a new restrict mode context object
func NewRestrictContext(unit structure.Unit, msgIndex uint32,
	assetMsg structure.IssueMessage,
	mci uint64, input *structure.ContractInput, prevOut *structure.TxUtxo) *Context {
	return &Context{
		TxUnit:     unit,
		TxMsgIndex: msgIndex,
		AssetMsg:   assetMsg,
		MCI:        mci,
		Input:      input,
		PrevOut:    prevOut,
	}
}

//NewConfig creates a new VM configuration object
func NewConfig(debug bool, maxstep uint32) *Config {
	return &Config{
		Debug:   debug,
		MaxStep: maxstep,
	}
}

//DefaultConfig creates a new default configuration object
func DefaultConfig() *Config {
	return NewConfig(false, 100)
}
