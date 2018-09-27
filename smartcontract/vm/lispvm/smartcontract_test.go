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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	rand2 "math/rand"
	"os"
	"testing"

	"github.com/SHDMT/gravity/infrastructure/crypto/asymmetric/ec/secp256k1"
	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/infrastructure/database"
	_ "github.com/SHDMT/gravity/infrastructure/database/badgerdb"
	"github.com/SHDMT/gravity/infrastructure/log"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/smartcontract"
	"github.com/SHDMT/gravity/platform/smartcontract/vm"
)

const (
	dbName = "badgerDB"
)

var (
	publisherAddr hash.HashType
	assetHash     hash.HashType
	cons          []*structure.Contract
	defs          []*structure.ContractDef
	rests         []*structure.Contract

	issueMessage *structure.IssueMessage
	utxos        []*structure.TxUtxo
	userEntries  []*AddressEntry

	defParams     []map[string][]byte
	amountParams  [][]uint64
	inputParams   []map[string][]byte
	outputParams  []map[string][]byte
	prevOutParams []map[string][]byte
	globalParams  []map[string][]byte
)

type AddressEntry struct {
	address hash.HashType
	pk      []byte
	privkey []byte
}

func TestSystemContract(t *testing.T) {
	_, err := FakeUnitWithSig()
	if err != nil {
		t.Fatal(err)
	}
	db, err := setupEnv()
	if err != nil {
		t.Fatal(err)
	}

	userEntries, err = FakeUserKeys(10)
	if err != nil {
		t.Fatal(err)
	}

	publisherAddr = userEntries[0].address
	setupParams()

	err = setupDB(db)
	if err != nil {
		t.Fatal(err)
	}

	for index := 0; index < len(cons); index++ {
		consAddr := cons[index].CalcAddress()

		in1, utxo1 := FakeContractInput(amountParams[index][0], userEntries[1].address, inputParams[index*2], prevOutParams[index*2], nil)
		in2, utxo2 := FakeContractInput(amountParams[index][1], userEntries[2].address, inputParams[index*2+1], prevOutParams[index*2+1], nil)
		out1 := FakeContractOutput(amountParams[index][2], userEntries[3].address, outputParams[0], nil)
		ins := []*structure.ContractInput{in1, in2}
		outs := []*structure.ContractOutput{out1}
		unit, err := FakeUnitWithInOut(index, consAddr, ins, outs, userEntries)
		if err != nil {
			t.Error(err.Error())
		}

		utxos = append(utxos, utxo1)
		utxos = append(utxos, utxo2)

		context := vm.Context{
			TxUnit:       *unit,
			TxMsgIndex:   0,
			AssetMsg:     *issueMessage,
			FetchPrevOut: fetchPrevOut,

			MCI:         20,
			ContractDef: defs[index],
		}
		lv := NewLispVM(context, vm.Config{Mode: vm.VMModeContract})
		log.Infof("to run contract[%d]", index)
		//lv.Exec(*cons[index])
		result := lv.Exec(cons[index])
		utxos = utxos[:0]
		if !result {
			t.Error("Exec lisp vm error.")
			break
		}
	}
}

