package physics2d

import (
	"math"
)

type AABB struct {
	Min, Max Vec2
}

type Circle struct {
	Box AABB
	Pos Vec2
	Rad float32
}

func NewCircle(Pos Vec2, Rad float32) *Circle {
	rvec := Vec2{Rad, Rad}
	box := AABB{Pos.Sub(rvec), Pos.Add(rvec)}
	return &Circle{box, Pos, Rad}
}

func (o *Circle) SetPos(newPos Vec2) {
	disp := newPos.Sub(o.Pos)
	o.Pos = newPos
	o.Box.Min = o.Box.Min.Add(disp)
	o.Box.Max = o.Box.Max.Add(disp)
}

func (o *Circle) Push(displacement Vec2) {
	o.Pos = o.Pos.Add(displacement)
	o.Box.Min = o.Box.Min.Add(displacement)
	o.Box.Max = o.Box.Max.Add(displacement)
}

func (a *Circle) Collides(b *Circle) (bool, *Collision) {

	if (a.Box.Max.X < b.Box.Min.X || a.Box.Min.X > b.Box.Max.X) ||
		(a.Box.Max.Y < b.Box.Min.Y || a.Box.Min.Y > b.Box.Max.Y) {
		// println("Boxes not collided")
		return false, nil
	}

	bothRad := a.Rad + b.Rad
	displacement := a.Pos.Sub(b.Pos)
	distSquared := displacement.LengthSquared()

	if distSquared > (bothRad * bothRad) {
		// println("Boxes collided, no actual collision")
		return false, nil
	}

	distance := float32(math.Sqrt(float64(distSquared)))

	if distance == 0 {
		// println("Boxes collided, circles overlapping perfectly")
		return true, &Collision{a, b, Vec2{1.0, 0.0}, a.Rad}
	}

	depth := distance - (a.Rad + b.Rad)
	normal := displacement.ScaleDivide(distance)

	// println("Boxes collided, circles overlapping")
	return true, &Collision{a, b, normal, depth}
}
