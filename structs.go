package main

import "fmt"

type Ponto struct { // HL
	x int
	y int
}

func main() {
	p := Ponto{-2, 3}
	fmt.Println(p)
}
