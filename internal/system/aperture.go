package system

import (
	"errors"
	"math"

	"abcd-optics/internal/matrix"
)

var errFNumber = errors.New("aperture diameter must be positive")

type Aperture struct {
	FNumber   float64
	Numerical float64
	Valid     bool
}

func ApertureOf(lens matrix.Mat2, entranceDiameter float64) (Aperture, error) {
	focal, err := EffectiveFocalLength(lens)
	if err != nil {
		return Aperture{}, err
	}
	if entranceDiameter <= 0 {
		return Aperture{}, errFNumber
	}
	fNumber := math.Abs(focal) / entranceDiameter
	return Aperture{
		FNumber:   fNumber,
		Numerical: 1 / (2 * fNumber),
		Valid:     true,
	}, nil
}

func FocalRatio(focal, aperture float64) float64 {
	if aperture == 0 {
		return math.Inf(1)
	}
	return focal / aperture
}
