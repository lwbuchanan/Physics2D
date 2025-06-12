package physics2d

// Technically, the game could manage all the physiscs objects on
// its own, but its convenient to have the physics world take care
// of updating itself. This means also means we can store global
// physics properties that affect all objects like gravity.
type World struct {
	Bodies     []*Body
	dimensions Vec2
	gravity    float64 // m/s/s
}

func NewWorld(bodies []*Body, dimensions Vec2, gravity float64) World {
	return World{Bodies: bodies, dimensions: dimensions, gravity: gravity}
}

// Call this every physics tick
func (w World) UpdatePhysics(dt float64) {
	for i := 0; i < len(w.Bodies); i++ {
		b1 := w.Bodies[i]

		// Update position
		// b1.Velocity -= w.gravity * dt
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
				if collides, collision := BallAndPolygonCollide(b1, b2); collides {
					collision.Resolve()
				}
			}

		}

		// For now, keep ball in bounds
		if b1.Position.y-b1.Dimensions.y/2 <= 0 {
			b1.Position = (Vec2{b1.Position.x, b1.Dimensions.y / 2})
		}
		if b1.Position.y+b1.Dimensions.y/2 >= w.dimensions.y {
			b1.Position = (Vec2{b1.Position.x, w.dimensions.y - b1.Dimensions.y/2})
		}
		if b1.Position.x-b1.Dimensions.x/2 <= 0 {
			b1.Position = (Vec2{b1.Dimensions.x / 2, b1.Position.y})
		}
		if b1.Position.x+b1.Dimensions.x/2 >= w.dimensions.X() {
			b1.Position = (Vec2{w.dimensions.x - b1.Dimensions.x/2, b1.Position.Y()})
		}
	}
}
