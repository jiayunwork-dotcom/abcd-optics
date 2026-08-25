package element

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"abcd-optics/internal/matrix"
)

type rawElement struct {
	Kind   string  `json:"kind"`
	Length float64 `json:"length"`
	Focal  float64 `json:"focal"`
	Radius float64 `json:"radius"`
	N1     float64 `json:"n1"`
	N2     float64 `json:"n2"`
}

type rawSpec struct {
	Name           string       `json:"name"`
	ObjectDistance float64      `json:"object_distance"`
	Ray            rawRay       `json:"ray"`
	Elements       []rawElement `json:"elements"`
}

type rawRay struct {
	Y float64 `json:"y"`
	U float64 `json:"u"`
}

func LoadSpecFile(path string) (Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Spec{}, fmt.Errorf("read spec: %w", err)
	}
	return ParseSpec(data)
}

func ParseSpec(data []byte) (Spec, error) {
	var raw rawSpec
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&raw); err != nil {
		return Spec{}, fmt.Errorf("decode spec: %w", err)
	}
	elements := make([]Element, 0, len(raw.Elements))
	for i, re := range raw.Elements {
		if err := checkRequiredFields(re); err != nil {
			return Spec{}, fmt.Errorf("element %d: %w", i, err)
		}
		e, err := BuildElement(re)
		if err != nil {
			return Spec{}, fmt.Errorf("element %d: %w", i, err)
		}
		elements = append(elements, e)
	}
	spec := NewSpec(raw.Name, raw.ObjectDistance, elements)
	spec.Incident = matrix.NewRay(raw.Ray.Y, raw.Ray.U)
	if err := spec.Validate(); err != nil {
		return Spec{}, err
	}
	return spec, nil
}
