package main

import "fmt"

func fibonacci() func() int {
	s1 := 0
	s2 := 0
	a := 1
	return func() int {
		a += s1
		s1 = s2
		s2 = a
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
