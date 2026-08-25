package system

import "abcd-optics/internal/matrix"

type Beam struct {
	Axis matrix.Ray
	Edge matrix.Ray
}

func NewBeam(axis, edge matrix.Ray) Beam {
	return Beam{Axis: axis, Edge: edge}
}

func (b Beam) Through(m matrix.Mat2) Beam {
	return Beam{
		Axis: matrix.Apply(m, b.Axis),
		Edge: matrix.Apply(m, b.Edge),
	}
}

func (b Beam) Width() float64 {
	return b.Edge.Y - b.Axis.Y
}

func (b Beam) Divergence() float64 {
	return b.Edge.U - b.Axis.U
}

func BeamAfterSystem(sys System, b Beam) (Beam, error) {
	return Beam{
		Axis: matrix.Apply(sys.Total, b.Axis),
		Edge: matrix.Apply(sys.Total, b.Edge),
	}, nil
}

func MarginalHeight(sys System, height float64) matrix.Ray {
	incident := matrix.NewRay(height, 0)
	return matrix.Apply(sys.Total, incident)
}

func ChiefRay(sys System, slope float64) matrix.Ray {
	incident := matrix.NewRay(0, slope)
	return matrix.Apply(sys.Total, incident)
}
