#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <iostream>

#include "../include/RenderWindow.hpp"
#include "../include/Entity.hpp"
#include "../include/Defs.hpp"

int main (int argc, char *argv[]) {

  if (SDL_Init(SDL_INIT_VIDEO) > 0) {
    std::cout << "SDL_INIT FAILED: " << SDL_GetError() << std::endl;
  }
  if (!IMG_Init(IMG_INIT_PNG)) {
    std::cout << "IMG_INIT FAILED: " << SDL_GetError() << std::endl;
  }

  RenderWindow window("Game v1.0", 1280, 6*32*4);

  SDL_Texture* grassTexture = window.loadTexture("res/gfx/ground_grass_1.png");
  Entity platform0(100, 50, grassTexture);
  Entity entities[5] = {
    Entity(0, 5*32, grassTexture),
    Entity(1*32, 5*32, grassTexture),
    Entity(2*32, 5*32, grassTexture),
    Entity(3*32, 5*32, grassTexture),
    Entity(4*32, 5*32, grassTexture)
  };

  


  bool gameRunning = true;
  SDL_Event event;
  while (gameRunning) 
  {
    while (SDL_PollEvent(&event)) {
      if (event.type == SDL_QUIT)
        gameRunning = false;
    }

    window.clear();

    for (int i = 0; i < 5; i++)
    {
      window.render(entities[i]);
    }

    window.display();
  }

  window.cleanUp();
  SDL_Quit();

  return 0;
}
