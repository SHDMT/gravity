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

	"github.com/SHDMT/crypto/bliss"
	"github.com/SHDMT/crypto/secp256k1"
	"github.com/SHDMT/gravity/infrastructure/crypto/asymmetric"
	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/infrastructure/log"
	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp"
)

//verify returns single signature result
func (lispvm *LispVM) verify(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 3 {
		return lisp.None, lisp.ErrParaNum
	}
	//get public key
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New("public key is not string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("public key is empty")
	}

	var pubKey asymmetric.PublicKey
	PK := []byte(x.Text.(string))
	PKHeader := PK[0]
	PKBody := PK[1:]
	log.Debugf(" verify pubkey head : %x \n", PKHeader)
	if PKHeader == 0 { //第一字节为0表示普通签名
		pubKey = new(secp256k1.PublicKey)
	} else {
		pubKey = new(bliss.PublicKey)
	}
	err = pubKey.UnmarshalP(PKBody)
	if err != nil {
		return lisp.None, errors.New("unmarshal public key failed")
	}
	//get content hash
	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}
	if y.Kind != lisp.String {
		return lisp.None, errors.New("contents is not string")
	}
	if len(y.Text.(string)) == 0 {
		return lisp.None, errors.New("contents is empty")
	}
	ConHash := ([]byte(y.Text.(string)))
	//get signature
	z, err := p.Exec(t[2])
	if err != nil {
		return lisp.None, err
	}
	if z.Kind != lisp.String {
		return lisp.None, errors.New("Sign is not string")
	}
	if len(z.Text.(string)) == 0 {
		return lisp.None, errors.New("sign is empty")
	}
	contSign := ([]byte(z.Text.(string)))
	//verify
	res := pubKey.Verify(ConHash, contSign)
	if !res {
		return lisp.False, errors.New("Verify sign failed")
	}
	return lisp.True, nil
}

//hash returns hash of content
func (lispvm *LispVM) hash(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.String {
		return lisp.None, errors.New("Contents is not string")
	}
	if len(x.Text.(string)) == 0 {
		return lisp.None, errors.New("Contents is empty")
	}
	contHashString := (hash.Sum256([]byte(x.Text.(string)))).String()
	return lisp.Token{Kind: lisp.String, Text: contHashString}, nil
}

//verifyMultiSign  returns multiple signature result
func (lispvm *LispVM) verifyMultiSign(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 4 {
		return lisp.None, lisp.ErrParaNum
	}
	//get threshold of signature
	threshold, err := p.Exec(t[3])
	if err != nil {
		return lisp.None, err
	}
	if threshold.Kind != lisp.Int {
		return lisp.None, lisp.ErrFitType
	}
	n := threshold.Text.(int64)
	//get public key list
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if x.Kind != lisp.List {
		return lisp.None, lisp.ErrFitType
	}
	pubKeys := make([]asymmetric.PublicKey, 0)
	for _, v := range x.Text.([]lisp.Token) {
		r, err := p.Exec(v)
		if err != nil {
			return lisp.None, err
		}
		if r.Kind != lisp.String {
			return lisp.None, errors.New("public key is not string")
		}
		if len(r.Text.(string)) == 0 {
			return lisp.None, errors.New("public key is empty")
		}
		var pubKey asymmetric.PublicKey
		PK := []byte(r.Text.(string))
		PKHeader := PK[0]
		PKBody := PK[1:]
		if PKHeader == 0 { //第一字节为0表示普通签名
			pubKey = new(secp256k1.PublicKey)
		} else {
			pubKey = new(bliss.PublicKey)
		}
		err = pubKey.UnmarshalP(PKBody)
		if err != nil {
			return lisp.None, errors.New("unmarshal pubKey failed")
		}
		pubKeys = append(pubKeys, pubKey)
	}
	//get content hash
	y, err := p.Exec(t[1])
	if err != nil {
		return lisp.None, err
	}
	if y.Kind != lisp.String {
		return lisp.None, errors.New("contents is not string")
	}
	if len(y.Text.(string)) == 0 {
		return lisp.None, errors.New("contents is empty")
	}
	ConHash := ([]byte(y.Text.(string)))
	//get signature list
	z, err := p.Exec(t[2])
	if err != nil {
		return lisp.None, err
	}
	if z.Kind != lisp.List {
		return lisp.None, lisp.ErrFitType
	}
	sigs := make([][]byte, 0)
	for _, v := range z.Text.([]lisp.Token) {
		r, err := p.Exec(v)
		if err != nil {
			return lisp.None, err
		}
		if r.Kind != lisp.String {
			return lisp.None, errors.New("Sig is not string")
		}
		if len(r.Text.(string)) == 0 {
			return lisp.None, errors.New("Sig is empty")
		}
		sig := []byte(r.Text.(string))
		sigs = append(sigs, sig)
	}
	//verify signature. if valid signatures' number is more than threshold,return true
	for _, sig := range sigs {
		for i, pubKey := range pubKeys {
			res := pubKey.Verify(ConHash, sig)
			if res {
				n--
				pubKeys = append(pubKeys[:i], pubKeys[i+1:]...)
				break
			}
		}
	}
	if n <= 0 {
		return lisp.True, nil
	}
	return lisp.False, nil
}

//countBytes returns bytes of content
func (lispvm *LispVM) countBytes(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	var bytesLen int64
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	x, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	bytesLen = int64(len([]byte(x.Text.(string))))
	return lisp.Token{Kind: lisp.Int, Text: bytesLen}, nil
}
