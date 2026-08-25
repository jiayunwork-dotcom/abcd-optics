package matrix

func Determinant(m Mat2) float64 {
	return m.A*m.D - m.B*m.C
}

func (m Mat2) Determinant() float64 {
	return Determinant(m)
}

func IsUnimodular(m Mat2) bool {
	return AlmostEqual(Determinant(m), 1)
}

func Trace(m Mat2) float64 {
	return m.A + m.D
}

func (m Mat2) Trace() float64 {
	return Trace(m)
}
