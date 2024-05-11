section .text
global _start

_start:
    mov al, 2
    mov bl, 3
    mul bl ; doesn't require two operand, its al*bl
    int 80h

;it will also expand the output size when vlaue exceeds for us
