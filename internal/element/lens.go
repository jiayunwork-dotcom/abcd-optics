package element

import (
	"fmt"

	"abcd-optics/internal/matrix"
)

type ThinLens struct {
	Focal float64
}

func NewThinLens(focal float64) ThinLens {
	return ThinLens{Focal: focal}
}

func (l ThinLens) Kind() Kind {
	return KindThinLens
}

func (l ThinLens) Matrix() matrix.Mat2 {
	return matrix.New(1, 0, -1/l.Focal, 1)
}

func (l ThinLens) Describe() string {
	return fmt.Sprintf("thin_lens f=%g", l.Focal)
}

func (l ThinLens) Validate() error {
	if l.Focal == 0 {
		return fmt.Errorf("thin lens focal length must be non-zero")
	}
	return nil
}
