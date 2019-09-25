package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// nil以外のエラー値を返却
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v is not support", float64(e))
}

func Sqrt(x float64) (float64, error) {
	// 負であれば0とnil以外のエラー値,正であれば値とnil
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
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
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
