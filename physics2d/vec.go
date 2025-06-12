package physics2d

import "math"

type Vec2 struct {
	x, y float64
}

func ZeroVec2() Vec2 {
	return Vec2{0, 0}
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) X() float64 {
	return v.x
}

func (v Vec2) Y() float64 {
	return v.y
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1.x + v2.x, v1.y + v2.y}
}

func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v1.x - v2.x, v1.y - v2.y}
}

func (v1 Vec2) ScaleMult(scalar float64) Vec2 {
	return Vec2{v1.x * scalar, v1.y * scalar}
}

func (v1 Vec2) ScaleDivide(scalar float64) Vec2 {
	return Vec2{v1.x / scalar, v1.y / scalar}
}

func (v1 Vec2) Dot(v2 Vec2) float64 {
	return v1.x*v2.x + v1.y*v2.y
}

// Z component of the cross product
func (v1 Vec2) Cross(v2 Vec2) float64 {
	return v1.x*v2.y - v1.y*v2.x
}

func (v Vec2) Transform(t Transform) Vec2 {
	rx := v.x*t.Cos - v.y*t.Sin
	ry := v.x*t.Sin + v.y*t.Cos
	// Rotate THEN translate
	return Vec2{rx, ry}.Add(t.Pos)
}

func (v1 Vec2) Equal(v2 Vec2) bool {
	return v1.x == v2.x && v1.y == v2.y
}

func (v Vec2) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y
}

// Uses sqrt, try not to use
func (v Vec2) Length() float64 {
	return float64(math.Sqrt(float64(v.LengthSquared())))
}

// Uses sqrt, try not to use
func (v Vec2) Normalize() Vec2 {
	return v.ScaleDivide(v.Length())
}
