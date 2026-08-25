package system

import (
	"fmt"

	"abcd-optics/internal/matrix"
)

func (s System) Through(ray matrix.Ray) matrix.Ray {
	return matrix.Apply(s.Total, ray)
}

func (s System) Determinant() float64 {
	return matrix.Determinant(s.Total)
}

func (s System) IsUnimodular() bool {
	return matrix.IsUnimodular(s.Total)
}

func (s System) EffectiveFocalLength() (float64, error) {
	return EffectiveFocalLength(s.Total)
}

func (s System) BackFocalDistance() (float64, error) {
	if matrix.AlmostZero(s.Total.C) {
		return 0, fmt.Errorf("afocal system: C=0, no finite back focal distance")
	}
	return s.Total.D / s.Total.C, nil
}

func (s System) FrontFocalDistance() (float64, error) {
	return FrontFocalDistance(s.Total)
}

func (s System) PrincipalPlanes() PrincipalPlanes {
	return LocatePrincipalPlanes(s.Total)
}

func (s System) IsAfocal() bool {
	return matrix.AlmostZero(s.Total.C)
}

func (s System) MagnificationAt(objectDistance float64) (float64, error) {
	sPrime, err := ImageDistance(s.Total, objectDistance)
	if err != nil {
		return 0, err
	}
	return Magnification(s.Total, sPrime), nil
}
