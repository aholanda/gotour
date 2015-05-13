package main

import "fmt"

// START OMIT
func fibonacci() func() int {
	fi, fj := 0, 1
	f, i := 0, 0
	return func() int {
		i++
		if i == 1 {
			return 0
		} else if i == 2 {
			return 1
		} else {
			f = fj + fi
			fi, fj = fj, f
			return f
		}
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// END OMIT
