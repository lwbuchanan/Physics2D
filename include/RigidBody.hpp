#pragma once

#include "Defs.hpp"
#include "Math.hpp"
#include <SDL2/SDL.h>
#include <vector>

enum Shape { CIRCLE, SQUARE };

class RigidBody {
public:
  RigidBody(Shape p_shape, Vector2f p_pos, Vector2f p_vel, float p_rot,
            float p_rotVel, float p_radius, float p_length, float p_width,
            float p_mass, float p_restitution);
  Vector2f &getPos() { return pos; }
  Vector2f &getVel() { return vel; }
  float getRadius() { return radius; }
  float getMass() { return mass; }
  void move(Vector2f newPos);
  void setVel(Vector2f newVel);
  void updatePhysics();
  void checkCollisions(std::vector<RigidBody> &rbs);
  void bounceOff(RigidBody rb, float cr);
  bool collidesWithGround();
  bool collidesWithCeiling();
  bool collidesWithLWall();
  bool collidesWithRWall();
  bool collidesWithRB(RigidBody &rb, Vector2f &normal, float &depth);
  std::vector<Uint8> color;

private:
  Shape shape;
  Vector2f pos;
  Vector2f vel;
  float rot;
  float rotVel;
  float radius;
  float length;
  float width;
  float mass;
  float restitution;
};
