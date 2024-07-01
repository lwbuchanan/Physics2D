#pragma once

#include "Math.hpp"
#include <SDL2/SDL.h>

class RigidBody {
public:
  RigidBody(Vector2f p_pos, Vector2f p_vel, float p_width, float p_height);
  Vector2f &getPos() { return pos; }
  Vector2f &getVel() { return vel; }
  SDL_Rect &getCurrentFrame() { return currentFrame; }
  void updateFrame();
  void updatePhysics(Vector2f force, float dt);

private:
  Vector2f pos;
  Vector2f vel;
  SDL_Rect currentFrame;
};
