package physics2d

type World struct {
	numCirlces int
	Circles    []*Circle
	dimensions Vec2
}

func NewWorld(numCircles int, circles []*Circle, dimensions Vec2) World {
	return World{numCirlces: numCircles, Circles: circles, dimensions: dimensions}
}

func (w World) UpdatePhysics(dt float32) {
	for i := 0; i < len(w.Circles); i++ {
		c1 := w.Circles[i]

		// Update position
		// c1.Vel.Y -= 2.5
		c1.UpdatePosition(dt)

		// Push collided balls apart
		for j := i + 1; j < len(w.Circles); j++ {
			c2 := w.Circles[j]

			if collides, collision := c1.Collides(c2); collides {
				collision.Resolve()
			}
		}

		// For now, keep ball in bounds
		if c1.Pos.Y()-c1.Rad <= 0 {
			c1.SetPos(Vec2{c1.Pos.X(), c1.Rad})
		}
		if c1.Pos.Y()+c1.Rad >= w.dimensions.Y() {
			c1.SetPos(Vec2{c1.Pos.X(), w.dimensions.Y() - c1.Rad})
		}
		if c1.Pos.X()-c1.Rad <= 0 {
			c1.SetPos(Vec2{c1.Rad, c1.Pos.Y()})
		}
		if c1.Pos.X()+c1.Rad >= w.dimensions.X() {
			c1.SetPos(Vec2{w.dimensions.X() - c1.Rad, c1.Pos.Y()})
		}
	}
}
