package physics2d

type World struct {
	numBodies  int
	Bodies     []*Body
	dimensions Vec2
}

func NewWorld(numBodies int, bodies []*Body, dimensions Vec2) World {
	return World{numBodies: numBodies, Bodies: bodies, dimensions: dimensions}
}

func (w World) UpdatePhysics(dt float64) {
	for i := 0; i < len(w.Bodies); i++ {
		b1 := w.Bodies[i]

		// Update position
		// b1.Vel.Y -= 2.5
		b1.Update(dt)

		// Push collided balls apart
		for j := i + 1; j < len(w.Bodies); j++ {
			b2 := w.Bodies[j]

			if b1.Shape == Ball && b2.Shape == Ball {
				if collides, collision := BallsCollide(b1, b2); collides {
					collision.Resolve()
				}
			} else if b1.Shape == Box && b2.Shape == Box {
				if collides, collision := BoxesCollide(b1, b2); collides {
					collision.Resolve()
				}
			} else if (b1.Shape == Ball && b2.Shape == Box) || (b1.Shape == Box && b2.Shape == Ball) {
				if collides, collision := BallAndBoxCollide(b1, b2); collides {
					collision.Resolve()
				}
			}

		}

		// For now, keep ball in bounds
		if b1.Position.Y()-b1.Radius <= 0 {
			b1.Position = (Vec2{b1.Position.X(), b1.Radius})
		}
		if b1.Position.Y()+b1.Radius >= w.dimensions.Y() {
			b1.Position = (Vec2{b1.Position.X(), w.dimensions.Y() - b1.Radius})
		}
		if b1.Position.X()-b1.Radius <= 0 {
			b1.Position = (Vec2{b1.Radius, b1.Position.Y()})
		}
		if b1.Position.X()+b1.Radius >= w.dimensions.X() {
			b1.Position = (Vec2{w.dimensions.X() - b1.Radius, b1.Position.Y()})
		}
	}
}
