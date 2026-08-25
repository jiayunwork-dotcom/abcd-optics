package system

import (
	"fmt"

	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

type Telescope struct {
	F1             float64
	F2             float64
	Spacing        float64
	Total          matrix.Mat2
	Afocal         bool
	Magnification  float64
	IdealSpacing   float64
}

func AnalyzeTelescope(f1, f2, spacing float64) (Telescope, error) {
	if f1 == 0 || f2 == 0 {
		return Telescope{}, fmt.Errorf("telescope focal lengths must be non-zero")
	}
	elements := []element.Element{
		element.NewThinLens(f1),
		element.NewSpace(spacing),
		element.NewThinLens(f2),
	}
	spec := element.NewSpec("telescope", 100, elements)
	sys, err := Compose(spec)
	if err != nil {
		return Telescope{}, err
	}
	afocal := matrix.AlmostZero(sys.Total.C)
	mag := Magnification(sys.Total, 0)
	if afocal {
		mag = -f2 / f1
	}
	return Telescope{
		F1:            f1,
		F2:            f2,
		Spacing:       spacing,
		Total:         sys.Total,
		Afocal:        afocal,
		Magnification: mag,
		IdealSpacing:  f1 + f2,
	}, nil
}

func (t Telescope) AfocalError() float64 {
	return t.Spacing - t.IdealSpacing
}

func (t Telescope) String() string {
	out := fmt.Sprintf("f1=%.6f f2=%.6f spacing=%.6f ideal=%.6f\n", t.F1, t.F2, t.Spacing, t.IdealSpacing)
	out += matrix.FormatMatrix("M", t.Total)
	out += fmt.Sprintf("afocal=%v\n", t.Afocal)
	out += fmt.Sprintf("magnification=%.6f (ideal -f2/f1 = %.6f)\n", t.Magnification, -t.F2/t.F1)
	return out
}
