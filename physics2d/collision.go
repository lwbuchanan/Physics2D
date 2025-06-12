package physics2d

import "math"

// Normal is in the a->b direction

type Collision struct {
	a      *Body
	b      *Body
	normal Vec2
	depth  float64
}

func (c *Collision) Resolve() {
	c.a.Move(c.normal.ScaleMult(-c.depth / 2))
	c.b.Move(c.normal.ScaleMult(c.depth / 2))
}

func BallsCollide(a, b *Body) (bool, *Collision) {
	bothRad := a.Radius + b.Radius
	displacement := a.Position.Sub(b.Position)
	distSquared := displacement.LengthSquared()

	if distSquared > (bothRad * bothRad) {
		// println("Boxes collided, no actual collision")
		return false, nil
	}

	distance := math.Sqrt(distSquared)

	if distance == 0 {
		// println("Boxes collided, Bodys overlapping perfectly")
		return true, &Collision{a, b, Vec2{1.0, 0.0}, a.Radius}
	}

	depth := distance - (a.Radius + b.Radius)
	normal := displacement.ScaleDivide(distance)

	// println("Boxes collided, Bodys overlapping")
	return true, &Collision{a, b, normal, depth}
}

func BoxesCollide(a, b *Body) (bool, *Collision) {
	return false, nil
}

func BallAndBoxCollide(a, b *Body) (bool, *Collision) {
	return false, nil
}
