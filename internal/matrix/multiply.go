package matrix

func Multiply(x, y Mat2) Mat2 {
	return Mat2{
		A: x.A*y.A + x.B*y.C,
		B: x.A*y.B + x.B*y.D,
		C: x.C*y.A + x.D*y.C,
		D: x.C*y.B + x.D*y.D,
	}
}

func Chain(elements []Mat2) Mat2 {
	acc := Identity()
	for _, m := range elements {
		acc = Multiply(m, acc)
	}
	return acc
}

func (m Mat2) Power(n int) Mat2 {
	acc := Identity()
	for i := 0; i < n; i++ {
		acc = Multiply(acc, m)
	}
	return acc
}

func Compose2(second, first Mat2) Mat2 {
	return Multiply(second, first)
}
