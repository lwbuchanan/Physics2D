package physics2d

import "math"

type Transform struct {
	Pos Vec2
	Sin float64
	Cos float64
}

func ZeroTransform() Transform {
	return NewTransform(ZeroVec2(), 0)
}

func NewTransform(pos Vec2, angle float64) Transform {
	return Transform{
		Pos: pos,
		Sin: math.Sin(angle),
		Cos: math.Cos(angle),
	}
}
