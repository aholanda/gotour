package main

import "fmt"

func count(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i * 3
	}
}

func main() {
	//  aceita até 3 inteiros no canal sem bloquear
	ch := make(chan int, 3) // HL

	go count(cap(ch), ch)

	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
}
