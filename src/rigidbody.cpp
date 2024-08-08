#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <cmath>
#include <cstdio>
#include <unistd.h>
#include <vector>

#include "../include/Defs.hpp"
#include "../include/RigidBody.hpp"

RigidBody::RigidBody(Shape p_shape, Vector2f p_pos, Vector2f p_vel, float p_rot,
                     float p_rotVel, float p_radius, float p_length,
                     float p_width, float p_mass, float p_restitution)
    : shape(p_shape), pos(p_pos), vel(p_vel), rot(p_rot), rotVel(p_rotVel),
      radius(p_radius), length(p_length), width(p_width), mass(p_mass),
      restitution(p_restitution) {
  color = {255, 255, 255, 255};
}

void RigidBody::move(Vector2f newPos) {
  pos.x = newPos.x;
  pos.y = newPos.y;
}

void RigidBody::setVel(Vector2f newVel) {
  vel.x = newVel.x;
  vel.y = newVel.y;
}

bool RigidBody::collidesWithGround() { return pos.y < radius; }
bool RigidBody::collidesWithCeiling() { return pos.y > SCREEN_HEIGHT - radius; }
bool RigidBody::collidesWithLWall() { return pos.x < radius; }
bool RigidBody::collidesWithRWall() { return pos.x > SCREEN_WIDTH - radius; }
bool RigidBody::collidesWithRB(RigidBody &rb, Vector2f &normal, float &depth) {
  float dist = (pos - rb.pos).norm();
  float radii = (radius + rb.radius);
  normal = Vector2f();
  depth = 0.0f;
  if (dist >= radii)
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
      pos = pos - (normal * (depth / 2.0f));
      rb.pos = rb.pos + (normal * (depth / 2.0f));
    }
  }
}

void RigidBody::updatePhysics() {
  if (vel.y <= 0 && collidesWithGround()) {
    pos.y = radius;
    vel.y *= -1;
  }

  if (vel.y >= 0 && collidesWithCeiling()) {
    pos.y = SCREEN_HEIGHT - radius;
    vel.y *= -1;
  }

  if (vel.x <= 0 && collidesWithLWall()) {
    pos.x = radius;
    vel.x *= -1;
  }

  if (vel.x >= 0 && collidesWithRWall()) {
    pos.x = SCREEN_WIDTH - radius;
    vel.x *= -1;
  }

  // Move body according to current velocity
  pos.x += vel.x / FRAME_TIME;
  pos.y += vel.y / FRAME_TIME;
  // fflush(stdout);
  // printf("x: %f | y: %f | x': %f | y': %f\r", pos.x, pos.y, vel.x, vel.y);
}
