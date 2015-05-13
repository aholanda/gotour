package main

import (
	"fmt"
	"math"
)

func main() {
	hipot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hipot(3, 4))
}
