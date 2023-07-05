package mathutils

import "github.com/samber/lo"

// Min Returns the smaller of two int values.
func Min(a, b int) int {
	return lo.Ternary(a <= b, a, b)
}

// Max Returns the larger of two int values.
func Max(a, b int) int {
	return lo.Ternary(a >= b, a, b)
}
