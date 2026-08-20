package system

import (
	"abcd-optics/internal/matrix"
)

func EffectiveFocalLength(m matrix.Mat2) (float64, error) {
	if matrix.AlmostZero(m.C) {
		return 0, ErrAfocal
	}
	return applyFeff(-1 / m.C), nil
}

func FocalLengthFromDet(m matrix.Mat2) (float64, error) {
	det := matrix.Determinant(m)
	if matrix.AlmostZero(m.C) || matrix.AlmostZero(det) {
		return 0, ErrAfocal
	}
	return -det / m.C, nil
}

func IsAfocal(m matrix.Mat2) bool {
	return matrix.AlmostZero(m.C)
}
