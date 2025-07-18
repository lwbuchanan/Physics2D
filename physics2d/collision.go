package physics2d

import (
	"fmt"
	"math"
)

// Normal is normalized and in the a->b direction
type Collision struct {
	a      *Body
	b      *Body
	normal Vec2
	depth  float64
}

func (c *Collision) Resolve() {

	// Start by separating bodies
	if c.a.inverseMass == 0 {
		c.b.Move(c.normal.ScaleMult(c.depth))
	} else if c.b.inverseMass == 0 {
		c.a.Move(c.normal.ScaleMult(-c.depth))
	} else {
		c.a.Move(c.normal.ScaleMult(-c.depth / 2.0))
		c.b.Move(c.normal.ScaleMult(c.depth / 2.0))
	}

	// We can find accurate collision points now that they are barely touching
	cp := collisionPoints(c)[0]
	// If I come back to this, it might be nice have special cases for multiple collision points

	relativeVelocity := c.b.velocity.Sub(c.a.velocity)
	rVelDotNormal := relativeVelocity.Dot(c.normal)

	if rVelDotNormal > 0.0 {
		// objects are separating
		return
	}

	e := math.Min(c.a.restitution, c.b.restitution)

	rAP_perp := cp.Sub(c.a.position).Perpendicular()
	rBP_perp := cp.Sub(c.b.position).Perpendicular()

	j := (-(1.0 + e) * rVelDotNormal) / ((c.a.inverseMass + c.b.inverseMass) +
		(rAP_perp.Dot(c.normal) * rAP_perp.Dot(c.normal) * c.a.inverseMomentOfIntertia) +
		(rBP_perp.Dot(c.normal) * rBP_perp.Dot(c.normal) * c.b.inverseMomentOfIntertia))

	c.a.velocity = c.a.velocity.Add(c.normal.ScaleMult(-j * c.a.inverseMass))
	c.b.velocity = c.b.velocity.Add(c.normal.ScaleMult(j * c.b.inverseMass))

	c.a.rotationalVelocity += rAP_perp.Dot(c.normal.ScaleMult(-j)) * c.a.inverseMomentOfIntertia
	c.b.rotationalVelocity += rBP_perp.Dot(c.normal.ScaleMult(j)) * c.b.inverseMomentOfIntertia
}

func Collide(a, b *Body) (*Collision, error) {
	switch a.shape {
	case Ball:
		switch b.shape {
		case Ball:
			return ballsCollide(a, b)
		case Polygon:
			return ballAndPolygonCollide(a, b)
		}
		return nil, fmt.Errorf("collision: %d is not a valid body shape", b.shape)
	case Polygon:
		switch b.shape {
		case Polygon:
			return polygonsCollide(a, b)
		case Ball:
			return ballAndPolygonCollide(b, a)
		}
		return nil, fmt.Errorf("collision: %d is not a valid body shape", b.shape)
	}
	return nil, fmt.Errorf("collision: %d is not a valid body shape", a.shape)
}

func collisionPoints(c *Collision) []Vec2 {
	if c.a.shape == Ball { // If a is a ball, we dont care what b is
		// Balls can only contact other objects at one point
		return []Vec2{c.a.position.Add(c.normal.ScaleMult(c.a.radius))}
	}
	// Otherwise, both are definitely polygons
	cp1 := ZeroVec2()
	cp2 := ZeroVec2()
	numContacts := 0

	aVerts := c.a.Vertices()
	bVerts := c.b.Vertices()
	minDistSquared := math.Inf(1)
	for _, aVert := range aVerts {
		for j, bEdgeStart := range bVerts {
			bEdgeEnd := bVerts[(j+1)%len(bVerts)]
			cp, distSquared := ClosestPointOnSegment(aVert, bEdgeStart, bEdgeEnd)

			if distSquared == minDistSquared {
				if cp.CloseTo(cp1) {
					cp2 = cp
					numContacts = 2
				}
			} else if distSquared < minDistSquared {
				cp1 = cp
				numContacts = 1
				minDistSquared = distSquared
			}
		}
	}
	for _, bVert := range bVerts {
		for j, aEdgeStart := range aVerts {
			aEdgeEnd := aVerts[(j+1)%len(aVerts)]
			cp, distSquared := ClosestPointOnSegment(bVert, aEdgeStart, aEdgeEnd)

			if distSquared == minDistSquared {
				if cp.CloseTo(cp1) {
					cp2 = cp
					numContacts = 2
				}
			} else if distSquared < minDistSquared {
				cp1 = cp
				numContacts = 1
				minDistSquared = distSquared
			}
		}
	}

	if numContacts == 1 {
		return []Vec2{cp1}
	}
	return []Vec2{cp1, cp2}
}

