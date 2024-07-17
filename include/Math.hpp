#pragma once
#include <cmath>
#include <iostream>

struct Vector2f {
  Vector2f() : x(0.0f), y(0.0f) {}
  Vector2f(float p_x, float p_y) : x(p_x), y(p_y) {}

  Vector2f diff(Vector2f vec) { return Vector2f(x - vec.x, y - vec.y); }
  Vector2f sum(Vector2f vec) { return Vector2f(x + vec.x, y + vec.y); }
  Vector2f scale(float scalar) { return Vector2f(scalar * x, scalar * y); }

  float dotProd(Vector2f vec) { return x * vec.x + y * vec.y; }
  float normSq() { return x * x + y * y; }

  float norm() {
    std::cout << "IM STUPID DONT USE SQUARE ROOT\n";
    return sqrt(x * x + y * y);
  }

  void print() { std::cout << x << ", " << y << std::endl; }

  float x, y;
};
