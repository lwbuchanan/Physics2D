package physics2d

type Collision struct {
	a      *Circle
	b      *Circle
	Normal Vec2
	Depth  float32
}

func (c *Collision) Resolve() {
	c.a.Push(c.Normal.ScaleMult(-c.Depth / 2))
	c.b.Push(c.Normal.ScaleMult(c.Depth / 2))
}
