package physics

type AABB struct {
	min, max Vec2[float64]
}

type Object struct {
	box  AABB
	pos  Vec2[float64]
	size float64
}

func (a *AABB) collides(b *AABB) *Collision {
	return &Collision{}
}

type Collision struct {
	a      *AABB
	b      *AABB
	normal Vec2[float64]
	depth  float64
}
