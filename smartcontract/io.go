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
	"fmt"
	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/infrastructure/database"
	"github.com/SHDMT/gravity/infrastructure/log"
	"github.com/SHDMT/gravity/platform/smartcontract/internal/dbnamespace"
)

func dbPutContract(dbTx database.Tx, key, value []byte) error {
	contractBucket := dbTx.Data().Bucket(dbnamespace.ContractBucket)

	err := contractBucket.Put(key, value)
	if err != nil {
		errString := fmt.Sprintf("Failed to put contract %v", key)
		return NewSmartContractError(ErrPutDB, errString, err)
	}
	return nil
}

func dbFetchContract(dbTx database.Tx, key []byte) ([]byte, error) {
	contractBucket := dbTx.Data().Bucket([]byte(dbnamespace.ContractBucket))

	value := contractBucket.Get(key)
	if value == nil {
		errString := fmt.Sprintf("Failed to find contract %v", key)
		return nil, NewSmartContractError(ErrNotFoundFormDB, errString, nil)
	}

	return value, nil
}

func dbDeleteContract(dbTx database.Tx, key []byte) error {
	contractBucket := dbTx.Data().Bucket([]byte(dbnamespace.ContractBucket))

	err := contractBucket.Delete(key)
	if err != nil {
		errString := fmt.Sprintf("Failed to delete contract %v", key)
		return NewSmartContractError(ErrDeleteDB, errString, err)
	}

	return nil
}

func dbHasContract(dbTx database.Tx, key []byte) bool {
	contractBucket := dbTx.Data().Bucket(dbnamespace.ContractBucket)
	return contractBucket.KeyExists(key)
}

func dbListAllContracts(dbTx database.Tx) ([]hash.HashType, error) {
	contractBucket := dbTx.Data().Bucket(dbnamespace.ContractBucket)
	contracts := make([]hash.HashType, 0, 16)
	err := contractBucket.ForEach(func(k, v []byte) error {
		contracts = append(contracts, k)
		return nil
	})
	if err != nil {
		errString := fmt.Sprintf("Failed to list all contracts")
		return nil, NewSmartContractError(ErrForEachDB, errString, err)
	}
	return contracts, nil
}

func dbPutAssetContract(dbTx database.Tx, key, value []byte) error {
	assetContractBucket := dbTx.Data().Bucket(dbnamespace.AssetContractBucket)

	err := assetContractBucket.Put(key, value)
	if err != nil {
		errString := fmt.Sprintf("Failed to put assetContract %v", key)
		return NewSmartContractError(ErrPutDB, errString, err)
	}
	return nil
}

func dbFetchAssetContract(dbTx database.Tx, key []byte) ([]byte, error) {
	assetContractBucket := dbTx.Data().Bucket([]byte(dbnamespace.AssetContractBucket))

	value := assetContractBucket.Get(key)
	if value == nil {
		errString := fmt.Sprintf("Failed to find assetContract %v", key)
		return nil, NewSmartContractError(ErrNotFoundFormDB, errString, nil)
	}

	return value, nil
}

func dbDeleteAssetContract(dbTx database.Tx, key []byte) error {
	assetContractBucket := dbTx.Data().Bucket([]byte(dbnamespace.AssetContractBucket))

	err := assetContractBucket.Delete(key)
	if err != nil {
		errString := fmt.Sprintf("Failed to delete assetContract %v", key)
		return NewSmartContractError(ErrDeleteDB, errString, err)
	}

	return nil
}

func dbHasAssetContract(dbTx database.Tx, key []byte) bool {
	assetContractBucket := dbTx.Data().Bucket(dbnamespace.AssetContractBucket)
	return assetContractBucket.KeyExists(key)
}

func dbPutAsset(dbTx database.Tx, key, value []byte) error {
	assetBucket := dbTx.Data().Bucket(dbnamespace.AssetBucket)

	err := assetBucket.Put(key, value)
	if err != nil {
		errString := fmt.Sprintf("Failed to put asset %v", key)
		return NewSmartContractError(ErrPutDB, errString, err)
	}
	return nil
}

func dbFetchAsset(dbTx database.Tx, key []byte) ([]byte, error) {
	assetBucket := dbTx.Data().Bucket([]byte(dbnamespace.AssetBucket))

	value := assetBucket.Get(key)
	if value == nil {
		errString := fmt.Sprintf("Failed to find asset %v", key)
		return nil, NewSmartContractError(ErrNotFoundFormDB, errString, nil)
	}

	return value, nil
}

func dbDeleteAsset(dbTx database.Tx, key []byte) error {
	assetBucket := dbTx.Data().Bucket([]byte(dbnamespace.AssetBucket))

	err := assetBucket.Delete(key)
	if err != nil {
		errString := fmt.Sprintf("Failed to delete asset %v", key)
		return NewSmartContractError(ErrDeleteDB, errString, err)
	}

	return nil
}

func dbHasAsset(dbTx database.Tx, key []byte) bool {
	assetBucket := dbTx.Data().Bucket(dbnamespace.AssetBucket)
	return assetBucket.KeyExists(key)
}
//CreateSmartContractBucket  is the bucket associated with creating smart contracts
func CreateSmartContractBucket(db database.Db) error {
	err := db.Update(func(tx database.Tx) error {
		errs := make([]error, 3)
		_, errs[0] = tx.Data().CreateBucket(dbnamespace.ContractBucket)
		_, errs[1] = tx.Data().CreateBucket(dbnamespace.AssetContractBucket)
		_, errs[2] = tx.Data().CreateBucket(dbnamespace.AssetBucket)

		for _, err := range errs {
			if err != nil {
				dbErr := err.(*database.DbError)
				if !database.IsBucketAlreadyExistsError(dbErr) {
					log.Errorf("create bucket err : %v\n", err)
					return err
				}
			}
		}

		return nil
	})
	return err
}
