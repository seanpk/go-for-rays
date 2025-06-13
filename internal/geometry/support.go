package geometry

import (
	"math"
)

const EPSILON = 1e-6

func epsilonOrDefault(epsilon ...float64) float64 {
	if len(epsilon) > 0 && epsilon[0] > 0 {
		return epsilon[0]
	}
	return EPSILON
}

func IsNearTo(a, b float64, epsilon ...float64) bool {
	eps := epsilonOrDefault(epsilon...)
	return math.Abs(a-b) < eps
}
