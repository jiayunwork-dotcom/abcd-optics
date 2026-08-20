package matrix

import "math"

const DefaultTolerance = 1e-9

func AlmostZero(v float64) bool {
	return math.Abs(v) <= DefaultTolerance
}

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= DefaultTolerance
}

func CloseRay(a, b Ray) bool {
	return AlmostEqual(a.Y, b.Y) && AlmostEqual(a.U, b.U)
}

func CloseMat(a, b Mat2) bool {
	return AlmostEqual(a.A, b.A) &&
		AlmostEqual(a.B, b.B) &&
		AlmostEqual(a.C, b.C) &&
		AlmostEqual(a.D, b.D)
}

func isFinite(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0)
}
