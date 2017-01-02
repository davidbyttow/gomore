package math

import "math"

func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}
