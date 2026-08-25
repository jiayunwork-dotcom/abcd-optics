package system

import (
	"abcd-optics/internal/matrix"
)

type RayStep struct {
	Index   int
	Element string
	Ray     matrix.Ray
}

type RayPath struct {
	Incident matrix.Ray
	Steps    []RayStep
	Emitted  matrix.Ray
}

func TraceRayPath(sys System, incident matrix.Ray) (RayPath, error) {
	current := incident
	steps := make([]RayStep, 0, len(sys.Matrices))
	for i, m := range sys.Matrices {
		current = matrix.Apply(m, current)
		steps = append(steps, RayStep{
			Index:   i,
			Element: sys.Spec.Elements[i].Describe(),
			Ray:     current,
		})
	}
	return RayPath{
		Incident: incident,
		Steps:    steps,
		Emitted:  current,
	}, nil
}
