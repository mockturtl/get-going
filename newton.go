package main

import (
	"errors"
	"fmt"
	"math"
)

const precision = 20

func Sqrt(x int, z float64) (out float64, delta float64, err error) {
	if x < 0 {
		out, delta = math.NaN(), math.NaN()
		err = errors.New(fmt.Sprint("cannot Sqrt negative number: ", x))
		return
	}

	i := 0
	const epsilon = 1e-13
	delta = math.Abs(math.Sqrt(float64(x)) - z)
	for delta > epsilon {
		fmt.Printf("guess %d: %.*f\n", i, precision, z)
		i++
		z = (z*z + float64(x)) / (2 * z)
		delta = math.Abs(math.Sqrt2 - z)
	}
	out = z
	return
}

func main() {
	answer, delta, err := Sqrt(2, 1.5)

	if err != nil {
		panic(err)
	}

	fmt.Printf("computed: %.*f, actual: %.*f, delta: %.*f\n",
		precision,
		answer,
		precision,
		math.Sqrt2,
		precision,
		delta)
}
