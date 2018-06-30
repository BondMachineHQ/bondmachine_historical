
test.o:     file format elf64-x86-64


Disassembly of section .text:

0000000000000000 <main>:
#include "stdio.h"

int main() {
   0:	55                   	push   %rbp
   1:	48 89 e5             	mov    %rsp,%rbp
	int test;
	for (test=0; test < 5 ; test ++) {
   4:	c7 45 fc 00 00 00 00 	movl   $0x0,-0x4(%rbp)
   b:	eb 04                	jmp    11 <main+0x11>
   d:	83 45 fc 01          	addl   $0x1,-0x4(%rbp)
  11:	83 7d fc 04          	cmpl   $0x4,-0x4(%rbp)
  15:	7e f6                	jle    d <main+0xd>
	}
} 
  17:	5d                   	pop    %rbp
  18:	c3                   	retq   
