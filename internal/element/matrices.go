package element

import (
	"fmt"

	"abcd-optics/internal/matrix"
)

func ElementMatrices(elements []Element) ([]matrix.Mat2, error) {
	out := make([]matrix.Mat2, 0, len(elements))
	for i, e := range elements {
		m := e.Matrix()
		if !m.IsFinite() {
			return nil, fmt.Errorf("element %d produces non-finite matrix", i)
		}
		out = append(out, m)
	}
	return out, nil
}

func SystemMatrix(elements []Element) (matrix.Mat2, error) {
	ms, err := ElementMatrices(elements)
	if err != nil {
		return matrix.Mat2{}, err
	}
	return matrix.Chain(ms), nil
}

func FirstKind(elements []Element) Kind {
	if len(elements) == 0 {
		return ""
	}
	return elements[0].Kind()
}

func LastKind(elements []Element) Kind {
	if len(elements) == 0 {
		return ""
	}
	return elements[len(elements)-1].Kind()
}

func CountKind(elements []Element, k Kind) int {
	n := 0
	for _, e := range elements {
		if e.Kind() == k {
			n++
		}
	}
	return n
}
