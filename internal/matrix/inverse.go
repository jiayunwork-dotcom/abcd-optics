package matrix

import "errors"

var ErrSingular = errors.New("matrix is singular")

func Inverse(m Mat2) (Mat2, error) {
	det := Determinant(m)
	if AlmostZero(det) {
		return Mat2{}, ErrSingular
	}
	inv := 1 / det
	return Mat2{
		A: m.D * inv,
		B: -m.B * inv,
		C: -m.C * inv,
		D: m.A * inv,
	}, nil
}

func MustInverse(m Mat2) Mat2 {
	inv, err := Inverse(m)
	if err != nil {
		panic(err)
	}
	return inv
}

func InverseChain(elements []Mat2) (Mat2, error) {
	acc := Identity()
	for i := len(elements) - 1; i >= 0; i-- {
		inv, err := Inverse(elements[i])
		if err != nil {
			return Mat2{}, err
		}
		acc = Multiply(acc, inv)
	}
	return acc, nil
}
