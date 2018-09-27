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

import "fmt"

//SmartContractError is the wrong structure associated with smart contracts
type SmartContractError struct {
	errorCode   uint32 // Describes the kind of error
	description string // Human readable description of the issue
	err         error  // Underlying error
}

const (
	//ErrNotFoundFormDB is not found in the database
	ErrNotFoundFormDB = iota
	//ErrPutDB was an insert database failure
	ErrPutDB
	//ErrDeleteDB was a failed delete database
	ErrDeleteDB
	//ErrForEachDB is traversal failure
	ErrForEachDB
)

//Error is the error type converted to string type
func (e *SmartContractError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%v, %s:\n%s", e.errorCode, e.description, e.err)
	}
	return fmt.Sprintf("%v, %s", e.errorCode, e.description)
}

//NewSmartContractError is creating a new error object
func NewSmartContractError(code uint32, des string, err error) *SmartContractError {
	return &SmartContractError{
		errorCode:   code,
		description: des,
		err:         err,
	}
}

//IsNotFoundFormDBError  is that this error does not exist
func IsNotFoundFormDBError(smartContractError *SmartContractError) bool {
	return smartContractError.isErrorType(ErrNotFoundFormDB)
}

func (e *SmartContractError) isErrorType(code uint32) bool {
	if e.errorCode == code {
		return true
	}
	switch e.err.(type) {
	case *SmartContractError:
		return e.err.(*SmartContractError).isErrorType(code)
	default:
		return false
	}
}
