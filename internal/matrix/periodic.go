package matrix

func PeriodicChain(unit []Mat2, repeats int) Mat2 {
	acc := Identity()
	for i := 0; i < repeats; i++ {
		acc = Multiply(Chain(unit), acc)
	}
	return acc
}

func ConjugateTranspose(m Mat2) Mat2 {
	det := Determinant(m)
	return New(m.D/det, -m.B/det, -m.C/det, m.A/det)
}

func Scale(m Mat2, k float64) Mat2 {
	return New(m.A*k, m.B*k, m.C*k, m.D*k)
}

func Add(a, b Mat2) Mat2 {
	return New(a.A+b.A, a.B+b.B, a.C+b.C, a.D+b.D)
}
