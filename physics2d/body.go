package physics2d

type BodyShape uint8

const (
	Ball BodyShape = iota
	Box
	Polygon
)

type Body struct {
	// Shape
	Shape               BodyShape // Ball, Box, Polygon
	Dimensions          Vec2      // m
	Vertices            []Vec2    // m
	TransformedVertices []Vec2    // m
	Radius              float64   // m

	// Physics
	Position           Vec2    // m
	Velocity           Vec2    // m/s
	Rotation           float64 // rad
	RotationalVelocity float64 // rad/s
	Restitution        float64 // 0..1
	InverseMass        float64 // 1/kg
}

func NewBall(pos Vec2, rad float64, rest float64, mass float64) *Body {
	return &Body{
		Shape:               Ball,
		Radius:              rad,
		Vertices:            nil,
		TransformedVertices: nil,

		Position:           pos,
		Velocity:           Vec2{0, 0},
		Rotation:           0,
		RotationalVelocity: 0,
		Restitution:        rest,
		InverseMass:        1.0 / mass,
	}
}

func NewBox(pos Vec2, dim Vec2, rest float64, mass float64) *Body {
	return &Body{
		Shape:               Box,
		Radius:              0,
		Vertices:            boxVertieces(dim),
		TransformedVertices: make([]Vec2, 4),

		Position:    pos,
		Velocity:    Vec2{0, 0},
		Restitution: rest,
		InverseMass: 1.0 / mass,
	}
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

// Move by vel*dt meters
func (b *Body) Update(dt float64) {
	displacement := b.Velocity.ScaleMult(dt)
	b.Move(displacement)
}

func (b *Body) Move(displacement Vec2) {
	b.Position = b.Position.Add(displacement)
}
