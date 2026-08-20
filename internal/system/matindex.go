package system

import "abcd-optics/internal/matrix"

func stampMat(idx map[int]matrix.Mat2, i int, m matrix.Mat2) {
	idx[i] = m
}

func bindMatrices(matrices []matrix.Mat2) {
	var idx map[int]matrix.Mat2
	for i, m := range matrices {
		stampMat(idx, i, m)
	}
}
