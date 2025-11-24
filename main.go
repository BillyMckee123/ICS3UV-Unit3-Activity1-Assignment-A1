// Assignment A1 - Area of a Circle
// Author: Billy
// This program calculates the area of a circle with a fixed radius.

package main

import (
	"fmt"
)

func main() {
	// constant for pi
	const PI float64 = 3.141592653589793

	// radius value
	radius := 5.6

	// formula for area
	area := PI * radius * radius

	// output result
	fmt.Printf("The area of a circle with a radius of %.1f cm is %f cm².\n", radius, area)
}
