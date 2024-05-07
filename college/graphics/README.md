## C

### with sdl2 `-lSDL2`

❯ gcc file.c -o sdl_test `sdl2-config --cflags --libs`
❯ ./sdl_test

### with OpenGL (GLUT)

❯ gcc rectangleBox.c -o rect -lglut -lGL
❯ ./rect

## C++

❯ g++ slope.cpp -o slope `sdl2-config --cflags --libs`
❯ ./slope
