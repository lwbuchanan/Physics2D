#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>

#include "../include/Defs.hpp"
#include "../include/RigidBody.hpp"

RigidBody::RigidBody(Vector2f p_pos, Vector2f p_vel, float p_width,
                     float p_height)
    : pos(p_pos), vel(p_vel) {
  currentFrame.x = p_pos.x;
  currentFrame.y = SCREEN_HEIGHT - p_pos.y - p_height;
  currentFrame.w = p_width;
  currentFrame.h = p_height;
}

void RigidBody::updateFrame() {
  currentFrame.y = SCREEN_HEIGHT - pos.y - currentFrame.h;
  currentFrame.x = pos.x;
}

void RigidBody::updatePhysics(Vector2f force, float dt) {
  pos.x += vel.x;
  pos.y += vel.y;
  updateFrame();
}
