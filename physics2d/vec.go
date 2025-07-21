package physics2d

import (
	"math"
	"slices"
)

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

func (v Vec2) Transform(t transform) Vec2 {
	rx := v.x*t.Cos - v.y*t.Sin
	ry := v.x*t.Sin + v.y*t.Cos
	// Rotate THEN translate
	return Vec2{rx, ry}.Add(t.Pos)
}

// True if v1 and v2 are within 0.00025
func (v1 Vec2) CloseTo(v2 Vec2) bool {
	return v1.DistanceSquared(v2) < 0.00025
}

func (v Vec2) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y
}

func (v1 Vec2) DistanceSquared(v2 Vec2) float64 {
	xDist := v2.x - v1.x
	yDist := v2.y - v1.y
	return xDist*xDist + yDist*yDist
}

// Uses sqrt, use DistanceSquared if possible
func (v1 Vec2) Distance(v2 Vec2) float64 {
	return math.Sqrt(v1.DistanceSquared(v2))
}

// Uses sqrt, use LengthSquared if possible
func (v Vec2) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// Uses sqrt, avoid if possible
func (v Vec2) Normalize() Vec2 {
	return v.ScaleDivide(v.Length())
}

func (v Vec2) Perpendicular() Vec2 {
	return NewVec2(-v.y, v.x)
}

func Midpoint(v1, v2 Vec2) Vec2 {
	return NewVec2((v1.x+v2.x)/2, (v1.y+v2.y)/2)
}

func cmpX(a, b Vec2) int {
	if a.x > b.x {
		return 1
	}
	if a.x < b.x {
		return -1
	}
	return 0
}

func cmpY(a, b Vec2) int {
	if a.y > b.y {
		return 1
	}
	if a.y < b.y {
		return -1
	}
	return 0
}

func MinX(vecs []Vec2) float64 {
	return slices.MinFunc(vecs, cmpX).x
}

func MinY(vecs []Vec2) float64 {
	return slices.MinFunc(vecs, cmpY).y
}

func MaxX(vecs []Vec2) float64 {
	return slices.MaxFunc(vecs, cmpX).x
}

func MaxY(vecs []Vec2) float64 {
	return slices.MaxFunc(vecs, cmpY).y
}
