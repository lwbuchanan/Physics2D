package physics2d

import (
	"math"
)

// Normal is in the a->b direction

type Collision struct {
	a      *Body
	b      *Body
	normal Vec2
	depth  float64
}

// This should work to resolve any kind of collision
// as long as a normal and depth are produced
func (c *Collision) Resolve() {
	c.a.Move(c.normal.ScaleMult(-c.depth / 2))
	c.b.Move(c.normal.ScaleMult(c.depth / 2))
}

func projectVertecies(vertices []Vec2, axis Vec2) (float64, float64) {
	min := math.MaxFloat64
	max := -math.MaxFloat64

	for _, v := range vertices {
		proj := v.Dot(axis)
		min = math.Min(min, proj)
		max = math.Max(max, proj)
	}

	return min, max
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
	// Parallel axes can be skipped
	return PolygonsCollide(a, b)
}

func PolygonsCollide(a, b *Body) (bool, *Collision) {
	normal := ZeroVec2()
	depth := math.MaxFloat64

	aVertices := a.Vertices()
	bVertices := b.Vertices()

	// Vertecies are stored clockwise, so we test edges clockwise
	for i := range len(aVertices) {
		vCurr := aVertices[i]
		vNext := aVertices[(i+1)%len(aVertices)]
		edge := vNext.Sub(vCurr)
		axis := NewVec2(-edge.y, edge.x) /* Orthoganal to tested edge */

		aMin, aMax := projectVertecies(aVertices, axis)
		bMin, bMax := projectVertecies(bVertices, axis)

		if aMin >= bMax || bMin >= aMax {
			// Found separating axis
			return false, nil
		}

		axisDepth := math.Min(aMax-bMin, bMax-aMin)
		if axisDepth < depth {
			depth = axisDepth
			normal = axis
		}
	}

	for i := range len(bVertices) {
		vCurr := bVertices[i]
		vNext := bVertices[(i+1)%len(bVertices)]
		edge := vNext.Sub(vCurr)
		axis := NewVec2(-edge.y, edge.x) /* Orthoganal to tested edge */

		aMin, aMax := projectVertecies(aVertices, axis)
		bMin, bMax := projectVertecies(bVertices, axis)

		if aMin >= bMax || bMin >= aMax {
			// Found separating axis
			return false, nil
		}

		axisDepth := math.Min(aMax-bMin, bMax-aMin)
		if axisDepth < depth {
			depth = axisDepth
			normal = axis
		}
	}

	depth /= normal.Length()
	normal = normal.Normalize()
	return true, &Collision{a, b, normal, depth}
}

func BallAndPolygonCollide(a, b *Body) (bool, *Collision) {
	return false, nil
}
