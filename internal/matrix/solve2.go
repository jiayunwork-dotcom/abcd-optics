package matrix

import "errors"

var ErrInconsistent = errors.New("linear system has no unique solution")

func Solve2x2(m Mat2, b Ray) (Ray, error) {
	det := Determinant(m)
	if AlmostZero(det) {
		return Ray{}, ErrInconsistent
	}
	y := (m.D*b.Y - m.B*b.U) / det
	u := (m.A*b.U - m.C*b.Y) / det
	return Ray{Y: y, U: u}, nil
}

func RayThroughSystem(elements []Mat2, target Ray) (Ray, error) {
	inv, err := InverseChain(elements)
	if err != nil {
		return Ray{}, err
	}
	return Apply(inv, target), nil
}
