package matrix

func IsUnimodularPair(a, b Mat2) bool {
	return AlmostEqual(Determinant(a), 1) && AlmostEqual(Determinant(b), 1)
}

func IsSymmetric(m Mat2) bool {
	return AlmostEqual(m.B, m.C)
}

func IsAfocalSystem(m Mat2) bool {
	return AlmostZero(m.C)
}

func HasRealEigenvalues(m Mat2) bool {
	disc := Trace(m)*Trace(m) - 4*Determinant(m)
	return disc >= 0
}

func IsIdentity(m Mat2) bool {
	return AlmostEqual(m.A, 1) &&
		AlmostEqual(m.B, 0) &&
		AlmostEqual(m.C, 0) &&
		AlmostEqual(m.D, 1)
}
