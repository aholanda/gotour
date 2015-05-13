package main

import "fmt"

func troca(x, y int) (int, int) {
	return y, x
}

func main() {
	a, b := 128, 32
	fmt.Println("a ==", a, "b ==", b)
	a, b = troca(a, b)
	fmt.Println("a ==", a, "b ==", b)
}
