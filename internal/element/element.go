package element

import (
	"fmt"
	"strings"

	"abcd-optics/internal/matrix"
)

type Kind string

const (
	KindSpace    Kind = "space"
	KindThinLens Kind = "thin_lens"
	KindRefract  Kind = "refraction"
)

type Element interface {
	Kind() Kind
	Matrix() matrix.Mat2
	Describe() string
	Validate() error
}

func KnownKinds() []string {
	return []string{string(KindSpace), string(KindThinLens), string(KindRefract)}
}

func DescribeElements(elements []Element) string {
	var b strings.Builder
	for i, e := range elements {
		b.WriteString(fmt.Sprintf("%2d: %s\n", i, e.Describe()))
	}
	return b.String()
}
