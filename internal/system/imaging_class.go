package system

import (
	"fmt"

	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

type ImageClass int

const (
	ImageUnknown ImageClass = iota
	ImageReal
	ImageVirtual
	ImageAtInfinity
)

func (c ImageClass) String() string {
	switch c {
	case ImageReal:
		return "real"
	case ImageVirtual:
		return "virtual"
	case ImageAtInfinity:
		return "at_infinity"
	default:
		return "unknown"
	}
}

func ClassifyImage(m matrix.Mat2, objectDistance float64) ImageClass {
	sPrime, err := ImageDistance(m, objectDistance)
	if err != nil {
		return ImageAtInfinity
	}
	if sPrime > 0 {
		return ImageReal
	}
	return ImageVirtual
}

func MagnificationSign(m matrix.Mat2, objectDistance float64) string {
	sPrime, err := ImageDistance(m, objectDistance)
	if err != nil {
		return "undefined"
	}
	mag := Magnification(m, sPrime)
	if mag > 0 {
		return "upright"
	}
	if mag < 0 {
		return "inverted"
	}
	return "zero"
}

func ClassifySystem(m matrix.Mat2) string {
	if matrix.AlmostZero(m.C) {
		return "afocal"
	}
	focal, _ := EffectiveFocalLength(m)
	if focal > 0 {
		return "converging"
	}
	return "diverging"
}

func DescribeImaging(spec element.Spec) (string, error) {
	sys, err := Compose(spec)
	if err != nil {
		return "", err
	}
	cls := ClassifyImage(sys.Total, spec.ObjectDistance)
	sign := MagnificationSign(sys.Total, spec.ObjectDistance)
	kind := ClassifySystem(sys.Total)
	return fmt.Sprintf(
		"system=%s image=%s magnification=%s",
		kind, cls, sign,
	), nil
}
