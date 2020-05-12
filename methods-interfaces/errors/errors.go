package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt to be raised when negative input
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

const delta = 1e-10

// Sqrt calculates square root
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	for {
		n := z - (z*z-x)/(2*z)
		if math.Abs(n-z) < delta {
			break
		}
		z = n
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
