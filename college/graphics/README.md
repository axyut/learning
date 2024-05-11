## C

### with sdl2 `-lSDL2`

```
❯ gcc file.c -o sdl_test `sdl2-config --cflags --libs`
❯ ./sdl_test
```

`for kde plasma, for other DEs, SDL_Delay(5000) should work, without the while loop & SDL_PumpEvents()`

```
    const Uint32 startMs = SDL_GetTicks();
    while( SDL_GetTicks() - startMs < 5000 )
    {
        SDL_PumpEvents();
        // other stuff
    }
    // SDL_Delay(5000);
```

`❯ gcc slope.c -o slope -lSDL2 && ./slope`
`SDL2 seems to be a better option than other libraries`

### with OpenGL (GLUT)

```
❯ gcc rectangleBox.c -o rect -lglut -lGL
❯ ./rect
```

## C++

```
❯ g++ slope.cpp -o slope `sdl2-config --cflags --libs`
❯ ./slope
```

```
❯ g++ redRect.cpp -o red -lglut -lGL
❯ ./red
```
