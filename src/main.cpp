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

  RenderWindow window("Physics2D", SCREEN_WIDTH, SCREEN_HEIGHT);
  std::cout << window.getRefreshRate() << std::endl;

  // SDL_Texture *grassTexture =
  // window.loadTexture("res/gfx/ground_grass_1.png");
  std::vector<RigidBody> objects = {};
  {
    RigidBody rb(Vector2f(300, 500), Vector2f(20, 0), 50, 500);
    RigidBody rb2(Vector2f(600, 500), Vector2f(-10, 0), 50, 500);
    objects.push_back(rb);
    objects.push_back(rb2);
  }

  bool gameRunning = true;
  SDL_Event event;
  const float frameTime = FRAME_TIME;

  while (gameRunning) {
    float startTime = utils::hireTimeInSeconds();
    // int startTicks = SDL_GetTicks();

    // Check to see if the game is over
    while (SDL_PollEvent(&event)) {
      if (event.type == SDL_QUIT)
        gameRunning = false;
    }

    for (RigidBody &e : objects) {
      e.updatePhysics(Vector2f(0, 0), frameTime);
      e.checkCollisions(objects);
    }

    // Draw Everything
    window.clear();
    for (RigidBody &e : objects) {
      window.render(e);
    }
    window.display();

    SDL_Delay(startTime + frameTime - utils::hireTimeInSeconds());
  }

  window.cleanUp();
  SDL_Quit();
  return 0;
}
