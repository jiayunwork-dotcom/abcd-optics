package element

import "fmt"

func CheckIndexContinuity(spec Spec) error {
	last := 1.0
	for i, e := range spec.Elements {
		switch r := e.(type) {
		case Refraction:
			if !closeFloat(r.N1, last) {
				return fmt.Errorf(
					"element %d: n1=%g does not continue previous medium %g",
					i, r.N1, last,
				)
			}
			last = r.N2
		case ThinLens:
		case Space:
		}
	}
	return nil
}

func closeFloat(a, b float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < 1e-9
}

func MediaEndToEnd(spec Spec) (float64, float64) {
	first := 1.0
	last := 1.0
	for _, e := range spec.Elements {
		if r, ok := e.(Refraction); ok {
			first = r.N1
			last = r.N2
		}
	}
	return first, last
}
