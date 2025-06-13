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

// Gets the min and max of all points projected onto the axis
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

// Gets the min and max points of the circle edge projected onto the axis
func projectCircle(position Vec2, radius float64, axis Vec2) (float64, float64) {
	centerProj := position.Dot(axis)
	min := centerProj - radius
	max := centerProj + radius

	return min, max
}

// Gets the index into the slice of the closest vertex to the given point
func closestVertexIdx(position Vec2, vertices []Vec2) int {
	closestIndex := -1
	minDistSquared := math.MaxFloat64
	for i, v := range vertices {
		distSquared := position.DistanceSquared(v)
		if distSquared < minDistSquared {
			minDistSquared = distSquared
			closestIndex = i
		}
	}

	return closestIndex
}

func BallsCollide(a, b *Body) (bool, *Collision) {
	bothRad := a.Radius + b.Radius
	displacement := a.Position.Sub(b.Position)
	distSquared := displacement.LengthSquared()

	if distSquared > (bothRad * bothRad) {
		return false, nil
	}

	distance := math.Sqrt(distSquared)

	if distance == 0 {
		return true, &Collision{a, b, Vec2{1.0, 0.0}, a.Radius}
	}

	depth := distance - (a.Radius + b.Radius)
	normal := displacement.ScaleDivide(distance)

	return true, &Collision{a, b, normal, depth}
}

func BoxesCollide(a, b *Body) (bool, *Collision) {
	// TODO: Parallel axes can be skipped
	return PolygonsCollide(a, b)
}

// SAT only works for convex polygons
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

		// The Two-Bit Coding videos I've been following say we need to normalize the axis.
		// Not sure exactly sure why... I will come back here and maybe try to optomize this.
		axis := NewVec2(-edge.y, edge.x).Normalize() /* Orthoganal to tested edge */

		aMin, aMax := projectVertecies(aVertices, axis)
		bMin, bMax := projectVertecies(bVertices, axis)

		if aMin >= bMax || bMin >= aMax {
			// Found separating axis
			return false, nil
		}

		// We need to eep track of the minimum non-separating axis, if a and b can't
		// be separated, the minimum depth collision will determine the normal.

		// The normal needs to point away from a, so if b is closer to the origin, the
		// normal will be negative. We will compare depth magnitude with abs().
		aPositiveDepth := aMax - bMin
		aNegativeDepth := aMin - bMax

		var axisDepth float64
		if aPositiveDepth < math.Abs(aNegativeDepth) {
			axisDepth = aPositiveDepth
		} else {
			axisDepth = aNegativeDepth
		}

		if math.Abs(axisDepth) < math.Abs(depth) {
			depth = axisDepth
			normal = axis
		}
	}

	// This part is the same, except we check all of b's axes instead
	for i := range len(bVertices) {
		vCurr := bVertices[i]
		vNext := bVertices[(i+1)%len(bVertices)]
		edge := vNext.Sub(vCurr)
		axis := NewVec2(-edge.y, edge.x).Normalize()

		aMin, aMax := projectVertecies(aVertices, axis)
		bMin, bMax := projectVertecies(bVertices, axis)

		if aMin >= bMax || bMin >= aMax {
			return false, nil
		}

		aPositiveDepth := aMax - bMin
		aNegativeDepth := aMin - bMax

		var axisDepth float64
		if aPositiveDepth < math.Abs(aNegativeDepth) {
			axisDepth = aPositiveDepth
		} else {
			axisDepth = aNegativeDepth
		}

		if math.Abs(axisDepth) < math.Abs(depth) {
			depth = axisDepth
			normal = axis
		}

	}

	return true, &Collision{a, b, normal, depth}
}

func BallAndPolygonCollide(ball, polygon *Body) (bool, *Collision) {
	vertices := polygon.Vertices()

	normal := ZeroVec2()
	depth := math.MaxFloat64

	for i := range len(vertices) {
		vCurr := vertices[i]
		vNext := vertices[(i+1)%len(vertices)]
		edge := vNext.Sub(vCurr)
		axis := NewVec2(-edge.y, edge.x).Normalize()

		pMin, pMax := projectVertecies(vertices, axis)
		bMin, bMax := projectCircle(ball.Position, ball.Radius, axis)

		if pMin >= bMax || bMin >= pMax {
			// Found separating axis
			return false, nil
		}

		bPositiveDepth := bMax - pMin
		bNegativeDepth := bMin - pMax

		var axisDepth float64
		if bPositiveDepth < math.Abs(bNegativeDepth) {
			axisDepth = bPositiveDepth
		} else {
			axisDepth = bNegativeDepth
		}

		if math.Abs(axisDepth) < math.Abs(depth) {
			depth = axisDepth
			normal = axis
		}
	}

	// Now we check for a SA between the circles edge to closest vertex
	closestVertex := vertices[closestVertexIdx(ball.Position, vertices)]
	axis := closestVertex.Sub(ball.Position).Normalize()

	pMin, pMax := projectVertecies(vertices, axis)
	bMin, bMax := projectCircle(ball.Position, ball.Radius, axis)

	if pMin >= bMax || bMin >= pMax {
		// Found separating axis
		return false, nil
	}

	bPositiveDepth := bMax - pMin
	bNegativeDepth := bMin - pMax

	var axisDepth float64
	if bPositiveDepth < math.Abs(bNegativeDepth) {
		axisDepth = bPositiveDepth
	} else {
		axisDepth = bNegativeDepth
	}

	if math.Abs(axisDepth) < math.Abs(depth) {
		depth = axisDepth
		normal = axis
	}

	return true, &Collision{ball, polygon, normal, depth}
}
