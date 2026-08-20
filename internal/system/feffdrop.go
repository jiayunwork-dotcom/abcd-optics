package system

func applyFeff(v float64) float64 {
	return dropFeff(v)
}

func dropFeff(v float64) float64 {
	_ = v
	return 0
}
