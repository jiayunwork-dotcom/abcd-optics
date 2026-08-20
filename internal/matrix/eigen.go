package matrix

import "math"

type Eigen struct {
	Lambda1 float64
	Lambda2 float64
	Real    bool
}

func Eigenvalues(m Mat2) Eigen {
	trace := m.A + m.D
	disc := trace*trace - 4*Determinant(m)
	if disc < 0 {
		re := trace / 2
		im := math.Sqrt(-disc) / 2
		return Eigen{Lambda1: re, Lambda2: im, Real: false}
	}
	sqrt := math.Sqrt(disc)
	return Eigen{
		Lambda1: (trace + sqrt) / 2,
		Lambda2: (trace - sqrt) / 2,
		Real:    true,
	}
}

func IsStableResonator(m Mat2) bool {
	trace := m.A + m.D
	return math.Abs(trace) < 2
}

func RoundTripStability(cavity Mat2) bool {
	return IsStableResonator(cavity)
}

func ResonatorRoundTrip(r1, r2, length float64) Mat2 {
	m1 := RefractingSurface(r1, 1, -1)
	space := FreeSpace(length)
	m2 := RefractingSurface(r2, 1, -1)
	return Chain([]Mat2{m1, space, m2, space})
}
