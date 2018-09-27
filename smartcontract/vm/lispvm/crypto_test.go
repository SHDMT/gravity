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
	"crypto/rand"
	"testing"

	"github.com/SHDMT/gravity/infrastructure/crypto/asymmetric"
	"github.com/SHDMT/gravity/infrastructure/crypto/asymmetric/bliss"
	"github.com/SHDMT/gravity/infrastructure/crypto/asymmetric/ec/secp256k1"
	"github.com/SHDMT/gravity/infrastructure/crypto/hash"
	"github.com/SHDMT/gravity/platform/smartcontract/vm"
	"github.com/SHDMT/gravity/platform/smartcontract/vm/lispvm/lisp"
)

func TestVerify(t *testing.T) {
	//1-1 normal signature verification
	content := []byte("testcontent")
	hashedContent := hash.Sum256(content)

	test := secp256k1.NewCipherSuite()
	priK, err := test.GenerateKey(rand.Reader)
	if err != nil {
		t.Error("generate key failed")
	}
	pubK := priK.Public()
	pubKByteBody, _ := pubK.MarshalP()
	var pubKByte = []byte{0}
	pubKByte = append(pubKByte, pubKByteBody...)
	sign, err := priK.Sign([]byte(hashedContent))
	if err != nil {
		t.Error("sign failed")
	}
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	var tks []lisp.Token

	PK := lisp.Token{Kind: lisp.String, Text: string(pubKByte)}
	chash := lisp.Token{Kind: lisp.String, Text: string(hashedContent)}
	sig := lisp.Token{Kind: lisp.String, Text: string(sign)}
	tks = append(tks, PK, chash, sig)
	_, err = lvm.verify(tks, lvm.vm)
	if err != nil {
		t.Error("verify failed,err=", err)
	}
	//1-2 bliss verification
	blissTest := bliss.NewCipherSuite()
	blissPriK, err := blissTest.GenerateKey(rand.Reader)
	if err != nil {
		t.Error("generate bliss key failed")
	}
	blissPubK := blissPriK.Public()
	blissPubKBody, err := blissPubK.MarshalP()
	var blissPubKByte = []byte{1}
	blissPubKByte = append(blissPubKByte, blissPubKBody...)
	blissSign, err := blissPriK.Sign([]byte(hashedContent))
	if err != nil {
		t.Error("bliss sign failed")
	}
	blissPK := lisp.Token{Kind: lisp.String, Text: string(blissPubKByte)}
	blissSig := lisp.Token{Kind: lisp.String, Text: string(blissSign)}
	var blissTks []lisp.Token
	blissTks = append(blissTks, blissPK, chash, blissSig)
	_, err = lvm.verify(blissTks, lvm.vm)
	if err != nil {
		t.Error("bliss verify failed,err=", err)
	}

	//1-3 parameter count invalid
	_, err = lvm.vm.Eval([]byte(`
      (println (verify "aaa"))
    `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-4 get public key failed
	_, err = lvm.vm.Eval([]byte(`
      (println (verify (test) "xx" "xx"))
    `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-5 public key is invalid
	_, err = lvm.vm.Eval([]byte(`
      (println (verify (+ 1 2) "xx" "xx"))
    `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-6 public key is empty
	_, err = lvm.vm.Eval([]byte(`
     (each
  	 (define pk "")
     (println (verify pk "a" "xx"))
   ) `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-7 public key unmarshal failed
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define pk "test")
     (println (verify pk (test)  "xx"))
    )`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-8 get content hash failed
	lvm = NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	unit, err := FakeUnitWithSig()
	lvm.context = vm.Context{
		TxUnit:       *unit,
		FetchPrevOut: GetContractOutput,
	}
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define addr (getPrevOutParam 0 "addr"))
	 (define pk (getPKByAddr addr))
	 (println (verify pk a b))
    )`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}

	//1-9 get hash content failed
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define addr (getPrevOutParam 0 "addr"))
	 (define pk (getPKByAddr addr))
	 (println (verify pk (+ 1 2) b))
    )`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}

	//1-10 get content type failed
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define addr (getPrevOutParam 0 "addr"))
	 (define pk (getPKByAddr addr))
	 (println (verify pk (+ 1 2) b))
    )`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}

	//1-11 content is empty
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define addr (getPrevOutParam 0 "addr"))
	 (define pk (getPKByAddr addr))
	 (println (verify pk "" b))
    )`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-12 get signature failed
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define addr (getPrevOutParam 0 "addr"))
	 (define pk (getPKByAddr addr))
	 (define data (getCurUnitHashToSign))
	 (println (verify pk data b))
    )`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-13 signature type is invalid
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define addr (getPrevOutParam 0 "addr"))
	 (define pk (getPKByAddr addr))
	 (define data (getCurUnitHashToSign))
	 (println (verify pk data (+ 1 2)))
    )`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-14 sign content is empty
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define addr (getPrevOutParam 0 "addr"))
	 (define pk (getPKByAddr addr))
	 (define data (getCurUnitHashToSign))
	 (println (verify pk data ""))
    )`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//1-15 verify failed
	_, err = lvm.vm.Eval([]byte(`
	(each
  	 (define addr (getPrevOutParam 0 "addr"))
	 (define pk (getPKByAddr addr))
	 (define data (getCurUnitHashToSign))
	 (println (verify pk data "aaa"))
    )`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}

}

