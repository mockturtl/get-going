package main

import (
	"fmt"
	"math"
)

func Sqrt(x int, z float64) (float64, float64) {
	i := 0
	delta := math.Abs(math.Sqrt(float64(x)) - z)
	epsilon := 1e-13
	for delta > epsilon {
		fmt.Printf("guess %d: %.20f\n", i, z)
		i++
		z = (z*z + float64(x)) / (2 * z)
		delta = math.Abs(math.Sqrt2 - z)
	}
	return z, delta
}

func main() {
	answer, delta := Sqrt(2, 1.5)
	fmt.Printf("computed: %.20f, actual: %.20f, delta: %.20f\n",
		answer,
		math.Sqrt2,
		delta)
}
