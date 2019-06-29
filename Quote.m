ADD(A)
 NEW CURR
 ; CURR is current integer
 SET CURR=^Quotes(0)
 ; Write curr+1 to new quote
 SET ^Quotes(CURR+1)=A
 ; Update index
 SET ^Quotes(0)=CURR+1
 ; Clear temp variables and gtfo
 KILL CURR
 QUIT

READ(A)
 NEW CURR
 SET CURR=^Quotes(0)
 IF A>CURR WRITE !,"", ! QUIT
 WRITE ^Quotes(A)
 KILL CURR
 QUIT

RANDOM()
 SET VAL=$RAND(^Quotes(0))
 WRITE ^Quotes(VAL)
 KILL VAL
 QUIT