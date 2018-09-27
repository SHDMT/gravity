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
	"testing"
	"fmt"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/smartcontract/vm"
	"encoding/base64"
	"encoding/hex"
	"github.com/SHDMT/gravity/platform/consensus/genesis"
)

var addition =
`
(defun
    Addition ()
    (setq blc (calcBalance))
    (if (>= blc 0)
        (return-from Addition 0)
    ) 
    (setq memAddrs (getCurContractDefParamList "cosignerAddrs"))   
    (setq threshold (getCurContractDefParam "threshold")) 
    (update threshold (Int threshold))
    (setq count 0)
    (setq sCount (sigCount))
    (setq data (getCurUnitHashToSign))
    (loop (setq i 0) (< i sCount)                               
        (progn
            (setq pk (getPK i))
            (setq sig (getAuthorSig i))
            (setq addr (getAuthorAddr i))
            (setq ret (verify pk data sig))
            (if (not ret)
                (return-from Addition 0)
            )    
            (setq signal 0)
            (for memAddr memAddrs
                (if (eq memAddr addr)
                    (update signal 1)
                )
            )
            (if signal
                (update count (+ count 1))
            )
            (update i (+ i 1))
        )
    )
    (if (< count threshold)
        (return-from Addition 0)
    ) 
    (return-from Addition 1)
)

(setq res (Addition))


(print "Result of Contract[Addition]: ")
(println res)
`

var cosign =
`
(defun
    Cosign ()

    (setq blc (calcBalance))
    (setq inCount (inputCount))
    (if (or (/= blc 0) (= inCount 0))
        (return-from Cosign 0)
    )
    (setq cosignerAddrs (getCurContractDefParamList "cosignerAddrs"))
    (setq threshold (getCurContractDefParam "threshold"))
    (update threshold (Int threshold))
    (setq pks 0)
    (setq sigs 0)
    (for addr cosignerAddrs                                 
        (if (hasPKByAddr addr)
            (progn
                (setq pk (getPKByAddr addr))
                (setq sig (getSig pk))
                (if (not pks)
                    (update pks (list pk))
                    (update pks (cons pk pks))
                )
                (if (not sigs)
                    (update sigs (list sig))
                    (update sigs (cons sig sigs))
                )                
            )
        )
    )                                   
    (setq data (getCurUnitHashToSign))
    (setq ret (verifyMultiSign pks data sigs threshold))
    (if (not ret)
        (return-from Cosign 0)
    )
    (loop (setq i 0) (< i inCount) 
        (progn 
            (if (hasPrevOutParam i "addr")
                (progn       
                    (setq addr (getPrevOutParam i "addr"))          
                    (setq pk (getPKByAddr addr))
                    (setq sig (getSig pk))
                    (setq ret (verify pk data sig))
                    (if (not ret)
                        (return-from Cosign 0)
                    ) 
                ) 
            )
            (update i (+ i 1))
        )
    )
    (return-from Cosign 1)
)

(setq res (Cosign))

(print "Result of Contract[Cosign]: ")
(println res)
`

var simpleContract =
`
(print "simple Contract:OK")
(setq x 1)
`
func TestLispToBytes(t *testing.T){

	contract1 := structure.NewContract()

	contract1.Name = "addition"
	contract1.ScriptCode = vm.LispScriptCode
	contract1.IsRestrict = false
	contract1.Code = []byte(addition)
	contract1.AddParam("cosignerAddrs", "Official address")
	contract1.AddParam("threshold", "Sign limit")

	t.Logf("addition: %x\n", contract1.Serialize())
	t.Logf("addition addr: %x\n", contract1.CalcAddress())

	contract2 := structure.NewContract()

	contract2.Name = "cosign"
	contract2.ScriptCode = vm.LispScriptCode
	contract2.IsRestrict = false
	contract2.Code = []byte(cosign)
	contract2.AddParam("cosignerAddrs", "Official address")
	contract2.AddParam("threshold", "Sign limit")
	t.Logf("cosign: %x\n", contract2.Serialize())
	t.Logf("cosign addr: %x\n", contract2.CalcAddress())


	contract3 := structure.NewContract()

	contract3.Name = "SimpleContract"
	contract3.ScriptCode = vm.LispScriptCode
	contract3.IsRestrict = false
	contract3.Code = []byte(simpleContract)
	t.Logf("Simple contract: %x\n", contract3.Serialize())
	t.Logf("Simple contract addr: %v\n", hex.EncodeToString(contract3.CalcAddress()))
	t.Logf("Simple contract addr(base64): %s\n", base64.StdEncoding.EncodeToString(contract3.CalcAddress()))

	t.Logf("Simple Verify contract addr: %v\n", hex.EncodeToString(genesis.SimpleVerifyContract.CalcAddress()))

}

func TestEncodeDecode(t *testing.T){
	base64Str := "k4SzlT8MBVEqDHf0XJpSHQFVSByJUS9gFAoNcFApr34="
	bytes,_ := base64.StdEncoding.DecodeString(base64Str)

	t.Logf("hex : %x \n", bytes)

	var myJSONString = `{
    "outputs": [{
        "amount": 100018,
        "params": {
            "addr": "Mt6YsonHlOxlWvgl0Gjk7QRPC2qnLvWXiqe13Jsk56Y="
        }
    }]
}
`
	fmt.Printf(" hex string : %x \n", []byte(myJSONString))
}