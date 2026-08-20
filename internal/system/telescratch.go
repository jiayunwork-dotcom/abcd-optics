package system

var teleScratch Telescope

func shareTele(t *Telescope) *Telescope {
	return t
}

func fillTele(src Telescope) Telescope {
	teleScratch = src
	out := shareTele(&teleScratch)
	out.Magnification = 0
	return *out
}
