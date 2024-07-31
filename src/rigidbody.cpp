#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <cmath>
#include <cstdio>
#include <unistd.h>
#include <vector>

#include "../include/Defs.hpp"
#include "../include/RigidBody.hpp"

RigidBody::RigidBody(Vector2f p_pos, Vector2f p_vel, float p_radius,
                     float p_mass)
    : pos(p_pos), vel(p_vel), radius(p_radius), mass(p_mass) {}

// void RigidBody::updateFrame() {
//   currentFrame.y = SCREEN_HEIGHT - pos.y - currentFrame.h;
//   currentFrame.x = pos.x;
// }

bool RigidBody::collidesWithGround() { return pos.y < radius; }
bool RigidBody::collidesWithCeiling() { return pos.y > SCREEN_HEIGHT - radius; }
bool RigidBody::collidesWithLWall() { return pos.x < radius; }
bool RigidBody::collidesWithRWall() { return pos.x > SCREEN_WIDTH - radius; }
bool RigidBody::collidesWithRB(RigidBody &rb, Vector2f &normal, float &depth) {
  float dist = (pos - rb.pos).norm();
  float radii = (radius + rb.radius);
  normal = Vector2f();
  depth = 0.0f;
  if (radii >= dist)
    return false;
  normal = (rb.pos - pos).normalize();
  depth = radii - dist;
  return true;
}

void RigidBody::checkCollisions(std::vector<RigidBody> &rbs) {
  for (RigidBody &rb : rbs) {

    if (&rb == this) {
      continue;
    }
    Vector2f normal;
    float depth;
    if (collidesWithRB(rb, normal, depth)) {
      float restitution = 0.5f;
      pos = pos - (normal * (depth / 2.0f));
      rb.pos = rb.pos + (normal * (depth / 2.0f));
    }
  }
}

void RigidBody::updatePhysics(Vector2f force, float dt) {
  if (vel.y < 0 && collidesWithGround()) {
    pos.y = radius;
    vel.y *= -1;
  }

  if (vel.y > 0 && collidesWithCeiling()) {
    pos.y = SCREEN_HEIGHT - radius;
    vel.y *= -1;
  }

  if (vel.x < 0 && collidesWithLWall()) {
    pos.x = radius;
    vel.x *= -1;
  }

  if (vel.x > 0 && collidesWithRWall()) {
    pos.x = SCREEN_WIDTH - radius;
    vel.x *= -1;
  }

  // Move body according to current velocity
  pos.x += vel.x / FRAME_TIME;
  pos.y += vel.y / FRAME_TIME;
  // fflush(stdout);
  // printf("x: %f | y: %f | x': %f | y': %f\r", pos.x, pos.y, vel.x, vel.y);
}
