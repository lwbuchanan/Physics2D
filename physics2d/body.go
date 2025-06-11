package physics2d

import (
	"math"
)

type Body interface {
	GetPos() Vec2
	SetPos(pos Vec2)
	GetVel() Vec2
	SetVel(vel Vec2)
	GetRest() float32
	GetInvMass() float32

	UpdatePos(dt float32)
	Push(displacement Vec2)
	Collides(b Body)
}

// Units are m, m/s, and kg
type Circle struct {
	Rad     float32
	Pos     Vec2
	Vel     Vec2
	Rest    float32
	invMass float32
}

// Units are m, m/s, and kg
func NewCircle(pos Vec2, rad float32, mass float32) *Circle {
	return &Circle{
		Rad:     rad,
		Pos:     pos,
		Vel:     Vec2{0, 0},
		Rest:    1.0,
		invMass: 1.0 / mass,
	}
}

func (o *Circle) GetPos() Vec2 {
	return o.Pos
}
func (o *Circle) SetPos(newPos Vec2) {
	o.Pos = newPos
}

// Move by vel*dt meters
func (o *Circle) UpdatePosition(dt float32) {
	displacement := o.Vel.ScaleMult(dt)
	o.Push(displacement)
}

func (o *Circle) Push(displacement Vec2) {
	o.Pos = o.Pos.Add(displacement)
}

func (a *Circle) Collides(b *Circle) (bool, *Collision) {
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
