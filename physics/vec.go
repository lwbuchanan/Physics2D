package physics

type Vec2 struct {
	X, Y float32
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 Vec2) Dot(v2 Vec2) float32 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (v1 Vec2) ScaleMult(scalar float32) Vec2 {
	return Vec2{v1.X * scalar, v1.Y * scalar}
}

func (v1 Vec2) ScaleDivide(scalar float32) Vec2 {
	return Vec2{v1.X / scalar, v1.Y / scalar}
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
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v Vec2) LengthSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

// func (v Vec2) Length() float64 {
// 	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
// }
