#pragma once

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>

#include "Entity.hpp"
#include "RigidBody.hpp"

class RenderWindow {
public:
  RenderWindow(const char *p_title, int p_w, int p_h);
  SDL_Texture *loadTexture(const char *p_filePath);

  int getRefreshRate();
  void cleanUp();
  void clear();
  void render(RigidBody &p_rigidBody);
  void render(Entity &p_entity);
  void display();

private:
  SDL_Window *window;
  SDL_Renderer *renderer;
};
