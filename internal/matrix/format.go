package matrix

import (
	"fmt"
	"strings"
)

func (m Mat2) String() string {
	return fmt.Sprintf("[%g %g; %g %g]", m.A, m.B, m.C, m.D)
}

func (r Ray) String() string {
	return fmt.Sprintf("(y=%g, u=%g)", r.Y, r.U)
}

func FormatMatrix(name string, m Mat2) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s = [[%10.6f %10.6f]\n", name, m.A, m.B))
	b.WriteString(fmt.Sprintf("        [%10.6f %10.6f]]\n", m.C, m.D))
	return b.String()
}

func FormatRow(name string, values []float64) string {
	parts := make([]string, 0, len(values))
	for _, v := range values {
		parts = append(parts, fmt.Sprintf("%10.6f", v))
	}
	return fmt.Sprintf("%s = [%s]", name, strings.Join(parts, " "))
}
