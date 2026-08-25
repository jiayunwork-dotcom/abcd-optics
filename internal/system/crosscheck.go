package system

import (
	"fmt"

	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

type Invariants struct {
	Det           float64
	DetIsUnit     bool
	RoundTripOK   bool
	Recovered     matrix.Ray
	SingleLens    bool
	LensFormulaOK bool
}

func CheckInvariants(sys System, incident matrix.Ray) Invariants {
	det := matrix.Determinant(sys.Total)
	recovered, err := matrix.RoundTrip(sys.Matrices, incident)
	inv := Invariants{
		Det:         det,
		DetIsUnit:   matrix.AlmostEqual(det, 1),
		RoundTripOK: err == nil && matrix.CloseRay(recovered, incident),
		Recovered:   recovered,
	}
	if len(sys.Spec.Elements) == 1 {
		if _, ok := sys.Spec.Elements[0].(element.ThinLens); ok {
			inv.SingleLens = true
			inv.LensFormulaOK = CheckSingleLens(sys.Total, sys.Spec.ObjectDistance)
		}
	}
	return inv
}

func CheckSingleLens(m matrix.Mat2, objectDistance float64) bool {
	imageDistance, err := ImageDistance(m, objectDistance)
	if err != nil {
		return false
	}
	focal, err := EffectiveFocalLength(m)
	if err != nil {
		return false
	}
	lhs := 1/objectDistance + 1/imageDistance
	return matrix.AlmostEqual(lhs, 1/focal)
}

func (i Invariants) String() string {
	return fmt.Sprintf(
		"det=%g unit=%v roundtrip=%v single_lens=%v lens_formula=%v",
		i.Det, i.DetIsUnit, i.RoundTripOK, i.SingleLens, i.LensFormulaOK,
	)
}
