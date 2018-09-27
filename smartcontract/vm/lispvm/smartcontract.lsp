(define 
    (verifyContract)
    (each 
        (define inputCount (getInputCount))
        (if (> inputCount 0) 
            (define isRight 1)
            (define isRight 0)
        )
        
        (loop (define i 0)(and (< i inputCount) isRight)
            (each
                (define addr (getPrevOutParam i "addr"))
                (define pk (getPKByAddr addr))
                (define sig (getSig pk))
                (define data (getCurUnitHashToSign))
                #(define data "1")
                (define ret (verify pk data sig))
                (if ret 
                    (define i (+ i 1))
                    (define isRight ret)
                )
            )
        )
        (define ret isRight)
    )
)

