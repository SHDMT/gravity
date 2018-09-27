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
package smartcontract

import (
	"math/rand"
	"os"
	"reflect"
	"testing"

	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/infrastructure/database"
	_ "github.com/SHDMT/gravity/infrastructure/database/badgerdb"
	"github.com/SHDMT/gravity/platform/consensus/genesis"
	"github.com/SHDMT/gravity/platform/consensus/structure"
)

const (
	dbName = "badgerDB"
)

func FakeIssueMessage() *structure.IssueMessage {
	header := new(structure.MessageHeader)
	header.App = structure.IssueMessageType
	header.Version = structure.Version

	def := structure.NewContractDef()
	def.Address = genesis.SimpleVerifyContract.CalcAddress()

	msg := &structure.IssueMessage{
		Header:             header,
		Name:               "test asset",
		Cap:                1000000000,
		FixedDenominations: false,
		Contracts:          []*structure.ContractDef{def},
		AllocationAddr:     []hash.HashType{FakeRandomHash()},
		AllocationAmount:   []int64{1000000000},
		PublisherAddress:   FakeRandomHash(),
		Note:               []byte("this is a test asset"),
	}

	msg.Header.PayloadHash = msg.CalcPayloadHash()

	return msg
}
func FakeRandomHash() hash.HashType {
	hash := make([]byte, 32)

	for i := 0; i < 32; i++ {
		hash[i] = byte(rand.Intn(255))
	}
	return hash
}

func CreateContract0() *structure.Contract {
	contract := structure.NewContract()
	contract.Version = 1
	contract.Name = "Hello123"
	contract.ScriptCode = 1
	//contract.Address = FakeRandomHash()
	contract.Code = []byte("12345678901234567890")

	contract.AddParam("name1", "value1")
	contract.AddParam("name2", "value2")
	contract.AddParam("name3", "value3")

	return contract
}

func CreateContract1() *structure.Contract {
	contract := structure.NewContract()
	contract.Version = 1
	contract.Name = "123123"
	contract.ScriptCode = 1
	//contract.Address = FakeRandomHash()
	contract.Code = []byte("555555")

	contract.AddParam("name12", "value13")
	contract.AddParam("name22", "value23")
	contract.AddParam("name32", "value33")

	return contract
}

