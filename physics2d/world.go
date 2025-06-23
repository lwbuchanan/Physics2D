package physics2d

import (
	"fmt"
	"os"
	"slices"
)

// Technically, the game could manage all the physiscs objects on
// its own, but its convenient to have the physics world take care
// of updating itself. This means also means we can store global
// physics properties that affect all objects like gravity.
type World struct {
	Bodies     []*Body
	dimensions Vec2
	gravity    float64 // m/s/s
	timeSteps  int
}

func NewWorld(bodies []*Body, dimensions Vec2, gravity float64, timeSteps int) World {
	return World{Bodies: bodies, dimensions: dimensions, gravity: gravity, timeSteps: timeSteps}
}

// Call this every physics tick
func (w *World) UpdatePhysics(dt float64) {
	for range w.timeSteps {
	for i, b1 := range w.Bodies {

		// Accelerate due to gravity
		if b1.inverseMass > 0 {
			b1.Accelerate(NewVec2(0, -w.gravity))
		}

		// Resolve forces acting on body
		b1.Update(dt/float64(w.timeSteps))

		// Check collisions
		for j := i + 1; j < len(w.Bodies); j++ {
			b2 := w.Bodies[j]
			if b1.inverseMass+b2.inverseMass == 0 {
				continue
			}
			collision, err := Collide(b1, b2)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			if collision != nil {
				collision.Resolve()
			}
		}
	}
	}
}

func (w *World) AddBody(body *Body) {
	w.Bodies = append(w.Bodies, body)
}

func (w *World) DeleteBody(bodyIdx int) {
	w.Bodies = slices.Delete(w.Bodies, bodyIdx, bodyIdx+1)
}
