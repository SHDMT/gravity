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
	"encoding/binary"
	"math"

	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/infrastructure/database"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/pkg/errors"
)

//ContractLibrary is a database read-write class
type ContractLibrary struct {
	readOnly bool
	tx       database.Tx
}

const writePermissionsError = "don't have write permissions"

//SaveContract is used to store contract to database
func (library *ContractLibrary) SaveContract(contract *structure.Contract, mci uint64) error {
	if library.readOnly {
		return errors.Errorf(writePermissionsError)
	}
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, mci)

	address := contract.CalcAddress()
	contractBytes := contract.Serialize()

	buf = append(buf, contractBytes...)

	return dbPutContract(library.tx, address, buf)
}

//ListContracts list contracts according to hashes
func (library *ContractLibrary) ListContracts() ([]hash.HashType, error) {
	return dbListAllContracts(library.tx)
}

//LoadContract is to read contract according to the address
func (library *ContractLibrary) LoadContract(address hash.HashType) (*structure.Contract, uint64, error) {
	var buf []byte
	data, err := dbFetchContract(library.tx, address)
	buf = data

	if err != nil {
		return nil, math.MaxUint64, err
	}
	mci := binary.BigEndian.Uint64(data)

	contract := structure.NewContract()
	contract.Deserialize(buf[8:])

	return contract, mci, err
}

//HasContract returns if the contract is existed
func (library *ContractLibrary) HasContract(address hash.HashType) bool {
	return dbHasContract(library.tx, address)
}

//RemoveContract is to remove contract from database
func (library *ContractLibrary) RemoveContract(address hash.HashType) error {
	if library.readOnly {
		return errors.Errorf(writePermissionsError)
	}

	return dbDeleteContract(library.tx, address)

}

//SaveAssetContractDef is to store all contracts definition associated with asset
func (library *ContractLibrary) SaveAssetContractDef(asset hash.HashType, contractDef *structure.ContractDef) error {
	if library.readOnly {
		return errors.Errorf(writePermissionsError)
	}

	addr := contractDef.Address
	key := append(asset, addr...)
	value := contractDef.Serialize()

	return dbPutAssetContract(library.tx, key, value)
}

//LoadAssetContractDef is to read contract definition associate with asset
func (library *ContractLibrary) LoadAssetContractDef(asset hash.HashType, addr hash.HashType) (*structure.ContractDef, error) {
	key := append(asset, addr...)
	var buf []byte

	data, err := dbFetchAssetContract(library.tx, key)
	buf = data

	if err != nil {
		return nil, err
	}

	contractDef := new(structure.ContractDef)
	contractDef.Deserialize(buf)

	return contractDef, nil
}

//RemoveAssetContract is to remove a contract associate with asset
func (library *ContractLibrary) RemoveAssetContract(asset hash.HashType, address hash.HashType) error {
	if library.readOnly {
		return errors.Errorf(writePermissionsError)
	}

	key := append(asset, address...)

	err := dbDeleteAssetContract(library.tx, key)

	return err
}

//LoadAssetContract is to read a contract associate with current asset
func (library *ContractLibrary) LoadAssetContract(asset hash.HashType, addr hash.HashType) (*structure.ContractDef, *structure.Contract, error) {
	contractDef, err := library.LoadAssetContractDef(asset, addr)

	if err != nil {
		return nil, nil, err
	}

	if !library.HasContract(contractDef.Address) {
		return contractDef, nil, errors.Errorf("Can't find contract from Contract library.")
	}

	contract, _, err := library.LoadContract(contractDef.Address)

	if err != nil {
		return contractDef, nil, err
	}

	return contractDef, contract, err
}

//SaveAsset is to store asset
func (library *ContractLibrary) SaveAsset(unithash hash.HashType, asset *structure.IssueMessage, mci uint64) error {
	if library.readOnly {
		return errors.Errorf(writePermissionsError)
	}
	address := unithash
	//address := asset.CalcPayloadHash()

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, mci)

	assetBytes := asset.Serialize()
	buf = append(buf, assetBytes...)

	return dbPutAsset(library.tx, address, buf)

}

//LoadAsset is to read asset
func (library *ContractLibrary) LoadAsset(assetHash hash.HashType) (*structure.IssueMessage, uint64, error) {
	var buf []byte

	data, err := dbFetchAsset(library.tx, assetHash)
	buf = data

	if err != nil {
		return nil, 0, err
	}
	mci := binary.BigEndian.Uint64(data)

	issueMessage := structure.NewIssueMessage()
	issueMessage.Deserialize(buf[8:])

	return issueMessage, mci, err
}

//RemoveAsset is to remove asset
func (library *ContractLibrary) RemoveAsset(asset hash.HashType) error {
	if library.readOnly {
		return errors.Errorf(writePermissionsError)
	}

	err := dbDeleteAsset(library.tx, asset)

	return err
}

//HasAsset is to check if current asset is exist in the database
func (library *ContractLibrary) HasAsset(asset hash.HashType) bool {
	return dbHasAsset(library.tx, asset)
}

//HasAssetContract is to check if current asset contract is exist in the database
func (library *ContractLibrary) HasAssetContract(asset hash.HashType, addr hash.HashType) bool {
	key := append(asset, addr...)
	return dbHasAssetContract(library.tx, key)

}

//NewContractLibrary is to create a new object to access contract database
func NewContractLibrary(tx database.Tx, readOnly bool) *ContractLibrary {
	return &ContractLibrary{
		tx:       tx,
		readOnly: readOnly,
	}
}
