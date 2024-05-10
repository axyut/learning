section .data
    num DD 5

section .text
global _start

_start:
    mov eax, 1
    mov ebx, [num] ; echo $? will echo whats at ebx
    mov edx, num; if do this it stores address of num and not 5 to reg ebx
    int 80h

; in section .data -> name type() value
; DB->1byte DW->2bytes DD->4bytes(32 bits) DQ->8bytes(64 bits)
; DT->10byte (80bits)
; DB - Define Byte  DW - Declare word

; instead of moving the address, move the value in that address with []