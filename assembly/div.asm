section .data
global _start 

_start:
    mov eax, 11
    mov ecx, 2
    div ecx
    IDIV ecx
    int 80h
