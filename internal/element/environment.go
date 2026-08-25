package element

import "fmt"

type RefractiveEnvironment struct {
	Index float64
	Valid bool
}

func EnvironmentAt(spec Spec, position int) RefractiveEnvironment {
	if position <= 0 {
		return RefractiveEnvironment{Index: 1, Valid: true}
	}
	index := 1.0
	for i := 0; i < position && i < len(spec.Elements); i++ {
		switch e := spec.Elements[i].(type) {
		case Refraction:
			index = e.N2
		}
	}
	return RefractiveEnvironment{Index: index, Valid: true}
}

func IncidentRayValid(spec Spec) error {
	if !spec.Incident.IsFinite() {
		return fmt.Errorf("incident ray must be finite, got %v", spec.Incident)
	}
	return nil
}

func AllIndexPositive(spec Spec) error {
	for i, e := range spec.Elements {
		if r, ok := e.(Refraction); ok {
			if r.N1 <= 0 {
				return fmt.Errorf("element %d: n1=%g must be positive", i, r.N1)
			}
			if r.N2 <= 0 {
				return fmt.Errorf("element %d: n2=%g must be positive", i, r.N2)
			}
		}
	}
	return nil
}
