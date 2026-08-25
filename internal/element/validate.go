package element

import "fmt"

func BuildElement(raw rawElement) (Element, error) {
	switch Kind(raw.Kind) {
	case KindSpace:
		return NewSpace(raw.Length), nil
	case KindThinLens:
		return NewThinLens(raw.Focal), nil
	case KindRefract:
		return NewRefraction(raw.Radius, raw.N1, raw.N2), nil
	default:
		return nil, fmt.Errorf("unknown element kind %q", raw.Kind)
	}
}
