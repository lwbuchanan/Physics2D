package physics2d

// Normal is in the a->b direction

type Collision struct {
	a      *Circle
	b      *Circle
	normal Vec2
	depth  float32
}

func (c *Collision) Resolve() {
	c.a.Push(c.normal.ScaleMult(-c.depth / 2))
	c.b.Push(c.normal.ScaleMult(c.depth / 2))
}
