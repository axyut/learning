#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <SDL2/SDL.h>

void DDALine(SDL_Renderer *renderer, int x0, int y0, int x1, int y1) {
    int dx = x1 - x0;
    int dy = y1 - y0;
    int steps = abs(dx) > abs(dy) ? abs(dx) : abs(dy);

    float xIncrement = dx / (float)steps;
    float yIncrement = dy / (float)steps;

    float x = x0;
    float y = y0;

    for (int i = 0; i <= steps; ++i) {
        SDL_RenderDrawPoint(renderer, round(x), round(y));
        x += xIncrement;
        y += yIncrement;
    }
}

int main() {
    const int WINDOW_WIDTH = 640;
    const int WINDOW_HEIGHT = 480;

    if (SDL_Init(SDL_INIT_VIDEO) != 0) {
        fprintf(stderr, "SDL_Init Error: %s\n", SDL_GetError());
        return 1;
    }

    SDL_Window* window = SDL_CreateWindow("DDA Line",
                                          SDL_WINDOWPOS_CENTERED,
                                          SDL_WINDOWPOS_CENTERED,
                                          WINDOW_WIDTH,
                                          WINDOW_HEIGHT,
                                          SDL_WINDOW_SHOWN);

    if (!window) {
        fprintf(stderr, "SDL_CreateWindow Error: %s\n", SDL_GetError());
        SDL_Quit();
        return 1;
    }

    SDL_Renderer* renderer = SDL_CreateRenderer(window, -1,
                                                 SDL_RENDERER_ACCELERATED |
                                                 SDL_RENDERER_PRESENTVSYNC);

    if (!renderer) {
        fprintf(stderr, "SDL_CreateRenderer Error: %s\n", SDL_GetError());
        SDL_DestroyWindow(window);
        SDL_Quit();
        return 1;
    }

    const Uint32 startMs = SDL_GetTicks();
    while (SDL_GetTicks() - startMs < 5000) {
        SDL_PumpEvents();
        SDL_SetRenderDrawColor(renderer, 255, 255, 255, 255); // White background
        SDL_RenderClear(renderer); // Clear the renderer with white color

        SDL_SetRenderDrawColor(renderer, 255, 0, 0, 255); // Red color
        DDALine(renderer, 100, 100, 300, 200); // Draw line from (100, 100) to (300, 200)

        SDL_RenderPresent(renderer);
    }


    // SDL_Delay(10000); // Wait for 3 seconds before quitting

    SDL_DestroyRenderer(renderer);
    SDL_DestroyWindow(window);
    SDL_Quit();

    return 0;
}
