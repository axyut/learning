```
❯ nasm -f elf64 hello.asm
❯ ld -m elf_x86_64 -s -o hello_64 hello.o
❯ ./hello_64
Hello, world!
```

x86 is a 32bit processor, each register is 32 bits long. x86_64 is a 64bit processor, each register is 64 bits long. The x86_64 architecture is backward compatible with x86. The x86_64 architecture is also known as AMD64, x64, and Intel 64.

## x86 Assembly with NASM

`nasm` Netwide Assembler, an assembler for the x86 architecture.

-   portable assembler that supports a range of x86 platforms.
-   Good for learning assembly language programming.

-   Example 32-bit registers:
    -   `EAX` accumulator,
    -   `EBX` base, general purpose,
    -   `ECX` loop counter,
    -   `EDX` data, general purpose,
    -   `ESI` source index, high speed memory transfer,
    -   `EDI` destination index, high speed memory transfer,
    -   `EBP` base pointer: reference function arguments and local variables,
    -   `ESP` stack pointer: to current stack address.
-   access 16 bits of a register, by dropping `E`: `AX`, `BX`, `CX`, `DX`, `SI`, `DI`, `BP`, `SP`.
-   access 8 bits of a register, high `H` and low `L`: `AH`, `AL`, `BH`, `BL`, `CH`, `CL`, `DH`, `DL`.

### first program

```asm
section .data ; stores variables
section .text ; actual code being run
global _start ; entry point of the program
_start: ; label, everything runs after this
    mov eax, 4 ; move number 4 into eax
    mov ebx, 1 ; move number 1 into ebx
;here we have two different values into register
    int 80h ; interrupt with system call, 80h is the system call number, will exit program with eax value
```

`compile`

```bash
nasm -f elf first.asm
    OR
❯ nasm -f elf -o ./bin/first.o first.asm
❯ ld -m elf_i386 -s -o ./bin/first ./bin/first.o
❯ bin/first
❯ echo $?
1
```

`see with gdb`

```bash
gdb -q first
    OR
gdb first
layout asm # to see assembly code
layout regs # to see registers
    break _start
    run
    stepi # to step through the program
    stepi # info registers eax -> to see eax value
        info registers eax
        x/x $eax shows address and value
    stepi
layout regs # to see registers
```
