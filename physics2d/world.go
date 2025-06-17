package physics2d

import (
	"fmt"
	"os"
)

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

		// Accelerate due to gravity
		b1.acceleration.y -= w.gravity

		// Check collisions
		for j := i + 1; j < len(w.Bodies); j++ {
			b2 := w.Bodies[j]

			collision, err := Collide(b1, b2)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			if collision != nil {
				collision.Resolve()
			}
		}

		err := CollideWall(b1, w)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		// Resolve forces acting on body
		b1.Update(dt)
	}
}
