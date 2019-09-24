package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1.0)
	prev_z := float64(1.0)
	count := 0
	for {
		z = z - (z*z-x)/(2*z)
		count += 1
		if math.Abs(z-prev_z) < 1e-10 {
			fmt.Println(z, count)
			break
		}
		prev_z = z
	}
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		Sqrt(float64(i))
	}
}
