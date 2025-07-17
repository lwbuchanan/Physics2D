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
	Bodies          []*Body
	dimensions      Vec2
	gravity         float64 // m/s/s
	timeSteps       int
	collisionBuffer []*Collision
	CollisionEvents []*Collision
	Paused          bool
}

func NewWorld(bodies []*Body, dimensions Vec2, gravity float64, timeSteps int) World {
	return World{
		Bodies:          bodies,
		dimensions:      dimensions,
		gravity:         gravity,
		timeSteps:       timeSteps,
		collisionBuffer: make([]*Collision, len(bodies)),
		CollisionEvents: make([]*Collision, len(bodies)),
		Paused:          false,
	}
}

// Call this every physics tick
func (w *World) UpdatePhysics(dt float64) {
	if w.Paused {
		return
	}
	w.CollisionEvents = w.CollisionEvents[:0]
	for range w.timeSteps {
		for i, b1 := range w.Bodies {

			// Accelerate due to gravity
			if b1.inverseMass > 0 {
				b1.Accelerate(NewVec2(0, -w.gravity))
			}

			// Resolve forces acting on body
			b1.Update(dt / float64(w.timeSteps))

			// Check collisions
			w.collisionBuffer = w.collisionBuffer[:0]
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
					w.collisionBuffer = append(w.collisionBuffer, collision)
					w.CollisionEvents = append(w.CollisionEvents, collision)
				}
			}

			// Resolve collisions
			for _, c := range w.collisionBuffer {
				c.Resolve()
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

func (w *World) NumSteps() int {
	return w.timeSteps
}
