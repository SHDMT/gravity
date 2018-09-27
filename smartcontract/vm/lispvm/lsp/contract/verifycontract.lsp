(defun 
    VerifyContract ()
     
    (setq blc (calcBalance))
    (setq inCount (inputCount))
    (if (or (/= blc 0) (= inCount 0))
        (return-from VerifyContract 0)
    )
    (setq data (getCurUnitHashToSign))  
    (loop (setq i 0) (< i inCount) 
        (progn 
            (if (hasPrevOutParam i "addr")
                (progn       
                    (setq addr (getPrevOutParam i "addr"))          
                    (setq pk (getPKByAddr addr))
                    (setq sig (getSig pk))
                    (setq ret (verify pk data sig))
                    (if (not ret)
                        (return-from VerifyContract 0)
                    ) 
                ) 
            )
            (update i (+ i 1))
        )
    )
    (return-from VerifyContract 1)
)

(setq res (VerifyContract))

(print "Result of Contract[VerifyContract]: ")
(println res)