package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // aponta para i
	fmt.Println(*p) // lê i pelo ponteiro
	*p = 21         // atribui i pelo ponteiro
	fmt.Println(i)  // novo valor de i

	p = &j         // aponta para j
	*p = *p / 37   // divide j pelo ponteiro
	fmt.Println(j) // novo valor de j
}
