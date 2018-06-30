clr r0
r2m r0 0
clr r0
r2m r0 1
rset r0 1
r2m r0 0
clr r0
rset r0 1
rset r1 1
rset r2 1
je r1 r2 13
rset r3 0
j 14
rset r3 1
je r0 r3 23
rset r1 6
rset r2 5
je r1 r2 20
rset r3 0
j 21
rset r3 1
je r0 r3 23
j 26
rset r1 11
r2m r1 1
j 29
rset r1 4
r2m r1 1
j 29
