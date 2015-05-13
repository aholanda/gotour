package main

import "fmt"

func count(n int, c chan int) {
	for i := 0; i < n+1; i++ {
		c <- i
	}
	close(c) // HL
}

func main() {
	ch := make(chan int, 10)

	go count(cap(ch), ch)

	for i := range ch { // HL
		fmt.Println(i)
	}
}
