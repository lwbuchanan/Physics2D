#include <SDL2/SDL.h>
#include <SDL2/SDL_events.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_keycode.h>
#include <SDL2/SDL_mouse.h>
#include <SDL2/SDL_timer.h>
#include <iostream>
#include <vector>

#include "../include/Defs.hpp"
#include "../include/RenderWindow.hpp"
#include "../include/RigidBody.hpp"
#include "../include/Utils.hpp"

void populateWorld(int numObjects, std::vector<RigidBody> &objects) {
  for (int i = 0; i < numObjects; i++) {
    // srand(12);
    float x = std::rand() % SCREEN_WIDTH;
    float y = std::rand() % SCREEN_HEIGHT;
    objects.push_back(RigidBody(CIRCLE, Vector2f(x, y), Vector2f(0, 0), 0.0f,
                                0.0f, 30.0f, 0.0f, 0.0f, 5.0f, 1.0f));
  }
}

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
  populateWorld(10, objects);
  RigidBody playerObj = RigidBody(CIRCLE, Vector2f(500, 500), Vector2f(-5, 5),
                                  0.0f, 0.0f, 30.0f, 0.0f, 0.0f, 5.0f, 1.0f);
  playerObj.color = {255, 255, 0, 255};
  objects.push_back(playerObj);
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
      if (event.type == SDL_KEYDOWN) {
        switch (event.key.keysym.sym) {
        case SDLK_RIGHT:
          playerObj.setVel(Vector2f(10, 0));
          std::cout << "right\n";
          break;
        case SDLK_UP:
          playerObj.setVel(Vector2f(0, 10));
          std::cout << "up\n";
          break;
        case SDLK_LEFT:
          playerObj.setVel(Vector2f(-10, 0));
          std::cout << "left\n";
          break;
        case SDLK_DOWN:
          playerObj.setVel(Vector2f(0, -10));
          std::cout << "down\n";
          break;
        default:
          break;
        }
      }
      // if (event.type == SDL_KEYUP) {
      //   playerObj.setVel(Vector2f(0, 0));
      // }
    }

    for (RigidBody &e : objects) {
      e.updatePhysics();
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
