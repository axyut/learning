#include <SDL2/SDL.h>

int main() {
    // Initialize SDL
    if (SDL_Init(SDL_INIT_VIDEO) != 0) {
        SDL_Log("Unable to initialize SDL: %s", SDL_GetError());
        return 1;
    }

    // Create a window
    SDL_Window* window = SDL_CreateWindow("SDL Line Drawing",
                                          SDL_WINDOWPOS_CENTERED,
                                          SDL_WINDOWPOS_CENTERED,
                                          640, 480,
                                          SDL_WINDOW_SHOWN);
    if (!window) {
        SDL_Log("Unable to create window: %s", SDL_GetError());
        return 1;
    }

    // Create a renderer
    SDL_Renderer* renderer = SDL_CreateRenderer(window, -1,
                                                 SDL_RENDERER_ACCELERATED |
                                                 SDL_RENDERER_PRESENTVSYNC);
    if (!renderer) {
        SDL_Log("Unable to create renderer: %s", SDL_GetError());
        return 1;
    }

    // Set draw color (red in this case)
    SDL_SetRenderDrawColor(renderer, 255, 0, 0, 255);

    // Draw a line
    SDL_RenderDrawLine(renderer, 100, 100, 300, 200);

    // Render the changes
    SDL_RenderPresent(renderer);

    // Wait for 10 seconds before quitting
    SDL_Delay(10 * 1000);

    // Cleanup
    SDL_DestroyRenderer(renderer);
    SDL_DestroyWindow(window);
    SDL_Quit();

    return 0;
}
