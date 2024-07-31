#pragma once
#include <cmath>
#include <iostream>

struct Vector2f {
  Vector2f() : x(0.0f), y(0.0f) {}
  Vector2f(float p_x, float p_y) : x(p_x), y(p_y) {}

  Vector2f operator-(Vector2f vec) { return Vector2f(x - vec.x, y - vec.y); }
  Vector2f operator+(Vector2f vec) { return Vector2f(x + vec.x, y + vec.y); }
  Vector2f operator*(float scalar) { return Vector2f(scalar * x, scalar * y); }
  Vector2f operator/(float scalar) { return Vector2f(x / scalar, y / scalar); }

  float norm() { return sqrt(x * x + y * y); }
  float normSquared() { return x * x + y * y; }
  Vector2f normalize() { return *this / this->norm(); }

  void print() { std::cout << x << ", " << y << std::endl; }

  float x, y;
};

class Math {
public:
  static float dotProd(Vector2f v1, Vector2f v2) {
    return v1.x * v2.x + v1.y * v2.y;
  }
};
