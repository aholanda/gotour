package main

import "fmt"

func count(c chan int) {
	for i := 0; i < 3; i++ {
		c <- i * 2 // envia i*2 para o canal
	}
}

func main() {
	rcv := make(chan int) // channel do tipo int

	go count(rcv)

	// canais x0, x1 e x2 recebem do canal snd
	x0, x1, x2 := <-rcv, <-rcv, <-rcv
	fmt.Println(x0, x1, x2)
}
