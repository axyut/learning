section .data
    char DB 'A'
    list db 1,2,3,4
    string db "ABC", 0

section .text
global _start

_start:
    mov bl, [list]
    mov eax, 1
    int 80h

; [char] echoes 65, ascii rather string
; without putting in eax it throws seg fault, idk why
; [list]
; null terminator -> put -1 at end of the list
; string is just list of numbers so it requires null terminator
; string db "abc", 0