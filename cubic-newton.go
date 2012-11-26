package main

import (
	"log"
	"math/cmplx"
)

const N = 2

func cubrt(x complex128) (out complex128) {
	const epsilon = 1e-10
	const oneThird = 0.333333333333333333333333

	var root complex128 = cmplx.Pow(x, oneThird)
	var z complex128 = 1

	delta := cmplx.Abs(z - root)

	for delta > epsilon {
		z = 2*z/3 + x/(3*z*z)
		log.Printf("guess: %g, actual: %g, delta: %g\n", z, root, delta)
		delta = cmplx.Abs(z - root)
	}

	out = z
	return
}

func main() {
	log.Printf("The cube root of %d is %g\n", N, cubrt(N))
}
