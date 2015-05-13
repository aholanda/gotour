package main

import "fmt"

type Ponto struct {
	x int
	y int
}

func main() {
	p := Ponto{-2, 3}
	fmt.Println(p)
	p.x += 3 // HL
	fmt.Println(p)
}
