package main

import (
	"fmt"
	"math"
)

func Sqrt(v float64) string {
	if v < 0 {
		return Sqrt(-v) + "i"
	}
	return fmt.Sprint(math.Sqrt(v))
}

// If with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if and else

func pow2(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("%g >= %g \n", v, lim)
	}
	return lim

}

func Sqrt1(x float64) float64 {

	var z float64 = 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z

}

func main() {
	fmt.Println(Sqrt(2), Sqrt(-4))

	fmt.Println(
		pow(2, 0.2, 10),
		pow(3, 3, 10),
	)

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 10),
	)
	fmt.Println(Sqrt1(2))
}

// z -= (z*z - x) / (2*z)
