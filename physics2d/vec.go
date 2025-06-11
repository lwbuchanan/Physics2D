package physics2d

import "math"

type Vec2 struct {
	x, y float32
}

func NewVec2(x, y float32) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) X() float32 {
	return v.x
}

func (v Vec2) Y() float32 {
	return v.y
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1.x + v2.x, v1.y + v2.y}
}

func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v1.x - v2.x, v1.y - v2.y}
}

func (v1 Vec2) Dot(v2 Vec2) float32 {
	return v1.x*v2.x + v1.y*v2.y
}

// Z component of the cross product
func (v1 Vec2) Cross(v2 Vec2) float32 {
	return v1.x*v2.y - v1.y*v2.x
}

func (v1 Vec2) ScaleMult(scalar float32) Vec2 {
	return Vec2{v1.x * scalar, v1.y * scalar}
}

func (v1 Vec2) ScaleDivide(scalar float32) Vec2 {
	return Vec2{v1.x / scalar, v1.y / scalar}
}

// func (v1 Vec2) Greater(v2 Vec2) bool {
// 	return v1.X > v2.X && v1.Y > v2.Y
// }

// func (v1 Vec2) GreaterEq(v2 Vec2) bool {
// 	return v1.X >= v2.X && v1.Y >= v2.Y
// }

// func (v1 Vec2) Less(v2 Vec2) bool {
// 	return v1.X < v2.X && v1.Y < v2.Y
// }

// func (v1 Vec2) LessEq(v2 Vec2) bool {
// 	return v1.X <= v2.X && v1.Y <= v2.Y
// }

func (v1 Vec2) Equal(v2 Vec2) bool {
	return v1.x == v2.x && v1.y == v2.y
}

func (v Vec2) LengthSquared() float32 {
	return v.x*v.x + v.y*v.y
}

func (v Vec2) Length() float32 {
	return float32(math.Sqrt(float64(v.LengthSquared())))
}

func (v Vec2) Normalize() Vec2 {
	return v.ScaleDivide(v.Length())
}
