package system

import (
	"fmt"

	"abcd-optics/internal/matrix"
)

type PrincipalPlanes struct {
	ImageSide    float64
	ObjectSide   float64
	BackFocal    float64
	FrontFocal   float64
	Valid        bool
}

func LocatePrincipalPlanes(m matrix.Mat2) PrincipalPlanes {
	if matrix.AlmostZero(m.C) {
		return PrincipalPlanes{Valid: false}
	}
	return PrincipalPlanes{
		ImageSide:  (1 - m.D) / m.C,
		ObjectSide: (1 - m.A) / m.C,
		BackFocal:  -m.A / m.C,
		FrontFocal: m.D / m.C,
		Valid:      true,
	}
}

func (p PrincipalPlanes) String() string {
	if !p.Valid {
		return "afocal: principal planes undefined"
	}
	return fmt.Sprintf(
		"image_side=%g object_side=%g back_focal=%g front_focal=%g",
		p.ImageSide, p.ObjectSide, p.BackFocal, p.FrontFocal,
	)
}
