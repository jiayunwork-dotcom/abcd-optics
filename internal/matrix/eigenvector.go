package matrix

import "errors"

var ErrComplexEigen = errors.New("complex eigenvalues: no real eigenvector")

func Eigenvectors(m Mat2) (Ray, Ray, error) {
	e := Eigenvalues(m)
	if !e.Real {
		return Ray{}, Ray{}, ErrComplexEigen
	}
	v1, ok := eigenvectorFor(m, e.Lambda1)
	if !ok {
		return Ray{}, Ray{}, ErrSingular
	}
	v2, ok := eigenvectorFor(m, e.Lambda2)
	if !ok {
		return Ray{}, Ray{}, ErrSingular
	}
	return v1, v2, nil
}

func eigenvectorFor(m Mat2, lambda float64) (Ray, bool) {
	a := m.A - lambda
	b := m.B
	if !AlmostZero(b) {
		return Ray{Y: 1, U: -a / b}, true
	}
	c := m.C
	if !AlmostZero(c) {
		return Ray{Y: -b / c, U: 1}, true
	}
	return Ray{}, false
}
