package matrix

func (m Mat2) IsValid() bool {
	return m.IsFinite()
}

func (m Mat2) ToSlice() [4]float64 {
	return [4]float64{m.A, m.B, m.C, m.D}
}

func FromSlice(v [4]float64) Mat2 {
	return New(v[0], v[1], v[2], v[3])
}

func (m Mat2) MulBy(r Ray) Ray {
	return Apply(m, r)
}

func RayHeight(r Ray) float64 {
	return r.Y
}

func RaySlope(r Ray) float64 {
	return r.U
}

func OnAxis(r Ray) bool {
	return AlmostZero(r.Y)
}

func Collimated(r Ray) bool {
	return AlmostZero(r.U)
}
