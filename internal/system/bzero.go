package system

import (
	"abcd-optics/internal/matrix"
)

func HasBZero(m matrix.Mat2) bool {
	return matrix.AlmostZero(m.B)
}

func ImagePlaneWhenBZero(m matrix.Mat2, objectDistance float64) float64 {
	if !matrix.AlmostZero(m.B) {
		return 0
	}
	den := m.C*objectDistance + m.D
	if matrix.AlmostZero(den) {
		return 0
	}
	return -(m.A*objectDistance) / den
}

func MagnificationFromRays(m matrix.Mat2, objectDistance float64) (float64, error) {
	sPrime, err := ImageDistance(m, objectDistance)
	if err != nil {
		return 0, err
	}
	objectRay := matrix.NewRay(1, 0)
	systemRay := matrix.Apply(m, objectRay)
	imageRay := matrix.NewRay(
		systemRay.Y+sPrime*systemRay.U,
		systemRay.U,
	)
	return imageRay.Y, nil
}

func SameConjugate(m matrix.Mat2, objectDistance float64) bool {
	sPrime, err := ImageDistance(m, objectDistance)
	if err != nil {
		return false
	}
	direct := Magnification(m, sPrime)
	viaRay, err := MagnificationFromRays(m, objectDistance)
	if err != nil {
		return false
	}
	return matrix.AlmostEqual(direct, viaRay)
}
