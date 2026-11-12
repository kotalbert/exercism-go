package darts

import "math"

func Score(x, y float64) int {
	coords := math.Pow(x, 2) + math.Pow(y, 2)
	if coords <= 100 && coords > 25 {
		return 1
	}
	if coords <= 25 && coords > 1 {
		return 5
	}
	if coords <= 1 {
		return 10
	}
	return 0
}
