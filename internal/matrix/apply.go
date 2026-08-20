package matrix

type Ray struct {
	Y float64
	U float64
}

func NewRay(y, u float64) Ray {
	return Ray{Y: y, U: u}
}

func Apply(m Mat2, r Ray) Ray {
	return Ray{
		Y: m.A*r.Y + m.B*r.U,
		U: m.C*r.Y + m.D*r.U,
	}
}

func ApplyChain(elements []Mat2, r Ray) Ray {
	for _, m := range elements {
		r = Apply(m, r)
	}
	return r
}

func (r Ray) IsFinite() bool {
	return isFinite(r.Y) && isFinite(r.U)
}

func (r Ray) Scale(k float64) Ray {
	return Ray{Y: r.Y * k, U: r.U * k}
}
