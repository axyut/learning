section .data

section .text
global _start

_start:
    mov al, 3
    mov bl, 2
    sub al, bl
    int 80h
