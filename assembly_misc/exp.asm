i2r r0 i0
rset r1 0f1
rset r2 0d5
rset r3 0f1
rset r4 0f1
rset r5 0f0
rset r6 0f1
rset r7 0f1
addf r5 r6
multf r3 r0
multf r4 r5
cpy r7 r3
divf r7 r4
addf r1 r7
dec r2
jz r2 17
j 8
r2o r1 o0
j 0
