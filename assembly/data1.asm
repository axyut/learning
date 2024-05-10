section .data
    num1 DB 3
    num2 DB 2

section .text
global _start

_start:
    mov bl, [num1] ; echo $? -> 3
    mov cl, [num2]
    mov eax, 1
    int 80h

; ebx means get whole of 32 bits of register, bx 16 bits, 
;   here we have 2 (8 bit) halves in b-x, each side(8bit) bl, bh
; using whole ecx for 2nd data here results in seg fault

; if used bh, ch there also doesn't matter, just choose right side bits
;   of register

;❯ nasm -f elf -o ./bin/data1.o data1.asm
;❯ ld -m elf_i386 -s -o ./bin/data1 ./bin/data1.o
;❯ bin/data1
;❯ echo $?