#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>

#include "../include/Defs.hpp"
#include "../include/Entity.hpp"

Entity::Entity(Vector2f p_pos, SDL_Texture *p_texture)
    : pos(p_pos), texture(p_texture) {
  currentFrame.x = 0;
  currentFrame.y = 0;
  currentFrame.w = 32;
  currentFrame.h = 32;
}

Entity::Entity(Vector2f p_pos, float p_width, float p_height)
    : pos(p_pos), texture(NULL) {
  currentFrame.x = p_pos.x;
  currentFrame.y = SCREEN_HEIGHT - p_pos.y - p_height;
  currentFrame.w = p_width;
  currentFrame.h = p_height;
}

void Entity::updateFrame() {
  currentFrame.y = SCREEN_HEIGHT - pos.y - currentFrame.h;
  currentFrame.x = pos.x;
}
