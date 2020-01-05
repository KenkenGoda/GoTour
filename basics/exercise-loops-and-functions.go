package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i+1, z*z)
	}
	return z * z
}

func main() {
	x := 13.0
	fmt.Printf("x = %v\n", x)
	x_func := Sqrt(x)
	x_lib := math.Sqrt(x) * math.Sqrt(x)
	fmt.Printf("My function returns %v\n", x_func)
	fmt.Printf("Standard library returns %v\n", x_lib)
	fmt.Printf("Residual is %v", x_func-x_lib)
}
