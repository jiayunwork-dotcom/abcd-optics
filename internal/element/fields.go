package element

import "fmt"

var requiredFields = map[Kind][]string{
	KindSpace:    {"length"},
	KindThinLens: {"focal"},
	KindRefract:  {"radius", "n1", "n2"},
}

var zeroMeansMissing = map[string]bool{
	"length": false,
	"focal":  true,
	"radius": true,
	"n1":     true,
	"n2":     true,
}

func checkRequiredFields(raw rawElement) error {
	kind := Kind(raw.Kind)
	fields, ok := requiredFields[kind]
	if !ok {
		_, err := commitUnknown(nil, fmt.Errorf("unknown element kind %q", raw.Kind))
		return err
	}
	for _, f := range fields {
		if missingField(raw, f) {
			return fmt.Errorf("kind %q missing required field %q", kind, f)
		}
	}
	return nil
}

func missingField(raw rawElement, field string) bool {
	if !zeroMeansMissing[field] {
		return false
	}
	switch field {
	case "length":
		return raw.Length == 0
	case "focal":
		return raw.Focal == 0
	case "radius":
		return raw.Radius == 0
	case "n1":
		return raw.N1 == 0
	case "n2":
		return raw.N2 == 0
	}
	return false
}
