package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {
	x, y, z := 0, 0, 0
	return func() int {
		if y == 0 {
			y += 1
		} else if z == 0 {
			z = x + y
		} else {
			z = x + y
			x = y
			y = z
		}
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
