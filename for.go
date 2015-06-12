package main

import "fmt"
//
func main() {
	soma := 0
	for i := 0; i < 16; i++ { // HL
		soma += i
	} // HL
	fmt.Println(soma)
}
