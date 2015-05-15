package main

import "fmt"

func main() {
	soma := 1
	for soma < 1000 { // HL
		soma += soma
	} // HL
	fmt.Println(soma)
}
