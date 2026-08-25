package system

import (
	"abcd-optics/internal/element"
)

type ReducedSystem struct {
	Focal      float64
	FrontNode  float64
	BackNode   float64
	Afocal     bool
}

func Reduce(spec element.Spec) (ReducedSystem, error) {
	sys, err := Compose(spec)
	if err != nil {
		return ReducedSystem{}, err
	}
	planes := LocatePrincipalPlanes(sys.Total)
	if !planes.Valid {
		return ReducedSystem{Afocal: true}, nil
	}
	focal, _ := EffectiveFocalLength(sys.Total)
	return ReducedSystem{
		Focal:     focal,
		FrontNode: planes.FrontFocal,
		BackNode:  planes.BackFocal,
	}, nil
}

func (r ReducedSystem) Describe() string {
	if r.Afocal {
		return "afocal system: no reduced single-lens model"
	}
	return ""
}

func ObjectImageConjugate(sys System, s float64) (sPrime float64, err error) {
	return ImageDistance(sys.Total, s)
}
