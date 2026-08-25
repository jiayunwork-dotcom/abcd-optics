package matrix

import "math"

func StabilityClass(m Mat2) string {
	t := Trace(m)
	switch {
	case t > 2:
		return "unstable_hyperbolic"
	case t < -2:
		return "unstable_negative"
	case AlmostEqual(t, 2):
		return "stable_marginal"
	case AlmostEqual(t, -2):
		return "stable_inverted_marginal"
	default:
		return "stable_elliptic"
	}
}

func PeriodOf(m Mat2) int {
	if !IsStableResonator(m) {
		return 0
	}
	t := Trace(m)
	theta := math.Acos(clampUnit(t / 2))
	for n := 1; n <= 100; n++ {
		if AlmostEqual(cosNTheta(theta, n), 1) {
			return n
		}
	}
	return 0
}

func cosNTheta(theta float64, n int) float64 {
	t := 2 * math.Cos(theta)
	if n == 1 {
		return t / 2
	}
	prev := 1.0
	cur := t
	for i := 2; i <= n; i++ {
		next := t*cur - prev
		prev, cur = cur, next
	}
	return cur / 2
}

func clampUnit(v float64) float64 {
	if v > 1 {
		return 1
	}
	if v < -1 {
		return -1
	}
	return v
}
