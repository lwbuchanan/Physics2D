#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <cstdio>
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
bool RigidBody::collidesWithLWall() { return pos.x < radius; }
bool RigidBody::collidesWithRWall() { return pos.x > SCREEN_WIDTH - radius; }
bool RigidBody::collidesWithRB(RigidBody &rb) {
  return pos.diff(rb.pos).normSq() <
         (radius + rb.radius) * (radius + rb.radius);
}

void RigidBody::checkCollisions(std::vector<RigidBody> &rbs) {
  for (RigidBody &rb : rbs) {

    if (&rb == this) {
      continue;
    }
    if (collidesWithRB(rb)) {
      // Separate the objects

      // Impart new velocity onto this rb
      bounceOff(rb, 1);
      // Do the same to the other rb
      rb.bounceOff(*this, 1);
    }
  }
}

void RigidBody::bounceOff(RigidBody rb, float cR) {
  vel = vel.sum(
      pos.diff(rb.pos).scale((((1 + cR) * rb.mass) / (mass + rb.mass)) *
                             ((vel.diff(rb.vel).dotProd(pos.diff(rb.pos))) /
                              (pos.diff(rb.pos).normSq()))));
}

void RigidBody::updatePhysics(Vector2f force, float dt) {
  if (vel.y < 0 && collidesWithGround()) {
    pos.y = radius;
    vel.y *= -0.6;
  }

  vel.y -= 0.1;

  if (vel.x < 0 && collidesWithLWall()) {
    pos.x = radius;
    vel.x *= -0.2;
  }

  if (vel.x > 0 && collidesWithRWall()) {
    pos.x = SCREEN_WIDTH - radius;
    vel.x *= -0.2;
  }

  // Move body according to current velocity
  pos.x += vel.x / FRAME_TIME;
  pos.y += vel.y / FRAME_TIME;
  // fflush(stdout);
  // printf("x: %f | y: %f | x': %f | y': %f\r", pos.x, pos.y, vel.x, vel.y);
}
