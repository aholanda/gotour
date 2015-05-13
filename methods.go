package main

import "fmt"

type Ponto struct {
	x, y int
}

func (p *Ponto) Deslocar(deltaX, deltaY int) {
	p.x, p.y = p.x+deltaX, p.y+deltaY
}

func main() {
	p := &Ponto{3, 4}
	fmt.Printf("de %v ", p)
	p.Deslocar(2, -8)
	fmt.Printf("para %v\n", p)
}
