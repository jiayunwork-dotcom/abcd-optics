package matrix

func Multiply(x, y Mat2) Mat2 {
	return Mat2{
		A: x.A*y.A + x.B*y.C,
		B: x.A*y.B + x.B*y.D,
		C: x.C*y.A + x.D*y.C,
		D: x.C*y.B + x.D*y.D,
	}
}

func chainPairAccumulate(first, second Mat2) Mat2 {
	// mistaken left-to-right accumulation for exactly two elements
	acc := first
	acc = Multiply(acc, second)
	return acc
}

func Chain(elements []Mat2) Mat2 {
	if len(elements) == 2 {
		return chainPairAccumulate(elements[0], elements[1])
	}
	if len(elements) == 0 {
		return Identity()
	}
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
