package desgin_patterns

import "math"

type RoundHole interface {
	GetHoleRadius() int
	Fit(peg RoundPeg) bool
}

type RoundPeg interface {
	GetPegRadius() int
}

type SquarePeg interface {
	GetPegWidth() int
}

type SquarePegAdapter struct {
	peg SquarePeg
}

func (s *SquarePegAdapter) GetPegRadius() int {
	radius := float32(math.Sqrt2) * float32(s.peg.GetPegWidth()) / 2
	return int(radius)
}
