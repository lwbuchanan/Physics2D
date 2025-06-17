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
	shape                   BodyShape
	dimensions              Vec2
	radius                  float64
	vertices                []Vec2
	transformedVertices     []Vec2
	needTransformUpdate     bool
	position                Vec2    // m
	velocity                Vec2    // m/s
	acceleration            Vec2    // m/s2
	inverseMass             float64 // kg
	rotation                float64 // rad
	rotationalVelocity      float64 // rad/s
	rotationalAcceleration  float64 // rad/s2
	inverseMomentOfIntertia float64 // kg*m2
	restitution             float64
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
		shape:                   Ball,
		dimensions:              Vec2{radius * 2, radius * 2},
		radius:                  radius,
		vertices:                nil,
		transformedVertices:     nil,
		needTransformUpdate:     true,
		position:                position,
		velocity:                Vec2{0, 0},
		acceleration:            Vec2{0, 0},
		inverseMass:             1.0 / mass,
		rotation:                0,
		rotationalVelocity:      0,
		rotationalAcceleration:  0,
		inverseMomentOfIntertia: 1.0 / (0.5 * mass * radius * radius),
		restitution:             restitution,
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
		shape:                   Polygon,
		dimensions:              dimensions,
		radius:                  dimensions.x / 2,
		vertices:                boxVertieces(dimensions),
		transformedVertices:     make([]Vec2, 4),
		needTransformUpdate:     true,
		position:                position,
		velocity:                Vec2{0, 0},
		acceleration:            Vec2{0, 0},
		inverseMass:             1.0 / mass,
		rotation:                rotation,
		rotationalVelocity:      0,
		rotationalAcceleration:  0,
		inverseMomentOfIntertia: 1.0 / ((1.0 / 12.0) * mass * (dimensions.x*dimensions.x + dimensions.y*dimensions.y)),
		restitution:             restitution,
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

// Expose some getters so we can draw everything. They are read only outside the package
// since the physics should only be controlled from inside the engine

func (b *Body) Shape() BodyShape {
	return b.shape
}

func (b *Body) Radius() float64 {
	return b.radius
}

// Only expose the most recently transformed vertices
func (b *Body) Vertices() []Vec2 {
	if b.needTransformUpdate {
		transform := newTransform(b.position, b.rotation)

		for i, v := range b.vertices {
			b.transformedVertices[i] = v.Transform(transform)
		}
		b.needTransformUpdate = false
	}

	return b.transformedVertices
}

func (b *Body) Position() Vec2 {
	return b.position
}

func (b *Body) Velocity() Vec2 {
	return b.velocity
}

// Integrate the acceleration/velocity over time to determine new velocity and position
func (b *Body) Update(dt float64) {
	b.velocity = b.velocity.Add(b.acceleration.ScaleMult(dt))
	b.rotationalVelocity += b.rotationalAcceleration * dt

	b.Move(b.velocity.ScaleMult(dt))
	b.Rotate(b.rotationalVelocity * dt)

	// Acceleration is reevaluated every tick
	b.acceleration = ZeroVec2()
	b.rotationalAcceleration = 0
}

func (b *Body) Move(displacement Vec2) {
	b.position = b.position.Add(displacement)
	b.needTransformUpdate = true
}

func (b *Body) MoveTo(position Vec2) {
	b.position = position
	b.needTransformUpdate = true
}

func (b *Body) Accelerate(acceleration Vec2) {
	b.acceleration = b.acceleration.Add(acceleration)
}

// Instantanoues force in Newtons (mass is kg)
func (b *Body) ApplyForce(force Vec2) {
	b.acceleration = b.acceleration.Add(force.ScaleMult(b.inverseMass))
}

// Instantaneous torque in Newton-meters
func (b *Body) ApplyTorque(torque float64) {
	b.rotationalAcceleration += torque * b.inverseMomentOfIntertia
}

// Applies the linear component of a force and its moment
func (b *Body) ApplyPositionalForce(force Vec2, position Vec2) {
	b.acceleration = b.acceleration.Add(force.ScaleMult(b.inverseMass))
	b.rotationalAcceleration += position.Cross(force) * b.inverseMomentOfIntertia
}

// A positive rotation is counter-clockwise (positive Z by RHR)
func (b *Body) Rotate(rotationalDisplacement float64) {
	b.rotation += rotationalDisplacement
	b.needTransformUpdate = true
}

func (b *Body) RotateTo(rotation float64) {
	b.rotation = rotation
	b.needTransformUpdate = true
}