func TestHash(t *testing.T) {
	//2-1 normal case
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	_, err := lvm.vm.Eval([]byte(`
	(each
      (define content "testcontent")
      (println (hash content))
    )`))
	if err != nil {
		t.Errorf("Lisp hash test failed,err= %v", err)
	}
	//2-2 parameter number invalid
	_, err = lvm.vm.Eval([]byte(`
	(each
      (define content "testcontent")
      (println (hash content "test"))
    )`))
	if err == nil {
		t.Errorf("Lisp hash test failed,err= %v", err)
	}
	//2-3 first parameter execute failed
	_, err = lvm.vm.Eval([]byte(`
      (println (hash (test)))
    `))
	if err == nil {
		t.Errorf("Lisp hash test failed,err= %v", err)
	}
	//2-4 hash content is not string
	_, err = lvm.vm.Eval([]byte(`
      (println (hash (+ 1 2)))
    `))
	if err == nil {
		t.Errorf("Lisp hash test failed,err= %v", err)
	}
	//2-5 hash content is empty
	_, err = lvm.vm.Eval([]byte(`
      (each
      (define content "")
      (println (hash content))
    )`))
	if err == nil {
		t.Errorf("Lisp hash test failed,err= %v", err)
	}

}

func TestVerifyMultiSign(t *testing.T) {
	//3-1 normal case
	content := []byte("testcontent")
	hashedContent := hash.Sum256(content)
	N, threshold := 6, 6
	pubKeys := make([]string, N)
	signs := make([]string, N)
	test := secp256k1.NewCipherSuite()
	blissTest := bliss.NewCipherSuite()
	for i := 0; i < N; i++ {
		var priK asymmetric.PrivateKey
		var err error
		var pkHeader byte
		if i%2 == 0 {
			priK, err = test.GenerateKey(rand.Reader)
			if err != nil {
				t.Error("generate key failed")
			}
			pkHeader = 0
		} else {
			priK, err = blissTest.GenerateKey(rand.Reader)
			if err != nil {
				t.Error("generate bliss key failed")
			}
			pkHeader = 1
		}
		pubK := priK.Public()
		pkBody, err := pubK.MarshalP()
		var pkByte = []byte{pkHeader}
		pkByte = append(pkByte, pkBody...)
		pubKeys[i] = string(pkByte)
		if err != nil {
			t.Error("marshal pubKey failed")
		}
		signByte, err := priK.Sign([]byte(hashedContent))
		signs[i] = string(signByte)
		if err != nil {
			t.Error("sign failed")
		}
	}
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	pks := make([]lisp.Token, 0)
	sigs := make([]lisp.Token, 0)
	for i := 0; i < N; i++ {
		pks = append(pks, lisp.Token{Kind: lisp.String, Text: pubKeys[i]})
		sigs = append(sigs, lisp.Token{Kind: lisp.String, Text: signs[i]})
	}
	pkTokens := lisp.Token{Kind: lisp.Fold, Text: pks}
	data := lisp.Token{Kind: lisp.String, Text: string(hashedContent)}
	sigTokens := lisp.Token{Kind: lisp.Fold, Text: sigs}
	n := lisp.Token{Kind: lisp.Int, Text: int64(threshold)}
	var tks []lisp.Token
	tks = append(tks, pkTokens, data, sigTokens, n)
	res, err := lvm.verifyMultiSign(tks, lvm.vm)
	if err != nil {
		t.Error("verify failed,err=", err)
	}
	if res.Kind == lisp.List {
		t.Error("verify failed!")
	}
	//3-2 parameter number invalid
	_, err = lvm.vm.Eval([]byte(`
     (println (verifyMultiSign "aaa"))
   `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-3 get threshold failed
	_, err = lvm.vm.Eval([]byte(`
     (println (verifyMultiSign "aaa" "bbb" "ccc" "ddd"'))
   `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-4 get threshold type failed
	_, err = lvm.vm.Eval([]byte(`
     (println (verifyMultiSign "aaa" "bbb" "ccc" (quote "xxx")))
   `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-5 get public key failed
	_, err = lvm.vm.Eval([]byte(`
      (println (verifyMultiSign (test) "xx" "xx" 1))
    `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-6 get public key type failed（not label)
	_, err = lvm.vm.Eval([]byte(`
      (println (verifyMultiSign (+ 1 2) "xx" "xx" 1))
    `))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-7 public key type invalid（not string)
	lvm = NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeContract})
	unit, err := FakeUnitWithSig()
	lvm.context = vm.Context{
		TxUnit:       *unit,
		FetchPrevOut: GetContractOutput,
	}
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(println (verifyMultiSign '((+ 1 2)) a b 1))
	)`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//public key unmarshal failed
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(println (verifyMultiSign '(test) a b 1))
	)`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}

	//3-8 public key is empty
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(println (verifyMultiSign '("") a b 1))
	)`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-9 public key unmarshal failed
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(println (verifyMultiSign '("aa") a b 1))
	)`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-10 get content hash failed
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(println (verifyMultiSign '(pk) a b 1))
	)`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}

	//3-11 get content hash failed
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(println (verifyMultiSign '(pk) (+ 1 2) b 1))
	)`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-12 content hash is empty
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(println (verifyMultiSign '(pk) "" b 1))
	)`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-13 get signature list failed
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(define data (getCurUnitHashToSign))
	(println (verifyMultiSign '(pk) data b 1))
	)`))

	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}

	//3-14 get signatures failed
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(define data (getCurUnitHashToSign))
	(println (verifyMultiSign '(pk) data "" 1))
	)`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-15 get signature from signature list failed
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(define data (getCurUnitHashToSign))
	(println (verifyMultiSign '(pk) data '(test) 1))
	)`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-16 signature type is invalid(not string)
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(define data (getCurUnitHashToSign))
	(println (verifyMultiSign '(pk) data '((+ 1 2)) 1))
	)`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	////3-17 sign content is empty
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(define data (getCurUnitHashToSign))
	(println (verifyMultiSign '(pk) data '("") 1))
	)`))
	if err == nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
	//3-18 very failed (valid signature is less than threshold)
	_, err = lvm.vm.Eval([]byte(`
	(each
	(define addr (getPrevOutParam 0 "addr"))
	(define pk (getPKByAddr addr))
	(define sig (getSig pk))
	(define data (getCurUnitHashToSign))
	(println (verifyMultiSign '(pk) data '("xxx") 1))
	)`))
	if err != nil {
		t.Errorf("Lisp verify test failed,err= %v", err)
	}
}

func TestLispVM_countBytes(t *testing.T) {
	//normal case
	lvm := NewLispVM(vm.Context{}, vm.Config{Mode: vm.VMModeRestrict})
	_, err := lvm.vm.Eval([]byte(`
	(each
      (define content "teststringlen")
      (println (countBytes content))
    )`))
	if err != nil {
		t.Errorf("Lisp hash test failed,err= %v", err)
	}
	//parameter count invalid
	_, err = lvm.vm.Eval([]byte(`
	(each
      (define content "teststringlen")
      (println (countBytes ))
    )`))
	if err == nil {
		t.Errorf("Lisp hash test failed,err= %v", err)
	}

	//get count of content failed
	_, err = lvm.vm.Eval([]byte(`
	(each
      (define content "teststringlen")
      (println (countBytes (content)))
    )`))
	if err == nil {
		t.Errorf("Lisp hash test failed,err= %v", err)
	}

}