func setupParams() {
	cosignAddrs := make([]byte, 0)
	l := len(publisherAddr)
	cosignAddrs = append(cosignAddrs, byte(l))
	cosignAddrs = append(cosignAddrs, publisherAddr...)

	memberAddrs := make([]byte, 0)
	l = len(userEntries[2].address)
	memberAddrs = append(memberAddrs, byte(l))
	memberAddrs = append(memberAddrs, userEntries[2].address...)
	l = len(userEntries[5].address)
	memberAddrs = append(memberAddrs, byte(l))
	memberAddrs = append(memberAddrs, userEntries[5].address...)

	defParams = []map[string][]byte{
		{
			"cosignerAddrs": cosignAddrs,
			"threshold":     []byte{'1'},
		},
		{
			"cosignerAddrs": cosignAddrs,
			"threshold":     []byte{'1'},
		},
		{
			"cosignerAddrs": cosignAddrs,
			"threshold":     []byte{'1'},
		}, {},
		{
			"memberAddrs": memberAddrs,
		}, {},
	}
	amountParams = [][]uint64{
		{20, 10, 1500},
		{1000, 1000, 2000},
		{1000, 1000, 500},
		{1000, 1000, 2000},
		{1000, 1000, 2000},
		{1000, 1000, 2000},
	}
	inputParams = []map[string][]byte{
		{}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {},
	}
	outputParams = []map[string][]byte{
		{
			"addr": userEntries[4].address,
		},
		{},
		{},
		{},
		{},
		{
			"mulitSign": cosignAddrs,
			"threshold": []byte{'1'},
			"addrs":     cosignAddrs,
		},
		{
			"addr":                        userEntries[5].address,
			"restrictmci.mci":             []byte("40"),
			"restrictmultisign.threshold": []byte("2"),
			"restrictmultisign.addrs":     memberAddrs,
		},
	}
	prevOutParams = []map[string][]byte{
		{
			"addr": userEntries[2].address,
		},
		{
			//"addr": userEntries[5].address,
		}, //1
		{
			"addr": userEntries[2].address,
		},
		{
			//"addr": userEntries[5].address,
		}, //2
		{
			"addr": userEntries[2].address,
		},
		{
			//"addr": userEntries[5].address,
		}, //3
		{
			"addr": userEntries[2].address,
		},
		{
			//"addr": userEntries[5].address,
		}, //4
		{
			"addr": userEntries[2].address,
		},
		{
			"multiSign": []byte{'1'},
			"addrs":     memberAddrs,
			"threshold": []byte{'2'},
		}, //5
		{
			"addr": userEntries[2].address,
		},
		{
			"addr": userEntries[5].address,
		}, //6
		{
			"addr":                        userEntries[5].address,
			"restrictmci.mci":             []byte("40"),
			"restrictmultisign.threshold": []byte("2"),
			"restrictmultisign.addrs":     memberAddrs,
		},
	}

	globalParams = []map[string][]byte{
		{}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {},
	}
}

func FakeUserKeys(count int) ([]*AddressEntry, error) {
	addrEntries := make([]*AddressEntry, count)
	for i := 0; i < count; i++ {
		entry, err := FakeUserKey()
		if err != nil {
			return nil, err
		}
		addrEntries[i] = entry
	}
	return addrEntries, nil
}
func FakeUserKey() (*AddressEntry, error) {
	worker := secp256k1.NewCipherSuite()
	privateKey, err := worker.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	pubkey := privateKey.Public()

	pkBody, err := pubkey.MarshalP()
	if err != nil {
		return nil, err
	}
	addr := hash.Sum256(pkBody)

	var pkBytes = []byte{0}
	pkBytes = append(pkBytes, pkBody...)
	privKeyBytes, err := privateKey.MarshalP()
	if err != nil {
		return nil, err
	}

	return &AddressEntry{
		address: addr,
		pk:      pkBytes,
		privkey: privKeyBytes,
	}, nil
}

func FakeRandomHash() hash.HashType {
	hash := make([]byte, 32)

	for i := 0; i < 32; i++ {
		hash[i] = byte(rand2.Intn(255))
	}
	return hash
}

func TestSmartContractVerify(t *testing.T) {
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})

	unit, err := FakeUnitWithSig()
	if err != nil {
		t.Error("fake unit with err:", err)
	}
	lvm.context = vm.Context{
		TxUnit: *unit,
	}
	//load a contract(lsp) file
	//tk1, err := lvm.vm.Load("F:/myGo/src/github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/smartcontract.lsp")
	tk1, err := lvm.vm.Load("./smartcontract.lsp")
	if err != nil {
		t.Error("load error:\n", err)
	}
	t.Logf("load contract result: %v \n", tk1.Text)
}

