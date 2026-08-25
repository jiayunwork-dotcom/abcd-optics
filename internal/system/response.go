package system

import (
	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

func ResponseAt(sys System, objectDistance float64) (float64, error) {
	sPrime, err := ImageDistance(sys.Total, objectDistance)
	if err != nil {
		return 0, err
	}
	return Magnification(sys.Total, sPrime), nil
}

func ResponseOverDistances(spec element.Spec, distances []float64) ([]float64, error) {
	sys, err := Compose(spec)
	if err != nil {
		return nil, err
	}
	out := make([]float64, 0, len(distances))
	for _, d := range distances {
		m, err := ResponseAt(sys, d)
		if err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, nil
}

func BeamWaistRatio(sys System, m matrix.Ray) float64 {
	return sys.Total.A + sys.Total.B*m.U/m.Y
}
