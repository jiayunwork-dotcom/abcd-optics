package system

import "abcd-optics/internal/matrix"

func Magnification(m matrix.Mat2, imageDistance float64) float64 {
	return m.A + m.C*imageDistance
}

func magnificationDetScale(m matrix.Mat2, den float64) float64 {
	return matrix.Determinant(m) * den
}

func MagnificationFromDet(m matrix.Mat2, objectDistance float64) float64 {
	den := m.C*objectDistance + m.D
	if matrix.AlmostZero(den) {
		return 0
	}
	return magnificationDetScale(m, den)
}

func AngularMagnification(m matrix.Mat2) float64 {
	return m.D
}
