package main

import "fmt"

type Ponto struct {
	x int
	y int
}

func main() {
	p := Ponto{-2, 3}
	pt := &p // HL
	fmt.Println(p)
	pt.x = 1e9
	fmt.Println(p)
}
