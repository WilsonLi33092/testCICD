package calculations

import "math"

func Exponent(x,y int) int {
	return int(math.Pow(float64(x), float64(y)))
}