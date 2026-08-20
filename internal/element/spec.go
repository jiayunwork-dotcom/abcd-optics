package element

import (
	"fmt"

	"abcd-optics/internal/matrix"
)

type Spec struct {
	Name           string
	ObjectDistance float64
	Incident       matrix.Ray
	Elements       []Element
}

func NewSpec(name string, objectDistance float64, elements []Element) Spec {
	return Spec{Name: name, ObjectDistance: objectDistance, Elements: elements}
}

func (s Spec) Validate() error {
	if len(s.Elements) == 0 {
		return ErrNoElements
	}
	if s.ObjectDistance <= 0 {
		return fmt.Errorf("object distance must be positive, got %g", s.ObjectDistance)
	}
	if !s.Incident.IsFinite() {
		return fmt.Errorf("incident ray must be finite, got %v", s.Incident)
	}
	for i, e := range s.Elements {
		if err := e.Validate(); err != nil {
			return fmt.Errorf("element %d: %w", i, err)
		}
	}
	return nil
}

func (s Spec) Describe() string {
	out := "elements:\n"
	out += DescribeElements(s.Elements)
	return out
}
