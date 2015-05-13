package main

import "fmt"

type Boolean bool

func (b0 Boolean) And(b1 Boolean) bool {
	return bool(b0 && b1)
}

func main() {
	x, y := Boolean(true), Boolean(false)
	fmt.Printf("%v AND %v = %v \n", x, x, x.And(x))
	fmt.Printf("%v AND %v = %v \n", y, y, y.And(y))
	fmt.Printf("%v AND %v = %v \n", x, y, x.And(y))
}
