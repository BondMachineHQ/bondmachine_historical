clr r0
r2m r0 0
clr r0
r2m r0 1
rset r0 1
r2m r0 0
m2r r0 0
rset r1 1
je r0 r1 14
rset r1 2
je r0 r1 17
rset r1 3
je r0 r1 21
j 24
rset r1 11
r2m r1 1
j 24
rset r1 12
r2m r1 1
j 21
j 24
rset r1 13
r2m r1 1
j 24