// Gets the point on the segment VW that is closest to P, and its distance(squared) from P
func ClosestPointOnSegment(p, v, w Vec2) (Vec2, float64) {
	vw := w.Sub(v)
	vp := p.Sub(v)

	proj := vp.Dot(vw)
	lenVWSq := vw.LengthSquared()

	d := proj / lenVWSq

	var cp Vec2
	if d <= 0 {
		cp = v
	} else if d >= 1 {
		cp = w
	} else {
		cp = v.Add(vw.ScaleMult(d))
	}

	return cp, cp.DistanceSquared(p)
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

func ballsCollide(a, b *Body) (*Collision, error) {
	bothRad := a.radius + b.radius
	distSquared := a.position.DistanceSquared(b.position)

	if distSquared >= (bothRad * bothRad) {
		return nil, nil
	}

	// Only do expensive operations when collision is confirmed
	distance := math.Sqrt(distSquared)
	depth := bothRad - distance
	displacement := b.position.Sub(a.position)
	normal := displacement.Normalize()

	return &Collision{a, b, normal, depth}, nil
}

// SAT only works for convex polygons
func polygonsCollide(a, b *Body) (*Collision, error) {
	normal := ZeroVec2()
	depth := math.MaxFloat64

	aVertices := a.Vertices()
	bVertices := b.Vertices()

	// Vertecies are stored clockwise, so we test edges clockwise
	for i := range len(aVertices) {
		vCurr := aVertices[i]
		vNext := aVertices[(i+1)%len(aVertices)]
		edge := vNext.Sub(vCurr)

		axis := edge.Perpendicular().Normalize() /* Orthoganal to tested edge */

		aMin, aMax := projectVertecies(aVertices, axis)
		bMin, bMax := projectVertecies(bVertices, axis)

		if aMin >= bMax || bMin >= aMax {
			// Found separating axis
			return nil, nil
		}

		axisDepth := math.Min(bMax-aMin, aMax-bMin)

		if axisDepth < depth {
			depth = axisDepth
			normal = axis
		}
	}

	// This part is the same, except we check all of b's axes instead
	for i := range len(bVertices) {
		vCurr := bVertices[i]
		vNext := bVertices[(i+1)%len(bVertices)]
		edge := vNext.Sub(vCurr)
		axis := edge.Perpendicular().Normalize()

		aMin, aMax := projectVertecies(aVertices, axis)
		bMin, bMax := projectVertecies(bVertices, axis)

		if aMin >= bMax || bMin >= aMax {
			return nil, nil
		}

		if aMin >= bMax || bMin >= aMax {
			// Found separating axis
			return nil, nil
		}

		axisDepth := math.Min(bMax-aMin, aMax-bMin)
		if axisDepth < depth {
			depth = axisDepth
			normal = axis
		}
	}

	// Ensure that normal points a->b
	if b.position.Sub(a.position).Dot(normal) < 0.0 {
		normal = normal.ScaleMult(-1)
	}

	return &Collision{a, b, normal, depth}, nil
}

func ballAndPolygonCollide(ball, polygon *Body) (*Collision, error) {
	vertices := polygon.Vertices()

	normal := ZeroVec2()
	depth := math.MaxFloat64

	// Check for a SA between the circles edge to closest vertex
	closestVertex := vertices[closestVertexIdx(ball.position, vertices)]
	axis := closestVertex.Sub(ball.position).Normalize()

	pMin, pMax := projectVertecies(vertices, axis)
	bMin, bMax := projectCircle(ball.position, ball.radius, axis)

	if pMin >= bMax || bMin >= pMax {
		// Found separating axis
		return nil, nil
	}

	axisDepth := math.Min(bMax-pMin, pMax-bMin)
	if axisDepth < depth {
		depth = axisDepth
		normal = axis
	}

	// This code is mostly the same as 2 polygons. Theres a lot of code duplication, but its
	// tricky to make it more abstract while making sure that the normal still points in the
	// right direction and everything.
	for i := range len(vertices) {
		vCurr := vertices[i]
		vNext := vertices[(i+1)%len(vertices)]
		edge := vNext.Sub(vCurr)
		axis := edge.Perpendicular().Normalize()

		bMin, bMax := projectCircle(ball.position, ball.radius, axis)
		pMin, pMax := projectVertecies(vertices, axis)

		if pMin >= bMax || bMin >= pMax {
			// Found separating axis
			return nil, nil
		}

		axisDepth := math.Min(bMax-pMin, pMax-bMin)
		if axisDepth < depth {
			depth = axisDepth
			normal = axis
		}
	}

	// Ensure that normal points a->b
	if polygon.position.Sub(ball.position).Dot(normal) < 0.0 {
		normal = normal.ScaleMult(-1)
	}

	return &Collision{ball, polygon, normal, depth}, nil
}
