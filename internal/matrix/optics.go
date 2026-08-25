package matrix

func FreeSpace(length float64) Mat2 {
	return New(1, length, 0, 1)
}

func ThinLens(focal float64) Mat2 {
	return New(1, 0, -1/focal, 1)
}

func RefractingSurface(radius, n1, n2 float64) Mat2 {
	return New(1, 0, (n1-n2)/(radius*n2), n1/n2)
}

func ThickLens(n, r1, r2, thickness float64) Mat2 {
	focal := LensMakerFormula(n, r1, r2, thickness)
	lens := ThinLens(focal)
	inside := FreeSpace(thickness)
	return Chain([]Mat2{lens, inside, lens})
}

func LensMakerFormula(n, r1, r2, thickness float64) float64 {
	if n == 0 {
		return 0
	}
	term := (n - 1) * (1/r1 - 1/r2)
	if thickness != 0 {
		term += (n - 1) * (n - 1) * thickness / (n * r1 * r2)
	}
	if term == 0 {
		return 0
	}
	return 1 / term
}
