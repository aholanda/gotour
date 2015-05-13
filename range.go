package main

import "fmt"

func main() {
	fibo := []int{0, 1, 1, 2, 3, 5, 8, 13} // HL
	fmt.Println("fibonacci ==", fibo)

	for i := range fibo { // HL
		fmt.Printf("fibo[%d] == %d\n", i, fibo[i])
	}
}
