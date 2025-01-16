package physics

import "math"

type AABB struct {
	min, max Vec2[float64]
}

type Circle struct {
	box AABB
	Pos Vec2[float64]
	Rad float64
}

func NewCircle(Pos Vec2[float64], Rad float64) *Circle {
	rvec := Vec2[float64]{Rad, Rad}
	box := AABB{Pos.Sub(rvec), Pos.Add(rvec)}
	return &Circle{box, Pos, Rad}
}

func (o *Circle) Move(newPos Vec2[float64]) {
	o.Pos = newPos
}

func (a *Circle) collides(b *Circle) (bool, *Collision) {

	if a.box.min.Greater(b.box.max) || a.box.max.Less(b.box.min) {
		return false, nil
	}

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
