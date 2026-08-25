package system

import (
	"abcd-optics/internal/matrix"
)

type SystemPair struct {
	A matrix.Mat2
	B matrix.Mat2
}

func CombineSystems(a, b matrix.Mat2) matrix.Mat2 {
	return matrix.Multiply(b, a)
}

func CombinePairs(pairs []SystemPair) matrix.Mat2 {
	acc := matrix.Identity()
	for _, p := range pairs {
		acc = matrix.Multiply(CombineSystems(p.A, p.B), acc)
	}
	return acc
}

func SystemFromParams(a, b, c, d float64) matrix.Mat2 {
	return matrix.New(a, b, c, d)
}

func PupilPosition(sys System, reference matrix.Mat2) (float64, error) {
	combined := CombineSystems(reference, sys.Total)
	return ImageDistance(combined, 0)
}
