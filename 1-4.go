package main

import (
	"fmt"
)

const (
	pi = 3.14
)

func circumference(radius float32) float32 {
	return 2 * pi * radius
}

func main() {
	var radius float32

	fmt.Println("Input a radius value:")
	fmt.Scanf("%f", &radius)
	length := circumference(radius)

	fmt.Println("The circumference of the circle is:", length, "radians")
}
