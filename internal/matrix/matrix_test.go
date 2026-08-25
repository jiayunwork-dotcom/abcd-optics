package matrix

import "testing"

func TestMatrixMultiplicationOrder(t *testing.T) {
	lens := New(1, 0, -0.02, 1)
	space := New(1, 60, 0, 1)
	forward := Multiply(space, lens)
	reversed := Multiply(lens, space)
	if CloseMat(forward, reversed) {
		t.Errorf("forward and reversed multiplication must differ")
	}
	want := New(1-60*0.02, 60, -0.02, 1)
	if !CloseMat(forward, want) {
		t.Errorf("space*lens = %v, want %v", forward, want)
	}
}

func TestChainFollowsSystemOrder(t *testing.T) {
	space := New(1, 25, 0, 1)
	lens := New(1, 0, -0.05, 1)
	chain := Chain([]Mat2{space, lens})
	manual := Multiply(lens, space)
	if !CloseMat(chain, manual) {
		t.Errorf("chain = %v, want %v (first element applies first)", chain, manual)
	}
}

func TestDeterminantUnit(t *testing.T) {
	lens := New(1, 0, -0.02, 1)
	space := New(1, 60, 0, 1)
	cases := []struct {
		name string
		m    Mat2
	}{
		{"lens", lens},
		{"space", space},
		{"air system", Multiply(lens, space)},
	}
	for _, c := range cases {
		det := Determinant(c.m)
		if !AlmostEqual(det, 1) {
			t.Errorf("%s: det=%g, want 1", c.name, det)
		}
	}
}

func TestRefractionDetIsIndexRatio(t *testing.T) {
	refract := New(1, 0, 0.01, 1.0/1.5)
	if got := Determinant(refract); !AlmostEqual(got, 1.0/1.5) {
		t.Errorf("refraction det=%g, want 1/1.5", got)
	}
}

func TestInverseRoundTrip(t *testing.T) {
	elements := []Mat2{
		New(1, 10, 0, 1),
		New(1, 0, -0.02, 1),
		New(1, 30, 0, 1),
		New(1, 0, 0.015, 0.8),
	}
	r := NewRay(2, 0.05)
	recovered, err := RoundTrip(elements, r)
	if err != nil {
		t.Fatalf("round trip failed: %v", err)
	}
	if !CloseRay(recovered, r) {
		t.Errorf("round trip = %v, want incident %v", recovered, r)
	}
}

func TestSolve2x2(t *testing.T) {
	m := New(2, 1, 1, 2)
	ray, err := Solve2x2(m, NewRay(5, 5))
	if err != nil {
		t.Fatalf("solve failed: %v", err)
	}
	if !CloseRay(ray, NewRay(5.0/3, 5.0/3)) {
		t.Errorf("solve = %v, want (5/3, 5/3)", ray)
	}
}