func FakeUnitWithSig() (*structure.Unit, error) {

	contractInput := make([]*structure.ContractInput, 1)
	contractInput[0] = &structure.ContractInput{
		SourceOutput: uint32(1),
	}
	msg := structure.InvokeMessage{
		Header: new(structure.MessageHeader),
		Inputs: contractInput,
	}
	invMsg := make([]structure.Message, 1)
	invMsg[0] = &msg

	cipherSuite := secp256k1.NewCipherSuite()
	privKByte, err := hex.DecodeString("17457dc2ccdd91431ba4db72d17fcad5c2b2ae3a58f48488680b392e28d39924")
	if err != nil {
		return nil, err
	}
	privK, err := cipherSuite.UnmarshalPrivateKey(privKByte)
	if err != nil {
		return nil, err
	}

	pubK := privK.Public()
	pubKBody, err := pubK.MarshalP()
	if err != nil {
		return nil, err
	}
	var pubKByte = []byte{0}
	pubKByte = append(pubKByte, pubKBody...)
	addr := hash.Sum256(pubKBody)
	auths := make([]*structure.Author, 1)
	auths[0] = &structure.Author{
		Address:    addr,
		Definition: pubKByte,
	}

	unit := structure.Unit{
		Messages: invMsg,
		Authors:  auths,
	}

	data := unit.GetHashToSign()
	sig, err := privK.Sign(data)
	if err != nil {
		return nil, err
	}
	unit.Authors[0].Authentifiers = sig
	return &unit, nil
}

func TestGetPrivKey(t *testing.T) {
	cipherSuite := secp256k1.NewCipherSuite()
	privK, err := cipherSuite.GenerateKey(rand.Reader)
	if err != nil {
		t.Error(err)
	}
	data, err := cipherSuite.MarshalPrivateKey(privK)
	if err != nil {
		t.Error(err)
	}
	t.Logf("privK:%s\n", hex.EncodeToString(data))

	pubK := privK.Public()
	pubKByte, err := pubK.MarshalP()
	if err != nil {
		t.Error(err)
	}
	//addr := hash.Sum256(pubKByte)
	t.Logf("addr:%s\n", hex.EncodeToString(pubKByte))
}

func TestVer(t *testing.T) {
	//data := []byte("test")
	cipherSuite := secp256k1.NewCipherSuite()
	privKByte, err := hex.DecodeString("17457dc2ccdd91431ba4db72d17fcad5c2b2ae3a58f48488680b392e28d39924")
	if err != nil {
		t.Error(err)
	}
	privK, err := cipherSuite.UnmarshalPrivateKey(privKByte)
	if err != nil {
		t.Error(err)
	}
	pubKByte2, err := hex.DecodeString("0421af2c7f64c10a34fbf4891ac5862a71f0c5805e0e4b50ce03244943a7859b7dd26e5a9ffd9ac0bccb5aadc81552b5961cc7f361a244df703d248aa5c13fb013")
	if err != nil {
		t.Error(err)
	}
	pubK := privK.Public()
	pubKByte, err := pubK.MarshalP()
	if !bytes.Equal(pubKByte2, pubKByte) {
		t.Error("unmatch privKey and pubKey")
	}
}

func FakeRestrict(name string, code []byte) *structure.Contract {
	contract := structure.NewContract()

	contract.Name = name
	contract.ScriptCode = LispVMVersion
	contract.IsRestrict = true
	contract.Code = code
	contract.CalcAddress()
	return contract
}

func FakeContract(name string, code []byte) *structure.Contract {
	contract := structure.NewContract()

	contract.Name = name
	contract.ScriptCode = LispVMVersion
	contract.IsRestrict = false
	contract.Code = code
	contract.CalcAddress()

	return contract
}

