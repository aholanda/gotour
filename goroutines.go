package main

import (
	"fmt"
	"time"
)

// START OMIT
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(50 * time.Millisecond)
	}
}
func main() {
	f("síncrono") // síncrono

	go f("goroutine") // assíncrona e concorrente

	// função anônima, também assíncrona e concorrente
	go func(msg string) {
		fmt.Println(msg)
	}("anon")

	var input string
	fmt.Scanln(&input) // pressione tecla para sair
	fmt.Println("fim")
}

// END OMIT
