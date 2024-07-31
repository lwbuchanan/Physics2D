#pragma once

#include "Defs.hpp"
#include "Math.hpp"
#include <SDL2/SDL.h>
#include <vector>

class RigidBody {
public:
  RigidBody(Vector2f p_pos, Vector2f p_vel, float p_radius, float p_mass);
  Vector2f &getPos() { return pos; }
  Vector2f &getVel() { return vel; }
  float getRadius() { return radius; }
  float getMass() { return mass; }
  // SDL_Rect &getCurrentFrame() { return currentFrame; }
  // void updateFrame();
  void updatePhysics(Vector2f force, float dt);
  void checkCollisions(std::vector<RigidBody> &rbs);
  void bounceOff(RigidBody rb, float cr);
  bool collidesWithGround();
  bool collidesWithCeiling();
  bool collidesWithLWall();
  bool collidesWithRWall();
  bool collidesWithRB(RigidBody &rb, Vector2f &normal, float &depth);

private:
  Vector2f pos;
  Vector2f vel;
  float mass;
  float restitution;
  float radius;
};