func FakeContractDef(contract *structure.Contract, params map[string][]byte) *structure.ContractDef {

	contractDef := structure.NewContractDef()
	for k, v := range params {
		contractDef.AddParam(k, v)
	}
	contractDef.Address = contract.CalcAddress()
	return contractDef
}

func FakeIssueMessage(publisher hash.HashType, contractDefs []*structure.ContractDef) *structure.IssueMessage {
	header := new(structure.MessageHeader)
	header.App = structure.IssueMessageType
	header.Version = structure.Version

	locationAddrs := make([]hash.HashType, len(userEntries))
	locationAmounts := make([]int64, len(userEntries))
	for i, v := range userEntries {
		locationAddrs[i] = v.address
		locationAmounts[i] = 100000000
	}

	msg := &structure.IssueMessage{
		Header:             header,
		Name:               "test asset",
		Cap:                1000000000,
		FixedDenominations: false,
		Contracts:          contractDefs,
		AllocationAddr:     locationAddrs,
		AllocationAmount:   locationAmounts,
		PublisherAddress:   publisher,
		Note:               []byte("this is a test asset"),
	}

	msg.Header.PayloadHash = msg.CalcPayloadHash()

	return msg
}

func FakeContractInput(amount uint64, addr hash.HashType, inParams map[string][]byte, outParams map[string][]byte, restricts []hash.HashType) (*structure.ContractInput, *structure.TxUtxo) {
	input := structure.NewContractInput()
	for k, v := range inParams {
		input.AddParam(k, v)
	}

	utxo := new(structure.TxUtxo)
	for k, v := range outParams {
		utxo.OutputParamsKey = append(utxo.OutputParamsKey, k)
		utxo.OutputParamsValue = append(utxo.OutputParamsValue, v)
	}

	unitHash := FakeRandomHash()

	input.SourceUnit = unitHash
	input.SourceMessage = 1
	input.SourceOutput = 2

	utxo.Unit = unitHash
	utxo.Message = 1
	utxo.Output = 2
	utxo.Asset = assetHash
	utxo.Restricts = restricts
	utxo.UtxoHeader.Amount = amount
	utxo.UtxoHeader.Address = addr

	return input, utxo
}

func FakeContractOutput(amount uint64, addr hash.HashType, outParams map[string][]byte, restricts []hash.HashType) *structure.ContractOutput {
	output := structure.NewContractOutput()
	output.Amount = amount
	for k, v := range outParams {
		output.AddParam(k, v)
	}
	output.Restricts = restricts
	return output
}

