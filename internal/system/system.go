package system

import (
	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

type System struct {
	Spec     element.Spec
	Matrices []matrix.Mat2
	Total    matrix.Mat2
}

func Compose(spec element.Spec) (System, error) {
	if err := spec.Validate(); err != nil {
		return System{}, err
	}
	matrices := make([]matrix.Mat2, 0, len(spec.Elements))
	for _, e := range spec.Elements {
		matrices = append(matrices, e.Matrix())
	}
	return System{
		Spec:     spec,
		Matrices: matrices,
		Total:    matrix.Chain(matrices),
	}, nil
}

func (s System) A() float64 {
	return s.Total.A
}

func (s System) B() float64 {
	return s.Total.B
}

func (s System) C() float64 {
	return s.Total.C
}

func (s System) D() float64 {
	return s.Total.D
}
