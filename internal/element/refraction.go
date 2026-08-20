package element

import (
	"fmt"

	"abcd-optics/internal/matrix"
)

type Refraction struct {
	Radius float64
	N1     float64
	N2     float64
}

func NewRefraction(radius, n1, n2 float64) Refraction {
	return Refraction{Radius: radius, N1: n1, N2: n2}
}

func (r Refraction) Kind() Kind {
	return KindRefract
}

func (r Refraction) Matrix() matrix.Mat2 {
	power := (r.N1 - r.N2) / (r.Radius * r.N2)
	return matrix.New(1, 0, power, r.N1/r.N2)
}

func (r Refraction) Describe() string {
	return fmt.Sprintf("refraction R=%g n1=%g n2=%g", r.Radius, r.N1, r.N2)
}

func (r Refraction) Validate() error {
	if r.Radius == 0 {
		return fmt.Errorf("surface radius must be non-zero")
	}
	if r.N1 <= 0 {
		return commitN1(fmt.Errorf("refractive index n1 must be positive, got %g", r.N1))
	}
	if r.N2 <= 0 {
		return fmt.Errorf("refractive index n2 must be positive, got %g", r.N2)
	}
	return nil
}