func FakeUnitWithInOut(index int, contractAddr hash.HashType, inputs []*structure.ContractInput, outputs []*structure.ContractOutput, addrEntries []*AddressEntry) (*structure.Unit, error) {
	header := new(structure.MessageHeader)
	header.App = structure.InvokeMessageType
	header.Version = structure.Version
	msg := structure.InvokeMessage{
		Header:       header,
		ContractAddr: contractAddr,
		Asset:        assetHash,
		Inputs:       inputs,
		Outputs:      outputs,
	}

	for k, v := range globalParams[index] {
		msg.AddParam(k, v)
	}

	msg.Header.PayloadHash = msg.CalcPayloadHash()
	invMsg := make([]structure.Message, 1)
	invMsg[0] = &msg

	auths := make([]*structure.Author, len(addrEntries))
	for i, addrEntry := range addrEntries {
		pubKByte := addrEntry.pk

		addr := addrEntry.address
		auths[i] = &structure.Author{
			Address:    addr,
			Definition: pubKByte,
		}
	}

	unit := structure.Unit{
		Messages: invMsg,
		Authors:  auths,
	}

	for i, addrEntry := range addrEntries {
		cipherSuite := secp256k1.NewCipherSuite()
		privKByte := addrEntry.privkey

		privK, err := cipherSuite.UnmarshalPrivateKey(privKByte)
		if err != nil {
			return nil, err
		}

		data := unit.GetHashToSign()
		sig, err := privK.Sign(data)
		if err != nil {
			return nil, err
		}
		unit.Authors[i].Authentifiers = sig
	}

	return &unit, nil
}
func createOrOpenDB(dbPath string) (database.Db, error) {
	var db database.Db
	var err error

	if _, err = os.Stat(dbPath); os.IsNotExist(err) {
		db, err = database.Create(dbName, dbPath, dbPath)
	} else {
		db, err = database.Open(dbName, dbPath, dbPath)
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setupEnv() (database.Db, error) {
	dbPath := "./~temp/"
	os.RemoveAll(dbPath)

	db, err := createOrOpenDB(dbPath)
	if err != nil {
		return nil, err
	}
	smartcontract.CreateSmartContractBucket(db)

	return db, nil
}
func FakeSystemContractDefs(contracts []*structure.Contract) ([]*structure.ContractDef, error) {
	contractDefs := make([]*structure.ContractDef, len(contracts))
	for i, contract := range contracts {
		contractDef := FakeContractDef(contract, defParams[i])
		contractDefs[i] = contractDef
	}

	return contractDefs, nil
}

func FakeSystemRestricts() ([]*structure.Contract, error) {
	files, err := ioutil.ReadDir("./lsp/restrict/")
	if err != nil {
		return nil, err
	}

	contracts := make([]*structure.Contract, len(files))
	for i, file := range files {
		buf, err := ioutil.ReadFile("./lsp/restrict/" + file.Name())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		contract := FakeContract(file.Name(), buf)
		contracts[i] = contract
		log.Infof("add restrict in file[%s] to restrict slice", file.Name())
	}
	return contracts, nil
}

func FakeSystemContracts() ([]*structure.Contract, error) {
	files, err := ioutil.ReadDir("./lsp/contract")
	if err != nil {
		return nil, err
	}

	contracts := make([]*structure.Contract, len(files))
	for i, file := range files {
		buf, err := ioutil.ReadFile("./lsp/contract/" + file.Name())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		contract := FakeContract(file.Name(), buf)
		contracts[i] = contract
		log.Infof("add contract in file[%s] to contracts slice", file.Name())
	}
	return contracts, nil
}

func setupDB(db database.Db) error {
	contracts, err := FakeSystemContracts()
	if err != nil {
		return err
	}
	restricts, err := FakeSystemRestricts()
	if err != nil {
		return err
	}

	tx := db.Begin(true)
	library := smartcontract.NewContractLibrary(tx, false)
	for _, restrict := range restricts {
		err = library.SaveContract(restrict, 0)
		if err != nil {
			return err
		}
		rests = append(rests, restrict)
	}

	for _, contract := range contracts {
		err = library.SaveContract(contract, 0)
		if err != nil {
			return err
		}
		cons = append(cons, contract)
	}
	contractDefs, _ := FakeSystemContractDefs(contracts)

	issueMessage = FakeIssueMessage(publisherAddr, contractDefs)

	assetHash = FakeRandomHash()
	err = library.SaveAsset(assetHash, issueMessage, 0)
	if err != nil {
		return err
	}

	for _, contractDef := range contractDefs {
		err = library.SaveAssetContractDef(assetHash, contractDef)
		if err != nil {
			return err
		}
		defs = append(defs, contractDef)
	}

	tx.Commit()
	return nil
}

func fetchPrevOut(input *structure.ContractInput) *structure.ContractOutput {
	contractOut := &structure.ContractOutput{}
	for _, utxo := range utxos {
		if utxo.Message == input.SourceMessage &&
			utxo.Output == input.SourceOutput &&
			bytes.Equal(utxo.Unit, input.SourceUnit) {
			contractOut = &structure.ContractOutput{
				Amount:  utxo.Amount(),
				Extends: utxo.Extends,

				OutputParamsKey:   utxo.OutputParamsKey,
				OutputParamsValue: utxo.OutputParamsValue,

				Restricts: utxo.Restricts,
			}
		}
	}

	return contractOut
}
