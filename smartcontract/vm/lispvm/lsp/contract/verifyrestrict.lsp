(defun
    VerifyRestrict ()

    (setq blc (calcBalance))
    (setq inCount (inputCount))
    (if (or (/= blc 0) (= inCount 0))
        (return-from VerifyRestrict 0)
    )
    (setq data (getCurUnitHashToSign))
    (setq memAddrs (getCurContractDefParamList "memberAddrs"))
    
    (loop (setq i 0) (< i inCount)
        (progn
            (if (hasPrevOutParam i "addr")
                (progn
                    (setq addr (getPrevOutParam i "addr"))
                    (setq flag 0)
                    (for memAddr memAddrs
                        (if (eq memAddr addr)
                            (update flag 1)
                        )
                    )
                    (if (not flag)
                        (return-from VerifyRestrict 0)  
                    )
                    (setq pk (getPKByAddr addr))
                    (setq sig (getSig pk))
                    (setq ret (verify pk data sig))
                    (if (not ret) 
                        (return-from VerifyRestrict 0)
                    )
                )
            )
            (update i (+ i 1))
        )
    )
    (return-from VerifyRestrict 1)
)

(setq res (VerifyRestrict))

(print "Result of Contract[VerifyRestrict]: ")
(println res)