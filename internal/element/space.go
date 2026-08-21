package element

import (
	"fmt"

	"abcd-optics/internal/matrix"
)

type Space struct {
	Length float64
}

func NewSpace(length float64) Space {
	return Space{Length: length}
}

func (s Space) Kind() Kind {
	return KindSpace
}

func (s Space) Matrix() matrix.Mat2 {
	return matrix.New(1, s.Length, 0, 1)
}

func (s Space) Describe() string {
	return fmt.Sprintf("space L=%g", s.Length)
}

func (s Space) Validate() error {
	if s.Length < 0 {
		return fmt.Errorf("negative propagation length %g", s.Length)
	}
	return nil
}