func CreateRandomContractDef(addr hash.HashType) *structure.ContractDef {
	contractDef := structure.NewContractDef()
	contractDef.Address = addr

	contractDef.AddParam("addr", FakeRandomHash())
	contractDef.AddParam("addr2", FakeRandomHash())
	contractDef.AddParam("addr3", FakeRandomHash())
	contractDef.AddParam("addr4", FakeRandomHash())
	return contractDef
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

func TestContractLibrary_SaveContract(t *testing.T) {

	db, err := createOrOpenDB("./testSaveContract")
	if err != nil {
		t.Fatal(err)
	}

	CreateSmartContractBucket(db)

	db.Update(func(tx database.Tx) error {
		library := NewContractLibrary(tx, false)
		contract0 := CreateContract0()
		contract1 := CreateContract1()

		err = library.SaveContract(contract1, 8)
		if err != nil {
			t.Error("Can't save contract, ", err)
		}

		if !library.HasContract(contract1.CalcAddress()) {
			t.Error("Can't find contract1 in db.")
		}

		if library.HasContract(contract0.CalcAddress()) {
			t.Error("Find contract0, Should can't find contract0 in db.")
		}

		contract2, mci, err := library.LoadContract(contract1.Address)
		if err != nil {
			t.Error("Can't load contract, ", err)
		}

		if mci != 8 {
			t.Fatal("MCI read error.")
		}
		contract2.CalcAddress()

		if !reflect.DeepEqual(contract1, contract2) {
			t.Error("Save or load error, contract1 is not equal to contract2.")
		}

		return nil
	})

}

func TestContractLibrary_SaveAssetContractDef(t *testing.T) {
	db, err := createOrOpenDB("./testSaveAssetContractDef")
	if err != nil {
		t.Fatal(err)
	}

	CreateSmartContractBucket(db)

	db.Update(func(tx database.Tx) error {
		library := NewContractLibrary(tx, false)

		contract0 := CreateContract0()
		contract1 := CreateContract1()

		asset := FakeRandomHash()
		assetFake := FakeRandomHash()

		contractDef0 := CreateRandomContractDef(contract0.CalcAddress())
		contractDef1 := CreateRandomContractDef(contract1.CalcAddress())

		err = library.SaveAssetContractDef(asset, contractDef0)
		if err != nil {
			t.Log("Can't save asset contractdef, ", err)
		}

		err = library.SaveContract(contract0, 12)
		if err != nil {
			t.Log("Can't save contract, ", err)
		}

		if !library.HasAssetContract(asset, contractDef0.Address) {
			t.Error("Can't find contractDef0 in db.")
		}

		if library.HasAssetContract(assetFake, contractDef0.Address) {
			t.Error("Find fakeAsset contractDef0, Should can't find contractDef0 in db.")
		}

		if library.HasAssetContract(asset, contractDef1.Address) {
			t.Error("Find contractDef01, Should can't find contractDef1 in db.")
		}

		contractDef2, err := library.LoadAssetContractDef(asset, contractDef0.Address)
		if err != nil {
			t.Error("Can't load contractdef, ", err)
		}
		if !reflect.DeepEqual(contractDef0, contractDef2) {
			t.Error("Save or load error, contractdef0 is not equal to contractdef2.")
		}

		contractDef3, contract3, err := library.LoadAssetContract(asset, contractDef0.Address)
		if err != nil {
			t.Error("Can't load contract, ", err)
		}
		if !reflect.DeepEqual(contractDef0, contractDef3) {
			t.Error("Save or load error, contractdef0 is not equal to contractdef3.")
		}
		contract3.CalcAddress()
		if !reflect.DeepEqual(contract0, contract3) {
			t.Error("Save or load error, contract0 is not equal to contract3.")
		}

		return nil
	})
}

func TestContractLibrary_SaveAsset(t *testing.T) {

	db, err := createOrOpenDB("./testSaveAsset")
	if err != nil {
		t.Fatal(err)
	}

	CreateSmartContractBucket(db)

	db.Update(func(tx database.Tx) error {
		library := NewContractLibrary(tx, false)
		asset := FakeIssueMessage()
		fakeUnit := FakeRandomHash()

		err = library.SaveAsset(fakeUnit, asset, 12)
		if err != nil {
			t.Error("Can't save contract, ", err)
		}

		if !library.HasAsset(fakeUnit) {
			t.Error("Can't find contract1 in db.")
		}

		dbAsset, mci, err := library.LoadAsset(fakeUnit)
		if err != nil {
			t.Error("Can't load contract, ", err)
		}

		if mci != 12 {
			t.Fatal("MCI read error.")
		}

		dbAsset.Header.PayloadHash = dbAsset.CalcPayloadHash()
		if !reflect.DeepEqual(dbAsset.Header, asset.Header) {
			t.Error("Save or load error, dbAsset header is not equal to asset header.")
		}

		if !reflect.DeepEqual(dbAsset.Contracts, asset.Contracts) {
			t.Error("Save or load error, dbAsset Contracts is not equal to asset Contracts.")
		}
		if len(dbAsset.Denominations) != len(asset.Denominations) {
			t.Error("Save or load error, dbAsset Denominations is not equal to asset Denominations.")
		}
		for i := 0; i < len(dbAsset.Denominations); i++ {
			if asset.Denominations[i] != dbAsset.Denominations[i] {
				t.Error("Save or load error, dbAsset Denominations is not equal to asset Denominations.")
			}
		}

		dbAsset.Header = asset.Header
		dbAsset.Contracts = asset.Contracts
		dbAsset.Denominations = asset.Denominations
		if !reflect.DeepEqual(dbAsset, asset) {
			t.Error("Save or load error, dbAsset is not equal to asset.")
		}

		return nil
	})

}
