package element

import "fmt"

type FocalInfo struct {
	ImageSide  float64
	ObjectSide float64
	Valid      bool
}

func RefractionFocalInfo(r Refraction) FocalInfo {
	if r.N2 == r.N1 || r.Radius == 0 {
		return FocalInfo{Valid: false}
	}
	imageFocal := r.N2 * r.Radius / (r.N2 - r.N1)
	objectFocal := -r.N1 * r.Radius / (r.N2 - r.N1)
	return FocalInfo{
		ImageSide:  imageFocal,
		ObjectSide: objectFocal,
		Valid:      true,
	}
}

func (f FocalInfo) String() string {
	if !f.Valid {
		return "no single-surface focal power"
	}
	return fmt.Sprintf("f_image=%g f_object=%g", f.ImageSide, f.ObjectSide)
}

func ThinLensFocalInfo(l ThinLens) FocalInfo {
	return FocalInfo{
		ImageSide:  l.Focal,
		ObjectSide: -l.Focal,
		Valid:      true,
	}
}
