
test.o:     file format elf64-x86-64


Disassembly of section .text:

0000000000000000 <main>:
#include "stdio.h"

int main() {
   0:	55                   	push   rbp
   1:	48 89 e5             	mov    rbp,rsp
	int test;
	for (test=0; test < 5 ; test ++) {
   4:	c7 45 fc 00 00 00 00 	mov    DWORD PTR [rbp-0x4],0x0
   b:	eb 04                	jmp    11 <main+0x11>
   d:	83 45 fc 01          	add    DWORD PTR [rbp-0x4],0x1
  11:	83 7d fc 04          	cmp    DWORD PTR [rbp-0x4],0x4
  15:	7e f6                	jle    d <main+0xd>
	}
} 
  17:	5d                   	pop    rbp
  18:	c3                   	ret    
