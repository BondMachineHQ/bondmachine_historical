
prova.o:     file format elf32-littlearm


Disassembly of section .text:

00000000 <main>:
#include "stdio.h"

int main() {
   0:	e52db004 	push	{fp}		; (str fp, [sp, #-4]!)
   4:	e28db000 	add	fp, sp, #0
   8:	e24dd00c 	sub	sp, sp, #12
	int test;
	for (test=0; test < 5 ; test ++) {
   c:	e3a03000 	mov	r3, #0
  10:	e50b3008 	str	r3, [fp, #-8]
  14:	ea000002 	b	24 <main+0x24>
  18:	e51b3008 	ldr	r3, [fp, #-8]
  1c:	e2833001 	add	r3, r3, #1
  20:	e50b3008 	str	r3, [fp, #-8]
  24:	e51b3008 	ldr	r3, [fp, #-8]
  28:	e3530004 	cmp	r3, #4
  2c:	dafffff9 	ble	18 <main+0x18>
	}
} 
  30:	e1a00003 	mov	r0, r3
  34:	e28bd000 	add	sp, fp, #0
  38:	e8bd0800 	ldmfd	sp!, {fp}
  3c:	e12fff1e 	bx	lr
