#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_timer.h>
#include <iostream>
#include <vector>

#include "../include/Defs.hpp"
// #include "../include/Entity.hpp"
#include "../include/RenderWindow.hpp"
#include "../include/RigidBody.hpp"
#include "../include/Utils.hpp"

int main(int argc, char *argv[]) {

  if (SDL_Init(SDL_INIT_VIDEO) > 0) {
    std::cout << "SDL_INIT FAILED: " << SDL_GetError() << std::endl;
  }
  if (!IMG_Init(IMG_INIT_PNG)) {
    std::cout << "IMG_INIT FAILED: " << SDL_GetError() << std::endl;
  }

  RenderWindow window("Game v1.0", SCREEN_WIDTH, SCREEN_HEIGHT);
  std::cout << window.getRefreshRate() << std::endl;

  // SDL_Texture *grassTexture =
  // window.loadTexture("res/gfx/ground_grass_1.png");
  std::vector<RigidBody> objects = {};
  {
    RigidBody rb(Vector2f(100, 100), Vector2f(5, 5), 100, 100);
    objects.push_back(rb);
  }

  bool gameRunning = true;
  SDL_Event event;

  const float timeStep = 0.015f;
  float accumulator = 0.0f;
  float currentTime = utils::hireTimeInSeconds();

  while (gameRunning) {
    int startTicks = SDL_GetTicks();
    float newTime = utils::hireTimeInSeconds();
    float frameTime = newTime - currentTime;
    currentTime = newTime;
    accumulator += frameTime;

    while (accumulator >= timeStep) {
      while (SDL_PollEvent(&event)) {
        if (event.type == SDL_QUIT)
          gameRunning = false;
      }

      accumulator -= timeStep;
    }

    // const float alpha = accumulator / timeStep;

    window.clear();
    for (RigidBody &e : objects) {
      window.render(e);
    }
    window.display();

    int frameTicks = SDL_GetTicks() - startTicks;
    if (frameTicks < 1000 / window.getRefreshRate())
      SDL_Delay(1000 / window.getRefreshRate());
  }

  window.cleanUp();
  SDL_Quit();

  return 0;
}
