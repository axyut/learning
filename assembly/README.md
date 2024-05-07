❯ nasm -f elf64 hello.asm
❯ ld -m elf_x86_64 -s -o hello_64 hello.o
❯ ./hello_64
Hello, world!
