package system

import (
	"fmt"

	"abcd-optics/internal/matrix"
)

func ImageDistance(m matrix.Mat2, objectDistance float64) (float64, error) {
	den := m.C*objectDistance + m.D
	if matrix.AlmostZero(den) {
		return 0, fmt.Errorf("object plane at focal point: no finite image")
	}
	return -(m.A*objectDistance + m.B) / den, nil
}

func ObjectDistanceFor(m matrix.Mat2, imageDistance float64) (float64, error) {
	den := m.C*imageDistance + m.A
	if matrix.AlmostZero(den) {
		return 0, fmt.Errorf("image plane at focal point: no finite object")
	}
	return -(m.B + m.D*imageDistance) / den, nil
}

func BackFocalDistance(m matrix.Mat2) (float64, error) {
	if matrix.AlmostZero(m.C) {
		return 0, fmt.Errorf("afocal system: C=0, no finite back focal distance")
	}
	return -m.A / m.C, nil
}

func FrontFocalDistance(m matrix.Mat2) (float64, error) {
	if matrix.AlmostZero(m.C) {
		return 0, fmt.Errorf("afocal system: C=0, no finite front focal distance")
	}
	return m.D / m.C, nil
}
