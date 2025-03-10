#pragma once

#include "Math.hpp"
#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>

class Entity {
public:
  Entity(Vector2f p_pos, SDL_Texture *p_texture);
  Entity(Vector2f p_pos, float p_width, float p_height);
  Vector2f &getPos() { return pos; }
  SDL_Texture *getTexture() { return texture; }
  SDL_Rect &getCurrentFrame() { return currentFrame; }
  void updateFrame();

private:
  Vector2f pos;
  SDL_Rect currentFrame;
  SDL_Texture *texture;
};
