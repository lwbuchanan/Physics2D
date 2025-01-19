package physics

import (
	"math"
)

type AABB struct {
	Min, Max Vec2[float64]
}

type Circle struct {
	Box AABB
	Pos Vec2[float64]
	Rad float64
}

func NewCircle(Pos Vec2[float64], Rad float64) *Circle {
	rvec := Vec2[float64]{Rad, Rad}
	box := AABB{Pos.Sub(rvec), Pos.Add(rvec)}
	return &Circle{box, Pos, Rad}
}

func (o *Circle) Move(newPos Vec2[float64]) {
	disp := o.Pos.Sub(newPos)
	o.Pos = newPos
	o.Box.Min.Add(disp)
	o.Box.Max.Add(disp)
}

func (a *Circle) Collides(b *Circle) (bool, *Collision) {

	if a.Box.Min.GreaterEq(b.Box.Max) || a.Box.Max.LessEq(b.Box.Min) {
		println("Boxes not collided")
		return false, nil
	}
	println("Boxes collided")

	bothRad := a.Rad + b.Rad
	displacement := a.Pos.Sub(b.Pos)
	distSquared := displacement.LengthSquared()

	if distSquared > bothRad {
		return false, nil
	}

	distance := math.Sqrt(float64(distSquared))

	if distance == 0 {
		return true, &Collision{a, b, Vec2[float64]{1.0, 0.0}, a.Rad}
	}

	depth := distance - (a.Rad + b.Rad)
	normal := displacement.ScaleDivide(distance)

	return true, &Collision{a, b, normal, depth}
}

type Collision struct {
	a      *Circle
	b      *Circle
	normal Vec2[float64]
	depth  float64
}
