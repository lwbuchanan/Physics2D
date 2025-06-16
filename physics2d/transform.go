package physics2d

import "math"

type transform struct {
	Pos Vec2
	Sin float64
	Cos float64
}

func zeroTransform() transform {
	return newTransform(ZeroVec2(), 0)
}

func newTransform(pos Vec2, angle float64) transform {
	return transform{
		Pos: pos,
		Sin: math.Sin(angle),
		Cos: math.Cos(angle),
	}
}
