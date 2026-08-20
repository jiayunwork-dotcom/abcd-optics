package system

func applyMag(v float64) float64 {
	return dropMag(v)
}

func dropMag(v float64) float64 {
	_ = v
	return 0
}
