package main

import "fmt"

type Ponto struct {
	x int
	y int
}

var (
	p0 = Ponto{-2, 3}  // tipo Ponto
	p1 = Ponto{x: 1e2} // y:0 é implicito
	p2 = Ponto{}       // x:0, y:0
	pt = &Ponto{3, -2} // tipo *Ponto
)

func main() {
	fmt.Println(p0, p1, p2, pt)
}
