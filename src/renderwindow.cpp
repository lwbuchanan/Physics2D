#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <cstdint>
#include <iostream>

#include "../include/Defs.hpp"
#include "../include/Entity.hpp"
#include "../include/RenderWindow.hpp"
#include "../include/RigidBody.hpp"

RenderWindow::RenderWindow(const char *p_title, int p_w, int p_h)
    : window(NULL), renderer(NULL) {
  window =
      SDL_CreateWindow(p_title, SDL_WINDOWPOS_UNDEFINED,
                       SDL_WINDOWPOS_UNDEFINED, p_w, p_h, SDL_WINDOW_SHOWN);
  if (window == NULL)
    std::cout << "Failed to init window: " << SDL_GetError() << std::endl;

  renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED);
}

SDL_Texture *RenderWindow::loadTexture(const char *p_filePath) {
  SDL_Texture *texture = NULL;
  texture = IMG_LoadTexture(renderer, p_filePath);

  if (texture == NULL)
    std::cout << "Failed to load texture: " << SDL_GetError() << std::endl;

  return texture;
}

void RenderWindow::clear() {
  SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255);
  SDL_RenderClear(renderer);
}

void RenderDrawCircle(SDL_Renderer *renderer, Vector2f pos, int r) {
  int32_t centreX = pos.x;
  int32_t centreY = SCREEN_HEIGHT - pos.y;
  const int32_t diameter = (r * 2);
  int32_t x = (r - 1);
  int32_t y = 0;
  int32_t tx = 1;
  int32_t ty = 1;
  int32_t error = (tx - diameter);

  while (x >= y) {
    //  Each of the following renders an octant of the circle
    SDL_RenderDrawPoint(renderer, centreX + x, centreY - y);
    SDL_RenderDrawPoint(renderer, centreX + x, centreY + y);
    SDL_RenderDrawPoint(renderer, centreX - x, centreY - y);
    SDL_RenderDrawPoint(renderer, centreX - x, centreY + y);
    SDL_RenderDrawPoint(renderer, centreX + y, centreY - x);
    SDL_RenderDrawPoint(renderer, centreX + y, centreY + x);
    SDL_RenderDrawPoint(renderer, centreX - y, centreY - x);
    SDL_RenderDrawPoint(renderer, centreX - y, centreY + x);

    if (error <= 0) {
      ++y;
      error += ty;
      ty += 2;
    }

    if (error > 0) {
      --x;
      tx += 2;
      error += (tx - diameter);
    }
  }
}

void RenderWindow::render(RigidBody &p_rigidBody) {
  SDL_SetRenderDrawColor(renderer, 255, 255, 255, 255);
  RenderDrawCircle(renderer, p_rigidBody.getPos(), p_rigidBody.getRadius());
}

void RenderWindow::render(Entity &p_entity) {
  if (p_entity.getTexture() == NULL) {
    SDL_SetRenderDrawColor(renderer, 255, 255, 255, 255);
    SDL_RenderDrawRect(renderer, &p_entity.getCurrentFrame());

  } else {
    SDL_Rect src;
    src.x = p_entity.getCurrentFrame().x;
    src.y = p_entity.getCurrentFrame().y;
    src.w = p_entity.getCurrentFrame().w;
    src.h = p_entity.getCurrentFrame().h;

    SDL_Rect dst;
    dst.x = p_entity.getPos().x * 4;
    dst.y = p_entity.getPos().y * 4;
    dst.w = p_entity.getCurrentFrame().w * 4;
    dst.h = p_entity.getCurrentFrame().h * 4;

    SDL_RenderCopy(renderer, p_entity.getTexture(), &src, &dst);
  }
}

void RenderWindow::display() { SDL_RenderPresent(renderer); }

void RenderWindow::cleanUp() { SDL_DestroyWindow(window); }

int RenderWindow::getRefreshRate() {
  int displayIndex = SDL_GetWindowDisplayIndex(window);

  SDL_DisplayMode mode;

  SDL_GetDisplayMode(displayIndex, 0, &mode);

  return mode.refresh_rate;
}
