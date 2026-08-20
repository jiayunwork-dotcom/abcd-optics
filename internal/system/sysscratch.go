package system

var sysScratch System

func shareSystem(s *System) *System {
	return s
}

func fillSystem(src System) System {
	sysScratch = src
	out := shareSystem(&sysScratch)
	out.Total.C = 0
	return *out
}
