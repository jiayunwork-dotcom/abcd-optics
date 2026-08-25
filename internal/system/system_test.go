package system

import (
	"testing"

	"abcd-optics/internal/element"
	"abcd-optics/internal/matrix"
)

func TestThinLensImagingFormula(t *testing.T) {
	f := 100.0
	s := 300.0
	spec := element.NewSpec("lens", s, []element.Element{element.NewThinLens(f)})
	sys, err := Compose(spec)
	if err != nil {
		t.Fatalf("compose failed: %v", err)
	}
	sPrime, err := ImageDistance(sys.Total, s)
	if err != nil {
		t.Fatalf("image distance failed: %v", err)
	}
	want := s * f / (s - f)
	if !matrix.AlmostEqual(sPrime, want) {
		t.Errorf("s'=%g, want %g from 1/s+1/s'=1/f", sPrime, want)
	}
	got := 1/s + 1/sPrime
	if !matrix.AlmostEqual(got, 1/f) {
		t.Errorf("1/s+1/s'=%g, want 1/f=%g", got, 1/f)
	}
}

func TestTelescopeAfocalMagnification(t *testing.T) {
	f1, f2 := 100.0, 50.0
	d := f1 + f2
	tele, err := AnalyzeTelescope(f1, f2, d)
	if err != nil {
		t.Fatalf("analyze failed: %v", err)
	}
	if !tele.Afocal {
		t.Errorf("spacing=%g should be afocal, C=%g", d, tele.Total.C)
	}
	if !matrix.AlmostEqual(tele.Magnification, -f2/f1) {
		t.Errorf("magnification=%g, want -f2/f1=%g", tele.Magnification, -f2/f1)
	}
	if tele.Magnification >= 0 {
		t.Errorf("telescope magnification must be negative, got %g", tele.Magnification)
	}
}

func TestTelescopeOffFocusHasPower(t *testing.T) {
	f1, f2 := 100.0, 50.0
	tele, err := AnalyzeTelescope(f1, f2, f1+f2+20)
	if err != nil {
		t.Fatalf("analyze failed: %v", err)
	}
	if tele.Afocal {
		t.Error("off-focus spacing must not be afocal")
	}
	if matrix.AlmostZero(tele.Total.C) {
		t.Error("off-focus spacing must produce non-zero C (focal power)")
	}
}

func TestParallelBeamBackFocal(t *testing.T) {
	f := 50.0
	lens := matrix.New(1, 0, -1/f, 1)
	eff, err := EffectiveFocalLength(lens)
	if err != nil {
		t.Fatalf("effective focal failed: %v", err)
	}
	if !matrix.AlmostEqual(eff, f) {
		t.Errorf("f_eff=%g, want -1/C=%g", eff, f)
	}
	bfd, err := BackFocalDistance(lens)
	if err != nil {
		t.Fatalf("back focal failed: %v", err)
	}
	if !matrix.AlmostEqual(bfd, f) {
		t.Errorf("back focal=%g, want %g", bfd, f)
	}
}

func TestImageDistanceConjugate(t *testing.T) {
	f := 100.0
	s := 200.0
	spec := element.NewSpec("lens", s, []element.Element{element.NewThinLens(f)})
	res, err := Solve(spec)
	if err != nil {
		t.Fatalf("solve failed: %v", err)
	}
	sBack, err := ObjectDistanceFor(res.System.Total, res.ImageDistance)
	if err != nil {
		t.Fatalf("object solve failed: %v", err)
	}
	if !matrix.AlmostEqual(sBack, s) {
		t.Errorf("recovered s=%g, want %g", sBack, s)
	}
}

func TestRoundTripRestoresIncident(t *testing.T) {
	elements := []element.Element{
		element.NewSpace(30),
		element.NewThinLens(80),
		element.NewSpace(45),
	}
	spec := element.NewSpec("trip", 100, elements)
	sys, err := Compose(spec)
	if err != nil {
		t.Fatalf("compose failed: %v", err)
	}
	incident := matrix.NewRay(1.2, 0.03)
	recovered, err := matrix.RoundTrip(sys.Matrices, incident)
	if err != nil {
		t.Fatalf("round trip failed: %v", err)
	}
	if !matrix.CloseRay(recovered, incident) {
		t.Errorf("recovered %v, want incident %v", recovered, incident)
	}
}

func TestAirSystemDeterminantUnit(t *testing.T) {
	elements := []element.Element{
		element.NewSpace(20),
		element.NewThinLens(75),
		element.NewSpace(40),
		element.NewThinLens(120),
	}
	spec := element.NewSpec("air", 100, elements)
	sys, err := Compose(spec)
	if err != nil {
		t.Fatalf("compose failed: %v", err)
	}
	if got := matrix.Determinant(sys.Total); !matrix.AlmostEqual(got, 1) {
		t.Errorf("det=%g, want 1 for air-to-air system", got)
	}
}

func TestMagnificationEquivalentForms(t *testing.T) {
	f := 100.0
	s := 250.0
	spec := element.NewSpec("lens", s, []element.Element{element.NewThinLens(f)})
	res, err := Solve(spec)
	if err != nil {
		t.Fatalf("solve failed: %v", err)
	}
	direct := res.Magnification
	fromDet := MagnificationFromDet(res.System.Total, s)
	if !matrix.AlmostEqual(direct, fromDet) {
		t.Errorf("m by A+C s'=%g, by det/(Cs+D)=%g", direct, fromDet)
	}
	want := -res.ImageDistance / s
	if !matrix.AlmostEqual(direct, want) {
		t.Errorf("m=%g, want -s'/s=%g", direct, want)
	}
}
