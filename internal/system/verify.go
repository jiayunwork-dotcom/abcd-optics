package system

import (
	"fmt"
	"strings"

	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

func VerifyDetUnit(spec element.Spec) error {
	sys, err := Compose(spec)
	if err != nil {
		return err
	}
	det := matrix.Determinant(sys.Total)
	if !matrix.AlmostEqual(det, 1) {
		return fmt.Errorf("air-to-air system det=%g, want 1", det)
	}
	return nil
}

func VerifyThinLensFormula(spec element.Spec) error {
	if len(spec.Elements) != 1 {
		return nil
	}
	if _, ok := spec.Elements[0].(element.ThinLens); !ok {
		return nil
	}
	sys, err := Compose(spec)
	if err != nil {
		return err
	}
	ok := CheckSingleLens(sys.Total, spec.ObjectDistance)
	if !ok {
		return fmt.Errorf("thin lens 1/s+1/s'=1/f not reproduced from matrix")
	}
	return nil
}

func VerifyRoundTrip(spec element.Spec, incident matrix.Ray) error {
	sys, err := Compose(spec)
	if err != nil {
		return err
	}
	recovered, err := matrix.RoundTrip(sys.Matrices, incident)
	if err != nil {
		return err
	}
	if !matrix.CloseRay(recovered, incident) {
		return fmt.Errorf("round trip recovered %v, want %v", recovered, incident)
	}
	return nil
}

func VerifyAll(spec element.Spec, incident matrix.Ray) string {
	var b strings.Builder
	for _, fn := range []func() error{
		func() error { return VerifyDetUnit(spec) },
		func() error { return VerifyThinLensFormula(spec) },
		func() error { return VerifyRoundTrip(spec, incident) },
	} {
		if err := fn(); err != nil {
			b.WriteString("FAIL " + err.Error() + "\n")
		} else {
			b.WriteString("PASS\n")
		}
	}
	return b.String()
}
