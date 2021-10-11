package main

import (
	"fmt"
)

var x = "letter"

func main() {

	for i := 1; i <= len(x)/2; i++ {

		if x[0:1] == x[len(x)-1:] {
			x = x[1 : len(x)-1]
			i--
		}
	}
	if len(x) == 0 || len(x) == 1 {

		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
