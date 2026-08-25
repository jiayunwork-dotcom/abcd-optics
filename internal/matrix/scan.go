package matrix

type ScanResult struct {
	Heights []float64
	Slopes  []float64
	Rays    []Ray
}

func ScanHeights(m Mat2, yMin, yMax float64, steps int) ScanResult {
	if steps < 2 {
		steps = 2
	}
	out := ScanResult{
		Heights: make([]float64, 0, steps),
		Slopes:  make([]float64, 0, steps),
		Rays:    make([]Ray, 0, steps),
	}
	for i := 0; i < steps; i++ {
		t := float64(i) / float64(steps-1)
		y := yMin + t*(yMax-yMin)
		r := Apply(m, NewRay(y, 0))
		out.Heights = append(out.Heights, r.Y)
		out.Slopes = append(out.Slopes, r.U)
		out.Rays = append(out.Rays, r)
	}
	return out
}

func ScanSlopes(m Mat2, uMin, uMax float64, steps int) ScanResult {
	if steps < 2 {
		steps = 2
	}
	out := ScanResult{
		Heights: make([]float64, 0, steps),
		Slopes:  make([]float64, 0, steps),
		Rays:    make([]Ray, 0, steps),
	}
	for i := 0; i < steps; i++ {
		t := float64(i) / float64(steps-1)
		u := uMin + t*(uMax-uMin)
		r := Apply(m, NewRay(0, u))
		out.Heights = append(out.Heights, r.Y)
		out.Slopes = append(out.Slopes, r.U)
		out.Rays = append(out.Rays, r)
	}
	return out
}

func (s ScanResult) MaxHeight() float64 {
	max := 0.0
	for _, h := range s.Heights {
		abs := h
		if abs < 0 {
			abs = -abs
		}
		if abs > max {
			max = abs
		}
	}
	return max
}

func (s ScanResult) MinHeight() float64 {
	min := 0.0
	for _, h := range s.Heights {
		if h < min {
			min = h
		}
	}
	return min
}
