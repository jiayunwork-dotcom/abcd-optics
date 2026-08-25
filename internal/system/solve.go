package system

import (
	"fmt"

	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

type SolveResult struct {
	System         System
	ImageDistance  float64
	Magnification  float64
	EffectiveFocal float64
	Invariants     Invariants
	RayOut         matrix.Ray
}

func (r SolveResult) RecoveredObjectDistance() (float64, error) {
	return ObjectDistanceFor(r.System.Total, r.ImageDistance)
}

func Solve(spec element.Spec) (SolveResult, error) {
	sys, err := Compose(spec)
	if err != nil {
		return SolveResult{}, err
	}
	imageDistance, err := ImageDistance(sys.Total, spec.ObjectDistance)
	if err != nil {
		return SolveResult{}, err
	}
	mag := Magnification(sys.Total, imageDistance)
	eff, _ := EffectiveFocalLength(sys.Total)
	incident := matrix.NewRay(1, 0)
	rayOut := matrix.Apply(sys.Total, incident)
	inv := CheckInvariants(sys, incident)
	return SolveResult{
		System:         sys,
		ImageDistance:  imageDistance,
		Magnification:  mag,
		EffectiveFocal: eff,
		Invariants:     inv,
		RayOut:         rayOut,
	}, nil
}

func (r SolveResult) String() string {
	m := r.System.Total
	out := fmt.Sprintf("A=%.6f B=%.6f C=%.6f D=%.6f\n", m.A, m.B, m.C, m.D)
	out += fmt.Sprintf("det(M)=%.9f\n", r.Invariants.Det)
	if matrix.AlmostZero(m.C) {
		out += "effective_focal_length=inf (afocal, C=0)\n"
	} else {
		out += fmt.Sprintf("effective_focal_length=%.6f\n", r.EffectiveFocal)
	}
	out += fmt.Sprintf("object_distance=%.6f\n", r.System.Spec.ObjectDistance)
	out += fmt.Sprintf("image_distance=%.6f\n", r.ImageDistance)
	out += fmt.Sprintf("magnification=%.6f\n", r.Magnification)
	out += fmt.Sprintf("round_trip_restores=%v\n", r.Invariants.RoundTripOK)
	return out
}
