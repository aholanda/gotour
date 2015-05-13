package main

import "fmt"

// START OMIT
func main() {
	m := make(map[string]int)

	m["k1"] = 18
	m["k2"] = 21

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, ok := m["k2"]
	fmt.Printf("chave %q existe? %v\n", "k2", ok)
}

// END OMIT
