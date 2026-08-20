package matrix

type Mat2 struct {
	A float64
	B float64
	C float64
	D float64
}

func New(a, b, c, d float64) Mat2 {
	return Mat2{A: a, B: b, C: c, D: d}
}

func Identity() Mat2 {
	return Mat2{A: 1, D: 1}
}

func Zero() Mat2 {
	return Mat2{}
}

func (m Mat2) Get(row, col int) float64 {
	switch {
	case row == 0 && col == 0:
		return m.A
	case row == 0 && col == 1:
		return m.B
	case row == 1 && col == 0:
		return m.C
	default:
		return m.D
	}
}

func (m Mat2) With(row, col int, value float64) Mat2 {
	switch {
	case row == 0 && col == 0:
		m.A = value
	case row == 0 && col == 1:
		m.B = value
	case row == 1 && col == 0:
		m.C = value
	default:
		m.D = value
	}
	return m
}

func (m Mat2) IsFinite() bool {
	return isFinite(m.A) && isFinite(m.B) && isFinite(m.C) && isFinite(m.D)
}

func (m Mat2) Transposed() Mat2 {
	return Mat2{A: m.A, B: m.C, C: m.B, D: m.D}
}
