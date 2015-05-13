package main

import (
	"fmt"
	"time"
)

// START OMIT
func f1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "um"
}
func f2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "dois"
}
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go f1(c1)

	go f2(c2)

	for i := 0; i < 2; i++ {
		select { // HL
		case msg1 := <-c1:
			fmt.Printf("recebido %q\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("recebido %q\n", msg2)
		}
	}
}

// END OMIT
