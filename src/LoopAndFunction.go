package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	variant := 0.005

	z := x / 2

	dif := z*z - x


	for dif > variant || dif < -variant {
		z -= dif / (2*z)
		dif = z*z - x
	}

	return z
}

func main() {
	fmt.Println(Sqrt(8))
}