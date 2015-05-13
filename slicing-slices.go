package main

import "fmt"

// START OMIT
func main() {
	fibo := []int{0, 1, 1, 2, 3, 5, 8, 13}
	fmt.Println("fibonacci ==", fibo)

	fmt.Println("fibo ==", fibo)
	fmt.Println("fibo[1:4] ==", fibo[1:4]) // HL

	// omissão do limite inferior implica 0
	fmt.Println("fibo[:3] ==", fibo[:3]) // HL

	// omissão do limite superior implica len(s)
	fmt.Println("fibo[4:] ==", fibo[4:]) // HL
}

// END OMIT
