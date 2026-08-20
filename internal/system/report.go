package system

import (
	"fmt"
	"strings"

	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

type Report struct {
	Spec     element.Spec
	System   System
	Path     RayPath
	Solve    SolveResult
}

func BuildReport(spec element.Spec) (Report, error) {
	sys, err := Compose(spec)
	if err != nil {
		return Report{}, err
	}
	path, err := TraceRayPath(sys, spec.Incident)
	if err != nil {
		return Report{}, err
	}
	solve, err := Solve(spec)
	if err != nil {
		return Report{}, err
	}
	return Report{Spec: spec, System: sys, Path: path, Solve: solve}, nil
}

func (r Report) String() string {
	var b strings.Builder
	b.WriteString("spec: " + r.Spec.Name + "\n")
	b.WriteString(r.Spec.Describe())
	b.WriteString(matrix.FormatMatrix("M", r.System.Total))
	b.WriteString(fmt.Sprintf("det(M)      = %.9f\n", r.Solve.Invariants.Det))
	b.WriteString(fmt.Sprintf("object dist = %.6f\n", r.Spec.ObjectDistance))
	b.WriteString(fmt.Sprintf("image dist  = %.6f\n", r.Solve.ImageDistance))
	b.WriteString(fmt.Sprintf("magnif m    = %.6f\n", r.Solve.Magnification))
	if r.System.IsAfocal() {
		b.WriteString("f_eff       = inf (afocal, C=0)\n")
	} else {
		b.WriteString(fmt.Sprintf("f_eff       = %.6f\n", r.Solve.EffectiveFocal))
	}
	b.WriteString(fmt.Sprintf("ray in      = %v\n", r.Path.Incident))
	for _, s := range r.Path.Steps {
		b.WriteString(fmt.Sprintf("  after [%d] %-24s %v\n", s.Index, s.Element, s.Ray))
	}
	b.WriteString(fmt.Sprintf("ray out     = %v\n", r.Path.Emitted))
	b.WriteString("cross-checks: " + r.Solve.Invariants.String() + "\n")
	return b.String()
}
