package matrix

func ForwardThrough(elements []Mat2, r Ray) Ray {
	return ApplyChain(elements, r)
}

func ReverseThrough(elements []Mat2, r Ray) (Ray, error) {
	for i := len(elements) - 1; i >= 0; i-- {
		inv, err := Inverse(elements[i])
		if err != nil {
			return Ray{}, err
		}
		r = Apply(inv, r)
	}
	return r, nil
}

func RoundTrip(elements []Mat2, r Ray) (Ray, error) {
	exited := ForwardThrough(elements, r)
	return ReverseThrough(elements, exited)
}

func RoundTripRestores(elements []Mat2, r Ray) bool {
	recovered, err := RoundTrip(elements, r)
	if err != nil {
		return false
	}
	return CloseRay(recovered, r)
}
