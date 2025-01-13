package physics

type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

type Vec2[T Number] struct {
	X, Y T
}

func (v1 Vec2[T]) Add(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 Vec2[T]) Sub(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 Vec2[T]) Dot(v2 Vec2[T]) T {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (v1 Vec2[T]) Scale(scalar T) Vec2[T] {
	return Vec2[T]{v1.X * scalar, v1.Y * scalar}
}

func (v1 Vec2[T]) Greater(v2 Vec2[T]) bool {
	return v1.X > v2.X && v1.Y > v2.Y
}

func (v1 Vec2[T]) GreaterEq(v2 Vec2[T]) bool {
	return v1.X >= v2.X && v1.Y >= v2.Y
}

func (v1 Vec2[T]) Less(v2 Vec2[T]) bool {
	return v1.X < v2.X && v1.Y < v2.Y
}

func (v1 Vec2[T]) LessEq(v2 Vec2[T]) bool {
	return v1.X <= v2.X && v1.Y <= v2.Y
}

func (v1 Vec2[T]) Equal(v2 Vec2[T]) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}
