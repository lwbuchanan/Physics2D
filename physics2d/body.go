package physics2d

import (
	"errors"
)

type BodyShape uint8

const (
	Ball BodyShape = iota
	Box
	Polygon
)

type Body struct {
	Shape               BodyShape
	Dimensions          Vec2
	Radius              float64
	vertices            []Vec2
	transformedVertices []Vec2
	needTransformUpdate bool
	Position            Vec2
	Velocity            Vec2
	Rotation            float64
	RotationalVelocity  float64
	Restitution         float64
	InverseMass         float64
}

func NewBall(position Vec2, radius float64, restitution float64, mass float64) (*Body, error) {
	if radius <= 0 {
		return nil, errors.New("physics2d: ball must have positive radius")
	}
	if restitution < 0 || restitution > 1 {
		return nil, errors.New("physics2d: ball must have restitution in range 0..1")
	}
	if mass < 0 {
		return nil, errors.New("physics2d: ball must have nonnegative mass")
	}
	return &Body{
		Shape:               Ball,
		Dimensions:          Vec2{radius * 2, radius * 2},
		Radius:              radius,
		vertices:            nil,
		transformedVertices: nil,
		needTransformUpdate: true,
		Position:            position,
		Velocity:            Vec2{0, 0},
		Rotation:            0,
		RotationalVelocity:  0,
		Restitution:         restitution,
		InverseMass:         1.0 / mass,
	}, nil
}

func NewBox(position Vec2, dimensions Vec2, rotationalVelocity float64, restitution float64, mass float64) (*Body, error) {
	if dimensions.x <= 0 || dimensions.y <= 0 {
		return nil, errors.New("physics2d: box must have positive dimension")
	}
	if restitution < 0 || restitution > 1 {
		return nil, errors.New("physics2d: box must have restitution in range 0..1")
	}
	if mass < 0 {
		return nil, errors.New("physics2d: box must have nonnegative mass")
	}
	return &Body{
		Shape:               Box,
		Dimensions:          dimensions,
		Radius:              dimensions.x / 2,
		vertices:            boxVertieces(dimensions),
		transformedVertices: make([]Vec2, 4),
		needTransformUpdate: true,
		Position:            position,
		Velocity:            Vec2{0, 0},
		Rotation:            0,
		RotationalVelocity:  rotationalVelocity,
		Restitution:         restitution,
		InverseMass:         1.0 / mass,
	}, nil
}

func boxVertieces(dim Vec2) []Vec2 {
	l := -dim.x / 2
	r := l + dim.x
	b := -dim.y / 2
	t := b + dim.y
	return []Vec2{
		{l, t}, // Top left
		{r, t}, // Top right
		{r, b}, // Bottom right
		{l, b}, // Bottom left
	}
}

// Only expose the most recently transformed vertices
func (b *Body) Vertices() []Vec2 {
	if b.needTransformUpdate {
		transform := newTransform(b.Position, b.Rotation)

		for i, v := range b.vertices {
			b.transformedVertices[i] = v.Transform(transform)
		}
		b.needTransformUpdate = false
	}

	return b.transformedVertices
}

// Move by vel*dt meters and rotate by rvel*dt radians
func (b *Body) Update(dt float64) {
	displacement := b.Velocity.ScaleMult(dt)
	rotationalDisplacement := b.RotationalVelocity * dt
	b.Move(displacement)
	b.Rotate(rotationalDisplacement)
}

func (b *Body) Move(displacement Vec2) {
	b.Position = b.Position.Add(displacement)
	b.needTransformUpdate = true
}

func (b *Body) MoveTo(position Vec2) {
	b.Position = position
	b.needTransformUpdate = true
}

// A positive rotation is counter-clockwise (positive Z by RHR)
func (b *Body) Rotate(rotationalDisplacement float64) {
	b.Rotation += rotationalDisplacement
	b.needTransformUpdate = true
}

func (b *Body) RotateTo(rotation float64) {
	b.Rotation = rotation
	b.needTransformUpdate = true
}
