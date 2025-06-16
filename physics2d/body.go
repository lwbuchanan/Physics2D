package physics2d

import (
	"errors"
)

type BodyShape uint8

const (
	Ball BodyShape = iota
	Polygon
)

type Body struct {
	Shape                   BodyShape
	Dimensions              Vec2
	Radius                  float64
	vertices                []Vec2
	transformedVertices     []Vec2
	needTransformUpdate     bool
	Position                Vec2    // m
	Velocity                Vec2    // m/s
	Acceleration            Vec2    // m/s2
	InverseMass             float64 // kg
	Rotation                float64 // rad
	RotationalVelocity      float64 // rad/s
	RotationalAcceleration  float64 // rad/s2
	InverseMomentOfIntertia float64 // kg*m2
	Restitution             float64
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
		Shape:                   Ball,
		Dimensions:              Vec2{radius * 2, radius * 2},
		Radius:                  radius,
		vertices:                nil,
		transformedVertices:     nil,
		needTransformUpdate:     true,
		Position:                position,
		Velocity:                Vec2{0, 0},
		Acceleration:            Vec2{0, 0},
		InverseMass:             1.0 / mass,
		Rotation:                0,
		RotationalVelocity:      0,
		RotationalAcceleration:  0,
		InverseMomentOfIntertia: 1.0 / (0.5 * mass * radius * radius),
		Restitution:             restitution,
	}, nil
}

func NewBox(position Vec2, dimensions Vec2, rotation float64, restitution float64, mass float64) (*Body, error) {
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
		Shape:                   Polygon,
		Dimensions:              dimensions,
		Radius:                  dimensions.x / 2,
		vertices:                boxVertieces(dimensions),
		transformedVertices:     make([]Vec2, 4),
		needTransformUpdate:     true,
		Position:                position,
		Velocity:                Vec2{0, 0},
		Acceleration:            Vec2{0, 0},
		InverseMass:             1.0 / mass,
		Rotation:                rotation,
		RotationalVelocity:      0,
		RotationalAcceleration:  0,
		InverseMomentOfIntertia: 1.0 / ((1.0 / 12.0) * mass * (dimensions.x*dimensions.x + dimensions.y*dimensions.y)),
		Restitution:             restitution,
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

// Integrate the acceleration/velocity over time to determine new velocity and position
func (b *Body) Update(dt float64) {
	b.Velocity = b.Velocity.Add(b.Acceleration.ScaleMult(dt))
	b.RotationalVelocity += b.RotationalAcceleration * dt

	b.Move(b.Velocity.ScaleMult(dt))
	b.Rotate(b.RotationalVelocity * dt)

	// Acceleration is reevaluated every tick
	b.Acceleration = ZeroVec2()
	b.RotationalAcceleration = 0
}

func (b *Body) Move(displacement Vec2) {
	b.Position = b.Position.Add(displacement)
	b.needTransformUpdate = true
}

func (b *Body) MoveTo(position Vec2) {
	b.Position = position
	b.needTransformUpdate = true
}

func (b *Body) Accelerate(acceleration Vec2) {
	b.Acceleration = b.Acceleration.Add(acceleration)
}

// Instantanoues force in Newtons (mass is kg)
func (b *Body) ApplyForce(force Vec2) {
	b.Acceleration = b.Acceleration.Add(force.ScaleMult(b.InverseMass))
}

// Instantaneous torque in Newton-meters
func (b *Body) ApplyTorque(torque float64) {
	b.RotationalAcceleration += torque * b.InverseMomentOfIntertia
}

// Applies the linear component of a force and its moment
func (b *Body) ApplyPositionalForce(force Vec2, position Vec2) {
	b.Acceleration = b.Acceleration.Add(force.ScaleMult(b.InverseMass))
	b.RotationalAcceleration += position.Cross(force) * b.InverseMomentOfIntertia
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
