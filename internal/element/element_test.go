package element

import "testing"

func TestParseRejectsUnknownKind(t *testing.T) {
	data := []byte(`{"object_distance":50,"elements":[{"kind":"prism","angle":30}]}`)
	if _, err := ParseSpec(data); err == nil {
		t.Error("expected error for unknown element kind, got nil")
	}
}

func TestParseRejectsZeroFocal(t *testing.T) {
	data := []byte(`{"object_distance":50,"elements":[{"kind":"thin_lens","focal":0}]}`)
	if _, err := ParseSpec(data); err == nil {
		t.Error("expected error for zero focal length, got nil")
	}
}

func TestParseRejectsNegativeIndex(t *testing.T) {
	data := []byte(`{"object_distance":50,"elements":[{"kind":"refraction","radius":20,"n1":-1.5,"n2":1}]}`)
	if _, err := ParseSpec(data); err == nil {
		t.Error("expected error for negative refractive index, got nil")
	}
}

func TestParseRejectsNegativeLength(t *testing.T) {
	data := []byte(`{"object_distance":50,"elements":[{"kind":"space","length":-10}]}`)
	if _, err := ParseSpec(data); err == nil {
		t.Error("expected error for negative propagation length, got nil")
	}
}

func TestParseRejectsMissingField(t *testing.T) {
	data := []byte(`{"object_distance":50,"elements":[{"kind":"refraction","radius":20,"n1":1}]}`)
	if _, err := ParseSpec(data); err == nil {
		t.Error("expected error for missing n2 field, got nil")
	}
}

func TestParseAllowsZeroLengthSpace(t *testing.T) {
	data := []byte(`{"object_distance":50,"elements":[{"kind":"space","length":0}]}`)
	if _, err := ParseSpec(data); err != nil {
		t.Errorf("zero-length space should be valid, got %v", err)
	}
}

func TestParseValidTelescope(t *testing.T) {
	data := []byte(`{
		"name":"telescope",
		"object_distance":500,
		"elements":[
			{"kind":"thin_lens","focal":100},
			{"kind":"space","length":150},
			{"kind":"thin_lens","focal":50}
		]
	}`)
	spec, err := ParseSpec(data)
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if len(spec.Elements) != 3 {
		t.Errorf("got %d elements, want 3", len(spec.Elements))
	}
	if spec.ObjectDistance != 500 {
		t.Errorf("object distance %g, want 500", spec.ObjectDistance)
	}
}
