clr r0
r2m r0 0
clr r0
rset r0 1
rset r1 1
je r0 r1 9
clr r1
je r0 r1 12
j 15
rset r1 11
r2m r1 0
j 15
rset r1 12
r2m r1 0
j 15
